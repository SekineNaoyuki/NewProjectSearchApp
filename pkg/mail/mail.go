package mail

import (
	"net/smtp"
	"NewProjectSearchApp/constants"
)

func SendEmail(subject, body string) error {

	headers := map[string]string{
		"From": constants.FromEmail,
		"To": constants.ToEmail,
		"Subject": subject,
		"MIME-Version": "1.0",
		"Content-Type": "text/plain; charset=\"utf-8\"",
	}

	// ヘッダーを連結
	var msgHeaders string
	for key, value := range headers {
		msgHeaders += key + ": " + value + "\r\n"
	}

	// メール本文とヘッダーを結合
	message := msgHeaders + "\r\n" + body

	auth := smtp.PlainAuth("", constants.SmtpUsername, constants.SmtpPassword, constants.SmtpHost)

	err := smtp.SendMail(constants.SmtpHost+":"+constants.SmtpPort, auth, constants.FromEmail, []string{constants.ToEmail}, []byte(message))
	return err
}
