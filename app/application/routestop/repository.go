package routestop

import (
	"context"

	"github.com/SeyramWood/bookibus/app/adapters/gateways"
	"github.com/SeyramWood/bookibus/app/adapters/presenters"
	"github.com/SeyramWood/bookibus/app/application"
	requeststructs "github.com/SeyramWood/bookibus/app/domain/request_structs"
	"github.com/SeyramWood/bookibus/app/framework/database"
	"github.com/SeyramWood/bookibus/ent"
	"github.com/SeyramWood/bookibus/ent/company"
	"github.com/SeyramWood/bookibus/ent/routestop"
)

type repository struct {
	db  *ent.Client
	ctx context.Context
}

func NewRepository(db *database.Adapter) gateways.RouteStopRepo {
	return &repository{
		db:  db.DB,
		ctx: context.Background(),
	}
}

// Delete implements gateways.RouteStopRepo.
func (r *repository) Delete(id int) error {
	return r.db.RouteStop.DeleteOneID(id).Exec(r.ctx)
}

// Insert implements gateways.RouteStopRepo.
func (r *repository) Insert(companyId int, request *requeststructs.RouteStopRequest) (*ent.RouteStop, error) {
	result, err := r.db.RouteStop.Create().
		SetAddress(request.Address).
		SetLatitude(request.Latitude).
		SetLongitude(request.Longitude).
		SetCompanyID(companyId).
		Save(r.ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Read implements gateways.RouteStopRepo.
func (r *repository) Read(id int) (*ent.RouteStop, error) {
	result, err := r.db.RouteStop.Query().Where(routestop.ID(id)).Only(r.ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReadAll implements gateways.RouteStopRepo.
func (r *repository) ReadAll(limit int, offset int) (*presenters.PaginationResponse, error) {
	query := r.db.RouteStop.Query()
	results, err := query.
		Limit(limit).
		Offset(offset).
		Order(ent.Desc(routestop.FieldCreatedAt)).
		All(r.ctx)
	if err != nil {
		return nil, err
	}
	return application.Paginate(query.CountX(r.ctx), results)
}

// ReadAllByCompany implements gateways.RouteStopRepo.
func (r *repository) ReadAllByCompany(companyId int, limit int, offset int) (*presenters.PaginationResponse, error) {
	query := r.db.RouteStop.Query().Where(routestop.HasCompanyWith(company.ID(companyId)))
	count := query.CountX(r.ctx)
	results, err := query.
		Limit(limit).
		Offset(offset).
		Order(ent.Desc(routestop.FieldCreatedAt)).
		All(r.ctx)
	if err != nil {
		return nil, err
	}
	return application.Paginate(count, results)
}

// Update implements gateways.RouteStopRepo.
func (r *repository) Update(id int, request *requeststructs.RouteStopRequest) (*ent.RouteStop, error) {
	result, err := r.db.RouteStop.UpdateOneID(id).
		SetAddress(request.Address).
		SetLatitude(request.Latitude).
		SetLongitude(request.Longitude).
		Save(r.ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
