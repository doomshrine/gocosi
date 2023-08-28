###############################################################################
# First stage:
#   - Building the COSI OSP using the default Go image.
###############################################################################
FROM --platform=${BUILDPLATFORM} docker.io/library/golang:{{ .GoVersion }} AS builder

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
# NOTE: you should replace the latest/ubi9 with specific digest, to ensure,
#   that builds are consistent.
###############################################################################
{{ if .Rootless }}
FROM --platform=${BUILDPLATFORM} {{ .Docker.RootlessImage }} AS runtime
{{ else }}
FROM --platform=${BUILDPLATFORM} {{ .Docker.Image }} AS runtime
{{ end }}

# Set the working directory.
WORKDIR /cosi

# Set volume mount point for COSI socket.
VOLUME [ "/var/lib/cosi" ]

# TODO: add metrics and healthceck port;
# EXPOSE 80

# TODO: add healthcheck command;
HEALTHCHECK NONE

# Set the default environment.
ENV COSI_ENDPOINT="/var/lib/cosi/cosi.sock"
{{ if .Rootless }}
ENV X_COSI_ENDPOINT_PERMS="0755"
ENV X_COSI_ENDPOINT_USER="1001"
ENV X_COSI_ENDPOINT_GROUP="1001"

# Create a non-root user and set permissions on the files.
RUN echo "cosi:*:1001:cosi-user" >> /etc/group && \
    echo "cosi-user:*:1001:1001::/cosi:/bin/false" >> /etc/passwd && \
    chown 1001:1001 /cosi && \
    chmod 0550 /cosi && \
    mkdir -p /var/lib/cosi && \
    chown -R 1001:1001 /var/lib/cosi

# Copy the newly build binary to final image, and set the permissions.
COPY --from=builder --chown=1001:1001 /cosi-osp/build/cosi-osp /usr/bin/cosi-osp
RUN chmod 0550 /cosi
{{ else }}
COPY --from=builder /cosi-osp/build/cosi-osp /usr/bin/cosi-osp
{{ end }}

# Set the correct entrypoint and command arguments.
ENTRYPOINT [ "/usr/bin/cosi-osp" ]
CMD []
