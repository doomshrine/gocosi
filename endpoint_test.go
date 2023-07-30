package gocosi

import (
	"context"
	"net/url"
	"testing"

	"github.com/doomshrine/gocosi/testutils"
	"github.com/doomshrine/must"
	"github.com/doomshrine/testcontext"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncpointListener(t *testing.T) {
	t.Parallel()

	ctx, cancel := testcontext.FromT(context.Background(), t)
	defer cancel()

	for _, tc := range []struct {
		name           string
		endpoint       *Endpoint
		errorAssertion func(assert.TestingT, error, ...interface{}) bool
		valueAssertion func(assert.TestingT, interface{}, ...interface{}) bool
	}{
		{
			name:           "valid UNIX socket",
			errorAssertion: assert.NoError,
			valueAssertion: assert.NotNil,
			endpoint: &Endpoint{
				address: testutils.MustMkUnixTemp("cosi.sock"),
			},
		},
		{
			name:           "valid TCP socket",
			errorAssertion: assert.NoError,
			valueAssertion: assert.NotNil,
			endpoint: &Endpoint{
				address: must.Do(url.Parse("tcp://:0")),
			},
		},
	} {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			require.NotNil(t, tc.endpoint, "endpoint is required")
			defer tc.endpoint.Close() // no matter what, I need it to be called

			lis, err := tc.endpoint.Listener(ctx)
			tc.errorAssertion(t, err)
			tc.valueAssertion(t, lis)

			lis, err = tc.endpoint.Listener(ctx)
			tc.errorAssertion(t, err)
			tc.valueAssertion(t, lis)
		})
	}
}
