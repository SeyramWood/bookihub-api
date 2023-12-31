package api

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/bookibus/app/adapters/gateways"
	"github.com/SeyramWood/bookibus/app/adapters/presenters"
	"github.com/SeyramWood/bookibus/app/application"
	"github.com/SeyramWood/bookibus/app/application/auth"
	requeststructs "github.com/SeyramWood/bookibus/app/domain/request_structs"
	"github.com/SeyramWood/bookibus/app/framework/database"
	"github.com/SeyramWood/bookibus/ent"
	"github.com/SeyramWood/bookibus/utils/jwt"
)

type authHandler struct {
	service  gateways.AuthService
	producer gateways.EventProducer
	jwt      *jwt.JWT
}

func NewAuthHandler(db *database.Adapter, cacheSrv gateways.CacheService, jwt *jwt.JWT, producer gateways.EventProducer, storage gateways.StorageService) *authHandler {
	return &authHandler{
		service:  auth.NewService(auth.NewRepository(db), cacheSrv, jwt, producer, storage),
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
func (h *authHandler) UpdateAvatar() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requeststructs.AvatarUpdateRequest)
		if c.BodyParser(request) != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(auth.ErrBadRequest))
		}
		avatarFile, err := c.FormFile("avatar")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		request.Avatar = avatarFile
		userID, _ := c.ParamsInt("id")
		avatar, err := h.service.UpdateAvatar(userID, request)

		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(err))
			}
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenters.SuccessResponse(avatar))
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
