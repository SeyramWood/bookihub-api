package vehicle

import (
	"context"
	"fmt"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application"
	requeststructs "github.com/SeyramWood/app/domain/request_structs"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/ent"
	"github.com/SeyramWood/ent/company"
	"github.com/SeyramWood/ent/vehicle"
	"github.com/SeyramWood/ent/vehicleimage"
)

type repository struct {
	db  *ent.Client
	ctx context.Context
}

func NewRepository(db *database.Adapter) gateways.VehicleRepo {
	return &repository{
		db:  db.DB,
		ctx: context.Background(),
	}
}

// Delete implements gateways.VehicleRepo.
func (r *repository) Delete(id int) error {
	return r.db.Vehicle.DeleteOneID(id).Exec(r.ctx)
}

// DeleteImage implements gateways.VehicleRepo.
func (r *repository) DeleteImage(id int) error {
	return r.db.VehicleImage.DeleteOneID(id).Exec(r.ctx)
}

// Insert implements gateways.VehicleRepo.
func (r *repository) Insert(companyId int, request *requeststructs.VehicleFormattedRequest) (*ent.Vehicle, error) {
	tx, err := r.db.Tx(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("error starting a transaction: %w", err)
	}
	result, err := tx.Vehicle.Create().
		SetCompanyID(companyId).
		SetRegistrationNumber(request.RegistrationNumber).
		SetModel(request.Model).
		SetSeat(request.Seat).
		Save(r.ctx)
	if err != nil {
		return nil, application.Rollback(tx, fmt.Errorf("failed creating vehicle: %w", err))
	}
	_, err = tx.VehicleImage.MapCreateBulk(request.Image, func(create *ent.VehicleImageCreate, i int) {
		create.SetImage(request.Image[i]).SetVehicle(result)
	}).Save(r.ctx)
	if err != nil {
		return nil, application.Rollback(tx, fmt.Errorf("failed creating vehicle image: %w", err))
	}
	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed committing vehicle creation transaction: %w", err)
	}
	return r.Read(result.ID)
}

// InsertImage implements gateways.VehicleRepo.
func (r *repository) InsertImage(id int, request []string) (*ent.Vehicle, error) {
	_, err := r.db.VehicleImage.MapCreateBulk(request, func(create *ent.VehicleImageCreate, i int) {
		create.SetImage(request[i]).SetVehicleID(id)
	}).Save(r.ctx)
	if err != nil {
		return nil, err
	}
	return r.Read(id)
}

// Read implements gateways.VehicleRepo.
func (r *repository) Read(id int) (*ent.Vehicle, error) {
	result, err := r.db.Vehicle.Query().Where(vehicle.ID(id)).
		WithImages().
		Only(r.ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReadImage implements gateways.VehicleRepo.
func (r *repository) ReadImage(id int) (*ent.VehicleImage, error) {
	result, err := r.db.VehicleImage.Query().Where(vehicleimage.ID(id)).
		Only(r.ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReadAll implements gateways.VehicleRepo.
func (r *repository) ReadAll(limit int, offset int) (*presenters.PaginationResponse, error) {
	query := r.db.Vehicle.Query()
	count := query.CountX(r.ctx)
	results, err := query.
		Limit(limit).
		Offset(offset).
		Order(ent.Desc(vehicle.FieldCreatedAt)).
		WithImages().
		All(r.ctx)
	if err != nil {
		return nil, err
	}
	return application.Paginate(count, results)
}

// ReadAllByCompany implements gateways.VehicleRepo.
func (r *repository) ReadAllByCompany(companyId int, limit int, offset int) (*presenters.PaginationResponse, error) {
	query := r.db.Vehicle.Query().Where(vehicle.HasCompanyWith(company.ID(companyId)))
	count := query.CountX(r.ctx)
	results, err := query.
		Limit(limit).
		Offset(offset).
		Order(ent.Desc(vehicle.FieldCreatedAt)).
		WithImages().
		All(r.ctx)
	if err != nil {
		return nil, err
	}
	return application.Paginate(count, results)
}

// Update implements gateways.VehicleRepo.
func (r *repository) Update(id int, request *requeststructs.VehicleUpdateRequest) (*ent.Vehicle, error) {
	_, err := r.db.Vehicle.UpdateOneID(id).
		SetRegistrationNumber(request.RegistrationNumber).
		SetModel(request.Model).
		SetSeat(request.Seat).
		Save(r.ctx)
	if err != nil {
		return nil, err
	}
	return r.Read(id)
}

// UpdateImage implements gateways.VehicleRepo.
func (r *repository) UpdateImage(id int, request string) (*ent.VehicleImage, error) {
	_, err := r.db.VehicleImage.UpdateOneID(id).
		SetImage(request).
		Save(r.ctx)
	if err != nil {
		return nil, err
	}
	return r.ReadImage(id)
}
