package terraform

import (
	"encoding/json"
	"testing"

	terraform "github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thomas.obenaus/inframapper/trace"
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

var multiVpc = `{
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
		}, {
			"resources": {
				"aws_vpc.important": {
					"type": "aws_vpc",
					"depends_on": ["foo","bar"],
					"primary": {
							"id": "vpc-fa697123"
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

	tfStateList := make([]*terraform.State, 1)
	tfStateList[0] = &terraform.State{}
	infra, err = NewInfra(tfStateList)
	assert.NotNil(t, infra)
	assert.NoError(t, err)
}

func TestCreateResourcesByIdMap(t *testing.T) {
	tracer := trace.Off()

	tfstate := &terraform.State{}
	err := json.Unmarshal([]byte(vpc), tfstate)
	require.NoError(t, err)

	resourcesById, err := createResourcesByIdMap(tfstate, tracer)
	assert.NotEmpty(t, resourcesById)
	assert.NoError(t, err)
	assert.NotNil(t, resourcesById["vpc-ff5fec97"])
}

func TestCreateResourcesByNameMap(t *testing.T) {
	tracer := trace.Off()

	tfstate := &terraform.State{}
	err := json.Unmarshal([]byte(vpc), tfstate)
	require.NoError(t, err)

	resourcesById, err := createResourcesByNameMap(tfstate, tracer)
	assert.NotEmpty(t, resourcesById)
	assert.NoError(t, err)
	assert.NotNil(t, resourcesById["aws_vpc.default"])
}

func TestCreateResourcesByXMap(t *testing.T) {
	tracer := trace.Off()

	resourcesByName, err := createResourcesByXMap(nil, filterCriteria_Name, tracer)
	assert.Empty(t, resourcesByName)
	assert.Error(t, err)

	tfstate := &terraform.State{}
	resourcesByName, err = createResourcesByXMap(tfstate, filterCriteria_Name, tracer)
	assert.Empty(t, resourcesByName)
	assert.NoError(t, err)

	tfstate = &terraform.State{}
	err = json.Unmarshal([]byte(vpc), tfstate)
	require.NoError(t, err)

	resourcesByName, err = createResourcesByXMap(tfstate, filterCriteria_Name, tracer)
	assert.NotEmpty(t, resourcesByName)
	assert.NoError(t, err)
	assert.NotNil(t, resourcesByName["aws_vpc.default"])

	tfstate = &terraform.State{}
	err = json.Unmarshal([]byte(multiVpc), tfstate)
	require.NoError(t, err)

	resourcesByName, err = createResourcesByXMap(tfstate, filterCriteria_Name, tracer)
	assert.NotEmpty(t, resourcesByName)
	assert.NoError(t, err)
	assert.NotNil(t, resourcesByName["aws_vpc.default"])
	assert.NotNil(t, resourcesByName["aws_vpc.important"])

	resourcesByName, err = createResourcesByXMap(tfstate, filterCriteria_Id, tracer)
	assert.NotEmpty(t, resourcesByName)
	assert.NoError(t, err)
	assert.NotNil(t, resourcesByName["vpc-ff5fec97"])
	assert.NotNil(t, resourcesByName["vpc-fa697123"])
}

func TestFindResourceByName(t *testing.T) {

	tfStateList := make([]*terraform.State, 1)
	tfStateList[0] = &terraform.State{}
	err := json.Unmarshal([]byte(vpc), tfStateList[0])
	require.NoError(t, err)

	infra, err := NewInfra(tfStateList)
	require.NotNil(t, infra)
	require.NoError(t, err)

	resource := infra.FindByName("aws_vpc.default")
	require.NotNil(t, resource)
	assert.Equal(t, "aws_vpc.default", resource.Name())
}

func TestFindResourceById(t *testing.T) {
	tfStateList := make([]*terraform.State, 1)
	tfStateList[0] = &terraform.State{}
	err := json.Unmarshal([]byte(vpc), tfStateList[0])
	require.NoError(t, err)

	infra, err := NewInfra(tfStateList)
	require.NotNil(t, infra)
	require.NoError(t, err)

	resource := infra.FindById("vpc-ff5fec97")
	require.NotNil(t, resource)
	assert.Equal(t, "vpc-ff5fec97", resource.Id())
}
