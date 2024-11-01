package utils

import (
	"CA_Backend/config"
	"fmt"
	smtp "net/smtp"
)

func RecoveryMail(to, token string) error {
	password := config.Config("SMTP_PASSWORD")
	from := config.Config("SMTP_EMAIL")
	smtpHost := config.Config("SMTP_HOST")
	smtpPort := config.Config("SMTP_PORT")
	message := []byte(fmt.Sprintf("Subject: Recovery\r\n\r\nlink :  https://[FRONTEND_URL]/recovery?token=%s\r\n", token))
	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)
	if err != nil {
		return err
	}
	fmt.Println("Email Sent Successfully")
	return nil
}
