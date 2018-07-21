package awsstate

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSLNew(t *testing.T) {
	sl, err := NewStateLoader("playground", "eu-central-1")
	assert.NotNil(t, sl)
	assert.Nil(t, err)

	sl, err = NewStateLoader("", "")
	assert.NotNil(t, err)
	assert.Nil(t, sl)

	sl, err = NewStateLoader("blubb", "")
	assert.NotNil(t, err)
	assert.Nil(t, sl)
}

func TestSMLLoad(t *testing.T) {
	sl, err := NewStateLoader("playground", "eu-central-1")
	require.Nil(t, err)
	require.NotNil(t, sl)
	assert.Nil(t, sl.Load())

	sl, err = NewStateLoader("unknown", "unknown")
	require.Nil(t, err)
	require.NotNil(t, sl)
	assert.NotNil(t, sl.Load())

}
