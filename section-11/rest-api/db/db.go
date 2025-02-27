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
	createEventsTable := `CREATE TABLE IF NOT EXISTS routes (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
    	description TEXT NOT NULL,
        location TEXT NOT NULL,
		datetime DATETIME NOT NULL,
		user_id INTEGER                                
    )`

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		panic(err)
	}
}

func InsertEvent() {

}
