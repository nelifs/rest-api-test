package main

import (
	"awesomeProject/db"
	"awesomeProject/handler"
	"awesomeProject/middleware"
	"awesomeProject/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	port := ":8080"
	app := fiber.New(fiber.Config{ErrorHandler: handler.ErrorHandler})
	app.Use(middleware.LoggerMiddleware)
	app.Route("/auth", router.AuthorizationRouter)
	app.Route("/users", router.UserRouter)
	db.Init()

	log.Info("Starting server ", port)
	log.Fatal(app.Listen(port))
}
