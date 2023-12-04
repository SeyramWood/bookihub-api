package booking

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/SeyramWood/bookibus/app/adapters/gateways"
	"github.com/SeyramWood/bookibus/app/adapters/presenters"
	"github.com/SeyramWood/bookibus/app/application/payment"
	"github.com/SeyramWood/bookibus/app/domain"
	requeststructs "github.com/SeyramWood/bookibus/app/domain/request_structs"
	"github.com/SeyramWood/bookibus/config"
	"github.com/SeyramWood/bookibus/ent"
)

type service struct {
	repo     gateways.BookingRepo
	producer gateways.EventProducer
	cache    gateways.CacheService
	payment  gateways.PaymentService
}

func NewService(repo gateways.BookingRepo, producer gateways.EventProducer, cache gateways.CacheService, payment gateways.PaymentService) gateways.BookingService {
	return &service{
		repo:     repo,
		producer: producer,
		cache:    cache,
		payment:  payment,
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
func (s *service) Create(request *requeststructs.BookingRequest, transType string) (*ent.Booking, error) {
	if transType == "cash" {
		result, err := s.repo.Insert(request, &requeststructs.PaymentReferenceResponse{
			PaidAt:    time.Now().Format(time.RFC3339),
			TransType: strings.ToLower(transType),
		})
		if err != nil {
			return nil, err
		}
		if request.SMSNotification {
			s.producer.Queue("notification:sms", domain.SMSPayload{
				Message:    fmt.Sprintf("Your booking was successful, use the link to view and download ticket: \n%s/trips/ticket/%s/download", config.App().AppURL, result.BookingNumber),
				Recipients: []string{result.Edges.Contact.Phone},
			})
		} else {
			s.producer.Queue("notification:email", domain.MailerMessage{
				To:      result.Edges.Contact.Email,
				Subject: "Trip Ticket - BookiRide",
				Data: map[string]any{
					"data":   presenters.BookingTicketResponse(result),
					"url":    config.App().AppWebsiteURL,
					"appUrl": config.App().AppURL,
					"bookiContact": map[string]string{
						"email": config.Contact().Email,
						"phone": config.Contact().Phone,
					},
				},
				Template: "newbooking",
			})
		}

		return result, nil
	}

	resp, err := s.payment.Verify(request.Reference)
	if err != nil {
		return nil, err
	}
	if resp.Status && resp.Message == payment.VerificationSuccessful {
		result, err := s.repo.Insert(request, resp)
		if err != nil {
			return nil, err
		}
		// TODO process new booking notification: sms or email and db
		if result.SmsNotification {
			s.producer.Queue("notification:sms", domain.SMSPayload{
				Message:    fmt.Sprintf("Your booking was successful, use the link to view and download ticket: \n%s/trips/ticket/%s/download", config.App().AppURL, result.BookingNumber),
				Recipients: []string{result.Edges.Contact.Phone},
			})
		} else {
			s.producer.Queue("notification:email", domain.MailerMessage{
				To:      result.Edges.Contact.Email,
				Subject: "Trip Ticket - BookiRide",
				Data: map[string]any{
					"data":   presenters.BookingTicketResponse(result),
					"url":    config.App().AppWebsiteURL,
					"appUrl": config.App().AppURL,
					"bookiContact": map[string]string{
						"email": config.Contact().Email,
						"phone": config.Contact().Phone,
					},
				},
				Template: "newbooking",
			})
		}
		return result, nil
	}
	return nil, errors.New(strings.ToLower(resp.Message))
}

// Fetch implements gateways.BookingService.
func (s *service) Fetch(id int) (*ent.Booking, error) {
	return s.repo.Read(id)
}

// FetchByBookingNumber implements gateways.BookingService.
func (s *service) FetchByBookingNumber(id string) (*ent.Booking, error) {
	return s.repo.ReadByBookingNumber(id)
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
