package main

import (
	"context"
	"os"

	"github.com/doomshrine/gocosi"
	"github.com/go-logr/logr"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0" // FIXME: this might need manual update

	"{{ .ModPath }}/servers/identity"
	"{{ .ModPath }}/servers/provisioner"
)

var (
	driverName = "cosi.example.com" // FIXME: replace with your own driver name
	driverVersion = "v0.1.0" // FIXME: replace with your own driver version

	exporterKind = gocosi.HTTPExporter

	log logr.Logger
)

func init() {
	// Setup your logger here.
	// You can use one of multiple available implementation, like:
	//   - https://github.com/kubernetes/klog/tree/main/klogr
	//   - https://github.com/go-logr/logr/tree/master/slogr
	//   - https://github.com/go-logr/stdr
	//   - https://github.com/bombsimon/logrusr
}

func main() {
	gocosi.SetLogger(log)

	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceName(driverName),
		semconv.ServiceVersion(driverVersion),
	)

	// If there is any additional confifuration needed for your COSI Driver,
	// put it below this line.

	driver, err := gocosi.New(
		identity.New(driverName, log),
		provisioner.New(log),
		res,
		gocosi.WithDefaultGRPCOptions(),
		gocosi.WithDefaultMetricExporter(exporterKind),
		gocosi.WithDefaultTraceExporter(exporterKind),
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
