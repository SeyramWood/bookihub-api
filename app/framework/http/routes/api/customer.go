package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"

	"github.com/SeyramWood/bookibus/app/framework/http/handlers/api"
	"github.com/SeyramWood/bookibus/app/framework/http/requests"
)

func CustomerRoutes(r fiber.Router, router *apiRouter) {
	handler := api.NewCustomerHandler(router.Adapter)

	customerGroup := r.Group("/customers")
	customerGroup.Get("", handler.FetchAll())
	customerGroup.Get("/:id", handler.Fetch())
	customerGroup.Get("/session/:id", handler.FetchBySession())
	customerGroup.Put("/:id", adaptor.HTTPMiddleware(requests.ValidateCustomerUpdate), handler.Update())
	customerGroup.Delete("/:id", handler.Remove())

}
