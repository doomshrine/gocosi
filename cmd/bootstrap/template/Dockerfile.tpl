###############################################################################
# First stage:
#   - Building the COSI OSP using the default Go image.
###############################################################################
FROM --platform=${BUILDPLATFORM} docker.io/library/golang:1.21.1 AS builder

# Set the working directory.
WORKDIR /cosi-osp

# Copy the Go Modules manifests.
COPY go.mod go.mod
COPY go.sum go.sum

# Cache deps before building and copying source, so that we don't need to re-download
# as much and so theat source changes don't invalidate our downloaded layer.
RUN go mod download

# Copy the source.
COPY main.go main.go
COPY servers/ servers/
# TODO: add any packages you want to include.

# Disable CGO.
ENV CGO_ENABLED=0

# Build the image.
RUN go build -o build/cosi-osp main.go

###############################################################################
# Second Stage:
#   - Runtime image.
#
# NOTE: you should replace the latest with specific digest, to ensure that
#       builds are consistent.
###############################################################################
FROM --platform=${BUILDPLATFORM} gcr.io/distroless/static:latest AS runtime

# Set the working directory.
WORKDIR /cosi

# Set volume mount point for COSI socket.
VOLUME [ "/var/lib/cosi" ]

# TODO: add metrics and healthceck port;
# EXPOSE 80

# TODO: add healthcheck command;
HEALTHCHECK NONE

# Set the default environment.
ENV COSI_ENDPOINT="unix:///var/lib/cosi/cosi.sock"
COPY --from=builder /cosi-osp/build/cosi-osp /usr/bin/cosi-osp

# Set the correct entrypoint and command arguments.
ENTRYPOINT [ "/usr/bin/cosi-osp" ]
CMD []
