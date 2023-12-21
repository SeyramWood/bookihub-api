package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/bookibus/app/framework/http/handlers/api"
)

func StatisticRoutes(r fiber.Router, router *apiRouter) {
	handler := api.NewStatisticsHandler(router.Adapter)

	statGroup := r.Group("/statistics")
	statGroup.Get("/admin/revenue", handler.FetchAdminRevenue())
	statGroup.Get("/admin/revenue-overview", handler.FetchAdminRevenueOverview())
	statGroup.Get("/admin/company-overview/:id", handler.FetchAdminCompanyOverview())
	statGroup.Get("/admin/best-selling", handler.FetchAdminBestSelling())
	statGroup.Get("/company/:id/revenue-overview", handler.FetchCompanyRevenueOverview())
	statGroup.Get("/company/:id/incident-overview", handler.FetchCompanyIncidentOverview())
	statGroup.Get("/company/:id/trip-overview", handler.FetchCompanyTripOverview())
	statGroup.Get("/company/:id/satisfaction-overview", handler.FetchCompanyMonthRevenue())
}
