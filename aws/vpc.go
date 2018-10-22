package aws

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/thomasobenaus/inframapper/aws/iface"
	"github.com/thomasobenaus/inframapper/helper"
	"github.com/thomasobenaus/inframapper/trace"
)

// Vpc represents an AWS VPC
type Vpc struct {
	VpcID        string
	IsDefaultVPC bool
	CIDR         string
}

func (vpc Vpc) ID() string {
	return vpc.VpcID
}

func (vpc Vpc) Type() ResourceType {
	return Type_VPC
}

func (vpc Vpc) String() string {
	result := "id=" + vpc.VpcID + ", cidr=" + vpc.CIDR

	if vpc.IsDefaultVPC {
		result += " (default)"
	}

	return result
}

// LoadVpcs is a method used to load vpc information from AWS
func LoadVpcs(loader iface.EC2IF, tracer trace.Tracer) ([]*Vpc, error) {
	inDesc := &ec2.DescribeVpcsInput{DryRun: helper.NewFalse()}
	outDesc, err := loader.DescribeVpcs(inDesc)

	if err != nil {
		return nil, err
	}

	var vpcs []*Vpc

	for _, vpc := range outDesc.Vpcs {

		if vpc == nil {
			tracer.Trace("Got nil vpc, ignore it.")
			continue
		}

		vpc := &Vpc{
			VpcID:        *vpc.VpcId,
			IsDefaultVPC: *vpc.IsDefault,
			CIDR:         *vpc.CidrBlock,
		}

		vpcs = append(vpcs, vpc)
	}

	return vpcs, nil
}
