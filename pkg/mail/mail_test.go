package mail

import (
    "database/sql"
    "testing"
	"fmt"
	"NewProjectSearchApp/pkg/job"
	"NewProjectSearchApp/database"
)

func TestBuildEmailBody(t *testing.T) {
    // テスト用のデータを作成
    dataSources := []struct {
        JobInfoSlice *[]job.JobInfo
        ErrPtr *error
        FetchFunc func(*sql.DB) ([]job.JobInfo, error)
        Title string
    }{
        {&[]job.JobInfo{
            {"Job1", "URL1"},
            {"Job2", "URL2"},
        }, nil, nil, "Test Source 1"},
        {&[]job.JobInfo{
            {"Job3", "URL3"},
            {"Job4", "URL4"},
        }, nil, nil, "Test Source 2"},
    }

    // 期待される結果を用意
    expected := "■Test Source 1\n\n　Job1\n　URL1\n\n　Job2\n　URL2\n\n\n■Test Source 2\n\n　Job3\n　URL3\n\n　Job4\n　URL4\n\n\n"

    // BuildEmailBody 関数を呼び出す
    result := BuildEmailBody(dataSources)

    // 結果を検証
    if result != expected {
        t.Errorf("期待値:\n%s\n\n実際の結果:\n%s", expected, result)
    }
}

func TestGetMailInfoFromDB(t *testing.T) {
	// テストDB接続
	testDB, dbErr := database.TestConnect()
	if dbErr != nil {
		fmt.Println("DB接続に失敗しました:", dbErr)
	}

    // テスト対象の関数を呼び出す
    fromEmail, toEmail, subject, smtpUsername, smtpPassword, smtpHost, smtpPort := GetMailInfoFromDB(testDB)

    // テストケースを記述し、期待される結果と比較
	expectedFromEmail := "myProject.com"
    expectedToEmail := "test@gmail.com"
    expectedSubject := "新規案件リスト"
    expectedSmtpUsername := "test@gmail.com"
    expectedSmtpPassword := "testpass"
    expectedSmtpHost := "smtp.gmail.com"
    expectedSmtpPort := "587"

    if fromEmail != expectedFromEmail {
        t.Errorf("fromEmailが予想された値 %s ではなく、実際の値は %s です", expectedFromEmail, fromEmail)
    }
    
    if toEmail != expectedToEmail {
        t.Errorf("toEmailが予想された値 %s ではなく、実際の値は %s です", expectedToEmail, toEmail)
    }
    
    if subject != expectedSubject {
        t.Errorf("subjectが予想された値 %s ではなく、実際の値は %s です", expectedSubject, subject)
    }
    
    if smtpUsername != expectedSmtpUsername {
        t.Errorf("smtpUsernameが予想された値 %s ではなく、実際の値は %s です", expectedSmtpUsername, smtpUsername)
    }
    
    if smtpPassword != expectedSmtpPassword {
        t.Errorf("smtpPasswordが予想された値 %s ではなく、実際の値は %s です", expectedSmtpPassword, smtpPassword)
    }
    
    if smtpHost != expectedSmtpHost {
        t.Errorf("smtpHostが予想された値 %s ではなく、実際の値は %s です", expectedSmtpHost, smtpHost)
    }
    
    if smtpPort != expectedSmtpPort {
        t.Errorf("smtpPortが予想された値 %s ではなく、実際の値は %s です", expectedSmtpPort, smtpPort)
    }
}
