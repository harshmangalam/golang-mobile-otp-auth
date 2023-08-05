package util

import (
	"auth/config"
	"fmt"
	"log"

	"github.com/twilio/twilio-go"

	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendOTP(to string, otp string) error {
	client := twilio.NewRestClient()

	params := &openapi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(config.Config("TWILIO_PHONE_NUMBER"))

	msg := fmt.Sprintf("Your OTP is %s", otp)
	params.SetBody(msg)

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	log.Println("SMS sent successfully!")

	return nil
}
