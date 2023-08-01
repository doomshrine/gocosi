package log

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
)

type Logger struct {
	LoggerImpl logr.Logger
}

// Interface guard.
var _ logging.Logger = (*Logger)(nil)

func (l *Logger) Log(_ context.Context, level logging.Level, msg string, keysAndValues ...any) {
	l.LoggerImpl.V(int(level)).Info(msg, keysAndValues...)
}
