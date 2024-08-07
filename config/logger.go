package config

import "log/slog"

type LoggerConfigurator interface {
	ConfigureLogger(logLevel string)
}

type SlogConfigurator struct{}

func NewSlogConfigurator() LoggerConfigurator {
	return &SlogConfigurator{}
}

func (c *SlogConfigurator) ConfigureLogger(logLevel string) {
	switch logLevel {
	case "debug":
		slog.SetLogLoggerLevel(slog.LevelDebug)
	case "info":
		slog.SetLogLoggerLevel(slog.LevelInfo)
	default:
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}
}
