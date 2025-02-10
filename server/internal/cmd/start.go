package cmd

import (
	"context"
	"embed"
	"fmt"

	"github.com/GooruApp/gooru/server/internal/api"
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
			port := 8000

			logger, err := logger.New("start")
			if err != nil {
				return fmt.Errorf("couldn't create a new logger: %v", err)
			}

			migrator, err := migrator.New(migrations, "sqlite://booru.db")
			if err != nil {
				return err
			}

			err = migrator.Up()
			if err != nil {
				return err
			}

			api := api.NewAPI(ctx, logger)
			srv := api.Server(port)

			go func() { _ = srv.ListenAndServe() }()

			fmt.Printf("started api on port %d\n", port)

			// Blocks until a value is passed on the done ch
			<-ctx.Done()

			_ = srv.Shutdown(ctx)

			return nil
		},
	}

	return cmd
}
