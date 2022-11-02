package cmd

import (
	"data-export/app/cmd/template"
	"data-export/pkg/app"
	"data-export/pkg/console"
	"data-export/pkg/str"
	"fmt"
	"github.com/spf13/cobra"
)

type genMigrateOptions struct {
	Connect string
}

func genMigrateCommand() *cobra.Command {
	opt := genMigrateOptions{}
	migration := &cobra.Command{
		Use:   "migration",
		Short: "generate database migration file",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			runGenMigrate(args, opt)
		},
	}

	flags := migration.Flags()
	flags.StringVarP(&opt.Connect, "connect", "c", "default", "database connect")

	return migration
}

func runGenMigrate(args []string, opt genMigrateOptions) {
	timeStr := app.TimeNowInTimezone().Format("2006_01_02_150405")
	fileName := timeStr + "_" + str.Snake(args[0])
	filePath := fmt.Sprintf("database/migrations/%s.go", fileName)

	generateFile(filePath, template.TemplateMigration, map[string]string{
		"{{PackageName}}": app.Name(),
		"{{FileName}}":    fileName,
		"{{Connect}}":     opt.Connect,
	})

	console.Successp("Migration file created，after modify it, use `migrate up` to migrate database.")
}
