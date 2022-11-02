package cmd

import (
	"data-export/pkg/migrate"
	"github.com/spf13/cobra"
)

func migrateRefreshCommand() *cobra.Command {
	rollback := &cobra.Command{
		Use:   "refresh",
		Short: "first rollback all migrations that have been run, then run all migrations",
		Run: func(cmd *cobra.Command, args []string) {
			migrate.Run(migrate.Reset)
			migrate.Run(migrate.Up)
		},
	}

	return rollback
}
