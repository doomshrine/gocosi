// Package identity
package identity

import (
	"context"

	"github.com/go-logr/logr"
	cosi "sigs.k8s.io/container-object-storage-interface-spec"
)

// Server implements cosi.IdentityServer interface.
type Server struct {
  log  logr.Logger
	name string
}

// Interface guards.
var _ cosi.IdentityServer = (*Server)(nil)

// New returns identitu.Server with name set to the "name" argument.
func New(name string, logger logr.Logger) *Server {
	if name == "" {
		panic("empty name")
	}

	return &Server{
		log: logger,
		name: name,
	}
}

// DriverGetInfo call is meant to retrieve the unique provisioner Identity.
func (s *Server) DriverGetInfo(_ context.Context, _ *cosi.DriverGetInfoRequest) (*cosi.DriverGetInfoResponse, error) {
	return &cosi.DriverGetInfoResponse{
		Name: s.name,
	}, nil
}
