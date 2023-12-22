package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/bookibus/app/adapters/gateways"
	"github.com/SeyramWood/bookibus/app/adapters/presenters"
	"github.com/SeyramWood/bookibus/app/application/auth"
	"github.com/SeyramWood/bookibus/app/application/onboarding/bookibus_user"
	requeststructs "github.com/SeyramWood/bookibus/app/domain/request_structs"
	"github.com/SeyramWood/bookibus/app/framework/database"
	"github.com/SeyramWood/bookibus/ent"
)

type bookibushHandler struct {
	service gateways.BookibusUserService
}

func NewBookibusHandler(db *database.Adapter, producer gateways.EventProducer, storage gateways.StorageService) *bookibushHandler {
	return &bookibushHandler{
		service: bookibus_user.NewService(bookibus_user.NewRepository(db), producer, storage),
	}
}

func (h *bookibushHandler) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.BookibusUserRequest)
		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		result, err := h.service.Create(request)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.BookibusUserResponse(result))
	}
}
func (h *bookibushHandler) Fetch() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		result, err := h.service.Fetch(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.BookibusUserResponse(result))
	}
}
func (h *bookibushHandler) FetchAll() fiber.Handler {
	return func(c *fiber.Ctx) error {
		results, err := h.service.FetchAll(c.QueryInt("limit"), c.QueryInt("offset"))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.BookibusUsersResponse(results))
	}
}

func (h *bookibushHandler) Remove() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		if err := h.service.Remove(id); err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.MessageResponse("user deleted"))
	}
}
func (h *bookibushHandler) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.BookibusUserUpdateRequest)
		if c.BodyParser(request) != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(auth.ErrBadRequest))
		}
		id, _ := c.ParamsInt("id")
		result, err := h.service.Update(id, request)
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.BookibusUserResponse(result))
	}
}
