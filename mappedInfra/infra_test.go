package mappedInfra

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thomasobenaus/inframapper/aws"
	"github.com/thomasobenaus/inframapper/test/mock_terraform"
)

func TestNewInfra(t *testing.T) {

	awsVpc := &aws.Vpc{VpcID: "1234"}

	mappedResources := make([]MappedResource, 1)
	mappedResources[0] = NewVpc(awsVpc, nil)

	infra, err := NewInfra(mappedResources)
	assert.NotNil(t, infra)
	assert.NoError(t, err)
	require.Equal(t, 1, len(infra.Resources()))
	assert.Equal(t, "1234", infra.Resources()[0].Aws().ID())
}

func TestUnMappedAws(t *testing.T) {

	awsVpc := &aws.Vpc{VpcID: "1234"}

	mappedResources := make([]MappedResource, 1)
	mappedResources[0] = NewVpc(awsVpc, nil)

	infra, err := NewInfra(mappedResources)
	assert.NotNil(t, infra)
	assert.NoError(t, err)
	require.Equal(t, 1, len(infra.UnMappedAwsResources()))
	assert.Equal(t, "1234", infra.UnMappedAwsResources()[0].Aws().ID())
}

func TestMapped(t *testing.T) {

	mapper := NewMapper()
	require.NotNil(t, mapper)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockAwsInfraObj := mock_terraform.NewMockResource(mockCtrl)

	awsVpc := &aws.Vpc{VpcID: "1234"}

	mappedResources := make([]MappedResource, 1)
	mappedResources[0] = NewVpc(awsVpc, mockAwsInfraObj)

	infra, err := NewInfra(mappedResources)
	assert.NotNil(t, infra)
	assert.NoError(t, err)
	require.Equal(t, 1, len(infra.MappedResources()))
	assert.Equal(t, "1234", infra.MappedResources()[0].Aws().ID())

}

func TestStringFormat(t *testing.T) {

	mapper := NewMapper()
	require.NotNil(t, mapper)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockAwsInfraObj := mock_terraform.NewMockResource(mockCtrl)

	awsVpc := &aws.Vpc{VpcID: "1234"}

	mappedResources := make([]MappedResource, 1)
	mappedResources[0] = NewVpc(awsVpc, mockAwsInfraObj)

	infra, err := NewInfra(mappedResources)
	assert.NotNil(t, infra)
	assert.NoError(t, err)
	require.Equal(t, 1, len(infra.MappedResources()))
	assert.Equal(t, "1234", infra.MappedResources()[0].Aws().ID())

	assert.Equal(t, "#res=1, #mapped=1, #aws_res=0", infra.String())
}
