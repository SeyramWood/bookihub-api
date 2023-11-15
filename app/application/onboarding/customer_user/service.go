package customer_user

import (
	"github.com/SeyramWood/bookibus/app/adapters/gateways"
	"github.com/SeyramWood/bookibus/app/adapters/presenters"
	requeststructs "github.com/SeyramWood/bookibus/app/domain/request_structs"
	"github.com/SeyramWood/bookibus/ent"
)

type service struct {
	repo gateways.CustomerRepo
}

func NewService(repo gateways.CustomerRepo) gateways.CustomerService {
	return &service{
		repo: repo,
	}
}

// Create implements gateways.CustomerService.
func (s *service) Create(request *requeststructs.CustomerRequest) (*ent.Customer, error) {
	return s.repo.Insert(request)
}

// Fetch implements gateways.CustomerService.
func (s *service) Fetch(id int) (*ent.Customer, error) {
	return s.repo.Read(id)
}

// FetchAll implements gateways.CustomerService.
func (s *service) FetchAll(limit int, offset int) (*presenters.PaginationResponse, error) {
	return s.repo.ReadAll(limit, offset)
}

// Remove implements gateways.CustomerService.
func (s *service) Remove(id int) error {
	return s.repo.Delete(id)
}

// Update implements gateways.CustomerService.
func (s *service) Update(id int, request *requeststructs.CustomerUpdateRequest) (*ent.Customer, error) {
	return s.repo.Update(id, request)
}
