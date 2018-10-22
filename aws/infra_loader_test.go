package aws

import (
	"fmt"
	"log"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/thomasobenaus/inframapper/helper"
	"github.com/thomasobenaus/inframapper/test/mock_aws_iface"
	"github.com/thomasobenaus/inframapper/trace"
)

func TestRLNew(t *testing.T) {

	sl, err := NewInfraLoader("", "")
	assert.NotNil(t, err)
	assert.Nil(t, sl)

	sl, err = NewInfraLoader("blubb", "")
	assert.NotNil(t, err)
	assert.Nil(t, sl)
}

func TestRLLoadFail(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ec2IF := mock_iface.NewMockEC2IF(mockCtrl)
	ec2IF.EXPECT().DescribeVpcs(gomock.Any()).Return(nil, fmt.Errorf("N/A"))

	loader := infraLoaderImpl{
		tracer: trace.Off(),
		ec2IF:  ec2IF,
	}

	infra, err := loader.Load()
	assert.NoError(t, err)
	assert.NotNil(t, infra)
}

func TestRLLoadSuccess(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ec2IF := mock_iface.NewMockEC2IF(mockCtrl)
	awsVpcs := make([]*ec2.Vpc, 0)
	vpc := &ec2.Vpc{VpcId: aws.String("12345"), IsDefault: helper.NewTrue(), CidrBlock: aws.String("10.100")}
	awsVpcs = append(awsVpcs, vpc)
	result := &ec2.DescribeVpcsOutput{Vpcs: awsVpcs}
	ec2IF.EXPECT().DescribeVpcs(gomock.Any()).Return(result, nil)

	loader := infraLoaderImpl{
		tracer: trace.Off(),
		ec2IF:  ec2IF,
	}

	infra, err := loader.Load()
	assert.NoError(t, err)
	assert.NotNil(t, infra)
}

func ExampleNewInfraLoader() {

	// Create a loader that will read the AWS resources for a given account
	// Here the account is specified by the AWS profile 'playground'
	// The region from where to read in is 'eu-central-1'
	iLoader, err := NewInfraLoader("playground", "eu-central-1")
	if err != nil {
		log.Fatalf("Error, creating infra-loader: %s", err.Error())
	}

	// After loading the AWS resources all information is stored in infra.
	infra, err := iLoader.Load()
	if err != nil {
		log.Fatalf("Error, loading infra: %s", err.Error())
	}
	fmt.Println(infra.String())
}
