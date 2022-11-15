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
