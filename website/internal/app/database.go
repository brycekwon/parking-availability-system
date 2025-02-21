package app

import (
	"fmt"
	"log/slog"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

func (a *App) initDatabase() error {
	db, err := sqlx.Connect(a.cfg.DatabaseConfig.Driver, fmt.Sprintf(
		"postgresql://%s:%d/%s?sslmode=%s&user=%s&password=%s",
		a.cfg.DatabaseConfig.Host,
		a.cfg.DatabaseConfig.Port,
		a.cfg.DatabaseConfig.Name,
		a.cfg.DatabaseConfig.SSLMode,
		a.cfg.DatabaseConfig.Username,
		a.cfg.DatabaseConfig.Password,
	))
	if err != nil {
		return err
	}
	a.db = db

	createTableQuery := `
		CREATE TABLE IF NOT EXISTS events (
			timestamp TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
			lotname VARCHAR(255) NOT NULL,
			spotid INT NOT NULL,
			status INT NOT NULL
		);
	`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		return fmt.Errorf("failed to create events table: %w", err)
	}

	a.logger.Info("database initialized", slog.String("address", fmt.Sprintf("%s:%d", a.cfg.DatabaseConfig.Host, a.cfg.DatabaseConfig.Port)))

	return nil
}

func (a *App) initCache() error {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", a.cfg.CacheConfig.Host, a.cfg.CacheConfig.Port),
		// Username: a.cfg.CacheConfig.Username,
		// Password: a.cfg.CacheConfig.Password,
		DB: 1,
	})
	a.cache = client

	a.logger.Info("cache initialized", slog.String("address", fmt.Sprintf("%s:%d", a.cfg.CacheConfig.Host, a.cfg.CacheConfig.Port)))

	return nil
}
