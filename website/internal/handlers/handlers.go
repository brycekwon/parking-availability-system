package handlers

import (
	"html/template"
	"log/slog"
)

type Handler struct {
	page   *template.Template
	logger *slog.Logger
}

func New(logger *slog.Logger, page *template.Template) *Handler {
	return &Handler{
		page:   page,
		logger: logger,
	}
}
