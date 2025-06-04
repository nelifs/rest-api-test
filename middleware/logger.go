package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func LoggerMiddleware(c *fiber.Ctx) error {
	log.Info("Handling request ", c.BaseURL())

	return c.Next()
}
