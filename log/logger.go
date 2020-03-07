package log

import (
	"io"

	"github.com/sirupsen/logrus"
)

const AllLoggers = ""

type Logger struct {
	*logrus.Logger
	name string
}

func (l *Logger) Name() string {
	return l.name
}

func (l *Logger) SetOutput(name string, output io.Writer) *Logger {
	l.Logger.SetOutput(output)
	return l
}

func (l *Logger) SetFormatter(name string, formatter Formatter) *Logger {
	l.Logger.SetFormatter(formatter)
	return l
}

func (l *Logger) SetLevel(name string, level Level) *Logger {
	l.Logger.SetLevel(level)
	return l
}

func (l *Logger) AddHook(name string, hook Hook) *Logger {
	l.Logger.AddHook(hook)
	return l
}

func (l *Logger) GetLogger(name string) *Logger {
	logger := NewLogger(name)
	logger.AddHook(name, NewLoggerNameHook(name))
	return logger
}

func NewLogger(name string) *Logger {
	return &Logger{
		Logger: logrus.New(),
		name:   name,
	}
}

func NewLoggerWithHooks(name string, hooks ...Hook) (logger *Logger) {
	logger = &Logger{
		Logger: logrus.New(),
		name:   name,
	}
	for _, hook := range hooks {
		logger.AddHook(AllLoggers, hook)
	}
	return
}
