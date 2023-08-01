package gocosi

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"os"
	"os/user"

	"github.com/doomshrine/gocosi/grpc/handlers"
	grpclog "github.com/doomshrine/gocosi/grpc/log"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

// WithCOSIEndpoint overrides the default COSI endpoint.
func WithCOSIEndpoint(url *url.URL) Option {
	return func(d *Driver) error {
		if url.Scheme != SchemeUNIX && url.Scheme != SchemeTCP {
			return errors.New("scheme should be either unix or tcp")
		}

		d.endpoint.address = url

		return nil
	}
}

// WithSocketPermissions is used to override default permissions (0o660).
// Permissions that are being set must be between:
//   - 0o600 - the minimum permissions
//   - 0o666 - the maximum permissions
func WithSocketPermissions(perm os.FileMode) Option {
	return func(d *Driver) error {
		const (
			minPermissions os.FileMode = 0o600
			maxPermissions os.FileMode = 0o666
		)

		if perm < minPermissions || perm > maxPermissions {
			return fmt.Errorf("permissions out of range, minimum %d, maximum %d",
				minPermissions,
				maxPermissions)
		}

		d.endpoint.permissions = perm

		return nil
	}
}

// WithSocketUser is used to override default user owning the socket (current user).
func WithSocketUser(user *user.User) Option {
	return func(d *Driver) error {
		d.endpoint.user = user
		return nil
	}
}

// WithSocketGroup is used to override default group owning the socket (current user's group).
func WithSocketGroup(group *user.Group) Option {
	return func(d *Driver) error {
		d.endpoint.group = group
		return nil
	}
}

// WithGRPCOptions overrides all previously applied gRPC ServerOptions by a default options.
//
// Default gRPC SeverOptions are:
// - ChainUnaryInterceptor - consists of:
//   - grpc.UnaryServerInterceptor() - starts and configures tracer for each request,
//     records events for request and response (error is recorded as normal event);
//   - logging.UnaryServerInterceptor() - records and logs according to the global logger (wrapped around grpc/log.Logger);
//   - recovery.UnaryServerInterceptor() - records metric for panics, and recovers (a log is created for each panic);
func WithDefaultGRPCOptions() Option {
	return func(d *Driver) error {
		log := &grpclog.Logger{LoggerImpl: log}

		d.grpcOptions = []grpc.ServerOption{
			grpc.ChainUnaryInterceptor(
				otelgrpc.UnaryServerInterceptor(),
				logging.UnaryServerInterceptor(log),
				recovery.UnaryServerInterceptor(recovery.WithRecoveryHandler(handlers.PanicRecovery(log,
					func(ctx context.Context) { panicsTotal.Add(ctx, 1) }))),
			),
		}

		return nil
	}
}

// WithGRPCOptions overrides all previously applied gRPC ServerOptions by a set provided as argument to this call.
func WithGRPCOptions(opts ...grpc.ServerOption) Option {
	return func(d *Driver) error {
		var (
			joinedErrors error
			grpcOptions  []grpc.ServerOption
		)

		for _, opt := range opts {
			if opt != nil {
				joinedErrors = errors.Join(joinedErrors, errors.New("nil option provided"))
			} else {
				grpcOptions = append(grpcOptions, opt)
			}
		}

		d.grpcOptions = grpcOptions

		return joinedErrors
	}
}
