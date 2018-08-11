package aws

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thomasobenaus/inframapper/trace"
)

func TestLoadVPCFail(t *testing.T) {
	sess, err := newAWSSession("unknown", "eu-central-1")
	require.NotNil(t, sess)
	require.Nil(t, err)

	sl := infraLoaderImpl{
		session: sess,
		tracer:  trace.Off(),
	}

	vpcs, err := sl.loadVpcs()
	assert.Nil(t, vpcs)
	assert.NotNil(t, err)

}

func TestLoadVPC(t *testing.T) {
	sess, err := newAWSSession("playground", "eu-central-1")
	require.NotNil(t, sess)
	require.Nil(t, err)

	sl := infraLoaderImpl{
		session: sess,
		tracer:  trace.Off(),
	}

	vpcs, err := sl.loadVpcs()
	assert.NotNil(t, vpcs)
	assert.Nil(t, err)

}
