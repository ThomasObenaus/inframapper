package mappedInfra

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thomas.obenaus/terrastate/aws"
	"github.com/thomas.obenaus/terrastate/test/mock_aws"
	"github.com/thomas.obenaus/terrastate/test/mock_terraform"
)

func TestMapVpc(t *testing.T) {

	mapper := NewMapper()
	require.NotNil(t, mapper)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockAwsInfraObj := mock_aws.NewMockInfra(mockCtrl)
	mockTerraformInfraObj := mock_terraform.NewMockInfra(mockCtrl)
	mockTerraformResourceObj := mock_terraform.NewMockResource(mockCtrl)

	vpcs := make([]*aws.Vpc, 2)
	vpcs[0] = &aws.Vpc{
		VpcId:        "1234",
		IsDefaultVPC: false,
		CIDR:         "10.100.0.0/16",
	}
	mockAwsInfraObj.EXPECT().Vpcs().Return(vpcs)
	mockTerraformInfraObj.EXPECT().FindById("1234").Return(mockTerraformResourceObj)

	mappedInfra, err := mapper.Map(mockAwsInfraObj, mockTerraformInfraObj)
	assert.NotNil(t, mappedInfra)
	assert.NoError(t, err)
	assert.Equal(t, 1, mappedInfra.NumResources())

}

func TestNew(t *testing.T) {
	mapper := NewMapper()
	require.NotNil(t, mapper)
}
