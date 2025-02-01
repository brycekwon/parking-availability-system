package database

import "time"

type Event struct {
	timestamp time.Time
	lot_name  string
	status    int
}
