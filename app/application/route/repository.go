package route

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application"
	requeststructs "github.com/SeyramWood/app/domain/request_structs"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/ent"
	"github.com/SeyramWood/ent/company"
	"github.com/SeyramWood/ent/route"
)

type repository struct {
	db  *ent.Client
	ctx context.Context
}

func NewRepository(db *database.Adapter) gateways.RouteRepo {
	return &repository{
		db:  db.DB,
		ctx: context.Background(),
	}
}

// Delete implements gateways.RouteRepo.
func (r *repository) Delete(id int) error {
	return r.db.Route.DeleteOneID(id).Exec(r.ctx)
}

// DeleteStop implements gateways.RouteRepo.
func (r *repository) DeleteStop(id int) error {
	return r.db.RouteStop.DeleteOneID(id).Exec(r.ctx)
}

// Insert implements gateways.RouteRepo.
func (r *repository) Insert(companyId int, request *requeststructs.RouteRequest) (*ent.Route, error) {
	tx, err := r.db.Tx(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("error starting a transaction: %w", err)
	}
	result, err := tx.Route.Create().
		SetCompanyID(companyId).
		SetFromLocation(request.From).
		SetToLocation(request.To).
		SetFromLatitude(request.FromLatitude).
		SetFromLongitude(request.FromLongitude).
		SetToLatitude(request.ToLatitude).
		SetToLongitude(request.ToLongitude).
		SetRate(request.Rate).
		SetDiscount(request.Discount).
		Save(r.ctx)
	if err != nil {
		return nil, application.Rollback(tx, fmt.Errorf("failed creating route: %w", err))
	}
	_, err = tx.RouteStop.MapCreateBulk(request.Stops, func(create *ent.RouteStopCreate, i int) {
		create.SetLatitude(request.Stops[i].Latitude).SetLongitude(request.Stops[i].Longitude).SetRoute(result)
	}).Save(r.ctx)
	if err != nil {
		return nil, application.Rollback(tx, fmt.Errorf("failed creating route stop: %w", err))
	}
	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed committing route creation transaction: %w", err)
	}
	return r.Read(result.ID)
}

// InsertRoute implements gateways.RouteRepo.
func (r *repository) InsertRouteStop(id int, request *requeststructs.RouteStopRequest) (*ent.RouteStop, error) {
	result, err := r.db.RouteStop.Create().SetLatitude(request.Latitude).SetLongitude(request.Longitude).SetRouteID(id).Save(r.ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Read implements gateways.RouteRepo.
func (r *repository) Read(id int) (*ent.Route, error) {
	result, err := r.db.Route.Query().Where(route.ID(id)).WithStops().Only(r.ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReadAll implements gateways.RouteRepo.
func (r *repository) ReadAll(limit int, offset int) (*presenters.PaginationResponse, error) {
	query := r.db.Route.Query()
	results, err := query.
		Limit(limit).
		Offset(offset).
		Order(ent.Desc(route.FieldCreatedAt)).
		WithStops().
		All(r.ctx)
	if err != nil {
		return nil, err
	}

	return application.Paginate(query.CountX(r.ctx), results)
}

// FetchAllDistinct implements gateways.RouteRepo.
func (r *repository) ReadAllDistinct(limit int, offset int) (*presenters.PaginationResponse, error) {
	query := r.db.Route.Query().Modify(func(s *sql.Selector) {
		s.Select(fmt.Sprintf("DISTINCT %s, %s", route.FieldFromLocation, route.FieldToLocation))
	})
	results, err := query.
		Limit(limit).
		Offset(offset).
		Order(ent.Desc(route.FieldFromLocation, route.FieldToLocation)).
		All(r.ctx)
	if err != nil {
		return nil, err
	}
	return application.Paginate(query.CountX(r.ctx), results)
}

// ReadAllByCompany implements gateways.RouteRepo.
func (r *repository) ReadAllByCompany(companyId int, limit int, offset int) (*presenters.PaginationResponse, error) {
	query := r.db.Route.Query().Where(route.HasCompanyWith(company.ID(companyId)))
	count := query.CountX(r.ctx)
	results, err := query.
		Limit(limit).
		Offset(offset).
		Order(ent.Desc(route.FieldCreatedAt)).
		WithStops().
		All(r.ctx)
	if err != nil {
		return nil, err
	}
	return application.Paginate(count, results)
}

// Update implements gateways.RouteRepo.
func (r *repository) Update(id int, request *requeststructs.RouteUpdateRequest) (*ent.Route, error) {
	_, err := r.db.Route.UpdateOneID(id).
		SetFromLocation(request.From).
		SetToLocation(request.To).
		SetFromLatitude(request.FromLatitude).
		SetFromLongitude(request.FromLongitude).
		SetToLatitude(request.ToLatitude).
		SetToLongitude(request.ToLongitude).
		SetRate(request.Rate).
		SetDiscount(request.Discount).
		Save(r.ctx)
	if err != nil {
		return nil, err
	}
	return r.Read(id)
}

// UpdateStop implements gateways.RouteRepo.
func (r *repository) UpdateStop(id int, request *requeststructs.RouteStopRequest) (*ent.RouteStop, error) {
	result, err := r.db.RouteStop.UpdateOneID(id).SetLatitude(request.Latitude).SetLongitude(request.Longitude).Save(r.ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
