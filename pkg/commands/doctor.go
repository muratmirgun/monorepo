package commands

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/muratmirgun/monorepo/config"
	"os/exec"
	"regexp"
	"time"

	"github.com/chelnak/ysmrr"
	"github.com/chelnak/ysmrr/pkg/animations"
	"github.com/chelnak/ysmrr/pkg/colors"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var ApiDoctor = &cobra.Command{
	Use:   "doctor",
	Short: "doctor command for check the required dependencies",
	Long: `doctor command for check the required dependencies and
		software needed to run the application.
	Example:
		$ ./monorepo doctor
		$ ./monorepo doctor --help
    `,
	Run: doctor,
}

func doctor(cmd *cobra.Command, args []string) {
	cfg, err := config.LoadDoctorConfig()
	if err != nil {
		color.Red(err.Error())
		return
	}

	sm := ysmrr.NewSpinnerManager(
		ysmrr.WithAnimation(animations.Dots),
		ysmrr.WithSpinnerColor(colors.FgHiBlue),
	)

	color.Green("Checking dependencies...")
	check := sm.AddSpinner("Golang installation check...")

	sm.Start()
	time.Sleep(2 * time.Second)

	cmdGo := exec.Command("go", "version")
	output, err := cmdGo.CombinedOutput()

	if err != nil {
		check.UpdateMessagef("Go installation check... %s %s", "ERROR:", err.Error())
		check.Complete()

		sm.Stop()

		color.Red("Go is not installed. Please install Go and try again.")
		return
	}

	versionRegex := regexp.MustCompile(`go(\d+\.\d+\.\d+)`)
	matches := versionRegex.FindStringSubmatch(string(output))

	check.UpdateMessagef("Go installation check version is %s", matches[0])
	check.Complete()

	// check for postgres from config connection
	check = sm.AddSpinner("Postgres installation check...")
	time.Sleep(2 * time.Second)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Name)
	open, err := sql.Open("postgres", dsn)
	if err != nil {
		check.UpdateMessagef("Postgres installation check... %s %s", "ERROR:", err.Error())
		check.Complete()

		sm.Stop()

		color.Red("Postgres is not installed. Please install Postgres and try again.")
		return
	}

	err = open.Ping()
	if err != nil {
		check.UpdateMessagef("Postgres installation check... %s %s", "ERROR:", err.Error())
		check.Complete()

		sm.Stop()

		color.Red("Postgres is not installed. Please install Postgres and try again.")
		return
	}

	check.UpdateMessagef("Postgres installation check... %s", "OK")
	check.Complete()

	sm.Stop()

	color.Green("All dependencies are OK you can start the application")
}
