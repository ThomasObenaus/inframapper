package mappedInfra

import (
	"github.com/thomas.obenaus/inframapper/trace"
)

type Infra interface {
	NumResources() int
	Resources() []MappedResource
	ResourceById(id string) MappedResource
}

type infraImpl struct {
	mappedResourcesById map[string]MappedResource
	mappedResources     []MappedResource
	tracer              trace.Tracer
}

func (in *infraImpl) NumResources() int {
	return len(in.mappedResources)
}

func (in *infraImpl) Resources() []MappedResource {
	return in.mappedResources
}

func (in *infraImpl) ResourceById(id string) MappedResource {
	return in.mappedResourcesById[id]
}

func NewInfraWithTracer(mappedResources []MappedResource, tracer trace.Tracer) (Infra, error) {
	if tracer == nil {
		tracer = trace.Off()
	}

	var mappedResourcesById map[string]MappedResource
	for _, mResource := range mappedResources {

		if !mResource.HasAws() {
			tracer.Trace("Ignore resource (res by id), since aws data is missing: ", mResource.String())
			continue
		}

		id := mResource.Aws().Id()
		if len(id) == 0 {
			tracer.Trace("Ignore resource (res by id), since id is missing: ", mResource.String())
			continue
		}

		mappedResourcesById[mResource.Aws().Id()] = mResource
	}

	return &infraImpl{
		mappedResourcesById: mappedResourcesById,
		mappedResources:     mappedResources,
		tracer:              tracer}, nil
}

func NewInfra(mappedResources []MappedResource) (Infra, error) {
	return NewInfraWithTracer(mappedResources, nil)
}
