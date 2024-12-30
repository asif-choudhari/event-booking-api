package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var Connection *sql.DB

func InitDb() {
	var err error
	Connection, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	Connection.SetMaxOpenConns(10)
	Connection.SetMaxIdleConns(5)

	createTable()
}

func createTable() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		userId INTEGER,
		FOREIGN KEY(userId) REFERENCES users(id)
	)
	`

	_, err := Connection.Exec(createEventsTable)
	if err != nil {
		panic("Could not create events table")
	}

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`

	_, err = Connection.Exec(createUsersTable)
	if err != nil {
		panic("Could not create user table")
	}

}
