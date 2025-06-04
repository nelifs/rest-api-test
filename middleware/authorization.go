package middleware

import (
	"awesomeProject/constant"
	"awesomeProject/dto"
	"awesomeProject/util"
	"github.com/gofiber/fiber/v2"
)

func AuthorizationMiddleware(c *fiber.Ctx) error {
	t := c.Get("Authorization")

	if _, err := util.Validate(t); err == nil {
		return c.Next()
	} else {
		return c.Status(403).JSON(dto.Error{Success: false, Code: constant.Unauthorized, Message: "You need to authorize first"})
	}
}
