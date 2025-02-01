package main

import (
	"context"
	"embed"
	"log/slog"
	"os"
	"os/signal"

	"github.com/brycekwon/parking-availability-system/website/internal/app"
)

//go:embed all:dist
var assets embed.FS

//go:embed all:templates
var pages embed.FS

func main() {
	// structured logger for debugging and monitoring
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// gracefully handle termination
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	a := app.New(logger, assets, pages)
	if err := a.Start(ctx); err != nil {
		logger.Error("failed to start server", slog.Any("error", err))
	}
}
