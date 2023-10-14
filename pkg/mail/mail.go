package mail

import (
	"net/smtp"
	"NewProjectSearchApp/constants"
)

func SendEmail(subject, body string) error {
	auth := smtp.PlainAuth("", constants.SmtpUsername, constants.SmtpPassword, constants.SmtpHost)
	msg := []byte(body)

	err := smtp.SendMail(constants.SmtpHost+":"+constants.SmtpPort, auth, constants.FromEmail, []string{constants.ToEmail}, msg)
	return err
}
