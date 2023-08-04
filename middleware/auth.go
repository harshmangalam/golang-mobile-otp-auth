package middleware

import (
	"auth/config"
	"auth/utils"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// Protected protect routes
func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:     jwtware.SigningKey{Key: []byte(config.Config("SECRET"))},
		ErrorHandler:   jwtError,
		SuccessHandler: jwtSuccess,
		ContextKey:     "payload",
	})
}

func jwtSuccess(c *fiber.Ctx) error {
	payload := c.Locals("payload").(*jwt.Token)
	claims := payload.Claims.(jwt.MapClaims)
	userId := claims["userId"].(string)
	user, err := utils.FindUserById(userId)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "User not exists",
			"data":    nil,
		})
	}
	c.Locals("user", user)
	return c.Next()
}
func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"success": false, "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"success": false, "message": "Invalid or expired JWT", "data": nil})
}
