package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"

	"github.com/SeyramWood/bookibus/app/framework/http/handlers/api"
	"github.com/SeyramWood/bookibus/app/framework/http/requests"
)

func BookingRoutes(r fiber.Router, router *apiRouter) {
	handler := api.NewBookingHandler(router.Adapter, router.EventProducer, router.CacheSrv, router.Payment)

	bookingGroup := r.Group("/bookings")
	bookingGroup.Get("", handler.FetchAll())
	bookingGroup.Get("/company/:id", handler.FetchAllByCompany())
	bookingGroup.Put("/:id", adaptor.HTTPMiddleware(requests.ValidateBookingUpdate), handler.Update())
	bookingGroup.Delete("/:id", handler.Remove())

}
