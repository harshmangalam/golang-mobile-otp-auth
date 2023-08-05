package middleware

import (
	"auth/config"
	"auth/schema"
	"auth/util"

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
	user, err := util.FindUserById(userId)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(schema.ResponseHTTP{
			Success: false,
			Message: "User not exists",
			Data:    nil,
		})
	}
	c.Locals("user", user)
	return c.Next()
}
func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(schema.ResponseHTTP{Success: false, Message: "Missing or malformed JWT", Data: nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(schema.ResponseHTTP{Success: false, Message: "Invalid or expired JWT", Data: nil})
}
