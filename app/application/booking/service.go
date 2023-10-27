package booking

import (
	"time"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	requeststructs "github.com/SeyramWood/app/domain/request_structs"
	"github.com/SeyramWood/ent"
)

type service struct {
	repo     gateways.BookingRepo
	producer gateways.EventProducer
	cache    gateways.CacheService
}

func NewService(repo gateways.BookingRepo, producer gateways.EventProducer, cache gateways.CacheService) gateways.BookingService {
	return &service{
		repo:     repo,
		producer: producer,
		cache:    cache,
	}
}

// SaveToCache implements gateways.BookingService.
func (s *service) SaveToCache(reference string, request *requeststructs.BookingRequest) error {

	return s.cache.Set(reference, request, time.Second*60*30)
}

// CancelBooking implements gateways.BookingService.
func (s *service) CancelBooking(id int, request *requeststructs.BookingCancelRequest) (*ent.Booking, error) {
	// TODO process refund payment if necessary
	return s.repo.CancelBooking(id, request)
}

// Create implements gateways.BookingService.
func (s *service) Create(request *requeststructs.BookingRequest) (*ent.Booking, error) {
	return s.repo.Insert(request)
}

// Fetch implements gateways.BookingService.
func (s *service) Fetch(id int) (*ent.Booking, error) {
	return s.repo.Read(id)
}

// FetchAll implements gateways.BookingService.
func (s *service) FetchAll(limit int, offset int, filter *requeststructs.BookingFilterRequest) (*presenters.PaginationResponse, error) {
	return s.repo.ReadAll(limit, offset, filter)
}

// FetchAllByCompany implements gateways.BookingService.
func (s *service) FetchAllByCompany(companyId int, limit int, offset int, filter *requeststructs.BookingFilterRequest) (*presenters.PaginationResponse, error) {
	return s.repo.ReadAllByCompany(companyId, limit, offset, filter)
}

// FetchAllCustomer implements gateways.BookingService.
func (s *service) FetchAllCustomer(limit int, offset int, filter *requeststructs.BookingFilterRequest, customerId ...int) (*presenters.PaginationResponse, error) {
	return s.repo.ReadAllCustomer(limit, offset, filter, customerId...)
}

// Remove implements gateways.BookingService.
func (s *service) Remove(id int) error {
	return s.repo.Delete(id)
}

// Update implements gateways.BookingService.
func (s *service) Update(id int, request *requeststructs.BookingUpdateRequest) (*ent.Booking, error) {
	return s.repo.Update(id, request)
}
