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
	auth.Post("/login", handlers.Login)
	auth.Post("/verify_otp", handlers.VerifyOTP)
	auth.Post("/resend_otp", handlers.ResendOTP)
	auth.Post("/me", handlers.GetCurrentUser)

	return app
}
