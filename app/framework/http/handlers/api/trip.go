package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/bookibus/app/adapters/gateways"
	"github.com/SeyramWood/bookibus/app/adapters/presenters"
	"github.com/SeyramWood/bookibus/app/application"
	"github.com/SeyramWood/bookibus/app/application/trip"
	requeststructs "github.com/SeyramWood/bookibus/app/domain/request_structs"
	"github.com/SeyramWood/bookibus/app/framework/database"
	"github.com/SeyramWood/bookibus/ent"
)

type tripHandler struct {
	service gateways.TripService
}

func NewTripHandler(db *database.Adapter) *tripHandler {
	return &tripHandler{
		service: trip.NewService(trip.NewRepository(db)),
	}
}

func (h *tripHandler) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.TripRequest)
		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		companyId, _ := c.ParamsInt("id")
		result, err := h.service.Create(companyId, request)
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.TripResponse(result))
	}
}

func (h *tripHandler) Fetch() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		result, err := h.service.Fetch(id)
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("trip not found")))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.TripResponse(result))
	}
}
func (h *tripHandler) FetchAll() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Query("datetime") != "" && !application.IsRFC3339Datetime(c.Query("datetime")) {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("datetime is invalid")))
		}
		results, err := h.service.FetchAll(c.QueryInt("limit"), c.QueryInt("offset"), &requeststructs.TripFilterRequest{
			From:      c.Query("from"),
			To:        c.Query("to"),
			Datetime:  c.Query("datetime"),
			Today:     c.QueryBool("today"),
			Scheduled: c.QueryBool("scheduled"),
			Completed: c.QueryBool("completed"),
		})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.TripsResponse(results))
	}
}
func (h *tripHandler) FetchAllSearch() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Query("searchKey") == "" {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("searchKey is empty")))
		}
		if c.Query("datetime") != "" && !application.IsRFC3339Datetime(c.Query("datetime")) {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("datetime is invalid")))
		}
		results, err := h.service.FetchAllSearch(c.Query("searchKey"), c.QueryInt("limit"), c.QueryInt("offset"), &requeststructs.TripFilterRequest{
			From:      c.Query("from"),
			To:        c.Query("to"),
			Datetime:  c.Query("datetime"),
			Today:     c.QueryBool("today"),
			Scheduled: c.QueryBool("scheduled"),
			Completed: c.QueryBool("completed"),
		})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.TripsResponse(results))
	}
}
func (h *tripHandler) FetchAllByCompany() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Query("datetime") != "" && !application.IsRFC3339Datetime(c.Query("datetime")) {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("datetime is invalid")))
		}
		if (c.Query("minTime") != "" && !application.IsTime(c.Query("minTime"))) || (c.Query("maxTime") != "" && !application.IsTime(c.Query("minTime"))) {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("minTime or maxTime is invalid")))
		}
		id, _ := c.ParamsInt("id")
		results, err := h.service.FetchAllByCompany(id, c.QueryInt("limit"), c.QueryInt("offset"), &requeststructs.TripFilterRequest{
			From:       c.Query("from"),
			To:         c.Query("to"),
			Datetime:   c.Query("datetime"),
			Today:      c.QueryBool("today"),
			Scheduled:  c.QueryBool("scheduled"),
			Completed:  c.QueryBool("completed"),
			Passengers: c.QueryInt("passengers"),
			TimeRange: func() string {
				if c.Query("minTime") != "" && c.Query("maxTime") != "" {
					return fmt.Sprintf("%s_%s", c.Query("minTime"), c.Query("maxTime"))
				}
				return ""
			}(),
			TripType:   c.Query("trip-type"),
			ReturnDate: c.Query("return-date"),
		})
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("company not found")))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}

		return c.Status(fiber.StatusOK).JSON(presenters.TripsResponse(results))
	}
}
func (h *tripHandler) FetchAllSearchByCompany() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Query("searchKey") == "" {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("searchKey is empty")))
		}
		if c.Query("datetime") != "" && !application.IsRFC3339Datetime(c.Query("datetime")) {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("datetime is invalid")))
		}
		id, _ := c.ParamsInt("id")
		results, err := h.service.FetchAllSearchByCompany(c.Query("searchKey"), id, c.QueryInt("limit"), c.QueryInt("offset"), &requeststructs.TripFilterRequest{
			From:       c.Query("from"),
			To:         c.Query("to"),
			Datetime:   c.Query("datetime"),
			Today:      c.QueryBool("today"),
			Scheduled:  c.QueryBool("scheduled"),
			Completed:  c.QueryBool("completed"),
			TripType:   c.Query("trip-type"),
			ReturnDate: c.Query("return-date"),
		})
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("company not found")))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.TripsResponse(results))
	}
}
func (h *tripHandler) FetchAllByDriver() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		results, err := h.service.FetchAllByDriver(id, c.QueryInt("limit"), c.QueryInt("offset"), &requeststructs.TripFilterRequest{
			Today:     c.QueryBool("today"),
			Scheduled: c.QueryBool("scheduled"),
			Completed: c.QueryBool("completed"),
		})
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("company not found")))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.TripsResponse(results))
	}
}
func (h *tripHandler) FetchAllCustomer() fiber.Handler {
	return func(c *fiber.Ctx) error {
		results, err := h.service.FetchAllCustomer(c.QueryInt("limit"), c.QueryInt("offset"), &requeststructs.CustomerTripFilterRequest{
			CompanyID:     c.QueryInt("company-id"),
			TripType:      c.Query("trip-type"),
			From:          c.Query("from"),
			To:            c.Query("to"),
			DepartureDate: c.Query("departure-date"),
			ReturnDate:    c.Query("return-date"),
			Passengers:    c.QueryInt("passenger"),
		})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.TripsResponse(results))
	}
}
func (h *tripHandler) FetchAllPopular() fiber.Handler {
	return func(c *fiber.Ctx) error {
		results, err := h.service.FetchAllPopular(c.QueryInt("limit", 6), c.QueryInt("offset"), c.Query("localDatetime"))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.TripsResponse(results))
	}
}

func (h *tripHandler) Remove() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		if err := h.service.Remove(id); err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("trip not found")))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.MessageResponse("trip deleted"))
	}
}
func (h *tripHandler) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.TripUpdateRequest)
		if c.BodyParser(request) != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("bad request")))
		}
		id, _ := c.ParamsInt("id")
		result, err := h.service.Update(id, request)
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("trip not found")))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.TripResponse(result))
	}
}

func (h *tripHandler) UpdateInspection() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.TripInspectionStatusRequest)
		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("bad request")))
		}
		id, _ := c.ParamsInt("id")
		result, err := h.service.UpdateInspection(id, request)
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("trip not found")))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.TripResponse(result))
	}
}
func (h *tripHandler) UpdateStatus() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Query("status") == "" {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("bad request")))
		}
		id, _ := c.ParamsInt("id")
		result, err := h.service.UpdateStatus(id, c.Query("status"))
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("trip not found")))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.TripResponse(result))
	}
}

func (h *tripHandler) UpdateSchedule() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Query("status") == "" {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("bad request")))
		}
		id, _ := c.ParamsInt("id")
		result, err := h.service.UpdateSchedule(id, c.QueryBool("status"))
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("trip not found")))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.TripResponse(result))
	}
}
