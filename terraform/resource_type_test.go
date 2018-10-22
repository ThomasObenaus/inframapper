package terraform

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrToType(t *testing.T) {
	ty := StrToType("?")
	assert.Equal(t, Type_unknown, ty)

	ty = StrToType("aws_vpc")
	assert.Equal(t, Type_aws_vpc, ty)
}
