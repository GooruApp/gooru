package start

import (
	"context"
	"fmt"

	"github.com/GooruApp/gooru/server/internal/api"
	"github.com/GooruApp/gooru/server/internal/config"
	"github.com/GooruApp/gooru/server/internal/logger"
	"github.com/GooruApp/gooru/server/internal/migrator"
)

func Run(ctx context.Context) error {
	// env, err := env.AppEnv.Test
	// if err != nil {
	// 	return fmt.Errorf("couldn't get the env: %v", err)
	// }
	config.Settings.Mode.Get()

	logger, err := logger.New("start", config.Settings.Mode.Get())
	if err != nil {
		return fmt.Errorf("couldn't create a new logger: %v", err)
	}

	migrator, err := migrator.New(config.Settings.DBConnStr.Get())
	if err != nil {
		return fmt.Errorf("couldn't create a new migrator: %v", err)
	}

	err = migrator.Up()
	if err != nil {
		return fmt.Errorf("error occured when running migrations: %v", err)
	}

	api := api.NewAPI(ctx, logger)
	srv := api.Server(config.Settings.Port.Get())

	go func() { _ = srv.ListenAndServe() }()

	fmt.Printf("Started API on port: %d\n", config.Settings.Port.Get())

	// Blocks until a value is passed on the done ch
	<-ctx.Done()

	_ = srv.Shutdown(ctx)

	return nil
}
