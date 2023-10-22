package main

import (
	"fmt"
	"sync"
	"database/sql"
    _ "github.com/lib/pq"
	"NewProjectSearchApp/database"
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

	// DB接続
	db, dbErr := database.Connect()
    if dbErr != nil {
        fmt.Println("DB接続に失敗しました:", dbErr)
    }

	// 構造体作成
	dataSources := []struct {
		JobInfoSlice *[]job.JobInfo
		ErrPtr *error
		FetchFunc func(*sql.DB) ([]job.JobInfo, error)
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
		go fetchDataInParallel(&wg, source.JobInfoSlice, source.ErrPtr, source.FetchFunc, db)
	}
	wg.Wait()

	// 取得したデータのエラー判定
	for _, source := range dataSources {
		if checkAndPrintError(*source.ErrPtr, "Error fetching "+source.Title+" job details:") {
			return
		}
	}

	// 本文作成
	emailBody := mail.BuildEmailBody(dataSources)

	// メール送信
	sendErr := mail.SendEmail(emailBody, db)
	if sendErr != nil {
		fmt.Println("メール送信に失敗しました:", sendErr)
		return
	}
}

// 並列処理
func fetchDataInParallel(wg *sync.WaitGroup, resultSlice *[]job.JobInfo, errPtr *error, fetchFunc func(*sql.DB) ([]job.JobInfo, error), db *sql.DB) {
    defer wg.Done()

    dataSlice, err := fetchFunc(db)
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
