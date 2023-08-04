package routes

import (
	"auth/handler"

	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	app := fiber.New()
	api := app.Group("/api")
	auth := api.Group("/auth")

	auth.Post("/register", handler.Register)
	auth.Post("/login", handler.Login)
	auth.Post("/verify_otp", handler.VerifyOTP)
	auth.Post("/resend_otp", handler.ResendOTP)
	auth.Post("/me", handler.GetCurrentUser)

	return app
}
