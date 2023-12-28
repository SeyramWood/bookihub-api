package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"

	"github.com/SeyramWood/bookibus/app/framework/http/handlers/api"
	"github.com/SeyramWood/bookibus/app/framework/http/requests"
)

func AuthRoutes(r fiber.Router, router *apiRouter) {
	handler := api.NewAuthHandler(router.Adapter, router.CacheSrv, router.jwt, router.EventProducer, router.StorageSrv)
	authGroup := r.Group("/auth")
	authGroup.Get("/session", handler.GetSession())
	authGroup.Put("/update-avatar/:id", adaptor.HTTPMiddleware(requests.ValidateAvatarUpdate), handler.UpdateAvatar())
	authGroup.Put("/update-password", handler.UpdatePassword())
}
