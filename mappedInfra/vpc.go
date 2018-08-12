package mappedInfra

import (
	"github.com/thomasobenaus/inframapper/aws"
	"github.com/thomasobenaus/inframapper/terraform"
)

type vpc struct {
	awsVpc *aws.Vpc
	tfVpc  terraform.Resource
}

func (v *vpc) Id() string {
	return v.awsVpc.Id()
}

func (v *vpc) Type() aws.ResourceType {
	return v.awsVpc.Type()
}

func (v *vpc) String() string {
	hasTfStr := "-tf"
	if v.HasTerraform() {
		hasTfStr = "+tf"
	}
	return "[" + hasTfStr + "] " + v.awsVpc.String()
}

func (v *vpc) Aws() aws.Resource {
	return v.awsVpc
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
	return Type_VPC
}

func NewVpc(awsVpc *aws.Vpc, tfVpc terraform.Resource) MappedResource {
	return &vpc{awsVpc: awsVpc, tfVpc: tfVpc}
}
