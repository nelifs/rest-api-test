package handler

import (
	"awesomeProject/constant"
	"awesomeProject/dto"
	"errors"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	msg := "Internal server error occurred"

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
		msg = e.Message
	}

	if err != nil {
		return c.Status(code).JSON(dto.Error{Success: false, Code: constant.Internal, Message: msg})
	}

	return nil
}
