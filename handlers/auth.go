package handlers

import "github.com/gofiber/fiber/v2"

func HandleRegister(c *fiber.Ctx) error {
	return c.JSON(ResponseHTTP{
		Success: true,
		Data:    nil,
		Message: "Account registered successfully",
	})
}
