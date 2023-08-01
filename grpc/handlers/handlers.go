package handlers

import (
	"context"
	"runtime/debug"

	"github.com/doomshrine/gocosi/grpc/log"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
)

func PanicRecovery(logger *log.Logger, callbacks ...func(context.Context)) recovery.RecoveryHandlerFunc {
	return func(p any) (err error) {
		ctx := context.Background()

		for _, callback := range callbacks {
			callback(ctx)
		}

		logger.Log(ctx, 0,
			"recovered from panic",
			"panic", p,
			"stack", debug.Stack())

		return nil
	}
}
