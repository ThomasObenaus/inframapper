package tfstate

import (
	"fmt"
	"io/ioutil"

	terraform "github.com/hashicorp/terraform/terraform"
	"github.com/thomas.obenaus/inframapper/trace"
)

type StateLoader interface {
	Load(filename string) (*terraform.State, error)
}

type tfStateLoader struct {
	tracer trace.Tracer
}

func (sl *tfStateLoader) Validate() error {
	if sl.tracer == nil {
		return fmt.Errorf("Tracer is nil")
	}
	return nil
}

func (sl *tfStateLoader) Load(filename string) (*terraform.State, error) {

	if err := sl.Validate(); err != nil {
		return nil, err
	}

	sl.tracer.Trace("Loading tfstate from '", filename, "'...")
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Error reading file %s: %s", filename, err.Error())
	}

	sl.tracer.Trace("Parse state...")
	tfstate, err := Parse(data)
	sl.tracer.Trace("Parse state...done")

	sl.tracer.Trace("Loading tfstate from '", filename, "'...done")
	return tfstate, err
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
