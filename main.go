package main

import (
	"context"
	"os"
	"os/signal"

	"tg_star_miner/internal/app"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Kill)
	defer cancel()

	if err := app.Run(ctx); err != nil {
		panic(err)
	}
}
