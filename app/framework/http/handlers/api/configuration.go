package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/bookibus/app/adapters/gateways"
	"github.com/SeyramWood/bookibus/app/adapters/presenters"
	"github.com/SeyramWood/bookibus/app/application/configuration"
	requeststructs "github.com/SeyramWood/bookibus/app/domain/request_structs"
	"github.com/SeyramWood/bookibus/app/framework/database"
	"github.com/SeyramWood/bookibus/ent"
)

type configurationHandler struct {
	service gateways.ConfigurationService
}

func NewConfigurationHandler(db *database.Adapter) *configurationHandler {
	return &configurationHandler{
		service: configuration.NewService(configuration.NewRepository(db)),
	}
}

func (h *configurationHandler) FetchCharge() fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := h.service.FetchCharge()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.ChargeResponse(result))
	}
}

func (h *configurationHandler) CreateCharge() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.TransactionChargeRequest)
		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		result, err := h.service.CreateCharge(request)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.ChargeResponse(result))
	}
}

func (h *configurationHandler) UpdateCharge() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.TransactionChargeRequest)
		if c.BodyParser(request) != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("bad request")))
		}
		id, _ := c.ParamsInt("id")
		result, err := h.service.UpdateCharge(id, request)
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.ChargeResponse(result))
	}
}
