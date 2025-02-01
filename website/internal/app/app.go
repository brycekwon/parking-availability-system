package app

import (
	"context"
	"errors"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/brycekwon/parking-availability-system/website/internal/middlewares"
)

type App struct {
	router *http.ServeMux
	logger *slog.Logger
	assets fs.FS
	pages  fs.FS
}

func New(logger *slog.Logger, assets fs.FS, pages fs.FS) *App {
	return &App{
		router: http.NewServeMux(),
		logger: logger,
		assets: assets,
		pages:  pages,
	}
}

func (a *App) Start(ctx context.Context) error {
	if err := a.loadAssets(); err != nil {
		a.logger.Error("failed to load assets", slog.Any("error", err))
		os.Exit(1)
	}

	if err := a.loadPages(); err != nil {
		a.logger.Error("failed to load pages", slog.Any("error", err))
		os.Exit(1)
	}

	server := http.Server{
		Addr:    ":3000",
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

	a.logger.Info("server listening", slog.String("addr", server.Addr))
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
