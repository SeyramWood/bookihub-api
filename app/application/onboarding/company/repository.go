package company

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
	"github.com/SeyramWood/bookibus/ent/company"
	"github.com/SeyramWood/bookibus/ent/companyuser"
	"github.com/SeyramWood/bookibus/ent/schema"
	"github.com/SeyramWood/bookibus/ent/user"
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

	return r.Read(c.ID)
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

// Onboard implements gateways.CompanyRepo.
func (r *repository) Onboard(id int, request *requeststructs.CompanyOnboardingRequest) (*ent.Company, error) {
	_, err := r.db.Company.UpdateOneID(id).
		SetBankAccount(&schema.BankAccount{
			AccountName:   request.BankAccount.AccountName,
			AccountNumber: request.BankAccount.AccountNumber,
			Bank:          request.BankAccount.Bank,
			Branch:        request.BankAccount.Branch,
		}).
		SetContactPerson(&schema.ContactPerson{
			Name:     request.ContactPerson.Name,
			Position: request.ContactPerson.Position,
			Phone:    request.ContactPerson.Phone,
			Email:    request.ContactPerson.Email,
		}).
		Save(r.ctx)
	if err != nil {
		return nil, err
	}
	return r.Read(id)
}

// BookiOnboard implements gateways.CompanyRepo.
func (r *repository) BookiOnboard(adminPassword string, request *requeststructs.BookiOnboardingRequest) (*ent.Company, error) {
	tx, err := r.db.Tx(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("error starting a transaction: %w", err)
	}
	c, err := tx.Company.Create().
		SetName(request.CompanyName).
		SetPhone(request.CompanyPhone).
		SetEmail(request.CompanyEmail).
		SetBankAccount(&schema.BankAccount{
			AccountName:   request.BankAccount.AccountName,
			AccountNumber: request.BankAccount.AccountNumber,
			Bank:          request.BankAccount.Bank,
			Branch:        request.BankAccount.Branch,
		}).
		SetContactPerson(&schema.ContactPerson{
			Name:     request.ContactPerson.Name,
			Position: request.ContactPerson.Position,
			Phone:    request.ContactPerson.Phone,
			Email:    request.ContactPerson.Email,
		}).
		Save(r.ctx)
	if err != nil {
		return nil, application.Rollback(tx, fmt.Errorf("failed onboarding company: %w", err))
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
	password, _ := bcrypt.GenerateFromPassword([]byte(strings.TrimSpace(adminPassword)), 16)
	_, err = tx.User.Create().
		SetCompanyUser(admin).
		SetUsername(request.AdminUsername).
		SetPassword(password).
		SetType(user.TypeCompany).
		Save(r.ctx)
	if err != nil {
		return nil, application.Rollback(tx, fmt.Errorf("failed creating admin user credentials: %w", err))
	}
	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed committing company onboarding transaction: %w", err)
	}
	return r.Read(c.ID)
}

// UpdateBankAccount implements gateways.CompanyRepo.
func (r *repository) UpdateBankAccount(id int, request *requeststructs.CompanyBankAccountRequest) (*ent.Company, error) {
	_, err := r.db.Company.UpdateOneID(id).
		SetBankAccount(&schema.BankAccount{
			AccountName:   request.AccountName,
			AccountNumber: request.AccountNumber,
			Bank:          request.Bank,
			Branch:        request.Branch,
		}).
		Save(r.ctx)
	if err != nil {
		return nil, err
	}
	return r.Read(id)
}

// UpdateCertificate implements gateways.CompanyRepo.
func (r *repository) UpdateCertificate(id int, cert string) (*ent.Company, error) {
	_, err := r.db.Company.UpdateOneID(id).
		SetCertificate(cert).
		Save(r.ctx)
	if err != nil {
		return nil, err
	}
	return r.Read(id)
}

// UpdateContactPerson implements gateways.CompanyRepo.
func (r *repository) UpdateContactPerson(id int, request *requeststructs.CompanyContactPersonUpdateRequest) (*ent.Company, error) {
	_, err := r.db.Company.UpdateOneID(id).
		SetContactPerson(&schema.ContactPerson{
			Name:     request.Name,
			Position: request.Position,
			Phone:    request.Phone,
			Email:    request.Email,
		}).
		Save(r.ctx)
	if err != nil {
		return nil, err
	}
	return r.Read(id)
}

// UpdateStatus implements gateways.CompanyRepo.
func (r *repository) UpdateStatus(id int, status string) (*ent.Company, error) {
	_, err := r.db.Company.UpdateOneID(id).
		SetStatus(company.Status(status)).
		Save(r.ctx)
	if err != nil {
		return nil, err
	}
	return r.Read(id)
}

// Update implements gateways.CompanyRepo.
func (r *repository) Update(id int, request *requeststructs.CompanyUpdateRequest) (*ent.Company, error) {
	_, err := r.db.Company.UpdateOneID(id).
		SetName(request.CompanyName).
		SetPhone(request.CompanyPhone).
		SetEmail(request.CompanyEmail).
		Save(r.ctx)
	if err != nil {
		return nil, err
	}
	return r.Read(id)
}
