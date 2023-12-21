package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/bookibus/app/adapters/gateways"
	"github.com/SeyramWood/bookibus/app/adapters/presenters"
	"github.com/SeyramWood/bookibus/app/application/statistics"
	"github.com/SeyramWood/bookibus/app/framework/database"
)

type statisticsHandler struct {
	service gateways.StatisticsService
}

func NewStatisticsHandler(db *database.Adapter) *statisticsHandler {
	return &statisticsHandler{
		service: statistics.NewService(statistics.NewRepository(db)),
	}
}

func (h *statisticsHandler) FetchAdminRevenue() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// id, _ := c.ParamsInt("id")
		result := h.service.FetchAdminRevenue()
		return c.Status(fiber.StatusOK).JSON(presenters.SuccessResponse(result))
	}
}

func (h *statisticsHandler) FetchAdminRevenueOverview() fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, _ := h.service.FetchAdminRevenueOverview(c.Query("filter"))
		return c.Status(fiber.StatusOK).JSON(presenters.SuccessResponse(result))
	}
}
func (h *statisticsHandler) FetchAdminCompanyOverview() fiber.Handler {
	return func(c *fiber.Ctx) error {
		companyId, _ := c.ParamsInt("id")
		result, _ := h.service.FetchAdminCompanyOverview(companyId, c.Query("filter"))
		return c.Status(fiber.StatusOK).JSON(presenters.SuccessResponse(result))
	}
}

func (h *statisticsHandler) FetchAdminBestSelling() fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, _ := h.service.FetchAdminBestSelling(c.QueryInt("limit", 3), c.QueryInt("offset"), c.Query("minDate"), c.Query("maxDate"))
		return c.Status(fiber.StatusOK).JSON(presenters.SuccessResponse(result))
	}
}

func (h *statisticsHandler) FetchCompanyTripOverview() fiber.Handler {
	return func(c *fiber.Ctx) error {
		companyId, _ := c.ParamsInt("id")
		result, _ := h.service.FetchCompanyTripOverview(companyId, c.Query("filter"))
		return c.Status(fiber.StatusOK).JSON(presenters.SuccessResponse(result))
	}
}
func (h *statisticsHandler) FetchCompanyRevenueOverview() fiber.Handler {
	return func(c *fiber.Ctx) error {
		companyId, _ := c.ParamsInt("id")
		result, _ := h.service.FetchCompanyRevenueOverview(companyId, c.Query("filter"))
		return c.Status(fiber.StatusOK).JSON(presenters.SuccessResponse(result))
	}
}
func (h *statisticsHandler) FetchCompanyMonthRevenue() fiber.Handler {
	return func(c *fiber.Ctx) error {
		companyId, _ := c.ParamsInt("id")
		result, _ := h.service.FetchCompanyMonthRevenue(companyId)
		return c.Status(fiber.StatusOK).JSON(presenters.SuccessResponse(result))
	}
}
func (h *statisticsHandler) FetchCompanyIncidentOverview() fiber.Handler {
	return func(c *fiber.Ctx) error {
		companyId, _ := c.ParamsInt("id")
		result, _ := h.service.FetchCompanyIncidentOverview(companyId, c.Query("filter"))
		return c.Status(fiber.StatusOK).JSON(presenters.SuccessResponse(result))
	}
}
