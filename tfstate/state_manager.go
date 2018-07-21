package tfstate

import (
	"fmt"
	"io/ioutil"

	terraform "github.com/hashicorp/terraform/terraform"
	"github.com/thomas.obenaus/terrastate/trace"
)

type TFStateManager interface {
	Load(filename string) (*terraform.State, error)
}

type tfStateManager struct {
	tracer trace.Tracer
}

func (sm *tfStateManager) Load(filename string) (*terraform.State, error) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Error reading file %s: %s", filename, err.Error())
	}

	return Parse(data)
}

// NewStateManager creates a new instance of a TFStateManager without tracing
func NewStateManager() TFStateManager {

	return &tfStateManager{
		tracer: trace.Off(),
	}
}

// NewStateManagerWithTracer creates a new instance of a TFStateManager with tracing
func NewStateManagerWithTracer(tracer trace.Tracer) TFStateManager {

	return &tfStateManager{
		tracer: tracer,
	}
}
