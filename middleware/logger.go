package middleware

import (
	"log/slog"
	"os"
)

// for  creating new loggers

func NewLogger() *slog.Logger {
	return slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}),
	)
}
