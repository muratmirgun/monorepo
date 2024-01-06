package commands

import (
	"github.com/spf13/cobra"
)

var ApiDoctor = &cobra.Command{
	Use:   "doctor",
	Short: "doctor command for check the required dependencies",
	Long: `doctor command for check the required dependencies and
		software needed to run the application.
	Example:
		$ ./go-cli doctor
		$ ./go-cli doctor --help
    `,
	Run: doctor,
}

func doctor(cmd *cobra.Command, args []string) {
	//cfg, err := config.LoadDoctorConfig()

}
