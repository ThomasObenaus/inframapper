// Package mappedInfra contains code to map infrastructure resources available in an AWS account with
// the corresponding resources described in terraform code.
package mappedInfra

import (
	"strconv"

	"github.com/thomasobenaus/inframapper/trace"
)

// Infra is an interface reflecting the mapping between real infrastructure
// resources living on AWS with the according resource description in
// terraform code.
// To obtain an Infra object you have to load terraform state, read AWS resource
// information and use this as input for NewInfraWithTracer().
type Infra interface {
	// NumResources returns the number of resources.
	NumResources() int

	// Resources returns all ressources those that could be mapped and those that could not.
	Resources() []MappedResource

	// AwsResourceByID returns the AWS resource that matches the given id.
	AwsResourceByID(id string) MappedResource

	// MappedResources returns all ressources that could be mapped.
	MappedResources() []MappedResource

	// MappedResources returns all ressources that could NOT be mapped.
	UnMappedAwsResources() []MappedResource
	String() string
}

type infraImpl struct {
	awsResourcesByID     map[string]MappedResource
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

func (in *infraImpl) AwsResourceByID(id string) MappedResource {
	return in.awsResourcesByID[id]
}

// NewInfraWithTracer creates an Infra object reflecting the mapping between AWS resources
// and the corresponding terraform code.
func NewInfraWithTracer(resources []MappedResource, tracer trace.Tracer) (Infra, error) {
	if tracer == nil {
		tracer = trace.Off()
	}

	var mappedResources []MappedResource
	var unMappedAwsResources []MappedResource

	awsResourcesByID := make(map[string]MappedResource)
	for _, mResource := range resources {

		if mResource.IsMapped() {
			awsResourcesByID[mResource.Aws().ID()] = mResource
			mappedResources = append(mappedResources, mResource)
		} else if mResource.HasAws() {
			awsResourcesByID[mResource.Aws().ID()] = mResource
			unMappedAwsResources = append(unMappedAwsResources, mResource)
		} else {
			tracer.Trace("Ignore resource since it has no aws data: ", mResource.String())
		}
	}

	return &infraImpl{
		awsResourcesByID:     awsResourcesByID,
		resources:            resources,
		mappedResources:      mappedResources,
		unMappedAwsResources: unMappedAwsResources,
		tracer:               tracer}, nil
}

// NewInfra creates an Infra object reflecting the mapping between AWS resources
// and the corresponding terraform code.
func NewInfra(resources []MappedResource) (Infra, error) {
	return NewInfraWithTracer(resources, nil)
}
