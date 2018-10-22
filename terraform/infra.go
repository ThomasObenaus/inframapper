// Package terraform contains code representing and for interacting with resources that are described by terraform code.
package terraform

import (
	"fmt"
	"strconv"

	tf "github.com/hashicorp/terraform/terraform"
	"github.com/thomasobenaus/inframapper/trace"
)

// Infra is an interface providing access to all resources
// that where described by terraform code.
// To obtain a Infra object you have to obtain the terraform state
// using the tfstate.StateLoader first.
// Based on the read terraform state the Infra can be created via
// NewInfraWithTracer() or NewInfra()
type Infra interface {
	// FindById returns the resource that matches the given id string (i.e. 'vpc-f8168d93')
	FindById(id string) Resource

	// FindByName returns the resource that matches the given name as it is used in terraform code (i.e 'aws_vpc.vpc_main')
	FindByName(name string) Resource

	// NumResources returns the number of resources known.
	NumResources() int
	String() string
}

type infraImpl struct {
	tracer          trace.Tracer
	resourcesById   map[string]Resource
	resourcesByName map[string]Resource
}

func (infra *infraImpl) NumResources() int {
	return len(infra.resourcesByName)
}

func (infra *infraImpl) String() string {
	return "#resources=" + strconv.Itoa(infra.NumResources())
}

func (infra *infraImpl) FindById(id string) Resource {
	return infra.resourcesById[id]
}

func (infra *infraImpl) FindByName(name string) Resource {
	return infra.resourcesByName[name]
}

// NewInfraWithTracer creates a Infra object based on the given terraform state.
// The state can be read in using the tfstate.StateLoader.
func NewInfraWithTracer(data []*tf.State, tracer trace.Tracer) (Infra, error) {

	if data == nil {
		return nil, fmt.Errorf("terraform state is nil")
	}

	if tracer == nil {
		tracer = trace.Off()
	}

	resourcesById := make(map[string]Resource)
	resourcesByName := make(map[string]Resource)

	for _, tfState := range data {
		rById, err := createResourcesByIdMap(tfState, tracer)
		if err != nil {
			tracer.Trace("Error creating resourceByIdMap. Will skip. Err: ", err.Error())
		} else {
			resourcesById = mergeResourceMaps(resourcesById, rById)
		}

		rByName, err := createResourcesByNameMap(tfState, tracer)
		if err != nil {
			tracer.Trace("Error creating resourceByNameMap. Will skip. Err: ", err.Error())
		} else {
			resourcesByName = mergeResourceMaps(resourcesByName, rByName)
		}
	}

	return &infraImpl{
		tracer:          tracer,
		resourcesById:   resourcesById,
		resourcesByName: resourcesByName,
	}, nil

}

// NewInfra creates a Infra object based on the given terraform state.
// The state can be read in using the tfstate.StateLoader.
func NewInfra(data []*tf.State) (Infra, error) {
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

func mergeResourceMaps(map1 map[string]Resource, map2 map[string]Resource) map[string]Resource {
	for k, v := range map2 {
		map1[k] = v
	}
	return map1
}
