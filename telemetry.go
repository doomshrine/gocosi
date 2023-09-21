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

package gocosi

import (
	"context"

	"github.com/doomshrine/must"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.opentelemetry.io/otel/sdk/trace"
)

var (
	DefaultMeter = otel.Meter("github.com/doomshrine/gocosi")

	PanicsTotal = must.Do(DefaultMeter.Int64Counter("grpc_req_panics_recovered_total"))
)

type TraceExporter interface {
	ExportSpans(ctx context.Context, ss []trace.ReadOnlySpan) error
	MarshalLog() interface{}
	Shutdown(ctx context.Context) error
	Start(ctx context.Context) error
}

var _ TraceExporter = (*otlptrace.Exporter)(nil)

type MetricExporter interface {
	Aggregation(k metric.InstrumentKind) metric.Aggregation
	Export(ctx context.Context, rm *metricdata.ResourceMetrics) error
	ForceFlush(ctx context.Context) error
	MarshalLog() interface{}
	Shutdown(ctx context.Context) error
	Temporality(k metric.InstrumentKind) metricdata.Temporality
}

var (
	_ MetricExporter = (*otlpmetricgrpc.Exporter)(nil)
	_ MetricExporter = (*otlpmetrichttp.Exporter)(nil)
)
