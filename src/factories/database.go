package factories

import (
	"database/sql"
	"log"
)

func MakeConnectionDatabse() sql.DB {
	db, err := sql.Open("sqlite3", "./order.db")
	if err != nil {
		log.Fatalf("[MakeConnectionDatabse] Error in connect to database: %s", err)
	}
	return *db
}

func MakeTables() {
	db := MakeConnectionDatabse()

	query := `
		CREATE TABLE IF NOT EXISTS orders (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		price INTEGER NOT NULL,
		date DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Printf("Failed to prepare SQL: %s", err)
	}

	if _, err := stmt.Exec(); err != nil {
		log.Printf("Error in create table: %s", err)
	}
}
