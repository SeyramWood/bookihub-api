package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/framework/http/handlers/api"
	"github.com/SeyramWood/app/framework/http/requests"
)

func BookingRoutes(r fiber.Router, router *apiRouter) {
	handler := api.NewBookingHandler(router.Adapter, router.EventProducer, router.CacheSrv, router.Payment)

	bookingGroup := r.Group("/bookings")
	bookingGroup.Get("", handler.FetchAll())
	bookingGroup.Get("/company/:id", handler.FetchAllByCompany())
	bookingGroup.Get("/:id", handler.Fetch())
	bookingGroup.Post("", requests.ValidateBooking(), handler.Create())
	bookingGroup.Put("/:id", requests.ValidateBookingUpdate(), handler.Update())
	bookingGroup.Delete("/:id", handler.Remove())

}
