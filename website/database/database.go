package database

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func Connect() (*sql.DB, error) {
	file, err := os.Create("./_data/sqlite.db")
	if err != nil {
		return nil, err
	}
	file.Close()

	db, err := sql.Open("sqlite3", "./_data/sqlite.db")
	if err != nil {
		return nil, err
	}

	if err := initTables(db); err != nil {
		return nil, err
	}

	return db, nil
}

func initTables(db *sql.DB) error {
	eventsTable := `CREATE TABLE events (
		"timestamp" DATETIME DEFAULT CURRENT_TIMESTAMP,
		"parkinglotname" TEXT,
		"status" INTEGER
	);`
	statement, err := db.Prepare(eventsTable)
	if err != nil {
		return err
	}
	statement.Exec()
	statement.Close()

	return nil
}
