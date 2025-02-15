package handlers

import (
	"context"
	"log/slog"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type Handler struct {
	logger *slog.Logger
	ctx    context.Context

	cache *redis.Client
	db    *sqlx.DB
}

func New(ctx context.Context, logger *slog.Logger, db *sqlx.DB, cache *redis.Client) *Handler {
	return &Handler{
		logger: logger,
		ctx:    ctx,

		cache: cache,
		db:    db,
	}
}
