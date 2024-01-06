package cmd

import (
	"github.com/spf13/cobra"
	"monorepo/pkg/commands"
)

func Init() {

	rootCmd.AddCommand(commands.ApiDoctor)
}

var rootCmd = &cobra.Command{
	Use:   "monorepo",
	Short: "monorepo is a CLI tool for managing your monolithic application",
	//Long:  `go-cli is a CLI tool for managing your application`,
}
