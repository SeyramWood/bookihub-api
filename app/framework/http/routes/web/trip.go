package web

import (
	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/bookibus/app/framework/http/handlers/web"
)

func TripRoutes(r fiber.Router, router *webRouter) {
	handler := web.NewTripHandler(router.Adapter, router.EventProducer, router.CacheSrv, router.Payment)
	tripGroup := r.Group("/trips")
	tripGroup.Get("/ticket/:id/download", handler.Ticket())
}
