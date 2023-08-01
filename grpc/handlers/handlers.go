// Copyright © 2023 doomshrine and gocosi authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package handlers includes common HandlerFuncs that can be used around the gRPC environment.
package handlers

import (
	"context"
	"runtime/debug"

	"github.com/doomshrine/gocosi/grpc/log"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
)

// PanicRecovery returns handler of the panics, that logs the panic and call stack.
// It take optional argument called callbacks, that are functions, e.g. wrapping incrementing the panicMetric.
func PanicRecovery(logger *log.Logger, callbacks ...func(context.Context)) recovery.RecoveryHandlerFunc {
	return func(p any) (err error) {
		ctx := context.Background()

		for _, callback := range callbacks {
			callback(ctx)
		}

		if logger != nil {
			logger.Log(ctx, 0,
				"recovered from panic",
				"panic", p,
				"stack", debug.Stack())
		}

		return nil
	}
}
