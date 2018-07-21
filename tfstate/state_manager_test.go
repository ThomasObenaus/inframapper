package tfstate

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSMNew(t *testing.T) {
	sm := NewStateManager()
	require.NotNil(t, sm)
}

func TestSMLoad(t *testing.T) {
	sm := NewStateManager()
	require.NotNil(t, sm)

	tfstate, err := sm.Load("ssss")
	assert.Error(t, err)
	assert.Nil(t, tfstate)

	tfstate, err = sm.Load("instance.tfstate")
	assert.NoError(t, err)
	assert.NotNil(t, tfstate)

}
