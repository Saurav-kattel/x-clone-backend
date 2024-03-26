package user

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendMail(userEmail, message string) error {
	from := os.Getenv("EMAIL")
	password := os.Getenv("APP_PASSWORD")

	if from == "" || password == "" {
		return fmt.Errorf("unable to load necessary environment variables")
	}

	to := []string{userEmail}

	// SMTP configuration
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(message))
	if err != nil {
		return err
	}
	return nil
}
