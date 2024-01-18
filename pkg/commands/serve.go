package commands

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"monorepo/config"
	"os"
	"strconv"
)

var ApiServe = &cobra.Command{
	Use:   "serve",
	Short: "serve command for run the application",
	Long: `serve command for run the application.
	Example:
		$ ./monorepo serve
		$ ./monorepo serve --help
	`,
	Run: serve,
}

func serve(cmd *cobra.Command, args []string) {
	// init config
	cfg, err := config.LoadConfig("")
	if err != nil {
		panic(err)
		return
	}

	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		short := file
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				short = file[i+1:]
				break
			}
		}
		file = short
		return file + ":" + strconv.Itoa(line)
	}
	log.Logger = log.With().Caller().Logger()
	if cfg.Env == config.DEV {
		log.Level(zerolog.DebugLevel)
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	} else {
		log.Level(zerolog.InfoLevel)
	}

	log.Info().Msg("Starting application...")

}
