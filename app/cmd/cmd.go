package cmd

import (
	"github.com/spf13/cobra"
)

func RootCommand() *cobra.Command {
	cmd := &cobra.Command{}

	cmd.AddCommand(
		testCommand(),
		serverCommand(),
		genCommand(),
		migrateCommand(),
	)

	return cmd
}
