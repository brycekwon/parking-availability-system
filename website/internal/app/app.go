package app

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/brycekwon/parking-availability-system/website/internal/database"
	"github.com/brycekwon/parking-availability-system/website/internal/middleware"
)

type App struct {
	logger *slog.Logger
	router *http.ServeMux
	db     *sql.DB
}

func New(logger *slog.Logger) *App {
	return &App{
		logger: logger,
		router: http.NewServeMux(),
	}
}

func (a *App) Start(ctx context.Context, port uint16) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}
	a.db = db

	if err := a.loadPages(); err != nil {
		return err
	}

	if err := a.loadRoutes(); err != nil {
		return err
	}

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: middleware.Logging(a.logger, a.router),
	}

	done := make(chan struct{})
	go func() {
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.logger.Error("failed to listen and serve", slog.Any("error", err))
		}
		close(done)
	}()

	a.logger.Info(fmt.Sprintf("server is listening on port %d", port))
	select {
	case <-done:
		break
	case <-ctx.Done():
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		server.Shutdown(ctx)
		cancel()
	}

	return nil
}
