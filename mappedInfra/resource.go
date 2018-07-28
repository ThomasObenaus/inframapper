package mappedInfra

import (
	"github.com/thomas.obenaus/terrastate/aws"
	"github.com/thomas.obenaus/terrastate/terraform"
)

type MappedResource interface {
	Aws() aws.Resource
	IsAws() bool
	HasTerraform() bool
	Terraform() terraform.Resource
}
