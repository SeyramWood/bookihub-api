package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/framework/http/handlers/api"
	"github.com/SeyramWood/app/framework/http/requests"
)

func BookibusRoutes(r fiber.Router, router *apiRouter) {
	handler := api.NewBookibusHandler(router.Adapter, router.EventProducer)

	userGroup := r.Group("/users")
	userGroup.Get("", handler.FetchAll())
	userGroup.Get("/:id", handler.Fetch())
	userGroup.Post("", requests.ValidateBookibusUser(), handler.Create())
	userGroup.Put("/:id", requests.ValidateBookibusUserUpdate(), handler.Update())
	userGroup.Delete("/:id", handler.Remove())

}
