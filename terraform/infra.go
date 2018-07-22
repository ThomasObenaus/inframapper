package terraform

import (
	"fmt"

	tf "github.com/hashicorp/terraform/terraform"
	"github.com/thomas.obenaus/terrastate/trace"
)

type Infra interface {
	FindById(id string) (Resource, error)
}

type infraImpl struct {
	tracer trace.Tracer

	data          *tf.State
	resourcesById map[string]Resource
}

func (infra *infraImpl) FindById(id string) (Resource, error) {
	// FIXME implement this one
	return nil, fmt.Errorf("N/A")
}

func newInfraWithTracer(data *tf.State, tracer trace.Tracer) (Infra, error) {

	if data == nil {
		return nil, fmt.Errorf("terraform state is nil")
	}

	if tracer == nil {
		tracer = trace.Off()
	}

	resourcesById := createResourcesByIdMap(data, tracer)
	return &infraImpl{
		tracer:        tracer,
		data:          data,
		resourcesById: resourcesById,
	}, nil

}

func newInfra(data *tf.State) (Infra, error) {
	return newInfraWithTracer(data, nil)
}

func createResourcesByIdMap(data *tf.State, tracer trace.Tracer) map[string]Resource {
	// FIXME implement this one
	return make(map[string]Resource)
}
