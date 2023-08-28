package main

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
	//   - https://github.com/go-logr/logr/tree/master/slogr
	//   - https://github.com/go-logr/stdr
	//   - https://github.com/bombsimon/logrusr
}

func main() {
	gocosi.SetLogger(log)

	// If there is any additional confifuration needed for your COSI Driver,
	// put it below this line.

	driver, err := gocosi.New(
		identity.New(driverName, log),
		provisioner.New(log),
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
