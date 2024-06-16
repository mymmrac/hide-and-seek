package logger

import "github.com/charmbracelet/log"

type Logger struct {
	*log.Logger
}

func (l *Logger) With(keyValues ...any) *Logger {
	return &Logger{
		Logger: l.Logger.With(keyValues...),
	}
}
