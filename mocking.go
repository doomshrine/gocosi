package gocosi

import cosi "sigs.k8s.io/container-object-storage-interface-spec"

//go:generate go run github.com/vektra/mockery/v2@v2.32.0

type COSIProvisionerServer interface {
	cosi.ProvisionerServer
}

type COSIIdentityServer interface {
	cosi.IdentityServer
}
