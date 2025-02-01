package main

import (
	"context"
	"embed"
	"log/slog"
	"os"
	"os/signal"

	"github.com/brycekwon/parking-availability-system/website/app"
)

//go:embed all:static
var assets embed.FS

//go:embed all:views
var views embed.FS

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	a := app.New(logger, assets, views)
	if err := a.Start(ctx, 3000); err != nil {
		logger.Error("failed to start server", slog.Any("error", err))
	}
}
