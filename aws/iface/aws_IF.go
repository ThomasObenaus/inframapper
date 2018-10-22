package iface

import "github.com/aws/aws-sdk-go/service/ec2"

// EC2IF is a minimal interface providing access to AWS EC2.
type EC2IF interface {
	DescribeVpcs(input *ec2.DescribeVpcsInput) (*ec2.DescribeVpcsOutput, error)
}
