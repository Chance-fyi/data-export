package cmd

import (
	"data-export/boot"
	"data-export/pkg/console"
	"data-export/pkg/file"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

func genCommand() *cobra.Command {
	gen := &cobra.Command{
		Use:   "gen",
		Short: "generate file and code",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			boot.Init()
		},
	}

	gen.AddCommand(
		genMigrateCommand(),
	)

	return gen
}

func generateFile(filePath string, templateStr string, variables ...map[string]string) {
	replaces := make(map[string]string)
	if len(variables) > 0 {
		replaces = variables[0]
	}

	// check whether the file exists
	if file.Exists(filePath) {
		console.Exit(filePath + " already exists!")
	}

	// variable replacement for template content
	for search, replace := range replaces {
		templateStr = strings.ReplaceAll(templateStr, search, replace)
	}

	err := file.Put([]byte(templateStr), filePath)
	console.ExitIf(err)

	console.Successp(fmt.Sprintf("[%s] created.", filePath))
}
