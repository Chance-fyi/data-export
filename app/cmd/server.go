package cmd

import (
	"data-export/boot"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

type serverOptions struct {
	Mode string
}

func serverCommand() *cobra.Command {
	opt := serverOptions{}
	server := &cobra.Command{
		Use:   "server",
		Short: "start http service",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			boot.Init()
		},
		Run: func(cmd *cobra.Command, args []string) {
			runServer(opt)
		},
	}

	flags := server.Flags()
	flags.StringVarP(&opt.Mode, "mode", "m", gin.ReleaseMode, "indicates gin mode")

	return server
}

func runServer(opt serverOptions) {
	gin.SetMode(opt.Mode)
	boot.Route.Init()
}
