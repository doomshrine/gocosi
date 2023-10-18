// Copyright © 2023 gocosi authors. All Rights Reserved.
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

package log

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"go.opentelemetry.io/otel"
)

type Logger struct {
	LoggerImpl logr.Logger
}

// Interface guards.
var (
	_ logging.Logger    = (*Logger)(nil)
	_ otel.ErrorHandler = (*Logger)(nil)
)

func (l *Logger) Log(_ context.Context, level logging.Level, msg string, keysAndValues ...any) {
	l.LoggerImpl.V(int(level)).Info(msg, keysAndValues...)
}

func (l *Logger) Handle(err error) {
	l.LoggerImpl.Error(err, "Caught error.")
}
