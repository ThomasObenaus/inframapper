package mappedInfra

import (
	"github.com/thomasobenaus/inframapper/aws"
	"github.com/thomasobenaus/inframapper/terraform"
)

type MappedResource interface {
	Aws() aws.Resource
	HasAws() bool
	HasTerraform() bool
	IsMapped() bool
	Terraform() terraform.Resource
	ResourceType() ResourceType
	String() string
}
