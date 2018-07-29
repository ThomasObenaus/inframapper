package mappedInfra

import (
	"github.com/thomas.obenaus/inframapper/terraform"
	"github.com/thomas.obenaus/inframapper/trace"

	"github.com/thomas.obenaus/inframapper/aws"
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

func (m *mapperImpl) Map(aws aws.Infra, tf terraform.Infra) (Infra, error) {
	var mappedResources []MappedResource
	// handle vpcs
	for _, awsVpc := range aws.Vpcs() {
		if awsVpc == nil {
			m.tracer.Trace("Ignore vpc which is nil")
			continue
		}

		m.tracer.Trace("Map ", awsVpc.String())
		tfResource := tf.FindById(awsVpc.Id())

		mResource := NewVpc(awsVpc, tfResource)
		mappedResources = append(mappedResources, mResource)
	}
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
