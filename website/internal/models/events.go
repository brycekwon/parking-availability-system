package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Event struct {
	timestamp time.Time
	lot_name  string
	status    int
}

func InsertEvent(db *sqlx.DB, lotname string, status int) error {
	event := &Event{
		timestamp: time.Now().UTC(),
		lot_name:  lotname,
		status:    status,
	}

	insertEvent := `INSERT INTO events ("timestamp", "lotname", "status") VALUES ($1, $2, $3);`
	statement, err := db.Prepare(insertEvent)
	if err != nil {
		return err
	}
	_, err = statement.Exec(event.timestamp, event.lot_name, event.status)
	return err
}
