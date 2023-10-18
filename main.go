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

	var freelancestartJobInfoSlice []job.JobInfo
	var levtechJobInfoSlice []job.JobInfo
	var akkodisJobInfoSlice []job.JobInfo
	var freelancestartErr error
	var levtechErr error
	var akkodisErr error

	// 構造体作成
	dataSources := []struct {
		jobInfoSlice *[]job.JobInfo
		errPtr *error
		fetchFunc func() ([]job.JobInfo, error)
		title string
	}{
		{&freelancestartJobInfoSlice, &freelancestartErr, job.GetFreelanceStartDetails, "フリーランススタート"},
		{&levtechJobInfoSlice, &levtechErr, job.GetLevtechDetails, "レバテック"},
		{&akkodisJobInfoSlice, &akkodisErr, job.GetAkkodisDetails, "AKKODIS"},
	}

	// 並列でデータを取得
	wg.Add(len(dataSources))
	for _, source := range dataSources {
		go fetchDataInParallel(&wg, source.jobInfoSlice, source.errPtr, source.fetchFunc)
	}
	wg.Wait()

	// 取得したデータのエラー判定
	for _, source := range dataSources {
		if checkAndPrintError(*source.errPtr, "Error fetching "+source.title+" job details:") {
			return
		}
	}

	// 件名と本文作成
	emailSubject := "新規案件リスト"
	emailBody := buildEmailBody(dataSources)

	// メール送信
	sendErr := mail.SendEmail(emailSubject, emailBody)
	if sendErr != nil {
		fmt.Println("Error sending email:", sendErr)
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
func buildEmailBody(dataSources []struct {
	jobInfoSlice *[]job.JobInfo
	errPtr *error
	fetchFunc func() ([]job.JobInfo, error)
	title string }) string {
		
    var emailBody strings.Builder
    for _, source := range dataSources {
        emailBody.WriteString("■" + source.title + "\n\n")
        for _, jobInfo := range *source.jobInfoSlice {
            emailBody.WriteString("　" + jobInfo.Name + "\n")
            emailBody.WriteString("　" + jobInfo.URL + "\n\n")
        }
        emailBody.WriteString("\n")
    }
    return emailBody.String()
}
