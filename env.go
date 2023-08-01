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

	// EnvOTELCollectorEndpoint is ... TODO: continue the description.
	EnvOTELCollectorEndpoint = "X_OPEN_TELEMETRY_COLLECTOR_ENDPOINT"
)

const (
	cosiSocket = "unix:///var/lib/cosi/cosi.sock"
)
