package cmd

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "monorepo",
		Short: "monorepo is a CLI tool for managing your monolithic application",
		Run: func(cmd *cobra.Command, args []string) {
			figure.NewColorFigure("Monorepo", "isometric1", "green", true).Print()

			fmt.Println("monorepo is a CLI tool for managing your monolithic application")
		},
	}

}
