package mappedInfra

import (
	"github.com/thomasobenaus/inframapper/aws"
	"github.com/thomasobenaus/inframapper/terraform"
)

type Vpc struct {
	awsVpc *aws.Vpc
	tfVpc  terraform.Resource
}

func (v *Vpc) ID() string {
	return v.awsVpc.ID()
}

func (v *Vpc) Type() aws.ResourceType {
	return v.awsVpc.Type()
}

func (v *Vpc) String() string {
	tfStateStr := "no tf-state"
	if v.HasTerraform() {
		tfStateStr = v.Terraform().Name()
	}
	return "[" + tfStateStr + "] " + v.awsVpc.String()
}

func (v *Vpc) Aws() aws.Resource {
	return v.awsVpc
}

func (v *Vpc) IsMapped() bool {
	return v.HasTerraform() && v.HasAws()
}

func (v *Vpc) HasAws() bool {
	return v.awsVpc != nil
}

func (v *Vpc) HasTerraform() bool {
	return v.tfVpc != nil
}

func (v *Vpc) Terraform() terraform.Resource {
	return v.tfVpc
}

func (v *Vpc) ResourceType() ResourceType {
	return TypeVPC
}

// NewVpc creates a mapping between an AWS Vpc and the according terraform resource.
func NewVpc(awsVpc *aws.Vpc, tfVpc terraform.Resource) MappedResource {
	return &Vpc{awsVpc: awsVpc, tfVpc: tfVpc}
}
