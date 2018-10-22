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
