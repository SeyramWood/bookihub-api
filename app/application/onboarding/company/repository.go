package company

import (
	"context"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application"
	requeststructs "github.com/SeyramWood/app/domain/request_structs"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/ent"
	"github.com/SeyramWood/ent/company"
	"github.com/SeyramWood/ent/companyuser"
	"github.com/SeyramWood/ent/user"
)

type repository struct {
	ctx context.Context
	db  *ent.Client
}

func NewRepository(db *database.Adapter) gateways.CompanyRepo {
	return &repository{
		ctx: context.Background(),
		db:  db.DB,
	}
}

// Delete implements gateways.CompanyRepo.
func (r *repository) Delete(id int) error {
	return r.db.Company.DeleteOneID(id).Exec(r.ctx)
}

// Insert implements gateways.CompanyRepo.
func (r *repository) Insert(request *requeststructs.CompanyRequest) (*ent.Company, error) {
	tx, err := r.db.Tx(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("error starting a transaction: %w", err)
	}
	c, err := tx.Company.Create().
		SetName(request.CompanyName).
		SetPhone(request.CompanyPhone).
		SetEmail(request.CompanyEmail).
		Save(r.ctx)
	if err != nil {
		return nil, application.Rollback(tx, fmt.Errorf("failed creating company: %w", err))
	}
	admin, err := tx.CompanyUser.Create().
		SetCompany(c).
		SetLastName("Company").
		SetOtherName("Administrator").
		SetRole(companyuser.DefaultRole).
		Save(r.ctx)
	if err != nil {
		return nil, application.Rollback(tx, fmt.Errorf("failed creating company user: %w", err))
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(strings.TrimSpace(request.Password)), 16)
	_, err = tx.User.Create().
		SetCompanyUser(admin).
		SetUsername(request.Username).
		SetPassword(password).
		SetType(user.TypeCompany).
		Save(r.ctx)
	if err != nil {
		return nil, application.Rollback(tx, fmt.Errorf("failed creating user credentials: %w", err))
	}
	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed committing company creation transaction: %w", err)
	}

	return c, nil
}

// Read implements gateways.CompanyRepo.
func (r *repository) Read(id int) (*ent.Company, error) {
	result, err := r.db.Company.Query().Where(company.ID(id)).Only(r.ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReadAll implements gateways.CompanyRepo.
func (r *repository) ReadAll(limit int, offset int) (*presenters.PaginationResponse, error) {
	query := r.db.Company.Query()
	count := query.CountX(r.ctx)
	results, err := query.
		Limit(limit).
		Offset(offset).
		Order(ent.Desc(company.FieldCreatedAt)).
		All(r.ctx)
	if err != nil {
		return nil, err
	}
	return application.Paginate(count, results)
}

// Update implements gateways.CompanyRepo.
func (r *repository) Update(id int, request *requeststructs.CompanyUpdateRequest) (*ent.Company, error) {
	result, err := r.db.Company.UpdateOneID(id).
		SetName(request.CompanyName).
		SetPhone(request.CompanyPhone).
		SetOtherPhone(request.CompanyOtherPhone).
		SetEmail(request.CompanyEmail).
		Save(r.ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
