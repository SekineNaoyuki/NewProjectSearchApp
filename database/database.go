package database

import (
	"sync"
	"time"
    "database/sql"
    _ "github.com/lib/pq"
	"NewProjectSearchApp/constants"
)

var mutex sync.Mutex

func Connect() (*sql.DB, error) {
	connStr :=
		"user=" + constants.DbUser + 
		" dbname=" + constants.DbName + 
		" password=" + constants.DbPass + 
		" host=" + constants.DbHost + 
		" port=" + constants.DbPort +
		" sslmode=disable"

    // データベースに接続
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, err
    }

    // データベース接続を確立
    err = db.Ping()
    if err != nil {
        return nil, err
    }

    return db, nil
}

func TestConnect() (*sql.DB, error) {
	connStr :=
		"user=" + constants.DbUser + 
		" dbname=" + constants.TestDbName + 
		" password=" + constants.DbPass + 
		" host=" + constants.DbHost + 
		" port=" + constants.DbPort +
		" sslmode=disable"

    // データベースに接続
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, err
    }

    // データベース接続を確立
    err = db.Ping()
    if err != nil {
        return nil, err
    }

    return db, nil
}

func IsNewJobUnique(db *sql.DB, name string) bool {
	mutex.Lock()
	defer mutex.Unlock()

    query := "SELECT COUNT(*) FROM notified_jobs WHERE name = $1"

    var count int
    err := db.QueryRow(query, name).Scan(&count)
    if err != nil {
        return false
    }

    return count == 0
}

func InsertJob(db *sql.DB, name string, url string) {
    mutex.Lock()
    defer mutex.Unlock()

	notificationTime := time.Now()

    query := "INSERT INTO notified_jobs (name, url, notification_time) VALUES ($1, $2, $3)"

    db.Exec(query, name, url, notificationTime)
}
