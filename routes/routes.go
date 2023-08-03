package routes

import (
	"auth/handlers"

	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	app := fiber.New()
	api := app.Group("/api")
	auth := api.Group("/auth")

	auth.Post("/register", handlers.HandleRegister)
	app.Post("/login")

	return app
}
