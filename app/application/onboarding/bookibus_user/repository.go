package bookibus_user

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
	"github.com/SeyramWood/ent/bookibususer"
	"github.com/SeyramWood/ent/user"
)

type repository struct {
	ctx context.Context
	db  *ent.Client
}

func NewRepository(db *database.Adapter) gateways.BookibusUserRepo {
	return &repository{
		ctx: context.Background(),
		db:  db.DB,
	}
}

// Delete implements gateways.BookibusUserRepo.
func (r *repository) Delete(id int) error {
	return r.db.BookibusUser.DeleteOneID(id).Exec(r.ctx)
}

// Insert implements gateways.BookibusUserRepo.
func (r *repository) Insert(request *requeststructs.BookibusUserRequest, password string) (*ent.BookibusUser, error) {
	tx, err := r.db.Tx(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("error starting a transaction: %w", err)
	}
	result, err := tx.BookibusUser.Create().
		SetLastName(request.LastName).
		SetOtherName(request.OtherName).
		SetPhone(request.Phone).
		SetOtherPhone(request.OtherPhone).
		SetRole(bookibususer.Role(request.Role)).
		Save(r.ctx)
	if err != nil {
		return nil, application.Rollback(tx, fmt.Errorf("failed creating bookibus user: %w", err))
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(strings.TrimSpace(password)), 16)
	_, err = tx.User.Create().
		SetBookibusUser(result).
		SetUsername(request.Username).
		SetPassword(hash).
		SetType(user.TypeBookibus).
		Save(r.ctx)
	if err != nil {
		return nil, application.Rollback(tx, fmt.Errorf("failed creating user credentials: %w", err))
	}
	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed committing bookibus creation transaction: %w", err)
	}
	return result, nil
}

// Read implements gateways.BookibusUserRepo.
func (r *repository) Read(id int) (*ent.BookibusUser, error) {
	result, err := r.db.BookibusUser.Query().Where(bookibususer.ID(id)).Only(r.ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReadAll implements gateways.BookibusUserRepo.
func (r *repository) ReadAll(limit int, offset int) (*presenters.PaginationResponse, error) {
	query := r.db.BookibusUser.Query()
	count := query.CountX(r.ctx)
	results, err := query.
		Limit(limit).
		Offset(offset).
		Order(ent.Desc(bookibususer.FieldCreatedAt)).
		All(r.ctx)
	if err != nil {
		return nil, err
	}
	return application.Paginate(count, results)
}

// Update implements gateways.BookibusUserRepo.
func (r *repository) Update(id int, request *requeststructs.BookibusUserUpdateRequest) (*ent.BookibusUser, error) {
	result, err := r.db.BookibusUser.UpdateOneID(id).
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
