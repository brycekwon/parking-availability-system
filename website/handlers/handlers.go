package handlers

import (
	"database/sql"
	"html/template"
	"log/slog"
)

type Handler struct {
	logger *slog.Logger
	views  *template.Template
	db     *sql.DB
}

func New(logger *slog.Logger, views *template.Template, db *sql.DB) *Handler {
	return &Handler{
		logger: logger,
		views:  views,
		db:     db,
	}
}
