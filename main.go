package main

import (
	"fmt"
	"sync"

	"NewProjectSearchApp/pkg/mail"
	"NewProjectSearchApp/pkg/job"
)

func main() {
	var wg sync.WaitGroup

	var freelancestartJobInfoSlice []job.JobInfo
	var levtechJobInfoSlice []job.JobInfo
	var akkodisJobInfoSlice []job.JobInfo
	var geechsJobInfoSlice []job.JobInfo

	var freelancestartErr error
	var levtechErr error
	var akkodisErr error
	var geechsErr error

	// 構造体作成
	dataSources := []struct {
		JobInfoSlice *[]job.JobInfo
		ErrPtr *error
		FetchFunc func() ([]job.JobInfo, error)
		Title string
	}{
		{&freelancestartJobInfoSlice, &freelancestartErr, job.GetFreelanceStartDetails, "フリーランススタート"},
		{&levtechJobInfoSlice, &levtechErr, job.GetLevtechDetails, "レバテック"},
		{&akkodisJobInfoSlice, &akkodisErr, job.GetAkkodisDetails, "AKKODIS"},
		{&geechsJobInfoSlice, &geechsErr, job.GetGeechsDetails, "Geechs"},
	}

	// 並列でデータを取得
	wg.Add(len(dataSources))
	for _, source := range dataSources {
		go fetchDataInParallel(&wg, source.JobInfoSlice, source.ErrPtr, source.FetchFunc)
	}
	wg.Wait()

	// 取得したデータのエラー判定
	for _, source := range dataSources {
		if checkAndPrintError(*source.ErrPtr, "Error fetching "+source.Title+" job details:") {
			return
		}
	}

	emailBody := mail.BuildEmailBody(dataSources)

	// メール送信
	sendErr := mail.SendEmail(emailBody)
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
