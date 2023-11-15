package company_user

import (
	"github.com/SeyramWood/bookibus/app/adapters/gateways"
	"github.com/SeyramWood/bookibus/app/adapters/presenters"
	"github.com/SeyramWood/bookibus/app/application"
	"github.com/SeyramWood/bookibus/app/domain"
	requeststructs "github.com/SeyramWood/bookibus/app/domain/request_structs"
	"github.com/SeyramWood/bookibus/config"
	"github.com/SeyramWood/bookibus/ent"
)

type service struct {
	repo     gateways.CompanyUserRepo
	producer gateways.EventProducer
}

func NewService(repo gateways.CompanyUserRepo, producer gateways.EventProducer) gateways.CompanyUserService {
	return &service{
		repo:     repo,
		producer: producer,
	}
}

// Create implements gateways.CompanyUserService.
func (s *service) Create(request *requeststructs.CompanyUserRequest) (*ent.CompanyUser, error) {
	password := application.RandomString(12)
	result, err := s.repo.Insert(request, password)
	if err != nil {
		return nil, err
	}
	s.producer.Queue("notification:email", domain.MailerMessage{
		To:      request.Username,
		Subject: "NEW BookiBus ACCOUNT",
		Data: map[string]string{
			"username": request.Username,
			"password": password,
			"url":      config.App().AppCompanyURL,
		},
		Template: "newuser",
	})
	return result, nil
}

// Fetch implements gateways.CompanyUserService.
func (s *service) Fetch(id int) (*ent.CompanyUser, error) {
	return s.repo.Read(id)
}

// FetchAll implements gateways.CompanyUserService.
func (s *service) FetchAll(limit int, offset int) (*presenters.PaginationResponse, error) {
	return s.repo.ReadAll(limit, offset)
}

// FetchAllByCompany implements gateways.CompanyUserService.
func (s *service) FetchAllByCompany(companyId int, limit int, offset int) (*presenters.PaginationResponse, error) {
	return s.repo.FetchAllByCompany(companyId, limit, offset)
}

// Remove implements gateways.CompanyUserService.
func (s *service) Remove(id int) error {
	return s.repo.Delete(id)
}

// Update implements gateways.CompanyUserService.
func (s *service) Update(id int, request *requeststructs.CompanyUserUpdateRequest) (*ent.CompanyUser, error) {
	return s.repo.Update(id, request)
}
