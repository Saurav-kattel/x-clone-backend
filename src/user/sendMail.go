package user

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendMail(otp, userEmail string) error {
	from := os.Getenv("EMAIL")
	password := os.Getenv("APP_PASSWORD")

	if from == "" || password == "" {
		return fmt.Errorf("unable to load necessary environment variables")
	}

	to := []string{userEmail}
	subject := "Your OTP for Password Reset"

	// SMTP configuration
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := fmt.Sprintf("From: X-clone %s\r\n"+
		"Subject: %s\r\n"+
		"Content-Type: text/html; charset=UTF-8\r\n\r\n"+
		"<h1>OTP</h1>"+
		"<p>Your OTP for changing the password is: <em><strong>%s</strong></em>.</p>"+
		"<p>Please do not share this OTP with others.</p>", from, subject, otp)

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(message))
	if err != nil {
		return err
	}
	return nil
}
