package job

import (
    "testing"
	"fmt"
	"NewProjectSearchApp/database"
)

func TestGetFaworksDetails(t *testing.T) {
	// テストDB接続
	TestDb, dbErr := database.TestConnect()
    if dbErr != nil {
        fmt.Println("DB接続に失敗しました:", dbErr)
    }

    // テスト対象の関数を呼び出す
    jobInfoSlice, err := GetFaworksDetails(TestDb)

    // テストケースを記述し、期待される結果と比較
    if err != nil {
        t.Errorf("GetFaworksDetailsの実行に失敗しました: %v", err)
    }

    // テスト結果を評価
    if len(jobInfoSlice) == 0 {
        t.Errorf("GetFaworksDetailsの実行結果のスライスが空です")
    }
}

func TestGetLevtechDetails(t *testing.T) {
		// テストDB接続
		TestDb, dbErr := database.TestConnect()
		if dbErr != nil {
			fmt.Println("DB接続に失敗しました:", dbErr)
		}
	
		// テスト対象の関数を呼び出す
		jobInfoSlice, err := GetLevtechDetails(TestDb)
	
		// テストケースを記述し、期待される結果と比較
		if err != nil {
			t.Errorf("GetLevtechDetailsの実行に失敗しました: %v", err)
		}
	
		// テスト結果を評価
		if len(jobInfoSlice) == 0 {
			t.Errorf("GetLevtechDetailsの実行結果のスライスが空です")
		}	
}

func TestGetAkkodisDetails(t *testing.T) {
		// テストDB接続
		TestDb, dbErr := database.TestConnect()
		if dbErr != nil {
			fmt.Println("DB接続に失敗しました:", dbErr)
		}
	
		// テスト対象の関数を呼び出す
		jobInfoSlice, err := GetAkkodisDetails(TestDb)
	
		// テストケースを記述し、期待される結果と比較
		if err != nil {
			t.Errorf("GetAkkodisDetailsの実行に失敗しました: %v", err)
		}
	
		// テスト結果を評価
		if len(jobInfoSlice) == 0 {
			t.Errorf("GetAkkodisDetailsの実行結果のスライスが空です")
		}	
}

func TestGetGeechsDetails(t *testing.T) {
		// テストDB接続
		TestDb, dbErr := database.TestConnect()
		if dbErr != nil {
			fmt.Println("DB接続に失敗しました:", dbErr)
		}
	
		// テスト対象の関数を呼び出す
		jobInfoSlice, err := GetGeechsDetails(TestDb)
	
		// テストケースを記述し、期待される結果と比較
		if err != nil {
			t.Errorf("GetGeechsDetailsの実行に失敗しました: %v", err)
		}
	
		// テスト結果を評価
		if len(jobInfoSlice) == 0 {
			t.Errorf("GetGeechsDetailsの実行結果のスライスが空です")
		}	
}