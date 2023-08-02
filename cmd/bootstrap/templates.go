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

package main

const mainGo = `package main

import (
	"context"
	"os"

	"github.com/doomshrine/gocosi"
	"github.com/go-logr/logr"


	"{{ .ModPath }}/servers/identity"
	"{{ .ModPath }}/servers/provisioner"
)

var (
	driverName = "cosi.example.com" // FIXME: replace with your own driver name

	log logr.Logger
)

func init() {
	// Setup your logger here.
	// You can use one of multiple available implementation, like:
	//   - https://github.com/kubernetes/klog/tree/main/klogr
	//   - https://github.com/go-logr/stdr
	//   - https://github.com/bombsimon/logrusr
}

func main() {
	gocosi.SetLogger(log)

	// If there is any additional confifuration needed for your COSI Driver,
	// put it below this line.

	driver, err := gocosi.New(
		identity.New(driverName),
		provisioner.New(),
		gocosi.WithDefaultGRPCOptions(),
	)
	if err != nil {
		log.Error(err, "failed to create COSI Driver")
		os.Exit(1)
	}

	if err := driver.Run(context.Background()); err != nil {
		log.Error(err, "failed to run COSI Driver")
		os.Exit(1)
	}
}
`

const identityGo = `package identity

import (
	"context"

	cosi "sigs.k8s.io/container-object-storage-interface-spec"
)

// Server implements cosi.IdentityServer interface.
type Server struct {
	name string
}

// Interface guards.
var _ cosi.IdentityServer = (*Server)(nil)

// New returns identitu.Server with name set to the "name" argument.
func New(name string) *Server {
	if name == "" {
		panic("empty name")
	}

	return &Server{
		name: name,
	}
}

// DriverGetInfo call is meant to retrieve the unique provisioner Identity.
func (s *Server) DriverGetInfo(_ context.Context, _ *cosi.DriverGetInfoRequest) (*cosi.DriverGetInfoResponse, error) {
	return &cosi.DriverGetInfoResponse{
		Name: s.name,
	}, nil
}
`

const provisionerGo = `package provisioner

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
`

const goMod = `module {{ .ModPath }}

go {{ .GoVersion }}
`
