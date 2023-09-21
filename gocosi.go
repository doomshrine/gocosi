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

// log is a global instance of logr.Logger used in the gocosi package.
var log logr.Logger

// Driver represents a COSI driver implementation.
type Driver struct {
	identity    cosi.IdentityServer
	provisioner cosi.ProvisionerServer

	traceexporter  TraceExporter
	metricexporter MetricExporter

	endpoint    *Endpoint
	grpcOptions []grpc.ServerOption

	logger        logr.Logger
	otelCollector string
}

// Option represents a functional option to configure the Driver.
type Option func(*Driver) error

// New creates a new instance of the COSI driver.
func New(identity cosi.IdentityServer, provisioner cosi.ProvisionerServer, opts ...Option) (*Driver, error) {
	p := &Driver{
		identity:    identity,
		provisioner: provisioner,

		endpoint: &Endpoint{
			permissions: 0o755,
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

// SetLogger is used to set the default global logger for the gocosi library.
func SetLogger(l logr.Logger) {
	log = l
}

// Run starts the COSI driver and serves requests.
func (d *Driver) Run(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
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

	log.V(4).Info("starting driver", "address", lis.Addr())

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
