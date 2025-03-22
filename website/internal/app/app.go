package app

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/brycekwon/parking-availability-system/website/internal/config"
	"github.com/brycekwon/parking-availability-system/website/internal/middlewares"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type App struct {
	cfg *config.Config
	ctx context.Context

	db     *sqlx.DB
	cache  *redis.Client
	router *http.ServeMux
	logger *slog.Logger
}

func New() *App {
	return &App{
		cfg: config.New(),

		router: http.NewServeMux(),
		// logger: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})),
		logger: slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})),
	}
}

func (a *App) Start(ctx context.Context) {
	a.ctx = ctx

	// if err := a.initDatabase(); err != nil {
	// 	a.logger.Error("failed to initialize database", slog.Any("error", err))
	// 	return
	// }

	// if err := a.initCache(); err != nil {
	// 	a.logger.Error("failed to initialize cache", slog.Any("error", err))
	// 	return
	// }

	if err := a.loadPages(); err != nil {
		a.logger.Error("failed to load pages", slog.Any("error", err))
		return
	}

	if err := a.loadRoutes(); err != nil {
		a.logger.Error("failed to load routes", slog.Any("error", err))
		return
	}

	addr := fmt.Sprintf("%s:%d", a.cfg.ServerConfig.Host, a.cfg.ServerConfig.Port)
	server := &http.Server{
		Addr:    addr,
		Handler: middlewares.Logging(a.logger, a.router),
	}

	done := make(chan struct{})
	go func() {
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.logger.Error("failed to listen and serve", slog.Any("error", err))
		}
		close(done)
	}()

	a.logger.Info("server is listening", slog.String("address", addr))

	select {
	case <-done:
		break
	case <-ctx.Done():
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		server.Shutdown(ctx)
		cancel()
	}
}
