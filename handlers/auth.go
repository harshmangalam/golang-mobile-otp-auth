package handlers

import (
	"auth/schema"
	"auth/utils"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	// request body data
	body := new(schema.RegisterBody)
	if err := c.BodyParser(body); err != nil {
		return c.JSON(ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: err.Error(),
		})
	}

	// validate duplicate mobile number

	user, err := utils.FindUserByPhone(body.Phone)

	if err != nil {
		return c.JSON(ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: err.Error(),
		})
	}

	if user != nil {
		return c.JSON(ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: "Phone number already in use",
		})
	}

	// create new user

	id, err := utils.InsertUser(body)
	if err != nil {
		return c.JSON(ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: err.Error(),
		})
	}

	return c.JSON(ResponseHTTP{
		Success: true,
		Data: fiber.Map{
			"id": id,
		},
		Message: "Account registered successfully",
	})
}

func Login(c *fiber.Ctx) error {
	// request body data
	body := new(schema.LoginSchema)
	if err := c.BodyParser(body); err != nil {
		return c.JSON(ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: err.Error(),
		})
	}
	// find phone in database
	user, err := utils.FindUserByPhone(body.Phone)

	if err != nil {
		return c.JSON(ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: err.Error(),
		})
	}

	if user == nil {
		return c.JSON(ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: "Phone number not exists",
		})
	}

	otp := utils.GenerateRandomNumber()

	// save otp in database
	utils.UpdateUser(user.ID, map[string]any{
		"otp": otp,
	})
	// send otp to user phone

	return c.JSON(ResponseHTTP{
		Success: true,
		Data:    user,
		Message: "Login successfully",
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
