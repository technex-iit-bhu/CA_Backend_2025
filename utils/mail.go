package utils

import (
	"CA_Portal_backend/config"
	"fmt"
	smtp "net/smtp"
)

func RecoveryMail(to, token string) error {
	password := config.Config("SMTP_PASSWORD")
	from := config.Config("SMTP_EMAIL")
	smtpHost := config.Config("SMTP_HOST")
	smtpPort := config.Config("SMTP_PORT")
	frontendURL := config.Config("FRONTEND_URL")

	if frontendURL == "" {
		frontendURL = "http://localhost:3000"
	}

	subject := "Password Recovery - Technex'26 CA Portal"
	body := fmt.Sprintf(`Hello,

You have requested to reset your password for Technex'26 Campus Ambassador Portal.

Click the link below to reset your password:
%s/password/reset-password?token=%s

This link will expire in 10 minutes.

If you did not request this, please ignore this email.

Best regards,
Technex'26 Team`, frontendURL, token)

	// NOW USE THE BODY IN THE MESSAGE!
	message := []byte(fmt.Sprintf("Subject: %s\r\n"+
		"MIME-version: 1.0;\r\n"+
		"Content-Type: text/plain; charset=\"UTF-8\";\r\n"+
		"\r\n"+
		"%s\r\n", subject, body))

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)
	if err != nil {
		fmt.Printf("Error sending email: %v\n", err)
		return err
	}
	fmt.Println("Email Sent Successfully")
	return nil
}