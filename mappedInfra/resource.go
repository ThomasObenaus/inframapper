package mappedInfra

import (
	"github.com/thomas.obenaus/terrastate/aws"
	"github.com/thomas.obenaus/terrastate/terraform"
)

type MappedResource interface {
	Aws() aws.Resource
	HasAws() bool
	HasTerraform() bool
	Terraform() terraform.Resource
	ResourceType() Type
}

// Type represents the type of an aws resource
type Type int

const (
	Type_VPC Type = iota
)
