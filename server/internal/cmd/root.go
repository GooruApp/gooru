package cmd

import (
	"context"
	"embed"

	"github.com/spf13/cobra"
)

func Execute(ctx context.Context, migrations embed.FS) int {
	rootCmd := &cobra.Command{
		Use:   "gooru [command]",
		Short: "gooru",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cmd.Help()
			if err != nil {
				return err
			}

			return nil
		},
	}

	rootCmd.AddCommand(StartCmd(ctx, migrations))

	if err := rootCmd.Execute(); err != nil {
		return 1
	}

	return 0
}
