package mappedInfra

import (
	"github.com/thomasobenaus/inframapper/aws"
	"github.com/thomasobenaus/inframapper/terraform"
)

// MappedResource represents a resource that describes a mapping
// between an AWS resource and its corresponding definition in
// terraform code.
type MappedResource interface {
	// Aws returns the actual AWS resource
	Aws() aws.Resource

	// HasAws returns true in case a AWS resource is available, false otherwise.
	HasAws() bool

	// Terraform returns the actual terraform resource representation.
	Terraform() terraform.Resource

	// HasTerraform returns true in case the terraform code for this resource is available, false otherwise.
	HasTerraform() bool

	// IsMapped returns true in case a mapping (code to real resource) is available, false otherwise.
	IsMapped() bool

	// ResourceType returns the type of this resource (i.e. Vpc)
	ResourceType() ResourceType
	String() string
}
