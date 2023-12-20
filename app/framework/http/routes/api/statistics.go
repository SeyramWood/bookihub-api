package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/bookibus/app/framework/http/handlers/api"
)

func StatisticRoutes(r fiber.Router, router *apiRouter) {
	handler := api.NewStatisticsHandler(router.Adapter)

	statGroup := r.Group("/statistics")
	statGroup.Get("/admin/revenue", handler.FetchAdminRevenue())
	statGroup.Get("/admin/best-selling", handler.FetchAdminBestSelling())
}
