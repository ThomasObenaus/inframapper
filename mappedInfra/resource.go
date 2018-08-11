package mappedInfra

import (
	"github.com/thomasobenaus/inframapper/aws"
	"github.com/thomasobenaus/inframapper/terraform"
)

type MappedResource interface {
	Aws() aws.Resource
	HasAws() bool
	HasTerraform() bool
	Terraform() terraform.Resource
	ResourceType() Type
	String() string
}

// Type represents the type of an aws resource
type Type int

const (
	Type_VPC Type = iota
)
