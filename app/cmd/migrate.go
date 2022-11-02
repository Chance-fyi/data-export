package cmd

import (
	"data-export/boot"
	"data-export/database/migrations"
	"data-export/pkg/migrate"
	"github.com/spf13/cobra"
)

func migrateCommand() *cobra.Command {
	m := &cobra.Command{
		Use:   "migrate",
		Short: "run all unexecuted migrations",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			boot.Init()
			migrations.Init()
		},
		Run: func(cmd *cobra.Command, args []string) {
			migrate.Run(migrate.Up)
		},
	}

	m.AddCommand(
		migrateRollbackCommand(),
		migrateResetCommand(),
		migrateRefreshCommand(),
	)

	return m
}
