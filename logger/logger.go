package logger

import (
	"log/slog"
	"os"
)

// Declare loggers
var logger slog.Logger = SetupLogger()

func SetupLogger() slog.Logger {
	// Defined loggers
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	return *logger
}

func GetLogger() slog.Logger {
	return logger
}
