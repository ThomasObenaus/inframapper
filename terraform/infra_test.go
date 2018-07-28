package terraform

import (
	"encoding/json"
	"testing"

	terraform "github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thomas.obenaus/terrastate/trace"
)

var vpc = `{
	"modules": [ {
			"resources": {
				"aws_vpc.default": {
					"type": "aws_vpc",
					"depends_on": ["bla","blubb"],
					"primary": {
							"id": "vpc-ff5fec97"
					},
					"provider": "provider.aws"
				}
			}
		}
	]
}
`

func TestNew(t *testing.T) {

	infra, err := NewInfra(nil)
	assert.Nil(t, infra)
	assert.Error(t, err)

	tfstate := &terraform.State{}
	infra, err = NewInfra(tfstate)
	assert.NotNil(t, infra)
	assert.NoError(t, err)
}

func TestCreateResourcesByIdMap(t *testing.T) {

	tracer := trace.Off()

	resourcesById, err := createResourcesByIdMap(nil, tracer)
	assert.Empty(t, resourcesById)
	assert.Error(t, err)

	tfstate := &terraform.State{}
	resourcesById, err = createResourcesByIdMap(tfstate, tracer)
	assert.Empty(t, resourcesById)
	assert.NoError(t, err)

	tfstate = &terraform.State{}
	err = json.Unmarshal([]byte(vpc), tfstate)
	require.NoError(t, err)

	resourcesById, err = createResourcesByIdMap(tfstate, tracer)
	assert.NotEmpty(t, resourcesById)
	assert.NoError(t, err)
	assert.NotNil(t, resourcesById["vpc-ff5fec97"])
}

func TestCreateResourcesByNameMap(t *testing.T) {

	tracer := trace.Off()

	resourcesByName, err := createResourcesByNameMap(nil, tracer)
	assert.Empty(t, resourcesByName)
	assert.Error(t, err)

	tfstate := &terraform.State{}
	resourcesByName, err = createResourcesByNameMap(tfstate, tracer)
	assert.Empty(t, resourcesByName)
	assert.NoError(t, err)

	tfstate = &terraform.State{}
	err = json.Unmarshal([]byte(vpc), tfstate)
	require.NoError(t, err)

	resourcesByName, err = createResourcesByNameMap(tfstate, tracer)
	assert.NotEmpty(t, resourcesByName)
	assert.NoError(t, err)
	assert.NotNil(t, resourcesByName["aws_vpc.default"])
}

func TestFindResourceByName(t *testing.T) {
	tfstate := &terraform.State{}
	err := json.Unmarshal([]byte(vpc), tfstate)
	require.NoError(t, err)

	infra, err := NewInfra(tfstate)
	require.NotNil(t, infra)
	require.NoError(t, err)

	resource := infra.FindByName("aws_vpc.default")
	require.NotNil(t, resource)
	assert.Equal(t, "aws_vpc.default", resource.Name())
}

func TestFindResourceById(t *testing.T) {
	tfstate := &terraform.State{}
	err := json.Unmarshal([]byte(vpc), tfstate)
	require.NoError(t, err)

	infra, err := NewInfra(tfstate)
	require.NotNil(t, infra)
	require.NoError(t, err)

	resource := infra.FindById("vpc-ff5fec97")
	require.NotNil(t, resource)
	assert.Equal(t, "vpc-ff5fec97", resource.Id())
}
