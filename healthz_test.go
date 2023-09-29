package gocosi

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/doomshrine/must"
	"github.com/doomshrine/testcontext"
	"github.com/hellofresh/health-go/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHealthcheckProbe(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name      string
		assertion func(assert.TestingT, error, ...interface{}) bool
		port      int
		healthz   *health.Health
	}{
		{
			name:      "default",
			assertion: assert.NoError,
			port:      30001,
			healthz:   must.Do(health.New()),
		},
		{
			name:      "error",
			assertion: assert.Error,
			port:      30002,
			healthz: must.Do(health.New(health.WithChecks(
				health.Config{
					Name:    "test",
					Timeout: time.Second,
					Check: func(ctx context.Context) error {
						return fmt.Errorf("forced error")
					},
				},
			))),
		},
	} {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx, cancel := testcontext.FromTimeout(context.Background(), t, time.Second)
			defer cancel()

			mux := http.NewServeMux()
			mux.Handle(HealthcheckEndpoint, tc.healthz.Handler())

			server := &http.Server{
				Addr:              fmt.Sprintf(":%d", tc.port),
				Handler:           mux,
				ReadTimeout:       1 * time.Second,
				WriteTimeout:      1 * time.Second,
				IdleTimeout:       30 * time.Second,
				ReadHeaderTimeout: 2 * time.Second,
			}
			defer server.Shutdown(ctx) //nolint:errcheck

			go func() {
				err := server.ListenAndServe()
				require.NoError(t, err)
			}()

			err := HealthcheckFunc(ctx, fmt.Sprintf("http://localhost:%d/healthz", tc.port))
			tc.assertion(t, err)
		})
	}
}
