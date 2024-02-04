package commands

import (
	"embed"
	"fmt"
	"github.com/spf13/cobra"
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

//go:embed migration/*.sql
var postgresMigrations embed.FS

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

const (
	postgresMigrationDir = "migration"
	postgresDialect      = "postgres"
)

func migrationUp(cmd *cobra.Command, args []string) {
	//	var db *database.Postgres
	//
	//	db, err := database.NewPostgres(&config.Database{})
	//	if err != nil {
	//		log.Fatal().Err(err).Msg("database connection error")
	//	}
	//	//defer closeDB(db)
	//
	//	goose.SetTableName("migration")
	//
	//	if err = goose.SetDialect(postgresDialect); err != nil {
	//		log.Fatal().Err(err).Msg("goose dialect error")
	//	}
	//
	//	goose.SetBaseFS(postgresMigrations)
	//
	//	//if err = goose.Up(db.DB, postgresMigrationDir); err != nil {
	//	//	log.Fatal().Err(err).Msg("goose up error")
	//	//}

}

func migrationDown(cmd *cobra.Command, args []string) {

}
