package auth

import (
	"context"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/SeyramWood/bookibus/app/adapters/gateways"
	requeststructs "github.com/SeyramWood/bookibus/app/domain/request_structs"
	"github.com/SeyramWood/bookibus/app/framework/database"
	"github.com/SeyramWood/bookibus/ent"
	"github.com/SeyramWood/bookibus/ent/user"
)

type repository struct {
	ctx context.Context
	db  *ent.Client
}

func NewRepository(db *database.Adapter) gateways.AuthRepo {
	return &repository{
		ctx: context.Background(),
		db:  db.DB,
	}
}

// ReadByID implements gateways.AuthRepo.
func (r *repository) ReadByID(id int) (*ent.User, error) {
	result, err := r.db.User.Query().
		Where(user.ID(id)).
		WithBookibusUser().
		WithCompanyUser(func(cuq *ent.CompanyUserQuery) {
			cuq.WithCompany()
		}).
		WithCustomer().
		Only(r.ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReadByUsername implements gateways.AuthRepo.
func (r *repository) ReadByUsername(username string) (*ent.User, error) {
	result, err := r.db.User.Query().
		Where(user.Username(username)).
		WithBookibusUser().
		WithCompanyUser(func(cuq *ent.CompanyUserQuery) {
			cuq.WithCompany()
		}).
		WithCustomer().
		Only(r.ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ResetPassword implements gateways.AuthRepo.
func (r *repository) ResetPassword(request *requeststructs.ResetPasswordRequest) (*ent.User, error) {
	password, _ := bcrypt.GenerateFromPassword([]byte(strings.TrimSpace(request.NewPassword)), 16)
	_, err := r.db.User.Update().Where(user.Username(request.Username)).
		SetPassword(password).
		Save(r.ctx)
	if err != nil {
		return nil, err
	}
	return r.ReadByUsername(request.Username)
}

// UpdateAvatar implements gateways.AuthRepo.
func (r *repository) UpdateAvatar(userID int, avatar string) error {
	_, err := r.db.User.UpdateOneID(userID).
		SetAvatar(avatar).
		Save(r.ctx)
	if err != nil {
		return err
	}
	return nil
}

// UpdatePassword implements gateways.AuthRepo.
func (r *repository) UpdatePassword(sessionID int, request *requeststructs.UpdatePasswordRequest) (*ent.User, error) {
	password, _ := bcrypt.GenerateFromPassword([]byte(strings.TrimSpace(request.NewPassword)), 16)
	result, err := r.db.User.UpdateOneID(sessionID).
		SetPassword(password).
		Save(r.ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
