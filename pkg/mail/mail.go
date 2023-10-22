package mail

import (
	"net/smtp"
	"strings"
	"database/sql"
    _ "github.com/lib/pq"
	"NewProjectSearchApp/pkg/job"
)

func SendEmail(body string, db *sql.DB) error {
	fromEmail, toEmail, subject, smtpUsername, smtpPassword, smtpHost, smtpPort := GetMailInfoFromDB(db)

	headers := map[string]string{
		"From": fromEmail,
		"To": toEmail,
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

	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, fromEmail, []string{toEmail}, []byte(message))
	return err
}

// 本文作成
func BuildEmailBody(dataSources []struct {
	JobInfoSlice *[]job.JobInfo
	ErrPtr *error
	FetchFunc func(*sql.DB) ([]job.JobInfo, error)
	Title string }) string {

    var emailBody strings.Builder
    for _, source := range dataSources {
        emailBody.WriteString("■" + source.Title + "\n\n")
        for _, jobInfo := range *source.JobInfoSlice {
            emailBody.WriteString("　" + jobInfo.Name + "\n")
            emailBody.WriteString("　" + jobInfo.URL + "\n\n")
        }
        emailBody.WriteString("\n")
    }
	
    return emailBody.String()
}

func GetMailInfoFromDB(db *sql.DB) (string, string, string, string, string, string, string) {
    var fromEmail, toEmail, subject, smtpUsername, smtpPassword, smtpHost, smtpPort string
    db.QueryRow("SELECT from_email, to_email, subject, smtp_username, smtp_password, smtp_host, smtp_port FROM mail").Scan(&fromEmail, &toEmail, &subject, &smtpUsername, &smtpPassword, &smtpHost, &smtpPort)

    return fromEmail, toEmail, subject, smtpUsername, smtpPassword, smtpHost, smtpPort
}
