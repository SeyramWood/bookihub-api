package api

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application"
	"github.com/SeyramWood/app/application/auth"
	requeststructs "github.com/SeyramWood/app/domain/request_structs"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/ent"
	"github.com/SeyramWood/utils/jwt"
)

type authHandler struct {
	service  gateways.AuthService
	producer gateways.EventProducer
	jwt      *jwt.JWT
}

func NewAuthHandler(db *database.Adapter, cacheSrv gateways.CacheService, jwt *jwt.JWT, producer gateways.EventProducer) *authHandler {
	return &authHandler{
		service:  auth.NewService(auth.NewRepository(db), cacheSrv, jwt, producer),
		producer: producer,
		jwt:      jwt,
	}
}

func (h *authHandler) Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.UserLoginRequest)
		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
		}
		result, err := h.service.Login(request)
		if err != nil {
			if ent.IsNotFound(err) || errors.Is(err, auth.ErrBadRequest) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(auth.ErrBadRequest))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.SuccessResponse(result))
	}
}
func (h *authHandler) GetSession() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(presenters.SuccessResponse(c.Locals("user")))
	}
}
func (h *authHandler) RefreshToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := h.service.RefreshToken(c.Get("X-Refresh-Token"))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.SuccessResponse(result))
	}
}

func (h *authHandler) SendPasswordResetCode() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.UsernameRequest)
		if c.BodyParser(request) != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(auth.ErrBadRequest))
		}
		_, err := h.service.SendPasswordResetCode(request)
		if err != nil {
			if ent.IsNotFound(err) || errors.Is(err, auth.ErrBadRequest) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(auth.ErrBadRequest))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.MessageResponse("top sent"))
	}
}
func (h *authHandler) VerifyOTP() fiber.Handler {
	return func(c *fiber.Ctx) error {
		result := h.service.VerifyOTP(c.Params("otp"))
		if !result {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(auth.ErrBadRequest))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.MessageResponse("verified"))
	}
}
func (h *authHandler) UpdatePassword() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.UpdatePasswordRequest)
		if c.BodyParser(request) != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(auth.ErrBadRequest))
		}
		_, err := h.service.UpdatePassword(application.FormatSessionID(c.Locals("user")), request)
		if err != nil {
			if ent.IsNotFound(err) || errors.Is(err, auth.ErrBadRequest) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.MessageResponse("password updated"))
	}
}

func (h *authHandler) ResetPassword() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.ResetPasswordRequest)
		if c.BodyParser(request) != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(auth.ErrBadRequest))
		}
		_, err := h.service.ResetPassword(request)
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(auth.ErrBadRequest))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.MessageResponse("password updated"))
	}
}
