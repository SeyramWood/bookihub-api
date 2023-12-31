package customer_user

import (
	"context"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/SeyramWood/bookibus/app/adapters/gateways"
	"github.com/SeyramWood/bookibus/app/adapters/presenters"
	"github.com/SeyramWood/bookibus/app/application"
	requeststructs "github.com/SeyramWood/bookibus/app/domain/request_structs"
	"github.com/SeyramWood/bookibus/app/framework/database"
	"github.com/SeyramWood/bookibus/ent"
	"github.com/SeyramWood/bookibus/ent/customer"
	"github.com/SeyramWood/bookibus/ent/user"
)

type repository struct {
	ctx context.Context
	db  *ent.Client
}

func NewRepository(db *database.Adapter) gateways.CustomerRepo {
	return &repository{
		ctx: context.Background(),
		db:  db.DB,
	}
}

// Delete implements gateways.CustomerRepo.
func (r *repository) Delete(id int) error {
	return r.db.Customer.DeleteOneID(id).Exec(r.ctx)
}

// Insert implements gateways.CustomerRepo.
func (r *repository) Insert(request *requeststructs.CustomerRequest) (*ent.Customer, error) {
	tx, err := r.db.Tx(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("error starting a transaction: %w", err)
	}
	result, err := tx.Customer.Create().
		SetLastName(request.LastName).
		SetOtherName(request.OtherName).
		SetPhone(request.Phone).
		Save(r.ctx)
	if err != nil {
		return nil, application.Rollback(tx, fmt.Errorf("failed creating customer: %w", err))
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(strings.TrimSpace(request.Password)), 16)
	_, err = tx.User.Create().
		SetCustomer(result).
		SetUsername(request.Username).
		SetPassword(hash).
		SetType(user.TypeCustomer).
		Save(r.ctx)
	if err != nil {
		return nil, application.Rollback(tx, fmt.Errorf("failed creating user credentials: %w", err))
	}
	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed committing customer creation transaction: %w", err)
	}
	return r.Read(result.ID)
}

// Read implements gateways.CustomerRepo.
func (r *repository) Read(id int) (*ent.Customer, error) {
	result, err := r.db.Customer.Query().Where(customer.ID(id)).WithProfile().Only(r.ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReadBySession implements gateways.CustomerRepo.
func (r *repository) ReadBySession(id int) (*ent.Customer, error) {
	result, err := r.db.Customer.Query().Where(customer.HasProfileWith(user.ID(id))).WithProfile().Only(r.ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReadAll implements gateways.CustomerRepo.
func (r *repository) ReadAll(limit int, offset int) (*presenters.PaginationResponse, error) {
	query := r.db.Customer.Query()
	count := query.CountX(r.ctx)
	results, err := query.
		Limit(limit).
		Offset(offset).
		Order(ent.Desc(customer.FieldCreatedAt)).
		WithProfile().
		All(r.ctx)
	if err != nil {
		return nil, err
	}
	return application.Paginate(count, results)
}

// Update implements gateways.CustomerRepo.
func (r *repository) Update(id int, request *requeststructs.CustomerUpdateRequest) (*ent.Customer, error) {
	result, err := r.db.Customer.UpdateOneID(id).
		SetLastName(request.LastName).
		SetOtherName(request.OtherName).
		SetPhone(request.Phone).
		Save(r.ctx)
	if err != nil {
		return nil, err
	}
	return r.Read(result.ID)
}
