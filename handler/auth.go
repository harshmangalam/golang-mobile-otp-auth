package handler

import (
	"auth/model"
	"auth/schema"
	"auth/util"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	// request body data
	body := new(schema.RegisterBody)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schema.ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: err.Error(),
		})
	}

	// validate duplicate mobile number

	user, err := util.FindUserByPhone(body.Phone)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(schema.ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: err.Error(),
		})
	}

	if user != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schema.ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: "Phone number already in use",
		})
	}

	// create new user

	id, err := util.InsertUser(body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(schema.ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(schema.ResponseHTTP{
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
		return c.Status(fiber.StatusBadRequest).JSON(schema.ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: err.Error(),
		})
	}
	// find phone in database
	user, err := util.FindUserByPhone(body.Phone)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(schema.ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: err.Error(),
		})
	}

	if user == nil {
		return c.Status(fiber.StatusBadRequest).JSON(schema.ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: "Phone number not exists",
		})
	}

	otp := util.GenerateRandomNumber()

	// save otp in database
	util.UpdateUser(user.ID, map[string]any{
		"otp": otp,
	})
	// send otp to user phone

	err = util.SendOTP(user.Phone, otp)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(schema.ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(schema.ResponseHTTP{
		Success: true,
		Data:    nil,
		Message: "Otp sent to registered mobile number",
	})

}
func VerifyOTP(c *fiber.Ctx) error {
	// request body data
	body := new(schema.VerifyOTPSchema)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schema.ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: err.Error(),
		})
	}

	// find phone in database
	user, err := util.FindUserByPhone(body.Phone)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(schema.ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: err.Error(),
		})
	}

	if user == nil {
		return c.Status(fiber.StatusBadRequest).JSON(schema.ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: "Phone number not exists",
		})
	}

	if user.Otp != body.Otp {
		return c.Status(fiber.StatusBadRequest).JSON(schema.ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: "Incorrect Otp",
		})
	}

	// generate jwt token
	token, err := util.GenerateJWT(user.ID.Hex())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(schema.ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: err.Error(),
		})
	}

	// remove old otp from db
	util.UpdateUser(user.ID, map[string]any{
		"otp": "",
	})

	return c.Status(fiber.StatusCreated).JSON(schema.ResponseHTTP{
		Success: true,
		Data: fiber.Map{
			"token": "Bearer " + token,
		},
		Message: "Account login successfully",
	})
}

func ResendOTP(c *fiber.Ctx) error {
	// request body data
	body := new(schema.VerifyOTPSchema)
	if err := c.Status(fiber.StatusBadRequest).BodyParser(body); err != nil {
		return c.JSON(schema.ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: err.Error(),
		})
	}

	// find phone in database
	user, err := util.FindUserByPhone(body.Phone)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(schema.ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: err.Error(),
		})
	}

	if user == nil {
		return c.Status(fiber.StatusBadRequest).JSON(schema.ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: "Phone number not exists",
		})
	}

	otp := util.GenerateRandomNumber()

	// save otp in database
	util.UpdateUser(user.ID, map[string]any{
		"otp": otp,
	})
	// send otp to user phone

	err = util.SendOTP(user.Phone, otp)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(schema.ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(schema.ResponseHTTP{
		Success: true,
		Data:    nil,
		Message: "Sent otp to registered mobile number",
	})
}

func GetCurrentUser(c *fiber.Ctx) error {
	user := c.Locals("user").(*model.User)
	user.Otp = ""
	return c.Status(fiber.StatusOK).JSON(schema.ResponseHTTP{
		Success: true,
		Data:    user,
		Message: "Get current user",
	})
}
