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

func TestSetLogger(t *testing.T) {
	t.Parallel()

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
