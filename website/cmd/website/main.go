package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	"github.com/brycekwon/parking-availability-system/website/internal/app"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	a := app.New(logger)
	if err := a.Start(ctx, 3000); err != nil {
		logger.Error("failed to start server", slog.Any("error", err))
	}
}
