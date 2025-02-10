package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/GooruApp/gooru/server/internal/cmd"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	result := cmd.Execute(ctx)

	os.Exit(result)
}
