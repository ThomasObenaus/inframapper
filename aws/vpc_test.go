package aws

import (
	"fmt"
	"log"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thomasobenaus/inframapper/helper"
	"github.com/thomasobenaus/inframapper/test/mock_aws_iface"
	"github.com/thomasobenaus/inframapper/trace"
)

func TestLoadVPCFail(t *testing.T) {
	tracer := trace.Off()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ec2IF := mock_iface.NewMockEC2IF(mockCtrl)
	ec2IF.EXPECT().DescribeVpcs(gomock.Any()).Return(nil, fmt.Errorf("N/A"))

	vpcs, err := LoadVpcs(ec2IF, tracer)

	assert.Error(t, err)
	assert.Nil(t, vpcs)
	assert.Empty(t, vpcs)

	awsVpcs := make([]*ec2.Vpc, 0)
	ec2IF.EXPECT().DescribeVpcs(gomock.Any()).Return(&ec2.DescribeVpcsOutput{Vpcs: awsVpcs}, nil)

	vpcs, err = LoadVpcs(ec2IF, tracer)
	assert.NoError(t, err)
	assert.Nil(t, vpcs)
	assert.Empty(t, vpcs)
}

func TestLoadVPCSuccess(t *testing.T) {
	tracer := trace.Off()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ec2IF := mock_iface.NewMockEC2IF(mockCtrl)
	awsVpcs := make([]*ec2.Vpc, 0)
	vpc := &ec2.Vpc{VpcId: aws.String("12345"), IsDefault: helper.NewTrue(), CidrBlock: aws.String("10.100")}
	awsVpcs = append(awsVpcs, vpc)
	vpc = &ec2.Vpc{VpcId: aws.String("67890"), IsDefault: helper.NewTrue(), CidrBlock: aws.String("10.120")}
	awsVpcs = append(awsVpcs, vpc)
	awsVpcs = append(awsVpcs, nil)
	ec2IF.EXPECT().DescribeVpcs(gomock.Any()).Return(&ec2.DescribeVpcsOutput{Vpcs: awsVpcs}, nil)

	vpcs, err := LoadVpcs(ec2IF, tracer)

	require.NoError(t, err)
	require.NotNil(t, vpcs)
	require.NotEmpty(t, vpcs)
	require.Len(t, vpcs, 2)
	assert.Equal(t, "12345", vpcs[0].VpcID)
	assert.Equal(t, "67890", vpcs[1].VpcID)
}

func ExampleLoadVpcs() {

	tracer := trace.Off()

	// create aws session
	sess, err := newAWSSession("myprofile", "eu-central-1")
	if err != nil {
		log.Fatalf("Unable to create session: %s", err.Error())
	}

	// obtain ec2 interface
	ec2IF := ec2.New(sess)

	// load the vpcs
	vpcs, err := LoadVpcs(ec2IF, tracer)
	if err != nil {
		log.Fatalf("Unable to load vpcs: %s", err.Error())
	}

	log.Printf("Loaded %d vpcs", len(vpcs))
}
