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

func (h *statisticsHandler) FetchAdminBestSelling() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// id, _ := c.ParamsInt("id")
		result, _ := h.service.FetchAdminBestSelling(c.QueryInt("limit"), c.QueryInt("offset"), c.Query("minDate"), c.Query("maxDate"))
		return c.Status(fiber.StatusOK).JSON(presenters.SuccessResponse(result))
	}
}
