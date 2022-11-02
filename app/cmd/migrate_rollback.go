package cmd

import (
	"data-export/pkg/migrate"
	"github.com/spf13/cobra"
)

func migrateRollbackCommand() *cobra.Command {
	rollback := &cobra.Command{
		Use:   "rollback",
		Short: "rollback the last migration operation",
		Run: func(cmd *cobra.Command, args []string) {
			migrate.Run(migrate.Rollback)
		},
	}

	return rollback
}
