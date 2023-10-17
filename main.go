package main

import (
	"fmt"
	"strings"
	"sync"

	"NewProjectSearchApp/pkg/mail"
	"NewProjectSearchApp/pkg/job"
)

func main() {
	var wg sync.WaitGroup

	var freelancestartJobInfoSlice, levtechJobInfoSlice []job.JobInfo
    var freelancestartErr, levtechErr error

	// 並列でデータを取得
	wg.Add(2)
	go fetchDataInParallel(&wg, &freelancestartJobInfoSlice, &freelancestartErr, job.GetFreelanceStartDetails)
    go fetchDataInParallel(&wg, &levtechJobInfoSlice, &levtechErr, job.GetLevtechDetails)
    wg.Wait()
	
	// 取得したデータのエラー判定
	if checkAndPrintError(freelancestartErr, "Error fetching freelance job details:") {
		return
	}
	if checkAndPrintError(levtechErr, "Error fetching Levtech job details:") {
		return
	}

	// 件名と本文作成
	emailSubject := "新規案件リスト"
	jobInfoSlices := map[string][]job.JobInfo{
		"フリーランススタート": freelancestartJobInfoSlice,
		"レバテック": levtechJobInfoSlice,
	}
	emailBody := buildEmailBody(jobInfoSlices)

	// メール送信
	sendErr  := mail.SendEmail(emailSubject, emailBody)
	if sendErr  != nil {
		fmt.Println("Error sending email:", sendErr )
		return
	}
}

// 並列処理
func fetchDataInParallel(wg *sync.WaitGroup, resultSlice *[]job.JobInfo, errPtr *error, fetchFunc func() ([]job.JobInfo, error)) {
    defer wg.Done()

    dataSlice, err := fetchFunc()
    if err != nil {
        *errPtr = err
        return
    }

    *resultSlice = append(*resultSlice, dataSlice...)
}

// エラーチェック
func checkAndPrintError(err error, message string) bool {
    if err != nil {
        fmt.Println(message, err)
        return true
    }
    return false
}

// 本文作成
func buildEmailBody(jobInfoSlices map[string][]job.JobInfo) string {
    var emailBody strings.Builder
    for category, jobInfoSlice := range jobInfoSlices {
        emailBody.WriteString("■" + category + "\n\n")
        for _, jobInfo := range jobInfoSlice {
            emailBody.WriteString("　" + jobInfo.Name + "\n")
            emailBody.WriteString("　" + jobInfo.URL + "\n\n")
        }
		emailBody.WriteString("\n")
    }
    return emailBody.String()
}
