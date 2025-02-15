package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/brycekwon/parking-availability-system/website/internal/app"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	a := app.New()
	a.Start(ctx)
}
