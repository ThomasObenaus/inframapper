package mappedInfra

import (
	"github.com/thomas.obenaus/terrastate/aws"
	"github.com/thomas.obenaus/terrastate/terraform"
)

type MappedResource interface {
	Aws() (MappedAwsResource, error)
}

type MappedAwsResource interface {
	Aws() aws.Resource
	HasTerraform() bool
	Terraform() terraform.Resource
}
