package mappedInfra

import (
	"testing"

	"github.com/golang/mock/gomock"
	terraform "github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thomas.obenaus/terrastate/test/mock_aws"
)

func TestNew(t *testing.T) {

	mapper, err := NewMapper(nil, nil)
	assert.Nil(t, mapper)
	assert.Error(t, err)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockObj := mock_aws.NewMockInfra(mockCtrl)

	tfstate := &terraform.State{}
	mapper, err = NewMapper(mockObj, tfstate)
	assert.NotNil(t, mapper)
	assert.NoError(t, err)
}

func TestMap(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockObj := mock_aws.NewMockInfra(mockCtrl)

	tfstate := &terraform.State{}
	mapper, err := NewMapper(mockObj, tfstate)
	require.NotNil(t, mapper)
	require.NoError(t, err)

}
