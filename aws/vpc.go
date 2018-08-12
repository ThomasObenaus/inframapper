package aws

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/thomasobenaus/inframapper/aws/iface"
	"github.com/thomasobenaus/inframapper/helper"
	"github.com/thomasobenaus/inframapper/trace"
)

type Vpc struct {
	VpcId        string
	IsDefaultVPC bool
	CIDR         string
}

type VpcLoader interface {
	DescribeVpcs(input *ec2.DescribeVpcsInput) (*ec2.DescribeVpcsOutput, error)
}

func (vpc Vpc) Id() string {
	return vpc.VpcId
}

func (vpc Vpc) Type() ResourceType {
	return Type_VPC
}

func (vpc Vpc) String() string {
	result := "id=" + vpc.VpcId + ", cidr=" + vpc.CIDR

	if vpc.IsDefaultVPC {
		result += " (default)"
	}

	return result
}

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
			VpcId:        *vpc.VpcId,
			IsDefaultVPC: *vpc.IsDefault,
			CIDR:         *vpc.CidrBlock,
		}

		vpcs = append(vpcs, vpc)
	}

	return vpcs, nil
}
