package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDb() {
	var err error
	DB, err = sql.Open("sqlite3", "./db/api.db")

	if err != nil {
		panic("Can't connect to database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date_time DATETIME DEFAULT CURRENT_TIMESTAMP,
		user_id INTEGER
	)
	`

	_, err := DB.Exec(createEventsTable)

	if err != nil {
		panic("Can't create events table")
	}
}
