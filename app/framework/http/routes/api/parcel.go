package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"

	"github.com/SeyramWood/bookibus/app/framework/http/handlers/api"
	"github.com/SeyramWood/bookibus/app/framework/http/requests"
)

func ParcelRoutes(r fiber.Router, router *apiRouter) {
	handler := api.NewParcelHandler(router.Adapter, router.StorageSrv, router.Payment, router.EventProducer)

	packageGroup := r.Group("/packages")
	packageGroup.Get("", handler.FetchAll())
	packageGroup.Get("/:id", handler.Fetch())
	packageGroup.Get("/company/:id", handler.FetchAllByCompany())
	packageGroup.Get("/driver/:id", handler.FetchAllByDriver())
	packageGroup.Post("/company/:id", adaptor.HTTPMiddleware(requests.ValidateParcel), handler.Create())
	packageGroup.Post("/:id/add-image", adaptor.HTTPMiddleware(requests.ValidateParcelImage), handler.AddImage())
	packageGroup.Put("/:id", adaptor.HTTPMiddleware(requests.ValidateParcelUpdate), handler.Update())
	packageGroup.Put("/:id/update-image", adaptor.HTTPMiddleware(requests.ValidateParcelImageUpdate), handler.UpdateImage())
	packageGroup.Put("/:id/update-status", adaptor.HTTPMiddleware(requests.ValidateParcelDeliveredUpdate), handler.UpdateStatus())
	packageGroup.Delete("/:id", handler.Remove())
	packageGroup.Delete("/:id/delete-image", handler.RemoveImage())

}
