package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to the DB. ")
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()
}

func createTables() {
	createEventTable := `
		CREATE TABLE IF NOT EXISTS events (
			id integer primary key autoincrement,
			name varchar(20) not null,
			description varchar(50) not null,
			location varchar(50) not null,
			dateTime DATETIME,
			user_id integer
		)
	`
	_, err := DB.Exec(createEventTable)

	if err != nil {
		panic("Could not create event table")
	}
}
