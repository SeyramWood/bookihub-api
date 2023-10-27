package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application/auth"
	"github.com/SeyramWood/app/application/onboarding/company_user"
	requeststructs "github.com/SeyramWood/app/domain/request_structs"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/ent"
)

type companyUserHandler struct {
	service gateways.CompanyUserService
}

func NewCompanyUserHandler(db *database.Adapter, producer gateways.EventProducer) *companyUserHandler {
	return &companyUserHandler{
		service: company_user.NewService(company_user.NewRepository(db), producer),
	}
}

func (h *companyUserHandler) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.CompanyUserRequest)
		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		result, err := h.service.Create(request)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.CompanyUserResponse(result))
	}
}
func (h *companyUserHandler) Fetch() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		result, err := h.service.Fetch(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.CompanyUserResponse(result))
	}
}
func (h *companyUserHandler) FetchAll() fiber.Handler {
	return func(c *fiber.Ctx) error {
		results, err := h.service.FetchAll(c.QueryInt("limit"), c.QueryInt("offset"))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.CompanyUsersResponse(results))
	}
}

func (h *companyUserHandler) FetchAllByCompany() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		results, err := h.service.FetchAllByCompany(id, c.QueryInt("limit"), c.QueryInt("offset"))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.CompanyUsersResponse(results))
	}
}

func (h *companyUserHandler) Remove() fiber.Handler {
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
func (h *companyUserHandler) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.CompanyUserUpdateRequest)
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
		return c.Status(fiber.StatusOK).JSON(presenters.CompanyUserResponse(result))
	}
}
