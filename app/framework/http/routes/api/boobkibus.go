package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"

	"github.com/SeyramWood/bookibus/app/framework/http/handlers/api"
	"github.com/SeyramWood/bookibus/app/framework/http/requests"
)

func BookibusRoutes(r fiber.Router, router *apiRouter) {
	handler := api.NewBookibusHandler(router.Adapter, router.EventProducer)

	userGroup := r.Group("/users")
	userGroup.Get("", handler.FetchAll())
	userGroup.Get("/:id", handler.Fetch())
	userGroup.Post("", adaptor.HTTPMiddleware(requests.ValidateBookibusUser), handler.Create())
	userGroup.Put("/:id", adaptor.HTTPMiddleware(requests.ValidateBookibusUserUpdate), handler.Update())
	userGroup.Delete("/:id", handler.Remove())

}
