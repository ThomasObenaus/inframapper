package mappedInfra

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thomas.obenaus/terrastate/test/mock_aws"
	"github.com/thomas.obenaus/terrastate/test/mock_terraform"
)

func TestMap(t *testing.T) {

	mapper := NewMapper()
	require.NotNil(t, mapper)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockAwsInfraObj := mock_aws.NewMockInfra(mockCtrl)
	mockTerraformInfraObj := mock_terraform.NewMockInfra(mockCtrl)

	mappedInfra, err := mapper.Map(mockAwsInfraObj, mockTerraformInfraObj)
	assert.NotNil(t, mappedInfra)
	assert.NoError(t, err)
}

func TestNew(t *testing.T) {
	mapper := NewMapper()
	require.NotNil(t, mapper)
}
