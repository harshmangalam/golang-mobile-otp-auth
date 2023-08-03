package handlers

import "github.com/gofiber/fiber/v2"

func Register(c *fiber.Ctx) error {
	return c.JSON(ResponseHTTP{
		Success: true,
		Data:    nil,
		Message: "Account registered successfully",
	})
}

func Login(c *fiber.Ctx) error {
	return c.JSON(ResponseHTTP{
		Success: true,
		Data:    nil,
		Message: "Account registered successfully",
	})
}
func VerifyOTP(c *fiber.Ctx) error {
	return c.JSON(ResponseHTTP{
		Success: true,
		Data:    nil,
		Message: "Account registered successfully",
	})
}

func ResendOTP(c *fiber.Ctx) error {
	return c.JSON(ResponseHTTP{
		Success: true,
		Data:    nil,
		Message: "Account registered successfully",
	})
}

func GetCurrentUser(c *fiber.Ctx) error {
	return c.JSON(ResponseHTTP{
		Success: true,
		Data:    nil,
		Message: "Account registered successfully",
	})
}
