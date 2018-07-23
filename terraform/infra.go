package terraform

import (
	"fmt"

	tf "github.com/hashicorp/terraform/terraform"
	"github.com/thomas.obenaus/terrastate/trace"
)

type Infra interface {
	FindById(id string) Resource
	FindByName(id string) Resource
}

type infraImpl struct {
	tracer trace.Tracer

	data            *tf.State
	resourcesById   map[string]Resource
	resourcesByName map[string]Resource
}

func (infra *infraImpl) FindById(id string) Resource {
	return infra.resourcesById[id]
}

func (infra *infraImpl) FindByName(name string) Resource {
	return infra.resourcesByName[name]
}

func newInfraWithTracer(data *tf.State, tracer trace.Tracer) (Infra, error) {

	if data == nil {
		return nil, fmt.Errorf("terraform state is nil")
	}

	if tracer == nil {
		tracer = trace.Off()
	}

	resourcesById, err := createResourcesByIdMap(data, tracer)
	if err != nil {
		return nil, err
	}

	resourcesByName, err := createResourcesByNameMap(data, tracer)
	if err != nil {
		return nil, err
	}

	return &infraImpl{
		tracer:          tracer,
		data:            data,
		resourcesById:   resourcesById,
		resourcesByName: resourcesByName,
	}, nil

}

func newInfra(data *tf.State) (Infra, error) {
	return newInfraWithTracer(data, nil)
}

func createResourcesByNameMap(data *tf.State, tracer trace.Tracer) (map[string]Resource, error) {

	var empty = make(map[string]Resource)

	if data == nil {
		return empty, fmt.Errorf("Data is nil")
	}

	if len(data.Modules) == 0 {
		tracer.Trace("No modules given")
		return empty, nil
	}

	module := data.Modules[0]
	resources := module.Resources
	if len(resources) == 0 {
		tracer.Trace("No resources given")
		return empty, nil
	}

	var result = make(map[string]Resource)

	for name, resource := range resources {

		r := &resourceImpl{
			id:        resource.Primary.ID,
			rType:     StrToType(resource.Type),
			name:      name,
			dependsOn: resource.Dependencies,
			provider:  resource.Provider,
		}

		tracer.Trace("Add resource ", r.String())
		result[r.Name()] = r
	}

	return result, nil
}

func createResourcesByIdMap(data *tf.State, tracer trace.Tracer) (map[string]Resource, error) {

	var empty = make(map[string]Resource)

	if data == nil {
		return empty, fmt.Errorf("Data is nil")
	}

	if len(data.Modules) == 0 {
		tracer.Trace("No modules given")
		return empty, nil
	}

	module := data.Modules[0]
	resources := module.Resources
	if len(resources) == 0 {
		tracer.Trace("No resources given")
		return empty, nil
	}

	var result = make(map[string]Resource)

	for name, resource := range resources {

		r := &resourceImpl{
			id:        resource.Primary.ID,
			rType:     StrToType(resource.Type),
			name:      name,
			dependsOn: resource.Dependencies,
			provider:  resource.Provider,
		}

		tracer.Trace("Add resource ", r.String())
		result[r.Id()] = r
	}

	return result, nil
}
