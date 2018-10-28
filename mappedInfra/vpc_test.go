package mappedInfra

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToVpc(t *testing.T) {

	mappedResource := &Vpc{}

	vpc, err := ToVpc(mappedResource)

	assert.NotNil(t, vpc)
	assert.NoError(t, err)
}
