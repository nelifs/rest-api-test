package router

import (
	"github.com/gofiber/fiber/v2"
	"rest-api-test/controller/user"
	"rest-api-test/middleware"
)

func UserRouter(api fiber.Router) {
	api.Use(middleware.AuthorizationMiddleware)
	api.Get("/", user.Get)
}
