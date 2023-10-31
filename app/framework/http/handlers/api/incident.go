package api

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application/incident"
	requeststructs "github.com/SeyramWood/app/domain/request_structs"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/ent"
)

type incidentHandler struct {
	service gateways.IncidentService
}

func NewIncidentHandler(db *database.Adapter, storage gateways.StorageService, produce gateways.EventProducer) *incidentHandler {
	return &incidentHandler{
		service: incident.NewService(incident.NewRepository(db), storage, produce),
	}
}

func (h *incidentHandler) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.IncidentRequest)
		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		errOptional := errors.New("there is no uploaded file associated with the given key")
		form, err := c.MultipartForm()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		audio, err := c.FormFile("voiceNote")
		if err != nil && err.Error() != errOptional.Error() {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		request.Image = form.File["image"]
		request.Audio = audio
		companyId, _ := c.ParamsInt("id")
		result, err := h.service.Create(companyId, request)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.IncidentResponse(result))
	}
}
func (h *incidentHandler) AddImage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.IncidentImageRequest)
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
		return c.Status(fiber.StatusOK).JSON(presenters.IncidentResponse(result))
	}
}
func (h *incidentHandler) Fetch() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		result, err := h.service.Fetch(id)
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("incident not found")))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.IncidentResponse(result))
	}
}
func (h *incidentHandler) FetchAll() fiber.Handler {
	return func(c *fiber.Ctx) error {
		results, err := h.service.FetchAll(c.QueryInt("limit"), c.QueryInt("offset"), &requeststructs.IncidentFilterRequest{
			Datetime: c.Query("datetime"),
		})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.IncidentsResponse(results))
	}
}

func (h *incidentHandler) FetchAllByCompany() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		results, err := h.service.FetchAllByCompany(id, c.QueryInt("limit"), c.QueryInt("offset"), &requeststructs.IncidentFilterRequest{
			Datetime: c.Query("datetime"),
		})
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("company not found")))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.IncidentsResponse(results))
	}
}

func (h *incidentHandler) FetchAllByDriver() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		results, err := h.service.FetchAllByDriver(id, c.QueryInt("limit"), c.QueryInt("offset"), &requeststructs.IncidentFilterRequest{
			Datetime: c.Query("datetime"),
		})
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("company not found")))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.IncidentsResponse(results))
	}
}

func (h *incidentHandler) Remove() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		if err := h.service.Remove(id); err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("incident not found")))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.MessageResponse("incident deleted"))
	}
}
func (h *incidentHandler) RemoveImage() fiber.Handler {
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
func (h *incidentHandler) RemoveAudio() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		if err := h.service.RemoveAudio(id); err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("incident not found")))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.MessageResponse("audio deleted"))
	}
}
func (h *incidentHandler) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.IncidentUpdateRequest)
		if c.BodyParser(request) != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("bad request")))
		}
		id, _ := c.ParamsInt("id")
		result, err := h.service.Update(id, request)
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("incident not found")))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.IncidentResponse(result))
	}
}
func (h *incidentHandler) UpdateImage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.IncidentImageUpdateRequest)
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
		return c.Status(fiber.StatusOK).JSON(presenters.IncidentImageResponse(result))
	}
}
func (h *incidentHandler) UpdateAudio() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.IncidentAudioUpdateRequest)
		if c.BodyParser(request) != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("bad request")))
		}
		audio, err := c.FormFile("voiceNote")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		id, _ := c.ParamsInt("id")
		request.Audio = audio
		result, err := h.service.UpdateAudio(id, request)
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("voice note not found")))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.SuccessResponse(result))
	}
}
