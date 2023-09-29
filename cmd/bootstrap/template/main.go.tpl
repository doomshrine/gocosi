package main

import (
	"context"
	"flag"
	"fmt"
	stdlog "log"
	"os"
	"os/signal"
	"syscall"

	"github.com/doomshrine/gocosi"
	"github.com/go-logr/logr"
	"github.com/go-logr/stdr"
	"github.com/hellofresh/health-go/v5"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0" // FIXME: this might need manual update

	"{{ .ModPath }}/servers/identity"
	"{{ .ModPath }}/servers/provisioner"
)

var (
	ospName    = "cosi.example.com" // FIXME: replace with your own OSP name
	ospVersion = "v0.1.0"           // FIXME: replace with your own OSP version

	exporterKind = gocosi.HTTPExporter

	log logr.Logger

	healthcheck bool
)

func main() {
	flag.BoolVar(&healthcheck, "healthcheck", false, "")
	flag.Parse()

	// Setup your logger here.
	// You can use one of multiple available implementation, like:
	//   - https://github.com/kubernetes/klog/tree/main/klogr
	//   - https://github.com/go-logr/logr/tree/master/slogr
	//   - https://github.com/go-logr/stdr
	//   - https://github.com/bombsimon/logrusr
	stdr.SetVerbosity(10)
	log = stdr.New(stdlog.New(os.Stdout, "", stdlog.LstdFlags))

	gocosi.SetLogger(log)

	if err := realMain(context.Background()); err != nil {
		log.Error(err, "critical failure")
		os.Exit(1)
	}
}

func realMain(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	if healthcheck {
		return runHealthcheck(ctx)
	}

	return runOSP(ctx)
}

func runHealthcheck(ctx context.Context) error {
	err := gocosi.HealthcheckFunc(ctx, gocosi.HealthcheckAddr)
	if err != nil {
		return fmt.Errorf("healthcheck call failed: %w", err)
	}

	return nil
}

func runOSP(ctx context.Context) error {
	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceName(ospName),
		semconv.ServiceVersion(ospVersion),
	)

	// If there is any additional confifuration needed for your COSI Driver,
	// put it below this line.

	driver, err := gocosi.New(
		identity.New(ospName, log),
		provisioner.New(log),
		res,
		gocosi.WithHealthcheck(
			health.WithComponent(health.Component{
				Name:    ospName,
				Version: ospVersion,
			}),
		),
		gocosi.WithDefaultGRPCOptions(),
		gocosi.WithDefaultMetricExporter(exporterKind),
		gocosi.WithDefaultTraceExporter(exporterKind),
	)
	if err != nil {
		return fmt.Errorf("failed to create COSI OSP: %w", err)
	}

	if err := driver.Run(ctx); err != nil {
		return fmt.Errorf("failed to run COSI OSP: %w", err)
	}

	return nil
}
