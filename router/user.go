package router

import (
	"awesomeProject/controller/user"
	"awesomeProject/middleware"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(api fiber.Router) {
	api.Use(middleware.AuthorizationMiddleware)
	api.Get("/", user.Get)
}
