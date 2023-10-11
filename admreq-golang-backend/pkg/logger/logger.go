package logger

import (
	"fmt"
	"os"

	"golang.org/x/exp/slog"
)

type Logger struct {
	*slog.Logger
}

func GetLogger(env string) (*Logger, error) {
	var log *slog.Logger
	switch env {
	case "dev":
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case "prod":
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelWarn}),
		)
	default:
		return nil, fmt.Errorf("Unspecified working environment")
	}

	return &Logger{log}, nil
}
