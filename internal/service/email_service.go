package service

import (
	"context"
	"log"
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
	// Mailgun configuration (replace with environment variables or config)
	domain := "your-domain.com"
	apiKey := "your-mailgun-api-key"
	sender := "sender@your-domain.com"

	mg := mailgun.NewMailgun(domain, apiKey)
	message := mg.NewMessage(sender, req.Subject, req.Body, req.Recipient)

	// Set a timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, id, err := mg.Send(ctx, message)
	if err != nil {
		return err
	}

	log.Printf("Email sent to %s. ID: %s\n", req.Recipient, id)
	return nil
}
