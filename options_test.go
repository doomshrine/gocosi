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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithDefaultMetricExporter(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name string
		kind ExporterKind
	}{
		{
			name: "default HTTP metric exporter",
			kind: HTTPExporter,
		},
		{
			name: "default GRPC metric exporter",
			kind: GRPCExporter,
		},
	} {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			opt := WithDefaultMetricExporter(tc.kind)

			d := &Driver{}

			err := opt(d)
			assert.NoError(t, err)
		})
	}
}

func TestWithHTTPMetricExporter(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name string
	}{} {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
		})
	}
}

func TestWithGRPCMetricExporter(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name string
	}{} {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
		})
	}
}

func TestWithDefaultTraceExporter(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name string
		kind ExporterKind
	}{
		{
			name: "default HTTP trace exporter",
			kind: HTTPExporter,
		},
		{
			name: "default GRPC trace exporter",
			kind: GRPCExporter,
		},
	} {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			opt := WithDefaultTraceExporter(tc.kind)

			d := &Driver{}

			err := opt(d)
			assert.NoError(t, err)
		})
	}
}

func TestWithHTTPTraceExporter(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name string
	}{} {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
		})
	}
}

func TestWithGRPCTraceExporter(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name string
	}{} {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
		})
	}
}
