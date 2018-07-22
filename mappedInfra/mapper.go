package mappedInfra

import (
	"fmt"

	"github.com/thomas.obenaus/terrastate/trace"

	terraform "github.com/hashicorp/terraform/terraform"
	"github.com/thomas.obenaus/terrastate/aws"
)

type Mapper interface {
	Map() (Infra, error)
}

type mapperImpl struct {
	aws     aws.Infra
	tfstate *terraform.State
	tracer  trace.Tracer
}

func (m *mapperImpl) Map() (Infra, error) {
	return nil, fmt.Errorf("N/A")
}

func NewMapperWithTracer(aws aws.Infra, tf *terraform.State, tracer trace.Tracer) (Mapper, error) {

	if aws == nil && tf == nil {
		return nil, fmt.Errorf("Both aws infra and terraform state are nil")
	}

	if tracer == nil {
		tracer = trace.Off()
	}

	return &mapperImpl{
		aws:     aws,
		tfstate: tf,
		tracer:  tracer,
	}, nil
}

func NewMapper(aws aws.Infra, tf *terraform.State) (Mapper, error) {
	return NewMapperWithTracer(aws, tf, nil)
}
