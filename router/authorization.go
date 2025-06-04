package router

import (
	"awesomeProject/controller/auth"
	"github.com/gofiber/fiber/v2"
)

func AuthorizationRouter(api fiber.Router) {
	api.Get("/register", auth.Register)
}
