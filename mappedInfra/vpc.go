package mappedInfra

import (
	"fmt"

	"github.com/thomasobenaus/inframapper/aws"
	"github.com/thomasobenaus/inframapper/terraform"
)

type Vpc struct {
	AwsVpc *aws.Vpc
	TfVpc  terraform.Resource
}

func (v *Vpc) ID() string {
	return v.AwsVpc.ID()
}

func (v *Vpc) Type() aws.ResourceType {
	return v.AwsVpc.Type()
}

func (v *Vpc) String() string {
	tfStateStr := "no tf-state"
	if v.HasTerraform() {
		tfStateStr = v.Terraform().Name()
	}
	return "[" + tfStateStr + "] " + v.AwsVpc.String()
}

func (v *Vpc) Aws() aws.Resource {
	return v.AwsVpc
}

func (v *Vpc) IsMapped() bool {
	return v.HasTerraform() && v.HasAws()
}

func (v *Vpc) HasAws() bool {
	return v.AwsVpc != nil
}

func (v *Vpc) HasTerraform() bool {
	return v.TfVpc != nil
}

func (v *Vpc) Terraform() terraform.Resource {
	return v.TfVpc
}

func (v *Vpc) ResourceType() ResourceType {
	return TypeVPC
}

// ToVpc casts the given resourc into a Vpc (if possible)
func ToVpc(r MappedResource) (*Vpc, error) {
	if r.ResourceType() != TypeVPC {
		return nil, fmt.Errorf("Not of type Vpc")
	}

	return r.(*Vpc), nil
}

// NewVpc creates a mapping between an AWS Vpc and the according terraform resource.
func NewVpc(AwsVpc *aws.Vpc, tfVpc terraform.Resource) MappedResource {
	return &Vpc{AwsVpc: AwsVpc, TfVpc: tfVpc}
}
