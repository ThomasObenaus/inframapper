package mappedInfra

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thomas.obenaus/inframapper/aws"
)

func TestNewInfra(t *testing.T) {

	awsVpc := &aws.Vpc{VpcId: "1234"}

	mappedResources := make([]MappedResource, 1)
	mappedResources[0] = NewVpc(awsVpc, nil)

	infra, err := NewInfra(mappedResources)
	assert.NotNil(t, infra)
	assert.NoError(t, err)
	require.Equal(t, 1, infra.Resources())
	assert.Equal(t, "1234", infra.Resources()[0].Aws().Id())
}
