package mappedInfra

import (
	"fmt"
	"log"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thomasobenaus/inframapper/aws"
	"github.com/thomasobenaus/inframapper/terraform"
	"github.com/thomasobenaus/inframapper/test/mock_aws"
	"github.com/thomasobenaus/inframapper/test/mock_terraform"
	"github.com/thomasobenaus/inframapper/tfstate"
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

	vpcID := "1234"
	vpcs := make([]*aws.Vpc, 2)
	vpcs[0] = &aws.Vpc{
		VpcID:        vpcID,
		IsDefaultVPC: false,
		CIDR:         "10.100.0.0/16",
	}
	mockAwsInfraObj.EXPECT().Vpcs().Return(vpcs)
	mockTerraformInfraObj.EXPECT().FindByID(vpcID).Return(mockTerraformResourceObj)
	mockTerraformResourceObj.EXPECT().ID().Return(vpcID)
	mockTerraformResourceObj.EXPECT().Name().Times(2).Return("aws_vpc.bla")
	mockTerraformResourceObj.EXPECT().Type().Return(terraform.TypeAwsVpc)

	mappedInfra, err := mapper.Map(mockAwsInfraObj, mockTerraformInfraObj)
	assert.NotNil(t, mappedInfra)
	assert.NoError(t, err)
	assert.Equal(t, 1, mappedInfra.NumResources())

	res := mappedInfra.AwsResourceByID(vpcID)
	require.NotNil(t, res)
	assert.Equal(t, true, res.HasAws())
	assert.Equal(t, vpcID, res.Aws().ID())
	assert.Equal(t, aws.TypeVPC, res.Aws().Type())
	assert.Equal(t, true, res.HasTerraform())
	assert.Equal(t, vpcID, res.Terraform().ID())
	assert.Equal(t, "aws_vpc.bla", res.Terraform().Name())
	assert.Equal(t, terraform.TypeAwsVpc, res.Terraform().Type())
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

	vpcID := "1234"
	vpcs := make([]*aws.Vpc, 2)
	vpcs[0] = &aws.Vpc{
		VpcID:        vpcID,
		IsDefaultVPC: false,
		CIDR:         "10.100.0.0/16",
	}

	mockTerraformInfraObj := mock_terraform.NewMockInfra(mockCtrl)
	mockTerraformResourceObj := mock_terraform.NewMockResource(mockCtrl)
	mockTerraformInfraObj.EXPECT().FindByID(vpcID).Return(mockTerraformResourceObj)
	mockTerraformResourceObj.EXPECT().Name().Times(2).Return("aws_vpc.bla")
	mockTerraformResourceObj.EXPECT().ID().Return(vpcID)
	mockTerraformResourceObj.EXPECT().Type().Return(terraform.TypeAwsVpc)

	mappedInfraList := mapper.mapVpcs(vpcs, mockTerraformInfraObj)
	assert.NotNil(t, mappedInfraList)
	assert.NotEmpty(t, mappedInfraList)
	assert.Equal(t, 1, len(mappedInfraList))

	res := mappedInfraList[0]
	require.NotNil(t, res)
	assert.Equal(t, true, res.HasAws())
	assert.Equal(t, vpcID, res.Aws().ID())
	assert.Equal(t, aws.TypeVPC, res.Aws().Type())
	assert.Equal(t, true, res.HasTerraform())
	assert.Equal(t, vpcID, res.Terraform().ID())
	assert.Equal(t, "aws_vpc.bla", res.Terraform().Name())
	assert.Equal(t, terraform.TypeAwsVpc, res.Terraform().Type())
}

func ExampleLoadAndMap_remote() {
	awsProfile := "develop"
	awsRegion := "eu-central-1"

	keys := make([]string, 2)
	keys[0] = "snapshot/base/networking/terraform.tfstate"
	keys[1] = "snapshot/base/common/terraform.tfstate"
	remoteCfg := tfstate.RemoteConfig{
		BucketName: "741125603121-tfstate",
		Keys:       keys,
		Profile:    "shared",
		Region:     "eu-central-1",
	}

	mappedInfra, err := LoadAndMap(awsProfile, awsRegion, remoteCfg, nil)
	if err != nil {
		log.Fatalf("Error loading/ mapping infrastructure: %s", err.Error())
	}

	// now you can do sth. with the resources
	fmt.Println("Mapped Resources [", len(mappedInfra.MappedResources()), "]:")
	fmt.Println("UnMapped AWS Resources [", len(mappedInfra.UnMappedAwsResources()), "]:")
}

func ExampleLoadAndMap_local() {
	awsProfile := "develop"
	awsRegion := "eu-central-1"

	files := []string{"../testdata/statefiles/instance.tfstate"}
	localConfig := tfstate.LocalConfig{
		Files: files,
	}

	mappedInfra, err := LoadAndMap(awsProfile, awsRegion, localConfig, nil)
	if err != nil {
		log.Fatalf("Error loading/ mapping infrastructure: %s", err.Error())
	}

	// now you can do sth. with the resources
	fmt.Println("Mapped Resources [", len(mappedInfra.MappedResources()), "]:")
	fmt.Println("UnMapped AWS Resources [", len(mappedInfra.UnMappedAwsResources()), "]:")
}
