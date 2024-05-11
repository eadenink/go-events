package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type createTableOperation struct {
	query     string
	tableName string
}

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
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date_time DATETIME DEFAULT CURRENT_TIMESTAMP,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	operations := []createTableOperation{
		{query: createUsersTable, tableName: "users"},
		{query: createEventsTable, tableName: "events"},
		{query: createRegistrationsTable, tableName: "registrations"},
	}

	done := make([]chan bool, len(operations))

	for i, operation := range operations {
		done[i] = make(chan bool)
		go createTable(done[i], operation.query, operation.tableName)
	}

	for _, done := range done {
		<-done
	}
}

func createTable(channel chan bool, query string, tableName string) {
	_, err := DB.Exec(query)
	if err != nil {
		panic(fmt.Sprintf("Can't create \"%s\" table", tableName))
	}

	channel <- true
}
