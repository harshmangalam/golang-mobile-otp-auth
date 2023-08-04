package router

import (
	"auth/handler"
	"auth/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func New() *fiber.App {
	app := fiber.New()
	// compression middleware
	app.Use(compress.New())
	// cors middleware
	app.Use(cors.New())
	api := app.Group("/api")
	auth := api.Group("/auth")

	auth.Post("/register", handler.Register)
	auth.Post("/login", handler.Login)
	auth.Post("/verify_otp", handler.VerifyOTP)
	auth.Post("/resend_otp", handler.ResendOTP)
	auth.Get("/me", middleware.Protected(), handler.GetCurrentUser)

	return app
}
