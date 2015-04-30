package sms

import (
	twilio "github.com/carlosdp/twiliogo"
	"github.com/sirsean/textback/config"
)

var Client twilio.Client

func Connect() {
	Client = twilio.NewClient(config.Get().Twilio.AccountSid, config.Get().Twilio.AuthToken)
}

func Send(to, body string) error {
	_, err := twilio.NewMessage(Client, config.Get().Twilio.From, to, twilio.Body(body))
	return err
}
