package email

import (
	"net/smtp"
	"unsafemail/config"
)

func Send(to, subject, body string) error {
	message := []byte(
		"Subject: " + subject + "\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
			"\r\n" +
			body +
			"\r\n",
	)
	auth := smtp.PlainAuth("", config.From, config.Password, config.SmtpHost)
	return smtp.SendMail(config.SmtpHost+":"+config.SmtpPort, auth, config.From, []string{to}, message)
}
