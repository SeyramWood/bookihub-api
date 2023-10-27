package web

import "github.com/gofiber/fiber/v2"

func Index() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	}
}
