package gocosi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hellofresh/health-go/v5"
)

const (
	// HealthcheckEndpoint is the HTTP endpoint path for the healthcheck service.
	HealthcheckEndpoint = "/healthz"

	// HealthcheckAddr.
	HealthcheckAddr = "http://localhost:8080" + HealthcheckEndpoint
)

// HealthcheckFunc.
func HealthcheckFunc(ctx context.Context, addr string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, addr, nil)
	if err != nil {
		return fmt.Errorf("unable to create new request: %w", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("healthcheck failed: %w", err)
	}
	defer res.Body.Close()

	c := &health.Check{}

	err = json.NewDecoder(res.Body).Decode(c)
	if err != nil {
		return fmt.Errorf("unable to decode healthcheck response: %w", err)
	}

	log.Info("healthcheck finished",
		"status", c.Status,
		"system", c.System,
		"failures", c.Failures,
		"component", c.Component,
	)

	switch res.StatusCode {
	case http.StatusOK:
		return nil

	case http.StatusServiceUnavailable:
		return &ErrHealthCheckFailure{failures: c.Failures}

	default:
		return ErrHealthcheckStatusUnknown
	}
}
