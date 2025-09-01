package main

import (
	"context"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Kill)
	defer cancel()

	if err := Run(ctx); err != nil {
		panic(err)
	}
}
