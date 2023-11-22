package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"

	"github.com/SeyramWood/bookibus/app/framework/http/handlers/api"
	"github.com/SeyramWood/bookibus/app/framework/http/middlewares"
	"github.com/SeyramWood/bookibus/app/framework/http/requests"
)

func UnauthorizedRoutes(r fiber.Router, router *apiRouter) {
	authHandler := api.NewAuthHandler(router.Adapter, router.CacheSrv, router.jwt, router.EventProducer, router.StorageSrv)
	companyHandler := api.NewCompanyHandler(router.Adapter, router.EventProducer)
	tripHandler := api.NewTripHandler(router.Adapter)
	customerHandler := api.NewCustomerHandler(router.Adapter)
	bookingHandler := api.NewBookingHandler(router.Adapter, router.EventProducer, router.CacheSrv, router.Payment)
	parcelHandler := api.NewParcelHandler(router.Adapter, router.StorageSrv, router.Payment, router.EventProducer)

	authGroup := r.Group("/auth")
	authGroup.Post("/login", authHandler.Login())
	authGroup.Post("/refresh", middlewares.ValidateRefreshToken(), authHandler.RefreshToken())
	authGroup.Post("/send-otp", authHandler.SendPasswordResetCode())
	authGroup.Post("/verify/:otp", authHandler.VerifyOTP())
	authGroup.Put("/reset-password", authHandler.ResetPassword())

	customerGroup := r.Group("/customers")
	customerGroup.Post("", adaptor.HTTPMiddleware(requests.ValidateCustomer), customerHandler.Create())

	companyGroup := r.Group("/companies")
	companyGroup.Post("", adaptor.HTTPMiddleware(requests.ValidateCompany), companyHandler.Create())

	tripGroup := r.Group("/trips")
	tripGroup.Get("/all", tripHandler.FetchAllCustomer())
	tripGroup.Get("/popular", tripHandler.FetchAllPopular())

	bookingGroup := r.Group("/bookings")
	bookingGroup.Get("/all", bookingHandler.FetchAllCustomer())
	bookingGroup.Post("", adaptor.HTTPMiddleware(requests.ValidateBooking), bookingHandler.Create())
	bookingGroup.Put("/:id/cancel", adaptor.HTTPMiddleware(requests.ValidateBookingCancel), bookingHandler.CancelBooking())

	packageGroup := r.Group("/packages")
	packageGroup.Get("/code/:code", parcelHandler.FetchByCode())

}
