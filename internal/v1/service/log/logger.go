package log

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

func InitializeLogger(logLevel zerolog.Level) *zerolog.Logger {
	prettyLogFormat := os.Getenv("PRETTY_LOG_FORMAT")
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	if prettyLogFormat == "true" {
		output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
		logger = zerolog.New(output).With().Timestamp().Logger()
	}

	zerolog.SetGlobalLevel(logLevel)
	return &logger
}
