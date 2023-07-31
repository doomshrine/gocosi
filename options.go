package gocosi

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"os/user"

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
func WithDefaultGRPCOptions() Option {
	return func(d *Driver) error {
		d.grpcOptions = []grpc.ServerOption{}

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
