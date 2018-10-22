package mappedInfra

import (
	"github.com/thomasobenaus/inframapper/aws"
	"github.com/thomasobenaus/inframapper/terraform"
)

type vpc struct {
	awsVpc *aws.Vpc
	tfVpc  terraform.Resource
}

func (v *vpc) ID() string {
	return v.awsVpc.ID()
}

func (v *vpc) Type() aws.ResourceType {
	return v.awsVpc.Type()
}

func (v *vpc) String() string {
	tfStateStr := "no tf-state"
	if v.HasTerraform() {
		tfStateStr = v.Terraform().Name()
	}
	return "[" + tfStateStr + "] " + v.awsVpc.String()
}

func (v *vpc) Aws() aws.Resource {
	return v.awsVpc
}

func (v *vpc) IsMapped() bool {
	return v.HasTerraform() && v.HasAws()
}

func (v *vpc) HasAws() bool {
	return v.awsVpc != nil
}

func (v *vpc) HasTerraform() bool {
	return v.tfVpc != nil
}

func (v *vpc) Terraform() terraform.Resource {
	return v.tfVpc
}

func (v *vpc) ResourceType() ResourceType {
	return TypeVPC
}

// NewVpc creates a mapping between an AWS vpc and the according terraform resource.
func NewVpc(awsVpc *aws.Vpc, tfVpc terraform.Resource) MappedResource {
	return &vpc{awsVpc: awsVpc, tfVpc: tfVpc}
}
