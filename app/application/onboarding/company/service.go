package company

import (
	"fmt"
	"strings"

	"github.com/SeyramWood/bookibus/app/adapters/gateways"
	"github.com/SeyramWood/bookibus/app/adapters/presenters"
	"github.com/SeyramWood/bookibus/app/application"
	"github.com/SeyramWood/bookibus/app/domain"
	requeststructs "github.com/SeyramWood/bookibus/app/domain/request_structs"
	"github.com/SeyramWood/bookibus/config"
	"github.com/SeyramWood/bookibus/ent"
)

type service struct {
	repo        gateways.CompanyRepo
	producer    gateways.EventProducer
	storage     gateways.StorageService
	storagePath string
}

func NewService(repo gateways.CompanyRepo, producer gateways.EventProducer, storage gateways.StorageService) gateways.CompanyService {
	return &service{
		repo:        repo,
		producer:    producer,
		storage:     storage,
		storagePath: "public/company/certificates",
	}
}

// Create implements gateways.CompanyService.
func (s *service) Create(request *requeststructs.CompanyRequest) (*ent.Company, error) {
	return s.repo.Insert(request)
}

// Fetch implements gateways.CompanyService.
func (s *service) Fetch(id int) (*ent.Company, error) {
	return s.repo.Read(id)
}

// FetchAll implements gateways.CompanyService.
func (s *service) FetchAll(limit int, offset int) (*presenters.PaginationResponse, error) {
	return s.repo.ReadAll(limit, offset)
}

// Remove implements gateways.CompanyService.
func (s *service) Remove(id int) error {
	return s.repo.Delete(id)
}

// BookiOnboard implements gateways.CompanyService.
func (s *service) BookiOnboard(request *requeststructs.BookiOnboardingRequest) (*ent.Company, error) {
	password := application.RandomString(12)
	result, err := s.repo.BookiOnboard(password, request)
	if err != nil {
		return nil, err
	}
	s.producer.Queue("notification:email", domain.MailerMessage{
		To:      request.AdminUsername,
		Subject: "COMPANY ADMIN ACCOUNT - Booki Rides",
		Data: map[string]string{
			"company":  request.CompanyName,
			"username": request.AdminUsername,
			"password": password,
			"url":      config.App().AppCompanyURL,
		},
		Template: "newcompanyadmin",
	})
	return result, nil
}

// Onboard implements gateways.CompanyService.
func (s *service) Onboard(id int, request *requeststructs.CompanyOnboardingRequest) (*ent.Company, error) {
	result, err := s.repo.Onboard(id, request)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateBankAccount implements gateways.CompanyService.
func (s *service) UpdateBankAccount(id int, request *requeststructs.CompanyBankAccountRequest) (*ent.Company, error) {
	return s.repo.UpdateBankAccount(id, request)
}

// UpdateCertificate implements gateways.CompanyService.
func (s *service) UpdateCertificate(id int, request *requeststructs.CompanyCertificateUpdateRequest) (*ent.Company, error) {
	com, err := s.repo.Read(id)
	if err != nil {
		return nil, err
	}
	certificate, err := s.storage.UploadFile(s.storagePath, request.BusinessCertificate)
	if err != nil {
		return nil, err
	}
	result, err := s.repo.UpdateCertificate(id, certificate)
	if err != nil {
		go s.storage.ExecuteTask(strings.Replace(certificate, config.App().AppURL, "public", 1), "delete_file")
		return nil, err
	}
	go s.storage.ExecuteTask(strings.Replace(com.Certificate, config.App().AppURL, "public", 1), "delete_file")
	return result, nil
}

// UpdateContactPerson implements gateways.CompanyService.
func (s *service) UpdateContactPerson(id int, request *requeststructs.CompanyContactPersonUpdateRequest) (*ent.Company, error) {
	return s.repo.UpdateContactPerson(id, request)
}

// UpdateStatus implements gateways.CompanyService.
func (s *service) UpdateStatus(id int, status string) (*ent.Company, error) {
	result, err := s.repo.Read(id)
	if err != nil {
		return nil, err
	}
	if result.Certificate == "" || result.BankAccount == nil {
		return nil, fmt.Errorf("invalid information: %s must update business certificate and bank account details", result.Name)
	}
	return s.repo.UpdateStatus(id, status)
}

// Update implements gateways.CompanyService.
func (s *service) Update(id int, request *requeststructs.CompanyUpdateRequest) (*ent.Company, error) {
	return s.repo.Update(id, request)
}
