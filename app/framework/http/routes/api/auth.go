package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/bookibus/app/framework/http/handlers/api"
)

func AuthRoutes(r fiber.Router, router *apiRouter) {
	handler := api.NewAuthHandler(router.Adapter, router.CacheSrv, router.jwt, router.EventProducer)
	authGroup := r.Group("/auth")
	authGroup.Get("/session", handler.GetSession())
	authGroup.Put("/update-password", handler.UpdatePassword())
}
