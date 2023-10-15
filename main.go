package main

import (
	"fmt"
	"strings"

	"NewProjectSearchApp/pkg/mail"
	"NewProjectSearchApp/pkg/job"
)

func main() {
	jobInfoSlice, err := job.GetLevtechDetails()
	if err != nil {
		fmt.Println("Error fetching job details:", err)
		return
	}

	emailSubject := "新規案件リスト"

	var emailBody strings.Builder
	for _, jobInfo := range jobInfoSlice {
		emailBody.WriteString(jobInfo.Name + "\n")
		emailBody.WriteString(jobInfo.URL + "\n\n")
	}

	sendErr  := mail.SendEmail(emailSubject, emailBody.String())
	if sendErr  != nil {
		fmt.Println("Error sending email:", sendErr )
		return
	}

	fmt.Println("OK")
}
