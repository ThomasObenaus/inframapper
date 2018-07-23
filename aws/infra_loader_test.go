package aws

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRLNew(t *testing.T) {
	sl, err := NewInfraLoader("playground", "eu-central-1")
	assert.NotNil(t, sl)
	assert.Nil(t, err)

	sl, err = NewInfraLoader("", "")
	assert.NotNil(t, err)
	assert.Nil(t, sl)

	sl, err = NewInfraLoader("blubb", "")
	assert.NotNil(t, err)
	assert.Nil(t, sl)
}

func TestRLLoad(t *testing.T) {
	sl, err := NewInfraLoader("playground", "eu-central-1")
	require.Nil(t, err)
	require.NotNil(t, sl)
	assert.Nil(t, sl.Load())

	sl, err = NewInfraLoader("unknown", "unknown")
	require.Nil(t, err)
	require.NotNil(t, sl)
	assert.NotNil(t, sl.Load())

}

func TestValidate(t *testing.T) {
	rl := infraLoaderImpl{}
	err := rl.Validate()
	assert.NotNil(t, err)

	session, err := newAWSSession("blubb", "bla")
	require.Nil(t, err)
	require.NotNil(t, session)
	rl = infraLoaderImpl{session: session}
	err = rl.Validate()
	assert.NotNil(t, err)

}

func ExampleNewInfraLoader() {
	iLoader, err := NewInfraLoader("playground", "eu-central-1")
	if err != nil {
		log.Fatalf("Error, creatging infra-loader: %s", err.Error())
	}

	if err := iLoader.Load(); err != nil {
		log.Fatalf("Error, loading infra: %s", err.Error())
	}

	infra := iLoader.GetLoadedInfra()
	if infra == nil {
		log.Fatalf("Error, obtaining loaded infra: %s", err.Error())
	}

	fmt.Println(infra.String())
}
