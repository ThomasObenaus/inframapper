package mappedInfra

import (
	"strconv"

	"github.com/thomasobenaus/inframapper/trace"
)

type Infra interface {
	NumResources() int
	Resources() []MappedResource
	AwsResourceById(id string) MappedResource
	MappedResources() []MappedResource
	UnMappedAwsResources() []MappedResource
	String() string
}

type infraImpl struct {
	awsResourcesById     map[string]MappedResource
	resources            []MappedResource
	mappedResources      []MappedResource
	unMappedAwsResources []MappedResource
	tracer               trace.Tracer
}

func (in *infraImpl) String() string {

	result := "#res=" + strconv.Itoa(in.NumResources())
	result += ", #mapped=" + strconv.Itoa(len(in.MappedResources()))
	result += ", #aws_res=" + strconv.Itoa(len(in.UnMappedAwsResources()))
	return result
}

func (in *infraImpl) UnMappedAwsResources() []MappedResource {
	return in.unMappedAwsResources
}

func (in *infraImpl) MappedResources() []MappedResource {
	return in.mappedResources
}

func (in *infraImpl) NumResources() int {
	return len(in.resources)
}

func (in *infraImpl) Resources() []MappedResource {
	return in.resources
}

func (in *infraImpl) AwsResourceById(id string) MappedResource {
	return in.awsResourcesById[id]
}

func NewInfraWithTracer(resources []MappedResource, tracer trace.Tracer) (Infra, error) {
	if tracer == nil {
		tracer = trace.Off()
	}

	var mappedResources []MappedResource
	var unMappedAwsResources []MappedResource

	awsResourcesById := make(map[string]MappedResource)
	for _, mResource := range resources {

		if mResource.IsMapped() {
			awsResourcesById[mResource.Aws().Id()] = mResource
			mappedResources = append(mappedResources, mResource)
		} else if mResource.HasAws() {
			awsResourcesById[mResource.Aws().Id()] = mResource
			unMappedAwsResources = append(unMappedAwsResources, mResource)
		} else {
			tracer.Trace("Ignore resource since it has no aws data: ", mResource.String())
		}
	}

	return &infraImpl{
		awsResourcesById:     awsResourcesById,
		resources:            resources,
		mappedResources:      mappedResources,
		unMappedAwsResources: unMappedAwsResources,
		tracer:               tracer}, nil
}

func NewInfra(resources []MappedResource) (Infra, error) {
	return NewInfraWithTracer(resources, nil)
}
