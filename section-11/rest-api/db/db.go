package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "sqlite3-data.db")
	if err != nil {
		// Panic if database connection failed
		panic(err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic(err)
	}

	createEventsTable := `CREATE TABLE IF NOT EXISTS events (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
    	description TEXT NOT NULL,
        location TEXT NOT NULL,
		datetime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE                                
    )`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic(err)
	}

	createRegistrationTable := `CREATE TABLE IF NOT EXISTS registrations (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
    	event_id INTEGER NOT NULL,
    	user_id INTEGER NOT NULL,
    	FOREIGN KEY(event_id) REFERENCES events(id) ON DELETE CASCADE,
    	FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
	)`

	_, err = DB.Exec(createRegistrationTable)
	if err != nil {
		panic(err)
	}
}
