package route

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	"github.com/SeyramWood/bookibus/app/adapters/gateways"
	"github.com/SeyramWood/bookibus/app/adapters/presenters"
	"github.com/SeyramWood/bookibus/app/application"
	requeststructs "github.com/SeyramWood/bookibus/app/domain/request_structs"
	"github.com/SeyramWood/bookibus/app/framework/database"
	"github.com/SeyramWood/bookibus/ent"
	"github.com/SeyramWood/bookibus/ent/company"
	"github.com/SeyramWood/bookibus/ent/route"
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

// Insert implements gateways.RouteRepo.
func (r *repository) Insert(companyId int, request *requeststructs.RouteRequest) (*ent.Route, error) {
	result, err := r.db.Route.Create().
		SetCompanyID(companyId).
		SetFromLocation(request.From).
		SetToLocation(request.To).
		Save(r.ctx)
	if err != nil {
		return nil, err
	}
	return r.Read(result.ID)
}

// Read implements gateways.RouteRepo.
func (r *repository) Read(id int) (*ent.Route, error) {
	result, err := r.db.Route.Query().Where(route.ID(id)).Only(r.ctx)
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
		All(r.ctx)
	if err != nil {
		return nil, err
	}
	return application.Paginate(count, results)
}

// Update implements gateways.RouteRepo.
func (r *repository) Update(id int, request *requeststructs.RouteRequest) (*ent.Route, error) {
	_, err := r.db.Route.UpdateOneID(id).
		SetFromLocation(request.From).
		SetToLocation(request.To).
		Save(r.ctx)
	if err != nil {
		return nil, err
	}
	return r.Read(id)
}
