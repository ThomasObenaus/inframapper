package tfstate

import (
	"fmt"
	"io/ioutil"

	terraform "github.com/hashicorp/terraform/terraform"
	"github.com/thomas.obenaus/terrastate/trace"
)

type StateLoader interface {
	Load(filename string) (*terraform.State, error)
}

type tfStateLoader struct {
	tracer trace.Tracer
}

func (sm *tfStateLoader) Load(filename string) (*terraform.State, error) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Error reading file %s: %s", filename, err.Error())
	}

	return Parse(data)
}

// NewStateLoader creates a new instance of a StateLoader without tracing
func NewStateLoader() StateLoader {
	return NewStateLoaderWithTracer(nil)
}

// NewStateLoaderWithTracer creates a new instance of a StateLoader with tracing
func NewStateLoaderWithTracer(tracer trace.Tracer) StateLoader {
	if tracer == nil {
		tracer = trace.Off()
	}
	return &tfStateLoader{
		tracer: tracer,
	}
}
