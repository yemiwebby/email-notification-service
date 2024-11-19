package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	MailgunDomain string
	MailgunAPIKey string
	MailgunSender string
)

// LoadConfig loads the configuration from environment variables
func LoadConfig() {

	if err := godotenv.Load(".env"); err != nil {
        log.Fatalf("Error loading .env file")
    }

	MailgunDomain = os.Getenv("MAILGUN_DOMAIN")
	MailgunAPIKey = os.Getenv("MAILGUN_API_KEY")
	MailgunSender = os.Getenv("MAILGUN_SENDER")

	if MailgunDomain == "" || MailgunAPIKey == "" || MailgunSender == "" {
		log.Fatal("Missing required environment variables for Mailgun configuration")
	}
}
