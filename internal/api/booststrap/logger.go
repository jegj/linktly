package bootstrap

import (
	"log/slog"
	"os"
)

var LogLevels = map[string]slog.Level{
	"DEBUG": slog.LevelDebug,
	"INFO":  slog.LevelInfo,
	"WARN":  slog.LevelWarn,
	"ERROR": slog.LevelError,
}

func SetUpLogger(env *EnvVar) *slog.LevelVar {
	programLevel := new(slog.LevelVar)

	level, exists := LogLevels[env.LogLevel]
	if !exists {
		level = slog.LevelWarn
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: programLevel}))
	programLevel.Set(level)
	slog.SetDefault(logger)
	return programLevel
}
