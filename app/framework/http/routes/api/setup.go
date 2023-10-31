package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/application/payment"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/utils/jwt"
)

type apiRouter struct {
	Adapter       *database.Adapter
	StorageSrv    gateways.StorageService
	CacheSrv      gateways.CacheService
	EventProducer gateways.EventProducer
	Payment       gateways.PaymentService
	jwt           *jwt.JWT
}

func NewAPIRouter(params []any) *apiRouter {
	instance := &apiRouter{
		jwt:     jwt.NewJWT(),
		Payment: payment.NewPaymentService(),
	}
	instance.jwt.GenerateKey()
	return instance.instantiate(params)
}

func (r *apiRouter) Router(app *fiber.App) {
	apiGroup := app.Group("/api")

	UnauthorizedRoutes(apiGroup, r)

	// apiGroup.Use(middlewares.Authenticate(r.jwt))

	AuthRoutes(apiGroup, r)
	BookibusRoutes(apiGroup, r)
	CompanyRoutes(apiGroup, r)
	CustomerRoutes(apiGroup, r)
	BookingRoutes(apiGroup, r)
	ParcelRoutes(apiGroup, r)
	IncidentRoutes(apiGroup, r)

}

func (r *apiRouter) instantiate(params []any) *apiRouter {
	for _, param := range params {
		if adapter, ok := param.(*database.Adapter); ok {
			r.Adapter = adapter
			continue
		}
		if storageService, ok := param.(gateways.StorageService); ok {
			r.StorageSrv = storageService
			continue
		}
		if cacheService, ok := param.(gateways.CacheService); ok {
			r.CacheSrv = cacheService
			continue
		}
		if eventProducer, ok := param.(gateways.EventProducer); ok {
			r.EventProducer = eventProducer
			continue
		}
	}
	return r
}
