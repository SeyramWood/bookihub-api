package company

import (
	"github.com/SeyramWood/bookibus/app/adapters/gateways"
	"github.com/SeyramWood/bookibus/app/adapters/presenters"
	requeststructs "github.com/SeyramWood/bookibus/app/domain/request_structs"
	"github.com/SeyramWood/bookibus/ent"
)

type service struct {
	repo gateways.CompanyRepo
}

func NewService(repo gateways.CompanyRepo) gateways.CompanyService {
	return &service{
		repo: repo,
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

// Update implements gateways.CompanyService.
func (s *service) Update(id int, request *requeststructs.CompanyUpdateRequest) (*ent.Company, error) {
	return s.repo.Update(id, request)
}
