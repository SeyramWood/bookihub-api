package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/bookibus/app/adapters/gateways"
	"github.com/SeyramWood/bookibus/app/adapters/presenters"
	"github.com/SeyramWood/bookibus/app/application"
	"github.com/SeyramWood/bookibus/app/application/booking"
	requeststructs "github.com/SeyramWood/bookibus/app/domain/request_structs"
	"github.com/SeyramWood/bookibus/app/framework/database"
	"github.com/SeyramWood/bookibus/ent"
)

type bookingHandler struct {
	service gateways.BookingService
}

func NewBookingHandler(db *database.Adapter, producer gateways.EventProducer, cache gateways.CacheService, payment gateways.PaymentService) *bookingHandler {
	return &bookingHandler{
		service: booking.NewService(booking.NewRepository(db), producer, cache, payment),
	}
}

func (h *bookingHandler) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.BookingRequest)
		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		result, err := h.service.Create(request, c.Query("trans-type"))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.BookingResponse(result))
	}
}
func (h *bookingHandler) Fetch() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		result, err := h.service.Fetch(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.BookingResponse(result))
	}
}
func (h *bookingHandler) FetchAll() fiber.Handler {
	return func(c *fiber.Ctx) error {
		results, err := h.service.FetchAll(c.QueryInt("limit"), c.QueryInt("offset"), &requeststructs.BookingFilterRequest{
			BookingNumber: c.Query("trip-id"),
			FullName:      c.Query("full-name"),
			Active:        c.QueryBool("active"),
			Completed:     c.QueryBool("completed"),
			Canceled:      c.QueryBool("canceled"),
		})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.BookingsResponse(results))
	}
}
func (h *bookingHandler) FetchAllByCompany() fiber.Handler {
	return func(c *fiber.Ctx) error {
		companyId, _ := c.ParamsInt("id")
		results, err := h.service.FetchAllByCompany(companyId, c.QueryInt("limit"), c.QueryInt("offset"), &requeststructs.BookingFilterRequest{
			BookingNumber: c.Query("trip-id"),
			FullName:      c.Query("full-name"),
			Active:        c.QueryBool("active"),
			Completed:     c.QueryBool("completed"),
			Canceled:      c.QueryBool("canceled"),
		})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.BookingsResponse(results))
	}
}
func (h *bookingHandler) FetchAllCustomer() fiber.Handler {
	return func(c *fiber.Ctx) error {
		customerId := application.FormatSessionID(c.Locals("user"))
		if customerId != 0 {
			results, err := h.service.FetchAllCustomer(c.QueryInt("limit"), c.QueryInt("offset"), &requeststructs.BookingFilterRequest{
				BookingNumber: c.Query("trip-id"),
				FullName:      c.Query("full-name"),
				Active:        c.QueryBool("active"),
				Completed:     c.QueryBool("completed"),
				Canceled:      c.QueryBool("canceled"),
			}, customerId)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
			}
			return c.Status(fiber.StatusOK).JSON(presenters.BookingsResponse(results))
		}

		results, err := h.service.FetchAllCustomer(c.QueryInt("limit"), c.QueryInt("offset"), &requeststructs.BookingFilterRequest{
			BookingNumber: c.Query("trip-id"),
			FullName:      c.Query("full-name"),
			Active:        c.QueryBool("active"),
			Completed:     c.QueryBool("completed"),
			Canceled:      c.QueryBool("canceled"),
		})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.BookingsResponse(results))
	}
}

func (h *bookingHandler) Remove() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		if err := h.service.Remove(id); err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.MessageResponse("booking deleted"))
	}
}

func (h *bookingHandler) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.BookingUpdateRequest)
		if c.BodyParser(request) != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("bad request")))
		}
		id, _ := c.ParamsInt("id")
		result, err := h.service.Update(id, request)
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.BookingResponse(result))
	}
}
func (h *bookingHandler) CancelBooking() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.BookingCancelRequest)
		if c.BodyParser(request) != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("bad request")))
		}
		id, _ := c.ParamsInt("id")
		result, err := h.service.CancelBooking(id, request)
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.BookingResponse(result))
	}
}
