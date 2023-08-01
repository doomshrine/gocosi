// Copyright © 2023 Dell Inc. or its subsidiaries. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gocosi

const (
	// CosiEndpoint is the name of the environment variable used to
	// specify the COSI endpoint.
	EnvCOSIEndpoint = "COSI_ENDPOINT"

	// EnvVarEndpointPerms is the name of the environment variable used
	// to specify the file permissions for the COSI endpoint when it is
	// a UNIX socket file. This setting has no effect if COSI_ENDPOINT
	// specifies a TCP socket. The default value is 0755.
	EnvCOSIEndpointPerms = "X_COSI_ENDPOINT_PERMS"

	// EnvVarEndpointUser is the name of the environment variable used
	// to specify the UID or name of the user that owns the endpoint's
	// UNIX socket file. This setting has no effect if COSI_ENDPOINT
	// specifies a TCP socket. The default value is the user that starts
	// the process.
	EnvCOSIEndpointUser = "X_COSI_ENDPOINT_USER"

	// EnvVarEndpointGroup is the name of the environment variable used
	// to specify the GID or name of the group that owns the endpoint's
	// UNIX socket file. This setting has no effect if COSI_ENDPOINT
	// specifies a TCP socket. The default value is the group that starts
	// the process.
	EnvCOSIEndpointGroup = "X_COSI_ENDPOINT_GROUP"
)

const (
	// cosiSocket is default location of COSI UNIX socket.
	cosiSocket = "unix:///var/lib/cosi/cosi.sock"
)
