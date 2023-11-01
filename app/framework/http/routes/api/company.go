package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/framework/http/handlers/api"
	"github.com/SeyramWood/app/framework/http/requests"
)

func CompanyRoutes(r fiber.Router, router *apiRouter) {
	handler := api.NewCompanyHandler(router.Adapter, router.EventProducer)
	userHandler := api.NewCompanyUserHandler(router.Adapter, router.EventProducer)
	vehicleHandler := api.NewVehicleHandler(router.Adapter, router.StorageSrv)
	routeHandler := api.NewRouteHandler(router.Adapter)
	tripHandler := api.NewTripHandler(router.Adapter)

	userGroup := r.Group("/company/users")
	userGroup.Get("", userHandler.FetchAll())
	userGroup.Get("/:id", userHandler.Fetch())
	userGroup.Post("", requests.ValidateCompanyUser(), userHandler.Create())
	userGroup.Put("/:id", requests.ValidateCompanyUserUpdate(), userHandler.Update())
	userGroup.Delete("/:id", userHandler.Remove())

	staffGroup := r.Group("/company")
	staffGroup.Get("/:id/users", userHandler.FetchAllByCompany())

	companyGroup := r.Group("/companies")
	companyGroup.Get("", handler.FetchAll())
	companyGroup.Get("/:id", handler.Fetch())
	companyGroup.Put("/:id", requests.ValidateCompanyUpdate(), handler.Update())
	companyGroup.Delete("/:id", handler.Remove())

	vehicleGroup := r.Group("/vehicles")
	vehicleGroup.Get("", vehicleHandler.FetchAll())
	vehicleGroup.Get("/company/:id", vehicleHandler.FetchAllByCompany())
	vehicleGroup.Get("/:id", vehicleHandler.Fetch())
	vehicleGroup.Post("/company/:id", requests.ValidateVehicle(), vehicleHandler.Create())
	vehicleGroup.Post("/:id/add-image", requests.ValidateVehicleImage(), vehicleHandler.AddImage())
	vehicleGroup.Put("/:id", requests.ValidateVehicleUpdate(), vehicleHandler.Update())
	vehicleGroup.Put("/:id/update-image", requests.ValidateVehicleImageUpdate(), vehicleHandler.UpdateImage())
	vehicleGroup.Delete("/:id", vehicleHandler.Remove())
	vehicleGroup.Delete("/:id/delete-image", vehicleHandler.RemoveImage())

	routeGroup := r.Group("/routes")
	routeGroup.Get("", routeHandler.FetchAll())
	routeGroup.Get("/distinct", routeHandler.FetchAllDistinct())
	routeGroup.Get("/company/:id", routeHandler.FetchAllByCompany())
	routeGroup.Get("/:id", routeHandler.Fetch())
	routeGroup.Post("/company/:id", requests.ValidateRoute(), routeHandler.Create())
	routeGroup.Post("/:id/add-stop", requests.ValidateRouteStop(), routeHandler.AddRouteStop())
	routeGroup.Put("/:id", requests.ValidateRouteUpdate(), routeHandler.Update())
	routeGroup.Put("/:id/update-stop", requests.ValidateRouteStop(), routeHandler.UpdateStop())
	routeGroup.Delete("/:id", routeHandler.Remove())
	routeGroup.Delete("/:id/delete-stop", routeHandler.RemoveStop())

	tripGroup := r.Group("/trips")
	tripGroup.Get("", tripHandler.FetchAll())
	tripGroup.Get("/popular", tripHandler.FetchAllPopular())
	tripGroup.Get("/company/:id", tripHandler.FetchAllByCompany())
	tripGroup.Get("/driver/:id", tripHandler.FetchAllByDriver())
	tripGroup.Get("/:id", tripHandler.Fetch())
	tripGroup.Post("/company/:id", requests.ValidateTrip(), tripHandler.Create())
	tripGroup.Put("/:id", requests.ValidateTripUpdate(), tripHandler.Update())
	tripGroup.Put("/:id/update-status", tripHandler.UpdateStatus())
	tripGroup.Put("/:id/update-schedule", tripHandler.UpdateSchedule())
	tripGroup.Put("/:id/update-inspection", tripHandler.UpdateInspection())
	tripGroup.Delete("/:id", tripHandler.Remove())

}
