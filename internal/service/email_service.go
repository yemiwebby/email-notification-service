package service

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

// EmailRequest represents the payload for sending an email
type EmailRequest struct {
	Recipient string `json:"recipient"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
}

// SendEmail sends an email using Mailgun
func SendEmail(req EmailRequest) error {
	log.Printf("Processing email for: %s\n", req.Recipient)

	simulationMode := os.Getenv("EMAIL_SIMULATION_MODE")
	if simulationMode == "true" {
		log.Printf("[SIMULATION] Email sent to %s with subject: %s\n", req.Recipient, req.Subject)
		return nil
	}

	domain := os.Getenv("MAILGUN_DOMAIN")
	apiKey := os.Getenv("MAILGUN_API_KEY")
	sender := os.Getenv("MAILGUN_SENDER")

	

	mg := mailgun.NewMailgun(domain, apiKey)
	message := mailgun.NewMessage(sender, req.Subject, req.Body, req.Recipient)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, id, err := mg.Send(ctx, message)
	if err != nil {
		log.Printf("Failed to send email to %s: %v\n", req.Recipient, err)
		return err
	}

	log.Printf("Email sent to %s. ID: %s\n", req.Recipient, id)
	return nil
}
