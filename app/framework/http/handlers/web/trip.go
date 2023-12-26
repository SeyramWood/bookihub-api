package web

import (
	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/bookibus/app/adapters/gateways"
	"github.com/SeyramWood/bookibus/app/adapters/presenters"
	"github.com/SeyramWood/bookibus/app/application/booking"
	"github.com/SeyramWood/bookibus/app/framework/database"
	"github.com/SeyramWood/bookibus/config"
)

type tripHandler struct {
	service gateways.BookingService
}

func NewTripHandler(db *database.Adapter, producer gateways.EventProducer, cache gateways.CacheService, payment gateways.PaymentService) *tripHandler {
	return &tripHandler{
		service: booking.NewService(booking.NewRepository(db), producer, cache, payment),
	}
}

func (h *tripHandler) Ticket() fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := h.service.FetchByBookingNumber(c.Params("id"))
		if err != nil {
			return c.Render("404", fiber.Map{
				"url": config.App().AppWebsiteURL,
			})
		}

		return c.Render("ticket", fiber.Map{
			"data":   presenters.BookingTicketResponse(result),
			"url":    config.App().AppWebsiteURL,
			"appUrl": config.App().AppURL,
			"bookiContact": map[string]string{
				"email": config.Contact().Email,
				"phone": config.Contact().Phone,
			},
		})
	}
}
