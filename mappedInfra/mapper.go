package mappedInfra

import (
	"github.com/thomasobenaus/inframapper/terraform"
	"github.com/thomasobenaus/inframapper/trace"

	"github.com/thomasobenaus/inframapper/aws"
)

type Mapper interface {
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

func NewMapperWithTracer(tracer trace.Tracer) Mapper {
	if tracer == nil {
		tracer = trace.Off()
	}

	return &mapperImpl{
		tracer: tracer,
	}
}

func NewMapper() Mapper {
	return NewMapperWithTracer(nil)
}
