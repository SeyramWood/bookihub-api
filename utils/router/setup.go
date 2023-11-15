package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/bookibus/app/framework/http/routes/api"
	"github.com/SeyramWood/bookibus/app/framework/http/routes/web"
)

type Router interface {
	Router(fiber *fiber.App)
}

func NewRouter(
	app *fiber.App, params ...any,
) {
	setup(app, api.NewAPIRouter(params), web.NewWebRouter(params))
}

func setup(app *fiber.App, routers ...Router) {
	for _, r := range routers {
		r.Router(app)
	}
}
