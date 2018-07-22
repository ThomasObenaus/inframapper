package aws

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thomas.obenaus/terrastate/trace"
)

func TestFindById(t *testing.T) {

	data := &infraData{}
	infra := &infraImpl{tracer: trace.Off(), data: data}

	resource, err := infra.FindById("ABCD")
	assert.Error(t, err)
	assert.Nil(t, resource)

	// TODO add positive tests
}

func TestFindVpc(t *testing.T) {

	data := &infraData{}
	infra := &infraImpl{tracer: trace.Off(), data: data}

	vpc, err := infra.FindVpc("ABCD")
	assert.Error(t, err)
	assert.Nil(t, vpc)

	// TODO add positive tests
}

func TestNew(t *testing.T) {
	infra, err := newInfra(nil)
	assert.Error(t, err)
	assert.Nil(t, infra)

	data := &infraData{}
	infra, err = newInfra(data)
	assert.NoError(t, err)
	assert.NotNil(t, infra)
}

func TestVpcs(t *testing.T) {

	data := &infraData{}
	infra := &infraImpl{tracer: trace.Off(), data: data}

	vpcs := infra.Vpcs()
	assert.Nil(t, vpcs)

	infra = &infraImpl{tracer: trace.Off()}
	vpcs = infra.Vpcs()
	assert.Nil(t, vpcs)

	// TODO add positive tests
}
