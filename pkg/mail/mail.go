package mail

import (
	"net/smtp"
	"strings"
	"NewProjectSearchApp/constants"
	"NewProjectSearchApp/pkg/job"
)

func SendEmail(body string) error {

	subject := "新規案件リスト"
		
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

// 本文作成
func BuildEmailBody(dataSources []struct {
	JobInfoSlice *[]job.JobInfo
	ErrPtr *error
	FetchFunc func() ([]job.JobInfo, error)
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
