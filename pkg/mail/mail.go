package mail

import (
	"net/smtp"
)

const (
	smtpHost     = "smtp.gmail.com" // SMTPサーバーホスト
	smtpPort     = "587"              // SMTPポート
	smtpUsername = "thierry.daniel.henry0302@gmail.com"    // SMTPユーザー名
	smtpPassword = "sggn qdma tnia mqqw"    // SMTPパスワード

	fromEmail    = "myProject.com" // 送信元メールアドレス
	toEmail      = "thierry.daniel.henry0302@gmail.com"  // 宛先メールアドレス
)

func SendEmail(subject, body string) error {
	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)
	msg := []byte(
		"To: " + toEmail + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + body
	)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, fromEmail, []string{toEmail}, msg)
	return err
}
