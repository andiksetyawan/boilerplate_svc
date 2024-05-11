package logger

import (
	"github.com/andiksetyawan/boilerplate_svc/internal/config"
	"github.com/andiksetyawan/log"
	"github.com/andiksetyawan/log/slog"
)

func NewSlog(config config.Config) log.Logger {
	level := slog.LevelInfo
	if config.ServiceEnv != "production" {
		level = slog.LevelDebug
	}

	logger, err := slog.New(slog.WithLevel(level))
	if err != nil {
		panic(err)
	}

	return logger
}
