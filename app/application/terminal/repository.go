package terminal

import (
	"context"

	"github.com/SeyramWood/bookibus/app/adapters/gateways"
	"github.com/SeyramWood/bookibus/app/adapters/presenters"
	"github.com/SeyramWood/bookibus/app/application"
	requeststructs "github.com/SeyramWood/bookibus/app/domain/request_structs"
	"github.com/SeyramWood/bookibus/app/framework/database"
	"github.com/SeyramWood/bookibus/ent"
	"github.com/SeyramWood/bookibus/ent/company"
	"github.com/SeyramWood/bookibus/ent/terminal"
)

type repository struct {
	db  *ent.Client
	ctx context.Context
}

func NewRepository(db *database.Adapter) gateways.TerminalRepo {
	return &repository{
		db:  db.DB,
		ctx: context.Background(),
	}
}

// Delete implements gateways.TerminalRepo.
func (r *repository) Delete(id int) error {
	return r.db.Terminal.DeleteOneID(id).Exec(r.ctx)
}

// Insert implements gateways.TerminalRepo.
func (r *repository) Insert(companyId int, request *requeststructs.TerminalRequest) (*ent.Terminal, error) {
	result, err := r.db.Terminal.Create().
		SetName(request.Name).
		SetCompanyID(companyId).
		Save(r.ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Read implements gateways.TerminalRepo.
func (r *repository) Read(id int) (*ent.Terminal, error) {
	return r.db.Terminal.Get(r.ctx, id)
}

// ReadAll implements gateways.TerminalRepo.
func (r *repository) ReadAll(limit int, offset int) (*presenters.PaginationResponse, error) {
	query := r.db.Terminal.Query()
	count := query.CountX(r.ctx)
	results, err := query.
		Limit(limit).
		Offset(offset).
		Order(ent.Desc(terminal.FieldCreatedAt)).
		All(r.ctx)
	if err != nil {
		return nil, err
	}
	return application.Paginate(count, results)
}

// ReadAllByCompany implements gateways.TerminalRepo.
func (r *repository) ReadAllByCompany(companyId int, limit int, offset int) (*presenters.PaginationResponse, error) {
	query := r.db.Terminal.Query().Where(terminal.HasCompanyWith(company.ID(companyId)))
	count := query.CountX(r.ctx)
	results, err := query.
		Limit(limit).
		Offset(offset).
		Order(ent.Desc(terminal.FieldCreatedAt)).
		All(r.ctx)
	if err != nil {
		return nil, err
	}
	return application.Paginate(count, results)
}

// Update implements gateways.TerminalRepo.
func (r *repository) Update(id int, request *requeststructs.TerminalRequest) (*ent.Terminal, error) {
	result, err := r.db.Terminal.UpdateOneID(id).
		SetName(request.Name).
		Save(r.ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
