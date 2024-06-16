package logger

import (
	"context"

	"github.com/charmbracelet/log"
)

type loggerCtxKey struct{}

func ToContext(ctx context.Context, logger *Logger) context.Context {
	return context.WithValue(ctx, loggerCtxKey{}, logger)
}

func FromContext(ctx context.Context) *Logger {
	logger, ok := ctx.Value(loggerCtxKey{}).(*Logger)
	if !ok || logger == nil {
		return &Logger{
			Logger: log.Default(),
		}
	}
	return logger
}
