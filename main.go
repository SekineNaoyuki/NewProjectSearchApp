package main

import (
	"fmt"
	"net/http"
	"net/smtp"
	"io/ioutil"
)

const (
	smtpHost     = "smtp.gmail.com" // SMTPサーバーホスト
	smtpPort     = "587"              // SMTPポート
	smtpUsername = "thierry.daniel.henry0302@gmail.com"    // SMTPユーザー名
	smtpPassword = "sggn qdma tnia mqqw"    // SMTPパスワード

	fromEmail    = "myProject.com" // 送信元メールアドレス
	toEmail      = "thierry.daniel.henry0302@gmail.com"  // 宛先メールアドレス
)

func getJobDetails() (string, error) {
	resp, err := http.Get("https://freelance.levtech.jp/project/skill-10/") // 案件の詳細URL
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func sendEmail(subject, body string) error {
	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)
	msg := []byte("To: " + toEmail + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, fromEmail, []string{toEmail}, msg)
	return err
}

func main() {
	jobDetails, err := getJobDetails()
	if err != nil {
		fmt.Println("Error fetching job details:", err)
		return
	}

	emailSubject := "Job Details"
	emailBody := "Here are the job details:\n" + jobDetails

	err = sendEmail(emailSubject, emailBody)
	if err != nil {
		fmt.Println("Error sending email:", err)
		return
	}

	fmt.Println("OK")
}