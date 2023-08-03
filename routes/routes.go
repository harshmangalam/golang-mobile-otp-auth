package routes

import (
	"auth/handlers"

	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	app := fiber.New()
	api := app.Group("/api")
	auth := api.Group("/auth")

	auth.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)
	app.Post("/verify_otp", handlers.VerifyOTP)
	app.Post("/resend_otp", handlers.ResendOTP)
	app.Post("/me", handlers.GetCurrentUser)

	return app
}
