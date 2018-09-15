package mappedInfra

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thomasobenaus/inframapper/aws"
	"github.com/thomasobenaus/inframapper/terraform"
	"github.com/thomasobenaus/inframapper/test/mock_aws"
	"github.com/thomasobenaus/inframapper/test/mock_terraform"
	"github.com/thomasobenaus/inframapper/trace"
)

func TestMap(t *testing.T) {

	mapper := NewMapper()
	require.NotNil(t, mapper)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockAwsInfraObj := mock_aws.NewMockInfra(mockCtrl)
	mockTerraformInfraObj := mock_terraform.NewMockInfra(mockCtrl)
	mockTerraformResourceObj := mock_terraform.NewMockResource(mockCtrl)

	vpcId := "1234"
	vpcs := make([]*aws.Vpc, 2)
	vpcs[0] = &aws.Vpc{
		VpcId:        vpcId,
		IsDefaultVPC: false,
		CIDR:         "10.100.0.0/16",
	}
	mockAwsInfraObj.EXPECT().Vpcs().Return(vpcs)
	mockTerraformInfraObj.EXPECT().FindById(vpcId).Return(mockTerraformResourceObj)
	mockTerraformResourceObj.EXPECT().Id().Return(vpcId)
	mockTerraformResourceObj.EXPECT().Name().Times(2).Return("aws_vpc.bla")
	mockTerraformResourceObj.EXPECT().Type().Return(terraform.Type_aws_vpc)

	mappedInfra, err := mapper.Map(mockAwsInfraObj, mockTerraformInfraObj)
	assert.NotNil(t, mappedInfra)
	assert.NoError(t, err)
	assert.Equal(t, 1, mappedInfra.NumResources())

	res := mappedInfra.AwsResourceById(vpcId)
	require.NotNil(t, res)
	assert.Equal(t, true, res.HasAws())
	assert.Equal(t, vpcId, res.Aws().Id())
	assert.Equal(t, aws.Type_VPC, res.Aws().Type())
	assert.Equal(t, true, res.HasTerraform())
	assert.Equal(t, vpcId, res.Terraform().Id())
	assert.Equal(t, "aws_vpc.bla", res.Terraform().Name())
	assert.Equal(t, terraform.Type_aws_vpc, res.Terraform().Type())
}

func TestNew(t *testing.T) {
	mapper := NewMapper()
	require.NotNil(t, mapper)
}

func TestMapVpc(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mapper := mapperImpl{tracer: trace.Off()}
	require.NotNil(t, mapper)

	vpcId := "1234"
	vpcs := make([]*aws.Vpc, 2)
	vpcs[0] = &aws.Vpc{
		VpcId:        vpcId,
		IsDefaultVPC: false,
		CIDR:         "10.100.0.0/16",
	}

	mockTerraformInfraObj := mock_terraform.NewMockInfra(mockCtrl)
	mockTerraformResourceObj := mock_terraform.NewMockResource(mockCtrl)
	mockTerraformInfraObj.EXPECT().FindById(vpcId).Return(mockTerraformResourceObj)
	mockTerraformResourceObj.EXPECT().Name().Times(2).Return("aws_vpc.bla")
	mockTerraformResourceObj.EXPECT().Id().Return(vpcId)
	mockTerraformResourceObj.EXPECT().Type().Return(terraform.Type_aws_vpc)

	mappedInfraList := mapper.mapVpcs(vpcs, mockTerraformInfraObj)
	assert.NotNil(t, mappedInfraList)
	assert.NotEmpty(t, mappedInfraList)
	assert.Equal(t, 1, len(mappedInfraList))

	res := mappedInfraList[0]
	require.NotNil(t, res)
	assert.Equal(t, true, res.HasAws())
	assert.Equal(t, vpcId, res.Aws().Id())
	assert.Equal(t, aws.Type_VPC, res.Aws().Type())
	assert.Equal(t, true, res.HasTerraform())
	assert.Equal(t, vpcId, res.Terraform().Id())
	assert.Equal(t, "aws_vpc.bla", res.Terraform().Name())
	assert.Equal(t, terraform.Type_aws_vpc, res.Terraform().Type())
}
