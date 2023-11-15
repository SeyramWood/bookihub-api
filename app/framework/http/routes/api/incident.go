package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"

	"github.com/SeyramWood/bookibus/app/framework/http/handlers/api"
	"github.com/SeyramWood/bookibus/app/framework/http/requests"
)

func IncidentRoutes(r fiber.Router, router *apiRouter) {
	handler := api.NewIncidentHandler(router.Adapter, router.StorageSrv, router.EventProducer)

	incidentGroup := r.Group("/incidents")
	incidentGroup.Get("", handler.FetchAll())
	incidentGroup.Get("/:id", handler.Fetch())
	incidentGroup.Get("/company/:id", handler.FetchAllByCompany())
	incidentGroup.Get("/driver/:id", handler.FetchAllByDriver())
	incidentGroup.Post("/company/:id", adaptor.HTTPMiddleware(requests.ValidateIncident), handler.Create())
	incidentGroup.Post("/:id/add-image", adaptor.HTTPMiddleware(requests.ValidateIncidentImage), handler.AddImage())
	incidentGroup.Put("/:id", adaptor.HTTPMiddleware(requests.ValidateIncidentUpdate), handler.Update())
	incidentGroup.Put("/:id/update-image", adaptor.HTTPMiddleware(requests.ValidateIncidentImageUpdate), handler.UpdateImage())
	incidentGroup.Put("/:id/update-audio", adaptor.HTTPMiddleware(requests.ValidateIncidentAudioUpdate), handler.UpdateAudio())
	incidentGroup.Put("/:id/update-status", handler.UpdateStatus())
	incidentGroup.Delete("/:id", handler.Remove())
	incidentGroup.Delete("/:id/delete-image", handler.RemoveImage())
	incidentGroup.Delete("/:id/delete-audio", handler.RemoveAudio())

}
