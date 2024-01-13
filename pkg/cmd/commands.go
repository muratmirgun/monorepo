package cmd

import (
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "monorepo",
		Short: "monorepo is a CLI tool for managing your monolithic application",
		//Long:  `go-cli is a CLI tool for managing your application`,
	}

}
