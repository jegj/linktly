package api

import (
	"log/slog"
	"os"

	"github.com/jegj/linktly/internal/config"
)

var LogLevels = map[string]slog.Level{
	"DEBUG": slog.LevelDebug,
	"INFO":  slog.LevelInfo,
	"WARN":  slog.LevelWarn,
	"ERROR": slog.LevelError,
}

func SetUpLogger(cfg config.Config) *slog.LevelVar {
	programLevel := new(slog.LevelVar)

	level, exists := LogLevels[cfg.LogLevel]
	if !exists {
		level = slog.LevelWarn
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: programLevel}))
	programLevel.Set(level)
	slog.SetDefault(logger)
	return programLevel
}
