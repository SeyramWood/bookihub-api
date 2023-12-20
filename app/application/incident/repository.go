package incident

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"golang.org/x/net/context"

	"github.com/SeyramWood/bookibus/app/adapters/gateways"
	"github.com/SeyramWood/bookibus/app/adapters/presenters"
	"github.com/SeyramWood/bookibus/app/application"
	requeststructs "github.com/SeyramWood/bookibus/app/domain/request_structs"
	"github.com/SeyramWood/bookibus/app/framework/database"
	"github.com/SeyramWood/bookibus/ent"
	"github.com/SeyramWood/bookibus/ent/company"
	"github.com/SeyramWood/bookibus/ent/companyuser"
	"github.com/SeyramWood/bookibus/ent/incident"
)

type repository struct {
	db  *ent.Client
	ctx context.Context
}

func NewRepository(db *database.Adapter) gateways.IncidentRepo {
	return &repository{
		db:  db.DB,
		ctx: context.Background(),
	}
}

// Delete implements gateways.IncidentRepo.
func (r *repository) Delete(id int) error {
	return r.db.Incident.DeleteOneID(id).Exec(r.ctx)
}

// DeleteImage implements gateways.IncidentRepo.
func (r *repository) DeleteImage(id int) error {
	return r.db.IncidentImage.DeleteOneID(id).Exec(r.ctx)
}

// DeleteAudio implements gateways.IncidentRepo.
func (r *repository) DeleteAudio(id int) error {
	_, err := r.db.Incident.UpdateOneID(id).SetAudio("").Save(r.ctx)
	return err
}

// Insert implements gateways.IncidentRepo.
func (r *repository) Insert(companyId int, request *requeststructs.IncidentRequest, images []string, audio ...string) (*ent.Incident, error) {
	tx, err := r.db.Tx(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("error starting a transaction: %w", err)
	}
	var result *ent.Incident
	if len(audio) > 0 {
		res, err := tx.Incident.Create().
			SetTime(application.ParseRFC3339Datetime(request.Time)).
			SetLocation(request.Location).
			SetDescription(request.Description).
			SetAudio(audio[0]).
			SetType(request.Type).
			SetTripID(request.TripID).
			SetDriverID(request.DriverID).
			SetCompanyID(companyId).
			Save(r.ctx)
		if err != nil {
			return nil, application.Rollback(tx, fmt.Errorf("failed incident: %w", err))
		}
		result = res
	} else {
		res, err := tx.Incident.Create().
			SetTime(application.ParseRFC3339Datetime(request.Time)).
			SetLocation(request.Location).
			SetDescription(request.Description).
			SetType(request.Type).
			SetTripID(request.TripID).
			SetDriverID(request.DriverID).
			SetCompanyID(companyId).
			Save(r.ctx)
		if err != nil {
			return nil, application.Rollback(tx, fmt.Errorf("failed incident : %w", err))
		}
		result = res
	}
	_, err = tx.IncidentImage.MapCreateBulk(images, func(create *ent.IncidentImageCreate, i int) {
		create.SetImage(images[i]).SetIncident(result)
	}).Save(r.ctx)
	if err != nil {
		return nil, application.Rollback(tx, fmt.Errorf("failed incident image: %w", err))
	}
	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed committing incident creation transaction: %w", err)
	}
	return r.Read(result.ID)
}

// InsertImage implements gateways.IncidentRepo.
func (r *repository) InsertImage(id int, request []string) (*ent.Incident, error) {
	tx, err := r.db.Tx(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("error starting a transaction: %w", err)
	}
	_, err = tx.IncidentImage.MapCreateBulk(request, func(create *ent.IncidentImageCreate, i int) {
		create.SetImage(request[i]).SetIncidentID(id)
	}).Save(r.ctx)
	if err != nil {
		return nil, application.Rollback(tx, fmt.Errorf("failed incident image: %w", err))
	}
	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed committing incident image creation transaction: %w", err)
	}
	return r.Read(id)
}

// Read implements gateways.IncidentRepo.
func (r *repository) Read(id int) (*ent.Incident, error) {
	result, err := r.db.Incident.Query().
		Where(incident.ID(id)).
		WithImages().
		WithTrip(func(tq *ent.TripQuery) {
			tq.WithVehicle(func(vq *ent.VehicleQuery) {
				vq.WithImages()
			})
			tq.WithRoute()
			tq.WithStops()
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
func (r *repository) ReadImage(id int) (*ent.IncidentImage, error) {
	return r.db.IncidentImage.Get(r.ctx, id)
}

// ReadAll implements gateways.IncidentRepo.
func (r *repository) ReadAll(limit int, offset int, filter *requeststructs.IncidentFilterRequest) (*presenters.PaginationResponse, error) {
	if filter.Datetime != "" {
		query := r.db.Incident.Query().Where(
			func(s *sql.Selector) {
				s.Where(sql.ExprP(fmt.Sprintf("DATE(%s) = ?", incident.FieldTime), application.ParseRFC3339MYSQLDatetime(filter.Datetime)))
			},
		)
		return r.filterIncident(query, limit, offset)
	}
	return r.filterIncident(r.db.Incident.Query(), limit, offset)
}

// ReadAllByCompany implements gateways.IncidentRepo.
func (r *repository) ReadAllByCompany(companyId int, limit int, offset int, filter *requeststructs.IncidentFilterRequest) (*presenters.PaginationResponse, error) {
	if filter.Datetime != "" {
		query := r.db.Incident.Query().Where(
			incident.And(
				incident.HasCompanyWith(company.ID(companyId)),
				func(s *sql.Selector) {
					s.Where(sql.ExprP(fmt.Sprintf("DATE(%s) = ?", incident.FieldTime), application.ParseRFC3339MYSQLDatetime(filter.Datetime)))
				},
			),
		)
		return r.filterIncident(query, limit, offset)
	}
	return r.filterIncident(r.db.Incident.Query().Where(incident.HasCompanyWith(company.ID(companyId))), limit, offset)
}

// ReadAllByDriver implements gateways.IncidentRepo.
func (r *repository) ReadAllByDriver(driverId int, limit int, offset int, filter *requeststructs.IncidentFilterRequest) (*presenters.PaginationResponse, error) {
	driverID := r.db.User.GetX(r.ctx, driverId).QueryCompanyUser().OnlyIDX(r.ctx)
	if filter.Datetime != "" {
		query := r.db.Incident.Query().Where(
			incident.And(
				incident.HasDriverWith(companyuser.ID(driverID)),
				func(s *sql.Selector) {
					s.Where(sql.ExprP(fmt.Sprintf("DATE(%s) = ?", incident.FieldTime), application.ParseRFC3339MYSQLDatetime(filter.Datetime)))
				},
			),
		)
		return r.filterIncident(query, limit, offset)
	}
	return r.filterIncident(r.db.Incident.Query().Where(incident.HasDriverWith(companyuser.ID(driverID))), limit, offset)
}

// Update implements gateways.IncidentRepo.
func (r *repository) Update(id int, request *requeststructs.IncidentUpdateRequest) (*ent.Incident, error) {
	_, err := r.db.Incident.UpdateOneID(id).
		SetLocation(request.Location).
		SetDescription(request.Description).
		SetType(request.Type).
		Save(r.ctx)
	if err != nil {
		return nil, err
	}
	return r.Read(id)
}

// UpdateStatus implements gateways.IncidentRepo.
func (r *repository) UpdateStatus(id int, status string) (*ent.Incident, error) {
	_, err := r.db.Incident.UpdateOneID(id).
		SetStatus(incident.Status(status)).
		Save(r.ctx)
	if err != nil {
		return nil, err
	}
	return r.Read(id)
}

// UpdateAudio implements gateways.IncidentRepo.
func (r *repository) UpdateAudio(id int, request string) (*ent.Incident, error) {
	_, err := r.db.Incident.UpdateOneID(id).
		SetAudio(request).
		Save(r.ctx)
	if err != nil {
		return nil, err
	}
	return r.Read(id)
}

// UpdateImage implements gateways.IncidentRepo.
func (r *repository) UpdateImage(id int, request string) (*ent.IncidentImage, error) {
	_, err := r.db.IncidentImage.UpdateOneID(id).
		SetImage(request).
		Save(r.ctx)
	if err != nil {
		return nil, err
	}
	return r.ReadImage(id)
}

func (r *repository) filterIncident(query *ent.IncidentQuery, limit, offset int) (*presenters.PaginationResponse, error) {
	count := query.CountX(r.ctx)
	results, err := query.
		Limit(limit).
		Offset(offset).
		Order(ent.Desc(incident.FieldCreatedAt)).
		WithImages().
		WithTrip(func(tq *ent.TripQuery) {
			tq.WithVehicle(func(vq *ent.VehicleQuery) {
				vq.WithImages()
			})
			tq.WithRoute()
			tq.WithStops()
			tq.WithDriver()
			tq.WithCompany()
		}).
		All(r.ctx)
	if err != nil {
		return nil, err
	}
	return application.Paginate(count, results)
}
