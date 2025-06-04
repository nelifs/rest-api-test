package router

import (
	"github.com/gofiber/fiber/v2"
	"rest-api-test/controller/auth"
)

func AuthorizationRouter(api fiber.Router) {
	api.Get("/register", auth.Register)
}
