# gocosi

[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-2.1-4baaaa.svg)](CODE_OF_CONDUCT.md)
[![GitHub](https://img.shields.io/github/license/doomshrine/gocosi)](LICENSE.txt)
[![Go Report Card](https://goreportcard.com/badge/github.com/doomshrine/gocosi)](https://goreportcard.com/report/github.com/doomshrine/gocosi)
[![codecov](https://codecov.io/gh/doomshrine/gocosi/branch/main/graph/badge.svg?token=UGGL63A65L)](https://codecov.io/gh/doomshrine/gocosi)
[![GitHub Workflow Status (with event)](https://img.shields.io/github/actions/workflow/status/doomshrine/gocosi/tests.yaml)](https://github.com/shanduur/gocosi/actions/workflows/tests.yaml)
[![Static Badge](https://img.shields.io/badge/COSI_Specification-v1alpha1-green)](https://github.com/kubernetes-sigs/container-object-storage-interface-spec/tree/v0.1.0)

A Container Object Storage Interface (COSI) library and other helpful utilities created with Go, that simplify process of writing your own Object Storage Plugin (OSP).

## Table of contents

- [gocosi](#gocosi)
  - [Table of contents](#table-of-contents)
  - [Quick Start](#quick-start)
  - [Features](#features)
    - [Automatic Starting of Traces](#automatic-starting-of-traces)
    - [Automatic Panic Recovery](#automatic-panic-recovery)
    - [Control Over Permissions by Default](#control-over-permissions-by-default)
    - [Configuration via Environment](#configuration-via-environment)
    - [Supports Any Logger Through logr](#supports-any-logger-through-logr)
  - [Configuration](#configuration)
    - [Environment variables](#environment-variables)
  - [Contributing](#contributing)
  - [Similar projects](#similar-projects)

## Quick Start

The following example illustrates using `gocosi` Object Storage Plugin bootstrapper to create a new COSI Object Storage Plugin from scratch:

```bash
go run \
  github.com/doomshrine/gocosi/cmd/bootstrap@main \
  -module example.com/your/cosi-osp \
  -dir cosi-osp
```

You will obtain the following file structure in the `cosi-osp` folder:

```
cosi-osp
├── go.mod
├── go.sum
├── main.go
└── servers
    ├── identity
    │   └── identity.go
    └── provisioner
        └── provisioner.go

4 directories, 5 files
```

## Features

### Automatic Starting of Traces

The gocosi package includes an automatic trace initialization feature using the `otelgrpc.UnaryServerInterceptor()` function. This functionality enables the automatic creation of traces for each incoming request. Alongside the tracing capability, this interceptor adds detailed events to the traces, enhancing the visibility into the request lifecycle.

### Automatic Panic Recovery

In the event of a panic occurring during the execution of your COSI driver, the package offers seamless recovery through the use of `recovery.UnaryServerInterceptor`. This interceptor employs a custom panic handler that captures the panic message and the associated debug stack trace. Every panic is logged for thorough error analysis. Additionally, the package provides a default OpenTelemetry (OTEL) Counter metric that tracks the occurrence of panics, aiding in monitoring and troubleshooting.

### Control Over Permissions by Default

The gocosi package introduces a built-in mechanism designed to grant users control over the permissions and access rights associated with the UNIX socket. This inherent feature empowers you to define specific user and group ownership for the socket, along with customizable permission settings. This level of control ensures that the socket's security and accessibility align with your application's requirements.

### Configuration via Environment

Streamlining the configuration process, the gocosi package exclusively employs environment variables for end-user configuration. This design choice simplifies setup and customization, as users can conveniently adjust various aspects of the COSI driver's behavior by modifying environment variables.

### Supports Any Logger Through logr

To cater to diverse logging preferences, the gocosi package seamlessly integrates with a wide range of loggers by leveraging the versatile [logr](github.com/go-logr/logr) library. This compatibility enables you to use your preferred logger implementation, ensuring consistency with your existing logging practices and enhancing your debugging and monitoring capabilities.

## Configuration

### Environment variables

`gocosi` defines several constants used to specify environment variables for configuring the COSI (Container Object Storage Interface) endpoint. The COSI endpoint is used to interact with an object storage plugin in a containerized environment.

- `COSI_ENDPOINT` (default: `unix:///var/lib/cosi/cosi.sock`) - it should be set to the address or location of the COSI endpoint, such as the URL or file path.
- `X_COSI_ENDPOINT_PERMS` (default: `0755`) - it should be set when the COSI endpoint is a UNIX socket file. It determines the file permissions (in octal notation) of the UNIX socket file. If the COSI endpoint is a TCP socket, this setting has no effect.
- `X_COSI_ENDPOINT_USER`  (default: *The user that starts the process*) - it should be set when the COSI endpoint is a UNIX socket file. It determines the owner (user) of the UNIX socket file. If the COSI endpoint is a TCP socket, this setting has no effect.
- `X_COSI_ENDPOINT_GROUP` (default: *The group that starts the process*) - it should be set when the COSI endpoint is a UNIX socket file. It determines the group ownership of the UNIX socket file. If the COSI endpoint is a TCP socket, this setting has no effect.

## Contributing

You want to contibute? Hop into the [CONTRIBUTING.md](CONTRIBUTING.md) and find how you can help the project!

## Similar projects

This project was inspired by [dell/gocsi](github.com/dell/gocsi).
