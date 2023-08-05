package util

import (
	"auth/config"

	"github.com/twilio/twilio-go"

	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendSMS(to string, msg string) (bool, error) {
	client := twilio.NewRestClient()

	params := &openapi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(config.Config("TWILIO_PHONE_NUMBER"))
	params.SetBody(msg)

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		return false, err
	}

	return true, nil
}
