package main

import (
	"context"
	"embed"
	"os"
	"os/signal"

	"github.com/GooruApp/gooru/server/internal/cmd"
)

//go:embed migrations/*/*.sql
var migrations embed.FS

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	result := cmd.Execute(ctx, migrations)

	os.Exit(result)
}
