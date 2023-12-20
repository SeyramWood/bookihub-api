package route

import (
	"github.com/SeyramWood/bookibus/app/adapters/gateways"
	"github.com/SeyramWood/bookibus/app/adapters/presenters"
	requeststructs "github.com/SeyramWood/bookibus/app/domain/request_structs"
	"github.com/SeyramWood/bookibus/ent"
)

type service struct {
	repo gateways.RouteRepo
}

func NewService(repo gateways.RouteRepo) gateways.RouteService {
	return &service{
		repo: repo,
	}
}

// Create implements gateways.RouteService.
func (s *service) Create(companyId int, request *requeststructs.RouteRequest) (*ent.Route, error) {
	return s.repo.Insert(companyId, request)
}

// Fetch implements gateways.RouteService.
func (s *service) Fetch(id int) (*ent.Route, error) {
	return s.repo.Read(id)
}

// FetchAll implements gateways.RouteService.
func (s *service) FetchAll(limit int, offset int) (*presenters.PaginationResponse, error) {
	return s.repo.ReadAll(limit, offset)
}

// FetchAllDistinct implements gateways.RouteService.
func (s *service) FetchAllDistinct(limit int, offset int) (*presenters.PaginationResponse, error) {
	return s.repo.ReadAllDistinct(limit, offset)
}

// FetchAllByCompany implements gateways.RouteService.
func (s *service) FetchAllByCompany(companyId int, limit int, offset int) (*presenters.PaginationResponse, error) {
	return s.repo.ReadAllByCompany(companyId, limit, offset)
}

// Remove implements gateways.RouteService.
func (s *service) Remove(id int) error {
	return s.repo.Delete(id)
}

// Update implements gateways.RouteService.
func (s *service) Update(id int, request *requeststructs.RouteRequest) (*ent.Route, error) {
	return s.repo.Update(id, request)
}
