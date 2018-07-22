package mappedInfra

import (
	"github.com/thomas.obenaus/terrastate/aws"
	"github.com/thomas.obenaus/terrastate/tfstate"
)

type MappedResource interface {
	Aws() (MappedAwsResource, error)
}

type MappedAwsResource interface {
	Aws() aws.Resource
	HasTerraform() bool
	Terraform() (tfstate.Resource, error)
}
