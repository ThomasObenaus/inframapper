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
	NameTag      string
}

// ID returns the id of the AWS resource (i.e. 'vpc-f8168d93')
func (vpc Vpc) ID() string {
	return vpc.VpcID
}

// Type returns the type of this resource (i.e. aws_vpc)
func (vpc Vpc) Type() ResourceType {
	return TypeVPC
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
			tracer.Warn("Got nil vpc, ignore it.")
			continue
		}

		nameTag := FindNameTag(vpc.Tags, "")

		vpc := &Vpc{
			VpcID:        *vpc.VpcId,
			IsDefaultVPC: *vpc.IsDefault,
			CIDR:         *vpc.CidrBlock,
			NameTag:      nameTag,
		}

		vpcs = append(vpcs, vpc)
	}

	return vpcs, nil
}

// FindNameTag returns the value of the tag with key 'Name'.
// If there is no such tag available the given default name will be returned.
func FindNameTag(tags []*ec2.Tag, defaultName string) string {
	for _, tag := range tags {
		if tag.Key == nil {
			continue
		}

		if *tag.Key == "Name" {
			if tag.Value == nil {
				continue
			}
			return *tag.Value
		}

	}

	return defaultName
}
