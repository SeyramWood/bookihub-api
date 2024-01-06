package trip

import (
	"github.com/SeyramWood/bookibus/app/adapters/gateways"
	"github.com/SeyramWood/bookibus/app/adapters/presenters"
	requeststructs "github.com/SeyramWood/bookibus/app/domain/request_structs"
	"github.com/SeyramWood/bookibus/ent"
)

type service struct {
	repo gateways.TripRepo
}

func NewService(repo gateways.TripRepo) gateways.TripService {
	return &service{
		repo: repo,
	}
}

// Create implements gateways.TripService.
func (s *service) Create(companyId int, request *requeststructs.TripRequest) (*ent.Trip, error) {
	return s.repo.Insert(companyId, request)
}

// Fetch implements gateways.TripService.
func (s *service) Fetch(id int) (*ent.Trip, error) {
	return s.repo.Read(id)
}

// FetchAll implements gateways.TripService.
func (s *service) FetchAll(limit int, offset int, filter *requeststructs.TripFilterRequest) (*presenters.PaginationResponse, error) {
	return s.repo.ReadAll(limit, offset, filter)
}

// FetchAllByCompany implements gateways.TripService.
func (s *service) FetchAllByCompany(companyId int, limit int, offset int, filter *requeststructs.TripFilterRequest) (*presenters.PaginationResponse, error) {
	return s.repo.ReadAllByCompany(companyId, limit, offset, filter)
}

// FetchAllByDriver implements gateways.TripService.
func (s *service) FetchAllByDriver(driverId int, limit int, offset int, filter *requeststructs.TripFilterRequest) (*presenters.PaginationResponse, error) {
	return s.repo.ReadAllByDriver(driverId, limit, offset, filter)
}

// FetchAllCustomer implements gateways.TripService.
func (s *service) FetchAllCustomer(limit int, offset int, filter *requeststructs.CustomerTripFilterRequest) (*presenters.PaginationResponse, error) {
	return s.repo.ReadAllCustomer(limit, offset, filter)
}

// FetchAllPopular implements gateways.TripService.
func (s *service) FetchAllPopular(limit int, offset int, localDatetime string) (*presenters.PaginationResponse, error) {
	return s.repo.ReadAllPopular(limit, offset, localDatetime)
}

// FetchAllSearch implements gateways.TripService.
func (s *service) FetchAllSearch(searchKey string, limit int, offset int, filter *requeststructs.TripFilterRequest) (*presenters.PaginationResponse, error) {
	return s.repo.ReadAllSearch(searchKey, limit, offset, filter)
}

// FetchAllSearchByCompany implements gateways.TripService.
func (s *service) FetchAllSearchByCompany(searchKey string, companyId int, limit int, offset int, filter *requeststructs.TripFilterRequest) (*presenters.PaginationResponse, error) {
	return s.repo.ReadAllSearchByCompany(searchKey, companyId, limit, offset, filter)
}

// Remove implements gateways.TripService.
func (s *service) Remove(id int) error {
	return s.repo.Delete(id)
}

// Update implements gateways.TripService.
func (s *service) Update(id int, request *requeststructs.TripUpdateRequest) (*ent.Trip, error) {
	return s.repo.Update(id, request)
}

// UpdateInspection implements gateways.TripService.
func (s *service) UpdateInspection(id int, request *requeststructs.TripInspectionStatusRequest) (*ent.Trip, error) {
	return s.repo.UpdateInspection(id, request)
}

// UpdateSchedule implements gateways.TripService.
func (s *service) UpdateSchedule(id int, status bool) (*ent.Trip, error) {
	return s.repo.UpdateSchedule(id, status)
}

// UpdateStatus implements gateways.TripService.
func (s *service) UpdateStatus(id int, status string) (*ent.Trip, error) {
	return s.repo.UpdateStatus(id, status)
}
