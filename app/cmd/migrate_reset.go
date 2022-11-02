package cmd

import (
	"data-export/pkg/migrate"
	"github.com/spf13/cobra"
)

func migrateResetCommand() *cobra.Command {
	rollback := &cobra.Command{
		Use:   "reset",
		Short: "rollback all migrations that have been run",
		Run: func(cmd *cobra.Command, args []string) {
			migrate.Run(migrate.Reset)
		},
	}

	return rollback
}
