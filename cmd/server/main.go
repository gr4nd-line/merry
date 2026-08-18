package main

import (
	"os"

	httpmerry "github.com/gr4nd-line/merry/internal/http"
	"github.com/gr4nd-line/merry/internal/logevents"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	e := echo.New()

	zlog := zerolog.New(os.Stdout).
		With().
		Timestamp().
		Logger()

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:       true,
		LogStatus:    true,
		LogMethod:    true,
		LogLatency:   true,
		LogRequestID: true,
		LogRemoteIP:  true,

		LogValuesFunc: func(c *echo.Context, v middleware.RequestLoggerValues) error {
			log.Info().
				Str("method", v.Method).
				Str("uri", v.URI).
				Str("ip", v.RemoteIP).
				Int("status", v.Status).
				Dur("latency", v.Latency).
				Str("request_id", v.RequestID).
				Msg("request")

			return nil
		},
	}))
	e.Use(middleware.Recover())

	httpmerry.AddRoutes(e)

	zlog.Info().Str("event", logevents.START_SERVER).Msg("starting server")
	if err := e.Start(":8081"); err != nil {
		zlog.Error().Str("event", logevents.START_SERVER).Err(err).Msg("failed to start server")
	}
}
