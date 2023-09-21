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
	"testing"
	"time"

	"github.com/doomshrine/gocosi/mock"
	"github.com/doomshrine/gocosi/testutils"
	"github.com/doomshrine/must"
	"github.com/doomshrine/testcontext"
	"github.com/go-logr/logr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	cosi "sigs.k8s.io/container-object-storage-interface-spec"
)

//nolint:paralleltest
func TestSetLogger(t *testing.T) {
	impl := logr.New(logr.Discard().GetSink())
	SetLogger(impl)
	assert.Equal(t, impl, log)
}

func TestNew(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name        string
		identity    cosi.IdentityServer
		provisioner cosi.ProvisionerServer
		options     []Option
	}{} {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			require.NotNil(t, tc.identity)
			require.NotNil(t, tc.provisioner)
			require.NotNil(t, tc.identity)

			_, err := New(
				tc.identity,
				tc.provisioner,
				nil,
				tc.options...,
			)
			assert.NoError(t, err)
		})
	}
}

func TestRun(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name   string
		driver *Driver
	}{
		{
			name: "run with defaults",
			driver: must.Do(func() (*Driver, error) {
				identity := &mock.MockCOSIIdentityServer{}
				provisioner := &mock.MockCOSIProvisionerServer{}

				return New(
					identity,
					provisioner,
					nil,
					WithCOSIEndpoint(testutils.MustMkUnixTemp("cosi.sock")),
				)
			}()),
		},
	} {
		tc := tc

		ctx, cancel := testcontext.FromT(context.Background(), t)
		defer cancel()

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctx, cancel := testcontext.FromTimeout(ctx, t, time.Second)
			defer cancel()

			require.NotNil(t, tc.driver)

			err := tc.driver.Run(ctx)
			assert.NoError(t, err)
		})
	}
}
