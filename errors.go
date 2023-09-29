package gocosi

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrNilMux                   = errors.New("nil mux")
	ErrHealthcheckStatusUnknown = errors.New("healthcheck status unknown")
)

type ErrHealthCheckFailure struct {
	failures map[string]string
}

var _ error = (*ErrHealthCheckFailure)(nil)

func (err *ErrHealthCheckFailure) Error() string {
	reasons := []string{}

	for service, reason := range err.failures {
		reasons = append(reasons, fmt.Sprintf("%s (reason: '%s')", service, reason))
	}

	return "healthcheck failed: " + strings.Join(reasons, ", ")
}
