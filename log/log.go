package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	l *Logger
)

func Set(out *Logger) {
	l = out
}

func L() *Logger {
	return l
}

// nolint:gochecknoinits
func init() {
	l, _ = NewLogger(WarnLevel)
}

type Logger struct {
	*zap.Logger
}

func NewLogger(lvl Level) (*Logger, error) {
	var logger *zap.Logger
	var err error
	if os.Getenv("LOGLEVEL") != "" {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction(zap.IncreaseLevel(zapcore.ErrorLevel))
	}

	return &Logger{
		Logger: logger,
	}, err
}

func (l *Logger) Named(s string) *Logger {
	return &Logger{
		Logger: l.Logger.Named(s),
	}
}
