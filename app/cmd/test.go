package cmd

import (
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
}
