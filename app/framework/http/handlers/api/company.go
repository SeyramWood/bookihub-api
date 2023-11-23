package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/bookibus/app/adapters/gateways"
	"github.com/SeyramWood/bookibus/app/adapters/presenters"
	"github.com/SeyramWood/bookibus/app/application/onboarding/company"
	requeststructs "github.com/SeyramWood/bookibus/app/domain/request_structs"
	"github.com/SeyramWood/bookibus/app/framework/database"
	"github.com/SeyramWood/bookibus/ent"
)

type companyHandler struct {
	service gateways.CompanyService
}

func NewCompanyHandler(db *database.Adapter, producer gateways.EventProducer, storage gateways.StorageService) *companyHandler {
	return &companyHandler{
		service: company.NewService(company.NewRepository(db), producer, storage),
	}
}

func (h *companyHandler) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.CompanyRequest)
		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		result, err := h.service.Create(request)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.CompanyResponse(result))
	}
}

func (h *companyHandler) BookiOnboard() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.BookiOnboardingRequest)
		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		result, err := h.service.BookiOnboard(request)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.CompanyResponse(result))
	}
}

func (h *companyHandler) Fetch() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		result, err := h.service.Fetch(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.CompanyResponse(result))
	}
}

func (h *companyHandler) FetchAll() fiber.Handler {
	return func(c *fiber.Ctx) error {
		results, err := h.service.FetchAll(c.QueryInt("limit"), c.QueryInt("offset"))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.CompaniesResponse(results))
	}
}

func (h *companyHandler) Remove() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		if err := h.service.Remove(id); err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.MessageResponse("company deleted"))
	}
}

func (h *companyHandler) Onboard() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.CompanyOnboardingRequest)
		if c.BodyParser(request) != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("bad request")))
		}
		id, _ := c.ParamsInt("id")
		result, err := h.service.Onboard(id, request)
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.CompanyResponse(result))
	}
}

func (h *companyHandler) UpdateBankAccount() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.CompanyBankAccountRequest)
		if c.BodyParser(request) != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("bad request")))
		}
		id, _ := c.ParamsInt("id")
		result, err := h.service.UpdateBankAccount(id, request)
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.CompanyResponse(result))
	}
}
func (h *companyHandler) UpdateContactPerson() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.CompanyContactPersonUpdateRequest)
		if c.BodyParser(request) != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("bad request")))
		}
		id, _ := c.ParamsInt("id")
		result, err := h.service.UpdateContactPerson(id, request)
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.CompanyResponse(result))
	}
}
func (h *companyHandler) UpdateCertificate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.CompanyCertificateUpdateRequest)
		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		certificate, err := c.FormFile("businessCertificate")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		request.BusinessCertificate = certificate
		id, _ := c.ParamsInt("id")
		result, err := h.service.UpdateCertificate(id, request)
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.CompanyResponse(result))
	}
}
func (h *companyHandler) UpdateStatus() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Query("status") == "" {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("bad request")))
		}
		id, _ := c.ParamsInt("id")
		result, err := h.service.UpdateStatus(id, c.Query("status"))
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.CompanyResponse(result))
	}
}

func (h *companyHandler) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.CompanyUpdateRequest)
		if c.BodyParser(request) != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("bad request")))
		}
		id, _ := c.ParamsInt("id")
		result, err := h.service.Update(id, request)
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.CompanyResponse(result))
	}
}
