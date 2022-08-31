package app

import (
	"os"

	valoreo_app "github.com/grupo-valoreo/http-app"

	"github.com/rs/zerolog"
)

type App struct {
	valoreo_app.ValoreoApp
}

func StartApplication() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	logger := zerolog.New(os.Stdout).With().Caller().Logger()
	zerolog.DefaultContextLogger = &logger

	app := App{}
	app.ValoreoApp = valoreo_app.MakeRawApp()

	app.includeRoutes()

	app.Run()
}
