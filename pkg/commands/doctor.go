package commands

import (
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
	//cfg, err := config.LoadDoctorConfig()
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
		check.UpdateMessagef("Go installation check... %s %s", "ERROR:"+err.Error())
		check.Complete()

		sm.Stop()

		color.Red("Go is not installed. Please install Go and try again.")
		return
	}

	versionRegex := regexp.MustCompile(`go(\d+\.\d+\.\d+)`)
	matches := versionRegex.FindStringSubmatch(string(output))

	check.UpdateMessagef("Go installation check version is %s", matches[0])
	check.Complete()

	sm.Stop()

	color.Green("All dependencies are OK you can start the application")
}
