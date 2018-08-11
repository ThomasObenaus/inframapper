package aws

import "github.com/aws/aws-sdk-go/service/ec2"

type EC2IF interface {
	DescribeVpcs(input *ec2.DescribeVpcsInput) (*ec2.DescribeVpcsOutput, error)
}
