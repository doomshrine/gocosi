// Package provisioner
package provisioner

import (
	"context"

	cosi "sigs.k8s.io/container-object-storage-interface-spec"
)

// Server implements cosi.ProvisionerServer interface.
type Server struct{}

// Interface guards.
var _ cosi.ProvisionerServer = (*Server)(nil)

// New returns provisioner.Server with default values.
func New() *Server {
	return &Server{}
}

// DriverCreateBucket call is made to create the bucket in the backend.
//
// NOTE: this call needs to be idempotent.
//  1. If a bucket that matches both name and parameters already exists, then OK (success) must be returned.
//  2. If a bucket by same name, but different parameters is provided, then the appropriate error code ALREADY_EXISTS must be returned.
func (s *Server) DriverCreateBucket(ctx context.Context, req *cosi.DriverCreateBucketRequest) (*cosi.DriverCreateBucketResponse, error) {
	// TODO: your implementation goes here.
	panic("unimplemented")
}

// DriverDeleteBucket call is made to delete the bucket in the backend.
//
// NOTE: this call needs to be idempotent.
// If the bucket has already been deleted, then no error should be returned.
func (s *Server) DriverDeleteBucket(ctx context.Context, req *cosi.DriverDeleteBucketRequest) (*cosi.DriverDeleteBucketResponse, error) {
	// TODO: your implementation goes here.
	panic("unimplemented")
}

// DriverGrantBucketAccess call grants access to an account.
// The account_name in the request shall be used as a unique identifier to create credentials.
//
// NOTE: this call needs to be idempotent.
// The account_id returned in the response will be used as the unique identifier for deleting this access when calling DriverRevokeBucketAccess.
// The returned secret does not need to be the same each call to achieve idempotency.
func (s *Server) DriverGrantBucketAccess(ctx context.Context, req *cosi.DriverGrantBucketAccessRequest) (*cosi.DriverGrantBucketAccessResponse, error) {
	// TODO: your implementation goes here.
	panic("unimplemented")
}

// DriverRevokeBucketAccess call revokes all access to a particular bucket from a principal.
//
// NOTE: this call needs to be idempotent.
func (s *Server) DriverRevokeBucketAccess(ctx context.Context, req *cosi.DriverRevokeBucketAccessRequest) (*cosi.DriverRevokeBucketAccessResponse, error) {
	// TODO: your implementation goes here.
	panic("unimplemented")
}
