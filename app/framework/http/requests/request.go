package requests

import (
	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/adapters/presenters"
	requeststructs "github.com/SeyramWood/app/domain/request_structs"
	"github.com/SeyramWood/utils/validator"
)

func ValidateAuthUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.UserLoginRequest)
		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		if er := validator.Validate(request); er != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(presenters.UnprocessableEntityResponse(er))
		}
		return c.Next()
	}
}
func ValidateUsername() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.UsernameRequest)
		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		if er := validator.Validate(request); er != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(presenters.UnprocessableEntityResponse(er))
		}
		return c.Next()
	}
}
func ValidateBookibusUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.BookibusUserRequest)
		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		if er := validator.Validate(request); er != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(presenters.UnprocessableEntityResponse(er))
		}
		return c.Next()
	}
}
func ValidateBookibusUserUpdate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.BookibusUserUpdateRequest)
		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		if er := validator.Validate(request); er != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(presenters.UnprocessableEntityResponse(er))
		}
		return c.Next()
	}
}
func ValidateCompanyUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.CompanyUserRequest)
		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		if er := validator.Validate(request); er != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(presenters.UnprocessableEntityResponse(er))
		}
		return c.Next()
	}
}
func ValidateCompanyUserUpdate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.CompanyUserUpdateRequest)
		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		if er := validator.Validate(request); er != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(presenters.UnprocessableEntityResponse(er))
		}
		return c.Next()
	}
}
func ValidateCustomer() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.CustomerRequest)
		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		if er := validator.Validate(request); er != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(presenters.UnprocessableEntityResponse(er))
		}
		return c.Next()
	}
}
func ValidateCustomerUpdate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.CustomerUpdateRequest)
		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		if er := validator.Validate(request); er != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(presenters.UnprocessableEntityResponse(er))
		}
		return c.Next()
	}
}
func ValidateCompany() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.CompanyRequest)
		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		if er := validator.Validate(request); er != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(presenters.UnprocessableEntityResponse(er))
		}
		return c.Next()
	}
}
func ValidateCompanyUpdate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.CompanyUpdateRequest)
		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		if er := validator.Validate(request); er != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(presenters.UnprocessableEntityResponse(er))
		}
		return c.Next()
	}
}
func ValidateVehicle() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.VehicleRequest)
		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		form, err := c.MultipartForm()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		request.Image = form.File["image"]
		if er := validator.Validate(request); er != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(presenters.UnprocessableEntityResponse(er))
		}
		return c.Next()
	}
}
func ValidateVehicleUpdate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.VehicleUpdateRequest)
		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		if er := validator.Validate(request); er != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(presenters.UnprocessableEntityResponse(er))
		}
		return c.Next()
	}
}
func ValidateVehicleImage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.VehicleImageRequest)
		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		form, err := c.MultipartForm()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		request.Image = form.File["image"]
		if er := validator.Validate(request); er != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(presenters.UnprocessableEntityResponse(er))
		}
		return c.Next()
	}
}
func ValidateVehicleImageUpdate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.VehicleImageUpdateRequest)
		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		file, err := c.FormFile("image")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		request.Image = file
		if er := validator.Validate(request); er != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(presenters.UnprocessableEntityResponse(er))
		}
		return c.Next()
	}
}
func ValidateRoute() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.RouteRequest)
		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		if er := validator.Validate(request); er != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(presenters.UnprocessableEntityResponse(er))
		}
		return c.Next()
	}
}
func ValidateRouteUpdate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.RouteUpdateRequest)
		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		if er := validator.Validate(request); er != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(presenters.UnprocessableEntityResponse(er))
		}
		return c.Next()
	}
}
func ValidateRouteStop() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.RouteStopRequest)
		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		if er := validator.Validate(request); er != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(presenters.UnprocessableEntityResponse(er))
		}
		return c.Next()
	}
}
func ValidateTrip() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.TripRequest)
		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		if er := validator.Validate(request); er != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(presenters.UnprocessableEntityResponse(er))
		}
		return c.Next()
	}
}
func ValidateTripBoardingPoint() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.TripNewBoardingPoint)
		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		if er := validator.Validate(request); er != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(presenters.UnprocessableEntityResponse(er))
		}
		return c.Next()
	}
}
func ValidateTripUpdate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.TripUpdateRequest)
		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		if er := validator.Validate(request); er != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(presenters.UnprocessableEntityResponse(er))
		}
		return c.Next()
	}
}

func ValidateBooking() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.BookingRequest)
		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		if er := validator.Validate(request); er != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(presenters.UnprocessableEntityResponse(er))
		}
		return c.Next()
	}
}
func ValidateBookingUpdate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.BookingUpdateRequest)
		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		if er := validator.Validate(request); er != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(presenters.UnprocessableEntityResponse(er))
		}
		return c.Next()
	}
}
func ValidateBookingCancel() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.BookingCancelRequest)
		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		if er := validator.Validate(request); er != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(presenters.UnprocessableEntityResponse(er))
		}
		return c.Next()
	}
}
