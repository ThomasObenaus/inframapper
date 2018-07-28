package mappedInfra

import (
	"fmt"

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
	return nil, fmt.Errorf("N/A")
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
