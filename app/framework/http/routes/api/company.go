package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"

	"github.com/SeyramWood/bookibus/app/framework/http/handlers/api"
	"github.com/SeyramWood/bookibus/app/framework/http/requests"
)

func CompanyRoutes(r fiber.Router, router *apiRouter) {
	handler := api.NewCompanyHandler(router.Adapter, router.EventProducer, router.StorageSrv)
	terminalHandler := api.NewTerminalHandler(router.Adapter)
	userHandler := api.NewCompanyUserHandler(router.Adapter, router.EventProducer)
	vehicleHandler := api.NewVehicleHandler(router.Adapter, router.StorageSrv)
	routeHandler := api.NewRouteHandler(router.Adapter)
	tripHandler := api.NewTripHandler(router.Adapter)

	userGroup := r.Group("/company/users")
	userGroup.Get("", userHandler.FetchAll())
	userGroup.Get("/:id", userHandler.Fetch())
	userGroup.Post("", adaptor.HTTPMiddleware(requests.ValidateCompanyUser), userHandler.Create())
	userGroup.Put("/:id", adaptor.HTTPMiddleware(requests.ValidateCompanyUserUpdate), userHandler.Update())
	userGroup.Delete("/:id", userHandler.Remove())

	staffGroup := r.Group("/company")
	staffGroup.Get("/:id/users", userHandler.FetchAllByCompany())

	companyGroup := r.Group("/companies")
	companyGroup.Get("", handler.FetchAll())
	companyGroup.Get("/:id", handler.Fetch())
	companyGroup.Post("/onboard-new", adaptor.HTTPMiddleware(requests.ValidateNewOnboarding), handler.BookiOnboard())
	companyGroup.Put("/:id/onboard", adaptor.HTTPMiddleware(requests.ValidateOnboarding), handler.Onboard())
	companyGroup.Put("/:id/update-bank-account", adaptor.HTTPMiddleware(requests.ValidateCompanyBankAccountUpdate), handler.UpdateBankAccount())
	companyGroup.Put("/:id/update-contact-person", adaptor.HTTPMiddleware(requests.ValidateCompanyContactPersonUpdate), handler.UpdateContactPerson())
	companyGroup.Put("/:id/update-logo", adaptor.HTTPMiddleware(requests.ValidateCompanyLogoUpdate), handler.UpdateLogo())
	companyGroup.Put("/:id/update-certificate", adaptor.HTTPMiddleware(requests.ValidateCompanyCertificateUpdate), handler.UpdateCertificate())
	companyGroup.Put("/:id/update-status", handler.UpdateStatus())
	companyGroup.Put("/:id", adaptor.HTTPMiddleware(requests.ValidateCompanyUpdate), handler.Update())
	companyGroup.Delete("/:id", handler.Remove())

	terminalGroup := r.Group("/terminals")
	terminalGroup.Get("", terminalHandler.FetchAll())
	terminalGroup.Get("/company/:id", terminalHandler.FetchAllByCompany())
	terminalGroup.Get("/:id", terminalHandler.Fetch())
	terminalGroup.Post("/company/:id", adaptor.HTTPMiddleware(requests.ValidateTerminal), terminalHandler.Create())
	terminalGroup.Put("/:id", adaptor.HTTPMiddleware(requests.ValidateTerminal), terminalHandler.Update())
	terminalGroup.Delete("/:id", terminalHandler.Remove())

	vehicleGroup := r.Group("/vehicles")
	vehicleGroup.Get("", vehicleHandler.FetchAll())
	vehicleGroup.Get("/company/:id", vehicleHandler.FetchAllByCompany())
	vehicleGroup.Get("/:id", vehicleHandler.Fetch())
	vehicleGroup.Post("/company/:id", adaptor.HTTPMiddleware(requests.ValidateVehicle), vehicleHandler.Create())
	vehicleGroup.Post("/:id/add-image", adaptor.HTTPMiddleware(requests.ValidateVehicleImage), vehicleHandler.AddImage())
	vehicleGroup.Put("/:id", adaptor.HTTPMiddleware(requests.ValidateVehicleUpdate), vehicleHandler.Update())
	vehicleGroup.Put("/:id/update-image", adaptor.HTTPMiddleware(requests.ValidateVehicleImageUpdate), vehicleHandler.UpdateImage())
	vehicleGroup.Delete("/:id", vehicleHandler.Remove())
	vehicleGroup.Delete("/:id/delete-image", vehicleHandler.RemoveImage())

	routeGroup := r.Group("/routes")
	routeGroup.Get("", routeHandler.FetchAll())
	routeGroup.Get("/company/:id", routeHandler.FetchAllByCompany())
	routeGroup.Get("/:id", routeHandler.Fetch())
	routeGroup.Post("/company/:id", adaptor.HTTPMiddleware(requests.ValidateRoute), routeHandler.Create())
	routeGroup.Post("/:id/add-stop", adaptor.HTTPMiddleware(requests.ValidateRouteStop), routeHandler.AddRouteStop())
	routeGroup.Put("/:id", adaptor.HTTPMiddleware(requests.ValidateRouteUpdate), routeHandler.Update())
	routeGroup.Put("/:id/update-stop", adaptor.HTTPMiddleware(requests.ValidateRouteStop), routeHandler.UpdateStop())
	routeGroup.Delete("/:id", routeHandler.Remove())
	routeGroup.Delete("/:id/delete-stop", routeHandler.RemoveStop())

	tripGroup := r.Group("/trips")
	tripGroup.Get("", tripHandler.FetchAll())
	tripGroup.Get("/search", tripHandler.FetchAllSearch())
	tripGroup.Get("/company/:id", tripHandler.FetchAllByCompany())
	tripGroup.Get("/company/:id/search", tripHandler.FetchAllSearchByCompany())
	tripGroup.Get("/driver/:id", tripHandler.FetchAllByDriver())
	tripGroup.Get("/:id", tripHandler.Fetch())
	tripGroup.Post("/company/:id", adaptor.HTTPMiddleware(requests.ValidateTrip), tripHandler.Create())
	tripGroup.Put("/:id", adaptor.HTTPMiddleware(requests.ValidateTripUpdate), tripHandler.Update())
	tripGroup.Put("/:id/update-status", tripHandler.UpdateStatus())
	tripGroup.Put("/:id/update-schedule", tripHandler.UpdateSchedule())
	tripGroup.Put("/:id/update-inspection", tripHandler.UpdateInspection())
	tripGroup.Delete("/:id", tripHandler.Remove())

}
