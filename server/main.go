package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/GooruApp/gooru/server/pkg/start"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	err := start.Run(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}

	os.Exit(0)
}
