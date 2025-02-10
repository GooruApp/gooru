package cmd

import (
	"context"
	"embed"
	"fmt"

	"github.com/GooruApp/gooru/server/internal/api"
	"github.com/GooruApp/gooru/server/internal/env"
	"github.com/GooruApp/gooru/server/internal/logger"
	"github.com/GooruApp/gooru/server/internal/migrator"
	"github.com/spf13/cobra"
)

func StartCmd(ctx context.Context, migrations embed.FS) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Args:  cobra.ExactArgs(0),
		Short: "Runs the REST API",
		RunE: func(cmd *cobra.Command, args []string) error {
			env, err := env.Get()
			if err != nil {
				return fmt.Errorf("couldn't get the env: %v", err)
			}

			logger, err := logger.New("start", env.AppEnv())
			if err != nil {
				return fmt.Errorf("couldn't create a new logger: %v", err)
			}

			migrator, err := migrator.New(migrations, env.DBConnStr())
			if err != nil {
				return fmt.Errorf("couldn't create a new migrator: %v", err)
			}

			err = migrator.Up()
			if err != nil {
				return fmt.Errorf("error occured when running migrations: %v", err)
			}

			api := api.NewAPI(ctx, logger)
			srv := api.Server(env.Port())

			go func() { _ = srv.ListenAndServe() }()

			fmt.Printf("Started API on port: %d\n", env.Port())

			// Blocks until a value is passed on the done ch
			<-ctx.Done()

			_ = srv.Shutdown(ctx)

			return nil
		},
	}

	return cmd
}
