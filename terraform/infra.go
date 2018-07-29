package terraform

import (
	"fmt"
	"strconv"

	tf "github.com/hashicorp/terraform/terraform"
	"github.com/thomas.obenaus/inframapper/trace"
)

type Infra interface {
	FindById(id string) Resource
	FindByName(id string) Resource
	NumResources() int
	String() string
}

type infraImpl struct {
	tracer trace.Tracer

	data            *tf.State
	resourcesById   map[string]Resource
	resourcesByName map[string]Resource
}

func (infra *infraImpl) NumResources() int {
	return len(infra.resourcesByName)
}

func (infra *infraImpl) String() string {
	return "#res=" + strconv.Itoa(infra.NumResources())
}

func (infra *infraImpl) FindById(id string) Resource {
	return infra.resourcesById[id]
}

func (infra *infraImpl) FindByName(name string) Resource {
	return infra.resourcesByName[name]
}

func NewInfraWithTracer(data *tf.State, tracer trace.Tracer) (Infra, error) {

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

func NewInfra(data *tf.State) (Infra, error) {
	return NewInfraWithTracer(data, nil)
}

// Type used to define which filter criteria should be used
type filterCriteria int

const (
	filterCriteria_Id filterCriteria = iota
	filterCriteria_Name
)

func createResourcesByXMap(data *tf.State, fCrit filterCriteria, tracer trace.Tracer) (map[string]Resource, error) {
	var empty = make(map[string]Resource)

	if data == nil {
		return empty, fmt.Errorf("Data is nil")
	}

	numModules := len(data.Modules)
	if numModules == 0 {
		tracer.Trace("No modules given")
		return empty, nil
	}
	tracer.Trace("#Modules=", strconv.Itoa(numModules))

	var result = make(map[string]Resource)
	// all modules
	for idx, module := range data.Modules {
		resources := module.Resources
		if len(resources) == 0 {
			tracer.Trace("No resources given for module ", idx, "/", numModules)
			continue
		}

		// all resources
		for name, resource := range resources {

			r := &resourceImpl{
				id:        resource.Primary.ID,
				rType:     StrToType(resource.Type),
				name:      name,
				dependsOn: resource.Dependencies,
				provider:  resource.Provider,
			}

			key := name
			if fCrit == filterCriteria_Id {
				key = r.Id()
			}

			tracer.Trace("Add resource ", r.String())
			result[key] = r
		}
	}

	return result, nil
}

func createResourcesByNameMap(data *tf.State, tracer trace.Tracer) (map[string]Resource, error) {
	return createResourcesByXMap(data, filterCriteria_Name, tracer)
}

func createResourcesByIdMap(data *tf.State, tracer trace.Tracer) (map[string]Resource, error) {
	return createResourcesByXMap(data, filterCriteria_Id, tracer)
}
