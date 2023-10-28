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

	var faworksJobInfoSlice []job.JobInfo
	var levtechJobInfoSlice []job.JobInfo
	var akkodisJobInfoSlice []job.JobInfo
	var geechsJobInfoSlice []job.JobInfo

	var faworksErr error
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
		{&faworksJobInfoSlice, &faworksErr, job.GetFaworksDetails, "FAworks"},
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
		if checkAndPrintError(*source.ErrPtr, source.Title+"のジョブ詳細を取得中にエラーが発生しました:") {
			return
		}
	}

	// JobInfoSliceが空の場合、要素を削除
	for i := 0; i < len(dataSources); {
		if len(*dataSources[i].JobInfoSlice) == 0 {
			dataSources = append(dataSources[:i], dataSources[i+1:]...)
		} else {
			i++
		}
	}

	if len(dataSources) > 0 {
		// 本文作成
		emailBody := mail.BuildEmailBody(dataSources)

		// メール送信
		sendErr := mail.SendEmail(emailBody, db)
		if sendErr != nil {
			fmt.Println("メール送信に失敗しました:", sendErr)
			return
		}
	} else {
		fmt.Println("新規案件がありません")
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

	if len(dataSlice) == 0 {
        return
    } else {
		*resultSlice = append(*resultSlice, dataSlice...)
	}
}

// エラーチェック
func checkAndPrintError(err error, message string) bool {
    if err != nil {
        fmt.Println(message, err)
        return true
    }
    return false
}
