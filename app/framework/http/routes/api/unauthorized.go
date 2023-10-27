package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/framework/http/handlers/api"
	"github.com/SeyramWood/app/framework/http/middlewares"
	"github.com/SeyramWood/app/framework/http/requests"
)

func UnauthorizedRoutes(r fiber.Router, router *apiRouter) {
	authHandler := api.NewAuthHandler(router.Adapter, router.CacheSrv, router.jwt, router.EventProducer)
	companyHandler := api.NewCompanyHandler(router.Adapter, router.EventProducer)
	tripHandler := api.NewTripHandler(router.Adapter)
	customerHandler := api.NewCustomerHandler(router.Adapter)
	bookingHandler := api.NewBookingHandler(router.Adapter, router.EventProducer, router.CacheSrv)

	authGroup := r.Group("/auth")
	authGroup.Post("/login", authHandler.Login())
	authGroup.Post("/refresh", middlewares.ValidateRefreshToken(), authHandler.RefreshToken())
	authGroup.Post("/send-otp", authHandler.SendPasswordResetCode())
	authGroup.Post("/verify/:otp", authHandler.VerifyOTP())
	authGroup.Put("/reset-password", authHandler.ResetPassword())

	customerGroup := r.Group("/customers")
	customerGroup.Post("", requests.ValidateCustomer(), customerHandler.Create())

	companyGroup := r.Group("/companies")
	companyGroup.Post("", requests.ValidateCompany(), companyHandler.Create())

	tripGroup := r.Group("/trips")
	tripGroup.Get("/all", tripHandler.FetchAllCustomer())

	bookingGroup := r.Group("/bookings")
	bookingGroup.Get("/all", bookingHandler.FetchAllCustomer())
	bookingGroup.Put("/:id/cancel", requests.ValidateBookingCancel(), bookingHandler.CancelBooking())

}
