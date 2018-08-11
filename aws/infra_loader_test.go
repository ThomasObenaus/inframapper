package aws

import (
	"fmt"
	"log"
	"testing"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/stretchr/testify/assert"
	"github.com/thomasobenaus/inframapper/trace"
)

type mockEC2IF struct {
}

func (m *mockEC2IF) DescribeVpcs(input *ec2.DescribeVpcsInput) (*ec2.DescribeVpcsOutput, error) {
	return nil, fmt.Errorf("N/A")
}

func TestRLNew(t *testing.T) {

	sl, err := NewInfraLoader("", "")
	assert.NotNil(t, err)
	assert.Nil(t, sl)

	sl, err = NewInfraLoader("blubb", "")
	assert.NotNil(t, err)
	assert.Nil(t, sl)
}

func TestRLLoad(t *testing.T) {

	loader := infraLoaderImpl{
		tracer: trace.Off(),
		ec2IF:  &mockEC2IF{},
	}

	infra, err := loader.Load()
	assert.NoError(t, err)
	assert.NotNil(t, infra)
}

func ExampleNewInfraLoader() {
	iLoader, err := NewInfraLoader("playground", "eu-central-1")
	if err != nil {
		log.Fatalf("Error, creating infra-loader: %s", err.Error())
	}

	infra, err := iLoader.Load()
	if err != nil {
		log.Fatalf("Error, loading infra: %s", err.Error())
	}
	fmt.Println(infra.String())
}
