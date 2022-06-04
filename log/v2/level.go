package log

import (
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

// Error returns a logger that includes a Key/ErrorValue pair.
func Error(logger log.Logger) log.Logger {
	return With(logger, level.Key(), level.ErrorValue())
}

// Warn returns a logger that includes a Key/WarnValue pair.
func Warn(logger log.Logger) log.Logger {
	return With(logger, level.Key(), level.WarnValue())
}

// Info returns a logger that includes a Key/InfoValue pair.
func Info(logger log.Logger) log.Logger {
	return With(logger, level.Key(), level.InfoValue())
}

// Debug returns a logger that includes a Key/DebugValue pair.
func Debug(logger log.Logger) log.Logger {
	return With(logger, level.Key(), level.DebugValue())
}
