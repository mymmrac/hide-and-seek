package logger

import (
	"context"

	"github.com/charmbracelet/log"
)

type loggerCtxKey struct{}

func ToContext(ctx context.Context, logger *log.Logger) context.Context {
	return context.WithValue(ctx, loggerCtxKey{}, logger)
}

func FromContext(ctx context.Context) *log.Logger {
	logger, ok := ctx.Value(loggerCtxKey{}).(*log.Logger)
	if !ok || logger == nil {
		return log.Default()
	}
	return logger
}
