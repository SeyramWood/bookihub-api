package middlewares

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slices"

	"github.com/SeyramWood/utils/jwt"
)

func ValidateRefreshToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Get("X-Refresh-Token") == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": false, "message": "X-Refresh-Token header missing"})
		}
		return c.Next()
	}
}
func ValidateOAuthToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Get("X-OAuth-Token") == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": false, "message": "X-OAuth-Token header missing"})
		}
		return c.Next()
	}
}

func Authenticate(jwt *jwt.JWT) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if isService(c.Request().Host()) && strings.SplitN(c.Get("User-Agent"), "/", 2)[0] == "Go-http-client" {
			return c.Next()
		}
		token, err := bearerToken(c.Get("Authorization"))
		if err != nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": false, "message": "Forbidden"})
		}
		claims, err := jwt.ValidateToken(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": false, "message": "Unauthorized"})
		}
		c.Locals("user", claims["session"])
		return c.Next()
	}
}

// bearerToken extracts the content from the header, striping the Bearer prefix
func bearerToken(rawToken string) (string, error) {
	pieces := strings.SplitN(rawToken, " ", 2)
	if len(pieces) < 2 {
		return "", fmt.Errorf("token with incorrect bearer format")
	}
	token := strings.TrimSpace(pieces[1])
	return token, nil
}
func isService(host []byte) bool {
	services := []string{
		"broker-service",
		"producer-service",
		"listener-service",
		"logger-service",
		"notification-service",
		"authentication-service",
		"user-service",
		"vehicle-service",
		"booking-service",
		"payment-service",
		"review-service",
		"report-service",
		"support-service",
		"devbroker-service",
		"devproducer-service",
		"devlistener-service",
		"devlogger-service",
		"devnotification-service",
		"devauthentication-service",
		"devuser-service",
		"devvehicle-service",
		"devbooking-service",
		"devpayment-service",
		"devreview-service",
		"devreport-service",
		"devsupport-service",
	}
	return slices.Contains(services, string(host))
}
