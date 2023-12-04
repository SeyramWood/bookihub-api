package configuration

import (
	"github.com/SeyramWood/bookibus/app/adapters/gateways"
	requeststructs "github.com/SeyramWood/bookibus/app/domain/request_structs"
	"github.com/SeyramWood/bookibus/ent"
)

type service struct {
	repo gateways.ConfigurationRepo
}

func NewService(repo gateways.ConfigurationRepo) gateways.ConfigurationService {
	return &service{
		repo: repo,
	}
}

// CreateCharge implements gateways.ConfigurationService.
func (s *service) CreateCharge(request *requeststructs.TransactionChargeRequest) (*ent.Configuration, error) {
	return s.repo.InsertCharge(request)
}

// Fetch implements gateways.ConfigurationService.
func (s *service) Fetch() (*ent.Configuration, error) {
	return s.repo.Read()
}

// FetchCharge implements gateways.ConfigurationService.
func (s *service) FetchCharge() (*ent.Configuration, error) {
	return s.repo.ReadCharge()
}

// UpdateCharge implements gateways.ConfigurationService.
func (s *service) UpdateCharge(id int, request *requeststructs.TransactionChargeRequest) (*ent.Configuration, error) {
	return s.repo.UpdateCharge(id, request)
}
