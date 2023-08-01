package gocosi

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"os/signal"
	"syscall"

	"github.com/doomshrine/must"
	"github.com/go-logr/logr"
	"google.golang.org/grpc"
	cosi "sigs.k8s.io/container-object-storage-interface-spec"
)

// log is a global instance of logr.Logger used in the gocosi.
var log logr.Logger

// Driver TODO: write description.
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

// New TODO: write description.
func New(identity cosi.IdentityServer, provisioner cosi.ProvisionerServer, opts ...Option) (*Driver, error) {
	p := &Driver{
		identity:    identity,
		provisioner: provisioner,

		endpoint: &Endpoint{
			permissions: 0o660,
			address:     must.Do(url.Parse(cosiSocket)),
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

// SetLogger is used to to set the default global logger for the gocosi library.
func SetLogger(l logr.Logger) {
	log = l
}

// Run TODO: write description.
func (d *Driver) Run(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx,
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer cancel()

	lis, err := d.endpoint.Listener(ctx)
	if err != nil {
		return fmt.Errorf("listener creation failed: %w", err)
	}
	defer d.endpoint.Close()

	srv, err := d.grpcServer()
	if err != nil {
		return fmt.Errorf("gRPC server creation failed: %w", err)
	}

	go func() {
		<-ctx.Done()
		srv.GracefulStop()
	}()

	err = srv.Serve(lis)
	if err != nil {
		return fmt.Errorf("gRPC server failed: %w", err)
	}

	return nil
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