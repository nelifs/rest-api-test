package auth

import (
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true})
}
