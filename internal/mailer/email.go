package mailer

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendWelcomeEmail(toEmail, prenom string) error {
	from := os.Getenv("EMAIL_USER")
	password := os.Getenv("EMAIL_PASS")
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	subject := "Bienvenue " + prenom + " !"
	body := fmt.Sprintf("Bonjour %s,\n\nMerci pour votre inscription à ShopEase !", prenom)

	message := []byte("Subject: " + subject + "\r\n" +
		"From: " + from + "\r\n" +
		"To: " + toEmail + "\r\n" +
		"\r\n" +
		body + "\r\n")

	auth := smtp.PlainAuth("", from, password, smtpHost)

	return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{toEmail}, message)
}
