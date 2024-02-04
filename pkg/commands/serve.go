package commands

import (
	"github.com/labstack/echo/v4"
	"github.com/muratmirgun/monorepo/app"
	"github.com/muratmirgun/monorepo/config"
	_ "github.com/muratmirgun/monorepo/docs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	echoSwagger "github.com/swaggo/echo-swagger"
	"os"
	"os/signal"
	"strconv"
	"syscall"
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

// @title Saas API
// @description Saas API is a REST API for the Saas application.
// @version 1.0.0
//
// @BasePath /api/v1
//
// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
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

	e := echo.New()

	e.GET("/docs/*", echoSwagger.WrapHandler)

	s, _ := app.NewServices(cfg)
	app.RegisterRoutes(e, s)

	go func() {
		info := e.Routes()
		for _, i := range info {
			log.Info().Str("method", i.Method).Str("path", i.Path).Msg("registered route")
		}
		e.Start(":" + cfg.Server.Port)
	}()

	// wait for interrupt signal to gracefully shut down the server.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Info().Msg("Shutting down application...")
	if err := e.Shutdown(cmd.Context()); err != nil {
		log.Fatal().Err(err).Msg("Server forced to shutdown")
	}
}
