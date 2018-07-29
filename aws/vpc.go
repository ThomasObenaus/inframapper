package aws

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/thomas.obenaus/inframapper/helper"
)

type Vpc struct {
	VpcId        string
	IsDefaultVPC bool
	CIDR         string
}

func (vpc *Vpc) Id() string {
	return vpc.VpcId
}

func (vpc *Vpc) Type() Type {
	return Type_VPC
}

func (vpc *Vpc) String() string {
	result := "id=" + vpc.VpcId + ", cidr=" + vpc.CIDR

	if vpc.IsDefaultVPC {
		result += " (default)"
	}

	return result
}

func (sl *infraLoaderImpl) loadVpcs() ([]*Vpc, error) {

	if err := sl.Validate(); err != nil {
		return nil, err
	}

	// load vpc data
	svc := ec2.New(sl.session)
	inDesc := &ec2.DescribeVpcsInput{DryRun: helper.NewFalse()}
	outDesc, err := svc.DescribeVpcs(inDesc)

	if err != nil {
		return nil, err
	}

	var vpcs []*Vpc

	for _, vpc := range outDesc.Vpcs {

		if vpc == nil {
			sl.tracer.Trace("Got nil vpc, ignore it.")
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
