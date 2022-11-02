package main

import (
	"data-export/app/cmd"
	"data-export/pkg/console"
)

func main() {
	c := cmd.RootCommand()
	err := c.Execute()
	console.ExitIf(err)
}
