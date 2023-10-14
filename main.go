package main

import (
	"fmt"
	"net/http"
	"io/ioutil"

	"NewProjectSearchApp/pkg/mail"
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

func main() {
	jobDetails, err := getJobDetails()
	if err != nil {
		fmt.Println("Error fetching job details:", err)
		return
	}

	emailSubject := "Job Details"
	emailBody := "Here are the job details:\n" + jobDetails

	err := mail.SendEmail("test", "test")
	if err != nil {
		fmt.Println("Error sending email:", err)
		return
	}

	fmt.Println("OK")
}
