package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/framework/http/handlers/api"
	"github.com/SeyramWood/app/framework/http/requests"
)

func ParcelRoutes(r fiber.Router, router *apiRouter) {
	handler := api.NewParcelHandler(router.Adapter, router.StorageSrv, router.Payment, router.EventProducer)

	packageGroup := r.Group("/packages")
	packageGroup.Get("", handler.FetchAll())
	packageGroup.Get("/:id", handler.Fetch())
	packageGroup.Get("/company/:id", handler.FetchAllByCompany())
	packageGroup.Get("/driver/:id", handler.FetchAllByDriver())
	packageGroup.Post("/company/:id", requests.ValidateParcel(), handler.Create())
	packageGroup.Post("/:id/add-image", requests.ValidateParcelImage(), handler.AddImage())
	packageGroup.Put("/:id", requests.ValidateParcelUpdate(), handler.Update())
	packageGroup.Put("/:id/update-image", requests.ValidateParcelImageUpdate(), handler.UpdateImage())
	packageGroup.Put("/:id/update-status", requests.ValidateParcelDeliveredUpdate(), handler.UpdateStatus())
	packageGroup.Delete("/:id", handler.Remove())
	packageGroup.Delete("/:id/delete-image", handler.RemoveImage())

}
