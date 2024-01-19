package main

import (
	"github.com/muratmirgun/monorepo/pkg/cmd"
	"github.com/muratmirgun/monorepo/pkg/commands"
	"os"
)

func main() {
	root := cmd.NewRootCommand()

	root.AddCommand(commands.ApiDoctor)
	root.AddCommand(commands.ApiServe)
	root.AddCommand(commands.ApiMigration)

	//root.AddCommand(version)

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
