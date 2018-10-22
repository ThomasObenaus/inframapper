package mappedInfra

import (
	"github.com/thomasobenaus/inframapper/terraform"
	"github.com/thomasobenaus/inframapper/trace"

	"github.com/thomasobenaus/inframapper/aws"
)

// Mapper is a interface for mapping AWS resources
// to the corresponding terraform code.
type Mapper interface {

	// Map maps the given AWS resources and the terraform state.
	Map(aws aws.Infra, tf terraform.Infra) (Infra, error)
	String() string
}

type mapperImpl struct {
	tracer trace.Tracer
}

func (m *mapperImpl) String() string {
	return "MappedInfra"
}

func (m *mapperImpl) mapVpcs(vpcs []*aws.Vpc, tf terraform.Infra) []MappedResource {
	var mappedResources []MappedResource
	// handle vpcs
	m.tracer.Trace("Map (", len(vpcs), ") vpcs:")
	for _, awsVpc := range vpcs {
		if awsVpc == nil {
			m.tracer.Trace("Ignore vpc which is nil")
			continue
		}

		tfResource := tf.FindById(awsVpc.Id())
		mapFrom := awsVpc.Id()
		mappedToTf := "N/A"
		if tfResource != nil {
			mappedToTf = tfResource.Name()
		}
		m.tracer.Trace("\t", mapFrom, "->", mappedToTf)
		mResource := NewVpc(awsVpc, tfResource)
		mappedResources = append(mappedResources, mResource)
	}
	return mappedResources
}

func (m *mapperImpl) Map(aws aws.Infra, tf terraform.Infra) (Infra, error) {
	var mappedResources []MappedResource
	// handle vpcs
	mappedResources = append(mappedResources, m.mapVpcs(aws.Vpcs(), tf)...)
	return NewInfraWithTracer(mappedResources, m.tracer)
}

// NewMapperWithTracer creates a mapper that can be used to map the AWS resources
// with the corresponding terraform code.
func NewMapperWithTracer(tracer trace.Tracer) Mapper {
	if tracer == nil {
		tracer = trace.Off()
	}

	return &mapperImpl{
		tracer: tracer,
	}
}

// NewMapper creates a mapper that can be used to map the AWS resources
// with the corresponding terraform code.
func NewMapper() Mapper {
	return NewMapperWithTracer(nil)
}
