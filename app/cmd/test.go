package cmd

import (
	"data-export/app/model"
	"data-export/pkg/console"
	"fmt"
	"github.com/spf13/cobra"
)

func testCommand() *cobra.Command {
	test := &cobra.Command{
		Use:   "test",
		Short: "run temporary test code",
		Run: func(cmd *cobra.Command, args []string) {
			runTest()
		},
	}

	return test
}

func runTest() {
	//测试代码调试
	console.Logln(fmt.Sprintf("%v", model.UserSql{}))
}
