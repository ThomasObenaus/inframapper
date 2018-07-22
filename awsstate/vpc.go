package awsstate

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/thomas.obenaus/terrastate/helper"
)

type AwsVpc struct {
	VpcId        string
	IsDefaultVPC bool
	CIDR         string
}

func (vpc *AwsVpc) Id() string {
	return vpc.VpcId
}

func (vpc *AwsVpc) Type() Type {
	return Type_VPC
}

func (vpc *AwsVpc) String() string {
	result := "id=" + vpc.VpcId + ", cidr=" + vpc.CIDR

	if vpc.IsDefaultVPC {
		result += " (default)"
	}

	return result
}

func (sl *resourceLoaderImpl) loadVpc() ([]AwsVpc, error) {

	if err := sl.Validate(); err != nil {
		return nil, err
	}

	// load vpc data
	svc := ec2.New(sl.session)
	inDesc := &ec2.DescribeVpcsInput{DryRun: helper.NewFalse()}
	vpcs, err := svc.DescribeVpcs(inDesc)

	if err != nil {
		return nil, err
	}

	var awsVpcs []AwsVpc

	for _, vpc := range vpcs.Vpcs {

		if vpc == nil {
			sl.tracer.Trace("Got nil vpc, ignore it.")
			continue
		}

		awsVpc := AwsVpc{
			VpcId:        *vpc.VpcId,
			IsDefaultVPC: *vpc.IsDefault,
			CIDR:         *vpc.CidrBlock,
		}

		awsVpcs = append(awsVpcs, awsVpc)
	}

	return awsVpcs, nil
}
