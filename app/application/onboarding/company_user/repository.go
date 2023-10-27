package company_user

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

func NewRepository(db *database.Adapter) gateways.CompanyUserRepo {
	return &repository{
		ctx: context.Background(),
		db:  db.DB,
	}
}

// Delete implements gateways.CompanyUserRepo.
func (r *repository) Delete(id int) error {
	return r.db.CompanyUser.DeleteOneID(id).Exec(r.ctx)
}

// Insert implements gateways.CompanyUserRepo.
func (r *repository) Insert(request *requeststructs.CompanyUserRequest, password string) (*ent.CompanyUser, error) {
	tx, err := r.db.Tx(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("error starting a transaction: %w", err)
	}
	result, err := tx.CompanyUser.Create().
		SetCompanyID(request.CompanyID).
		SetLastName(request.LastName).
		SetOtherName(request.OtherName).
		SetPhone(request.Phone).
		SetOtherPhone(request.OtherPhone).
		SetRole(companyuser.Role(request.Role)).
		Save(r.ctx)
	if err != nil {
		return nil, application.Rollback(tx, fmt.Errorf("failed creating company user user: %w", err))
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(strings.TrimSpace(password)), 16)
	_, err = tx.User.Create().
		SetCompanyUser(result).
		SetUsername(request.Username).
		SetPassword(hash).
		SetType(user.TypeCompany).
		Save(r.ctx)
	if err != nil {
		return nil, application.Rollback(tx, fmt.Errorf("failed creating company user credentials: %w", err))
	}
	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed committing company user creation transaction: %w", err)
	}
	return result, nil
}

// Read implements gateways.CompanyUserRepo.
func (r *repository) Read(id int) (*ent.CompanyUser, error) {
	result, err := r.db.CompanyUser.Query().Where(companyuser.ID(id)).Only(r.ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReadAll implements gateways.CompanyUserRepo.
func (r *repository) ReadAll(limit int, offset int) (*presenters.PaginationResponse, error) {
	query := r.db.CompanyUser.Query()
	count := query.CountX(r.ctx)
	results, err := query.
		Limit(limit).
		Offset(offset).
		Order(ent.Desc(companyuser.FieldCreatedAt)).
		All(r.ctx)
	if err != nil {
		return nil, err
	}
	return application.Paginate(count, results)
}

// FetchAllByCompany implements gateways.CompanyUserRepo.
func (r *repository) FetchAllByCompany(companyId int, limit int, offset int) (*presenters.PaginationResponse, error) {
	query := r.db.CompanyUser.Query().Where(companyuser.HasCompanyWith(company.ID(companyId)))
	count := query.CountX(r.ctx)
	results, err := query.
		Limit(limit).
		Offset(offset).
		Order(ent.Desc(companyuser.FieldCreatedAt)).
		All(r.ctx)
	if err != nil {
		return nil, err
	}
	return application.Paginate(count, results)
}

// Update implements gateways.CompanyUserRepo.
func (r *repository) Update(id int, request *requeststructs.CompanyUserUpdateRequest) (*ent.CompanyUser, error) {
	result, err := r.db.CompanyUser.UpdateOneID(id).
		SetLastName(request.LastName).
		SetOtherName(request.OtherName).
		SetPhone(request.Phone).
		SetOtherPhone(request.OtherPhone).
		Save(r.ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
