package terraform

import (
	"testing"

	terraform "github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {

	infra, err := newInfra(nil)
	assert.Nil(t, infra)
	assert.Error(t, err)

	tfstate := &terraform.State{}
	infra, err = newInfra(tfstate)
	assert.NotNil(t, infra)
	assert.NoError(t, err)
}
