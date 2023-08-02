# gocosi

[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-2.1-4baaaa.svg)](CODE_OF_CONDUCT.md)
[![GitHub](https://img.shields.io/github/license/doomshrine/gocosi)](LICENSE.txt)
[![Go Report Card](https://goreportcard.com/badge/github.com/doomshrine/gocosi)](https://goreportcard.com/report/github.com/doomshrine/gocosi)
[![codecov](https://codecov.io/gh/shanduur/gocosi/branch/main/graph/badge.svg?token=8NA6HYH122)](https://codecov.io/gh/shanduur/gocosi)
[![GitHub Workflow Status (with event)](https://img.shields.io/github/actions/workflow/status/doomshrine/gocosi/tests.yaml)](https://github.com/shanduur/gocosi/actions/workflows/tests.yaml)
[![Static Badge](https://img.shields.io/badge/COSI_Specification-v1alpha1-green)](https://github.com/kubernetes-sigs/container-object-storage-interface-spec/tree/v0.1.0)

A Container Object Storage Interface (COSI) library and other helpful utilities created with Go, that simplify process of writing your own Object Storage Plugin (OSP).

## Table of contents

- [gocosi](#gocosi)
  - [Table of contents](#table-of-contents)
  - [Quick Start](#quick-start)
  - [Configuration](#configuration)
    - [Environment variables](#environment-variables)

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

## Configuration

### Environment variables

`gocosi` defines several constants used to specify environment variables for configuring the COSI (Container Object Storage Interface) endpoint. The COSI endpoint is used to interact with an object storage plugin in a containerized environment.

- `COSI_ENDPOINT` (default: `unix:///var/lib/cosi/cosi.sock`) - it should be set to the address or location of the COSI endpoint, such as the URL or file path.
- `X_COSI_ENDPOINT_PERMS` (default: `0755`) - it should be set when the COSI endpoint is a UNIX socket file. It determines the file permissions (in octal notation) of the UNIX socket file. If the COSI endpoint is a TCP socket, this setting has no effect.
- `X_COSI_ENDPOINT_USER`  (default: *The user that starts the process*) - it should be set when the COSI endpoint is a UNIX socket file. It determines the owner (user) of the UNIX socket file. If the COSI endpoint is a TCP socket, this setting has no effect.
- `X_COSI_ENDPOINT_GROUP` (default: *The group that starts the process*) - it should be set when the COSI endpoint is a UNIX socket file. It determines the group ownership of the UNIX socket file. If the COSI endpoint is a TCP socket, this setting has no effect.
