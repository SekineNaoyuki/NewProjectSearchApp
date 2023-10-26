package database

import (
    "testing"
)

func TestIsNewJobUnique(t *testing.T) {
	// テストDB接続
	testDb, dbErr := TestConnect()
	if dbErr != nil {
		t.Errorf("DB接続に失敗しました: %v", dbErr)
	}
	defer testDb.Close()

    // テストデータ準備
	name := "テストサイトネーム"
	nonUniqueName := "テストサイトネーム２"
	url := "https://test.com"
	InsertJob(testDb, name, url)

	// テストケース1: 重複する名前
    unique := IsNewJobUnique(testDb, name)
    if unique {
        t.Errorf("重複する名前で実行した場合、期待した評価値ではありません: %t", unique)
    }

    // テストケース2: 重複のない名前
    nonUnique := IsNewJobUnique(testDb, nonUniqueName)
    if !nonUnique {
        t.Errorf("重複のない名前で実行した場合、期待した評価値ではありません: %t", nonUnique)
    }

	// テストデータ削除
	query := "DELETE FROM notified_jobs WHERE name = $1"
    _, err := testDb.Exec(query, name)
    if err != nil {
        t.Errorf("テストデータの削除中にエラーが発生しました: %v", err)
    }
}

func TestInsertJob(t *testing.T) {
    // テストDB接続
    testDb, dbErr := TestConnect()
    if dbErr != nil {
        t.Fatalf("DB接続に失敗しました: %v", dbErr)
    }
    defer testDb.Close()

    // テストデータ
    name := "テストサイトネーム"
    url := "https://test.com"

    // テスト対象の関数を呼び出す
    InsertJob(testDb, name, url)

    // データベースから該当のデータを取得
    var insertedName string
    err := testDb.QueryRow("SELECT name FROM notified_jobs WHERE name = $1", name).Scan(&insertedName)
    if err != nil {
        t.Errorf("データベースからのデータ取得中にエラーが発生しました: %v", err)
    }

    // テスト結果を評価
    if insertedName != name {
        t.Errorf("期待したデータがデータベースに挿入されていません: 期待値=%s, 実際の値=%s", name, insertedName)
    }
}