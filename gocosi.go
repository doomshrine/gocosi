package gocosi

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-logr/logr"
	"google.golang.org/grpc"
	cosi "sigs.k8s.io/container-object-storage-interface-spec"
)

// logger is a global instance of logr.Logger used in the gocosi.
var logger logr.Logger

// Driver.
type Driver struct {
	identity    cosi.IdentityServer
	provisioner cosi.ProvisionerServer

	endpoint    *Endpoint
	grpcOptions []grpc.ServerOption

	logger        logr.Logger
	otelCollector string
}

// Option.
type Option func(*Driver) error

// New.
func New(identity cosi.IdentityServer, provisioner cosi.ProvisionerServer, opts ...Option) (*Driver, error) {
	p := &Driver{
		identity:    identity,
		provisioner: provisioner,

		endpoint: &Endpoint{
			permissions: 0o755,
		},
	}

	var combinedErrors error

	for _, opt := range opts {
		if err := opt(p); err != nil {
			combinedErrors = errors.Join(combinedErrors, err)
		}
	}

	return p, combinedErrors
}

func SetLogger(log logr.Logger) {
	logger = log
}

// Run.
func (d *Driver) Run(ctx context.Context) error {
	// TODO: trap signals
	// TODO: configure listener
	// TODO: configure grpc server
	// TODO: start grpcServer
	lis, err := d.endpoint.Listener(ctx)
	if err != nil {
		return fmt.Errorf("listener creation failed: %w", err)
	}

	srv, err := d.grpcServer()
	if err != nil {
		return fmt.Errorf("gRPC server creation failed: %w", err)
	}

	err = srv.Serve(lis)
	if err != nil {
		return fmt.Errorf("gRPC server failed: %w", err)
	}

	panic("unimplemented")
}

func (d *Driver) grpcServer() (*grpc.Server, error) {
	server := grpc.NewServer(d.grpcOptions...)

	if d.provisioner == nil || d.identity == nil {
		return nil, errors.New("provisioner and identity servers cannot be nil")
	}

	cosi.RegisterIdentityServer(server, d.identity)
	cosi.RegisterProvisionerServer(server, d.provisioner)

	return server, nil
}
