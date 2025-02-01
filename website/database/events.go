package database

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func InsertEvent(db *sql.DB, parkinglotname string, status int) error {
	event := &Event{
		timestamp: time.Now().UTC(),
		lot_name:  parkinglotname,
		status:    status,
	}

	insertEvent := `INSERT INTO events (timestamp, parkinglotname, status) VALUES (?, ?, ?)`

	statement, err := db.Prepare(insertEvent)
	if err != nil {
		return err
	}
	_, err = statement.Exec(event.timestamp, event.lot_name, event.status)
	return err
}
