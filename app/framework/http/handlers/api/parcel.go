package api

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application/parcel"
	requeststructs "github.com/SeyramWood/app/domain/request_structs"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/ent"
)

type parcelHandler struct {
	service gateways.ParcelService
}

func NewParcelHandler(db *database.Adapter, storage gateways.StorageService, payment gateways.PaymentService, produce gateways.EventProducer) *parcelHandler {
	return &parcelHandler{
		service: parcel.NewService(parcel.NewRepository(db), storage, payment, produce),
	}
}

func (h *parcelHandler) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.ParcelRequest)
		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		form, err := c.MultipartForm()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		request.Image = form.File["image"]
		companyId, _ := c.ParamsInt("id")
		result, err := h.service.Create(companyId, request, c.Query("trans-type"))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.ParcelResponse(result))
	}
}
func (h *parcelHandler) AddImage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.ParcelImageRequest)
		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		form, err := c.MultipartForm()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		request.Image = form.File["image"]
		id, _ := c.ParamsInt("id")
		result, err := h.service.AddImage(id, request)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.ParcelResponse(result))
	}
}
func (h *parcelHandler) Fetch() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		result, err := h.service.Fetch(id)
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("package not found")))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.ParcelResponse(result))
	}
}
func (h *parcelHandler) FetchAll() fiber.Handler {
	return func(c *fiber.Ctx) error {
		results, err := h.service.FetchAll(c.QueryInt("limit"), c.QueryInt("offset"), &requeststructs.ParcelFilterRequest{
			Status: c.Query("status"),
		})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.ParcelsResponse(results))
	}
}

func (h *parcelHandler) FetchAllByCompany() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		results, err := h.service.FetchAllByCompany(id, c.QueryInt("limit"), c.QueryInt("offset"), &requeststructs.ParcelFilterRequest{
			Status: c.Query("status"),
		})
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("company not found")))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.ParcelsResponse(results))
	}
}

func (h *parcelHandler) FetchAllByDriver() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		results, err := h.service.FetchAllByDriver(id, c.QueryInt("limit"), c.QueryInt("offset"), &requeststructs.ParcelFilterRequest{
			Status: c.Query("status"),
		})
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("company not found")))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.ParcelsResponse(results))
	}
}

func (h *parcelHandler) Remove() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		if err := h.service.Remove(id); err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("package not found")))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.MessageResponse("package deleted"))
	}
}
func (h *parcelHandler) RemoveImage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		if err := h.service.RemoveImage(id); err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("image not found")))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.MessageResponse("image deleted"))
	}
}
func (h *parcelHandler) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.ParcelUpdateRequest)
		if c.BodyParser(request) != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("bad request")))
		}
		id, _ := c.ParamsInt("id")
		result, err := h.service.Update(id, request)
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("package not found")))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.ParcelResponse(result))
	}
}
func (h *parcelHandler) UpdateStatus() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.ParcelDeliveredRequest)
		if c.BodyParser(request) != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("bad request")))
		}
		form, err := c.MultipartForm()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		id, _ := c.ParamsInt("id")
		request.Image = form.File["image"]
		result, err := h.service.UpdateStatus(id, request)
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("package not found")))
			}
			if errors.Is(err, parcel.ErrWrongPackageCode) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(parcel.ErrWrongPackageCode))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.ParcelResponse(result))
	}
}

func (h *parcelHandler) UpdateImage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.ParcelImageUpdateRequest)
		if c.BodyParser(request) != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("bad request")))
		}
		image, err := c.FormFile("image")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		id, _ := c.ParamsInt("id")
		request.Image = image
		result, err := h.service.UpdateImage(id, request)
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("image not found")))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.ParcelImageResponse(result))
	}
}
