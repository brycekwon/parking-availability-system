package app

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"net/http"
	"time"

	"github.com/brycekwon/parking-availability-system/website/database"
	"github.com/brycekwon/parking-availability-system/website/middlewares"
)

type App struct {
	logger *slog.Logger
	assets fs.FS
	views  fs.FS
	router *http.ServeMux
	db     *sql.DB
}

func New(logger *slog.Logger, assets fs.FS, views fs.FS) *App {
	return &App{
		logger: logger,
		assets: assets,
		views:  views,
		router: http.NewServeMux(),
	}
}

func (a *App) Start(ctx context.Context, port uint16) error {
	db, err := database.Connect()
	if err != nil {
		a.logger.Error("could not connect to database", slog.Any("error", err))
		return err
	}
	a.db = db

	if err := a.loadStatic(); err != nil {
		a.logger.Error("could not load static assets", slog.Any("error", err))
		return err
	}

	if err := a.loadViews(); err != nil {
		a.logger.Error("could not load page views", slog.Any("error", err))
		return err
	}

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: middlewares.Logging(a.logger, middlewares.ErrorHandler(a.logger, a.router)),
	}

	done := make(chan struct{})
	go func() {
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.logger.Error("failed to listen and server", slog.Any("error", err))
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
