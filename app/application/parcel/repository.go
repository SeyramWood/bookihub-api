package parcel

import (
	"fmt"
	"strings"

	"golang.org/x/net/context"

	"github.com/SeyramWood/bookibus/app/adapters/gateways"
	"github.com/SeyramWood/bookibus/app/adapters/presenters"
	"github.com/SeyramWood/bookibus/app/application"
	requeststructs "github.com/SeyramWood/bookibus/app/domain/request_structs"
	"github.com/SeyramWood/bookibus/app/framework/database"
	"github.com/SeyramWood/bookibus/ent"
	"github.com/SeyramWood/bookibus/ent/company"
	"github.com/SeyramWood/bookibus/ent/companyuser"
	"github.com/SeyramWood/bookibus/ent/parcel"
	"github.com/SeyramWood/bookibus/ent/parcelimage"
)

type repository struct {
	db  *ent.Client
	ctx context.Context
}

func NewRepository(db *database.Adapter) gateways.ParcelRepo {
	return &repository{
		db:  db.DB,
		ctx: context.Background(),
	}
}

// Delete implements gateways.ParcelRepo.
func (r *repository) Delete(id int) error {
	return r.db.Parcel.DeleteOneID(id).Exec(r.ctx)
}

// DeleteImage implements gateways.ParcelRepo.
func (r *repository) DeleteImage(id int) error {
	return r.db.ParcelImage.DeleteOneID(id).Exec(r.ctx)
}

// Insert implements gateways.ParcelRepo.
func (r *repository) Insert(companyId int, request *requeststructs.ParcelRequest, refResponse *requeststructs.PaymentReferenceResponse, images []string) (*ent.Parcel, error) {
	tx, err := r.db.Tx(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("error starting a transaction: %w", err)
	}
	result, err := tx.Parcel.Create().
		SetParcelCode(fmt.Sprintf("PC_%s", application.OTP(9))).
		SetSenderName(request.SenderName).
		SetSenderPhone(request.SenderPhone).
		SetSenderEmail(request.SenderEmail).
		SetRecipientName(request.RecipientName).
		SetRecipientPhone(request.RecipientPhone).
		SetRecipientLocation(request.RecipientLocation).
		SetAmount(request.Amount).
		SetPaidAt(application.ParseRFC3339Datetime(refResponse.PaidAt)).
		SetTansType(parcel.TansType(refResponse.TransType)).
		SetWeight(request.Weight).
		SetType(request.Type).
		SetTripID(request.TripID).
		SetDriverID(request.DriverID).
		SetCompanyID(companyId).
		Save(r.ctx)
	if err != nil {
		return nil, application.Rollback(tx, fmt.Errorf("failed creating package: %w", err))
	}
	_, err = tx.ParcelImage.MapCreateBulk(images, func(create *ent.ParcelImageCreate, i int) {
		create.SetImage(images[i]).SetParcel(result)
	}).Save(r.ctx)
	if err != nil {
		return nil, application.Rollback(tx, fmt.Errorf("failed creating package image: %w", err))
	}
	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed committing package creation transaction: %w", err)
	}
	return r.Read(result.ID)
}

// InsertImage implements gateways.ParcelRepo.
func (r *repository) InsertImage(id int, request []string) (*ent.Parcel, error) {
	tx, err := r.db.Tx(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("error starting a transaction: %w", err)
	}
	_, err = tx.ParcelImage.MapCreateBulk(request, func(create *ent.ParcelImageCreate, i int) {
		create.SetImage(request[i]).SetParcelID(id)
	}).Save(r.ctx)
	if err != nil {
		return nil, application.Rollback(tx, fmt.Errorf("failed package image: %w", err))
	}
	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed committing package image transaction: %w", err)
	}
	return r.Read(id)
}

// Read implements gateways.ParcelRepo.
func (r *repository) Read(id int) (*ent.Parcel, error) {
	result, err := r.db.Parcel.Query().Where(parcel.ID(id)).
		WithImages().
		WithTrip(func(tq *ent.TripQuery) {
			tq.WithVehicle(func(vq *ent.VehicleQuery) {
				vq.WithImages()
			})
			tq.WithRoute(func(rq *ent.RouteQuery) {
				rq.WithStops()
			})
			tq.WithDriver()
			tq.WithCompany()
		}).
		Only(r.ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReadByCode implements gateways.ParcelRepo.
func (r *repository) ReadByCode(code string) (*ent.Parcel, error) {
	result, err := r.db.Parcel.Query().Where(parcel.ParcelCode(code)).
		WithImages().
		WithTrip(func(tq *ent.TripQuery) {
			tq.WithVehicle(func(vq *ent.VehicleQuery) {
				vq.WithImages()
			})
			tq.WithRoute(func(rq *ent.RouteQuery) {
				rq.WithStops()
			})
			tq.WithDriver()
			tq.WithCompany()
		}).
		Only(r.ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReadImage implements gateways.IncidentRepo.
func (r *repository) ReadImage(id int) (*ent.ParcelImage, error) {
	return r.db.ParcelImage.Get(r.ctx, id)
}

// ReadAll implements gateways.ParcelRepo.
func (r *repository) ReadAll(limit int, offset int, filter *requeststructs.ParcelFilterRequest) (*presenters.PaginationResponse, error) {
	if filter.Status != "" {
		query := r.db.Parcel.Query().Where(parcel.StatusEQ(parcel.Status(strings.ToLower(filter.Status))))
		return r.filterParcel(query, limit, offset)
	}

	return r.filterParcel(r.db.Parcel.Query(), limit, offset)
}

// ReadAllByCompany implements gateways.ParcelRepo.
func (r *repository) ReadAllByCompany(companyId int, limit int, offset int, filter *requeststructs.ParcelFilterRequest) (*presenters.PaginationResponse, error) {
	if filter.Status != "" {
		query := r.db.Parcel.Query().Where(
			parcel.And(
				parcel.HasCompanyWith(company.ID(companyId)),
				parcel.StatusEQ(parcel.Status(strings.ToLower(filter.Status))),
			),
		)
		return r.filterParcel(query, limit, offset)
	}

	return r.filterParcel(r.db.Parcel.Query().Where(parcel.HasCompanyWith(company.ID(companyId))), limit, offset)
}

// ReadAllByDriver implements gateways.ParcelRepo.
func (r *repository) ReadAllByDriver(driverId int, limit int, offset int, filter *requeststructs.ParcelFilterRequest) (*presenters.PaginationResponse, error) {
	if filter.Status != "" {
		query := r.db.Parcel.Query().Where(
			parcel.And(
				parcel.HasDriverWith(companyuser.ID(driverId)),
				parcel.StatusEQ(parcel.Status(strings.ToLower(filter.Status))),
			),
		)
		return r.filterParcel(query, limit, offset)
	}
	return r.filterParcel(r.db.Parcel.Query().Where(parcel.HasDriverWith(companyuser.ID(driverId))), limit, offset)
}

// Update implements gateways.ParcelRepo.
func (r *repository) Update(id int, request *requeststructs.ParcelUpdateRequest) (*ent.Parcel, error) {
	_, err := r.db.Parcel.UpdateOneID(id).
		SetSenderName(request.SenderName).
		SetSenderPhone(request.SenderPhone).
		SetSenderEmail(request.SenderEmail).
		SetRecipientName(request.RecipientName).
		SetRecipientPhone(request.RecipientPhone).
		SetRecipientLocation(request.RecipientLocation).
		SetType(request.Type).
		Save(r.ctx)
	if err != nil {
		return nil, err
	}
	return r.Read(id)
}

// UpdateImage implements gateways.ParcelRepo.
func (r *repository) UpdateImage(id int, request string) (*ent.ParcelImage, error) {
	_, err := r.db.ParcelImage.UpdateOneID(id).
		SetImage(request).
		Save(r.ctx)
	if err != nil {
		return nil, err
	}
	return r.ReadImage(id)
}

// UpdateStatus implements gateways.ParcelRepo.
func (r *repository) UpdateStatus(id int, images []string) (*ent.Parcel, error) {
	tx, err := r.db.Tx(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("error starting a transaction: %w", err)
	}
	_, err = tx.ParcelImage.MapCreateBulk(images, func(create *ent.ParcelImageCreate, i int) {
		create.SetImage(images[i]).SetKind(parcelimage.KindRecipient).SetParcelID(id)
	}).Save(r.ctx)
	if err != nil {
		return nil, application.Rollback(tx, fmt.Errorf("failed creating package image: %w", err))
	}
	_, err = tx.Parcel.UpdateOneID(id).
		SetStatus(parcel.StatusDelivered).
		Save(r.ctx)
	if err != nil {
		return nil, application.Rollback(tx, fmt.Errorf("failed updating package status: %w", err))
	}
	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed committing package update transaction: %w", err)
	}
	return r.Read(id)
}

func (r *repository) filterParcel(query *ent.ParcelQuery, limit, offset int) (*presenters.PaginationResponse, error) {
	count := query.CountX(r.ctx)
	results, err := query.
		Limit(limit).
		Offset(offset).
		Order(ent.Desc(parcel.FieldCreatedAt)).
		WithImages().
		WithTrip(func(tq *ent.TripQuery) {
			tq.WithVehicle(func(vq *ent.VehicleQuery) {
				vq.WithImages()
			})
			tq.WithRoute(func(rq *ent.RouteQuery) {
				rq.WithStops()
			})
			tq.WithDriver()
			tq.WithCompany()
		}).
		All(r.ctx)
	if err != nil {
		return nil, err
	}
	return application.Paginate(count, results)
}
