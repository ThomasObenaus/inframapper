package aws

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thomasobenaus/inframapper/helper"
	"github.com/thomasobenaus/inframapper/trace"
)

type vpcLoaderMock struct {
	vpcs []*ec2.Vpc
}

func (vpl *vpcLoaderMock) DescribeVpcs(input *ec2.DescribeVpcsInput) (*ec2.DescribeVpcsOutput, error) {
	if vpl.vpcs == nil || len(vpl.vpcs) == 0 {
		return nil, fmt.Errorf("No vpcs")
	}
	result := &ec2.DescribeVpcsOutput{Vpcs: vpl.vpcs}
	return result, nil
}

func TestLoadVPCFail(t *testing.T) {
	tracer := trace.Off()
	loader := &vpcLoaderMock{}
	vpcs, err := LoadVpcs(loader, tracer)

	assert.Error(t, err)
	assert.Nil(t, vpcs)
	assert.Empty(t, vpcs)

	awsVpcs := make([]*ec2.Vpc, 0)
	loader = &vpcLoaderMock{
		vpcs: awsVpcs,
	}

	vpcs, err = LoadVpcs(loader, tracer)
	assert.Error(t, err)
	assert.Nil(t, vpcs)
	assert.Empty(t, vpcs)
}

func TestLoadVPCSuccess(t *testing.T) {
	tracer := trace.Off()

	awsVpcs := make([]*ec2.Vpc, 0)
	vpc := &ec2.Vpc{VpcId: aws.String("12345"), IsDefault: helper.NewTrue(), CidrBlock: aws.String("10.100")}
	awsVpcs = append(awsVpcs, vpc)
	vpc = &ec2.Vpc{VpcId: aws.String("67890"), IsDefault: helper.NewTrue(), CidrBlock: aws.String("10.120")}
	awsVpcs = append(awsVpcs, vpc)
	awsVpcs = append(awsVpcs, nil)

	loader := &vpcLoaderMock{
		vpcs: awsVpcs,
	}

	vpcs, err := LoadVpcs(loader, tracer)
	require.NoError(t, err)
	require.NotNil(t, vpcs)
	require.NotEmpty(t, vpcs)
	require.Len(t, vpcs, 2)
	assert.Equal(t, "12345", vpcs[0].VpcId)
	assert.Equal(t, "67890", vpcs[1].VpcId)
}
