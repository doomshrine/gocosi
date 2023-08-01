package gocosi

import (
	"github.com/doomshrine/must"
	"go.opentelemetry.io/otel"
)

var (
	meter = otel.Meter("github.com/doomshrine/gocosi")

	panicsTotal = must.Do(meter.Int64Counter("grpc_req_panics_recovered_total"))
)
