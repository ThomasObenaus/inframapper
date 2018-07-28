package mappedInfra

import (
	"github.com/thomas.obenaus/terrastate/terraform"
	"github.com/thomas.obenaus/terrastate/trace"

	"github.com/thomas.obenaus/terrastate/aws"
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
	mappedResources := make([]MappedResource, 0)
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

	infra := &infraImpl{
		mappedResources: mappedResources,
	}

	return infra, nil
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
