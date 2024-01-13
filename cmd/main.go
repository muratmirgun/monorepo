package main

import (
	"monorepo/pkg/cmd"
	"monorepo/pkg/commands"
	"os"
)

func main() {
	root := cmd.NewRootCommand()

	root.AddCommand(commands.ApiDoctor)
	//root.AddCommand(commands.ApiServe)

	//version := cmd.NewVersionCommand()
	//root.AddCommand(version)

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
