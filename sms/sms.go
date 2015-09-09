package sms

import (
	twilio "github.com/carlosdp/twiliogo"
	"log"
	"os"
)

var Client twilio.Client

var twilioAccountSid string = os.Getenv("TWILIO_ACCOUNT_SID")
var twilioAuthToken string = os.Getenv("TWILIO_AUTH_TOKEN")
var from string = os.Getenv("TWILIO_FROM")

func init() {
	if twilioAccountSid == "" {
		log.Fatal("TWILIO_ACCOUNT_SID is required")
	}
	if twilioAuthToken == "" {
		log.Fatal("TWILIO_AUTH_TOKEN is required")
	}
	if from == "" {
		log.Fatal("FROM is required")
	}
}

func Connect() {
	Client = twilio.NewClient(twilioAccountSid, twilioAuthToken)
}

func Send(to, body string) error {
	_, err := twilio.NewMessage(Client, from, to, twilio.Body(body))
	return err
}
