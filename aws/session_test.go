package aws

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thomasobenaus/inframapper/helper"
)

func TestNewSession(t *testing.T) {
	session, err := newAWSSession("invalid", "my-region")
	assert.NoError(t, err)
	assert.NotNil(t, session)
	assert.Equal(t, session.Config.CredentialsChainVerboseErrors, helper.NewTrue())
	assert.Equal(t, *session.Config.Region, "my-region")
}
