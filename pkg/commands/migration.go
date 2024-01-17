package commands

import (
	"database/sql"
	"fmt"
	"github.com/spf13/cobra"

	"github.com/pressly/goose/v3"
)

var ApiMigration = &cobra.Command{
	Use:   "migration",
	Short: "serve command for run the application",
	Long: `serve command for run the application.
	Example:
		$ ./monorepo migration
		$ ./monorepo migration --help
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migration command use up or down sub command")
	},
}

func init() {
	ApiMigration.AddCommand(ApiMigrationUp)
	ApiMigration.AddCommand(ApiMigrationDown)
}

// ApiMigrationUp sub command add
var ApiMigrationUp = &cobra.Command{
	Use:   "up",
	Short: "up command for migration up the database",
	Long: `up command for migration up the database.

	Example:
		$ ./monorepo migration up
		$ ./monorepo migration up --help
	`,

	Run: migrationUp,
}

// ApiMigrationDown sub command add
var ApiMigrationDown = &cobra.Command{
	Use:   "down",
	Short: "down command for migration down the database",
	Long: `down command for migration down the database.

	Example:
		$ ./monorepo migration down
		$ ./monorepo migration down --help
	`,
	Run: migrationDown,
}

func migrationUp(cmd *cobra.Command, args []string) {
	goose.Up(&sql.DB{}, "./migrations")
}

func migrationDown(cmd *cobra.Command, args []string) {

}
