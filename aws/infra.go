package aws

import (
	"fmt"

	"github.com/thomas.obenaus/terrastate/trace"
)

type Infra interface {
	FindById(id string) (Resource, error)

	FindVPC(id string) (*Vpc, error)
	Vpcs() []*Vpc
}

type infraData struct {
	vpcs []*Vpc
}

type infraImpl struct {
	tracer trace.Tracer

	data          *infraData
	resourcesById map[string]Resource
}

func (infra *infraImpl) FindById(id string) (Resource, error) {
	return nil, fmt.Errorf("N/A")
}

func (infra *infraImpl) FindVPC(id string) (*Vpc, error) {
	return nil, fmt.Errorf("N/A")
}

func (infra *infraImpl) Vpcs() []*Vpc {
	return nil
}

func newInfraWithTracer(data *infraData, tracer trace.Tracer) (Infra, error) {

	if data == nil {
		return nil, fmt.Errorf("InfraData is nil")
	}

	if tracer == nil {
		tracer = trace.Off()
	}

	resourcesById := createResourcesByIdMap(data, tracer)

	return &infraImpl{
		tracer:        tracer,
		data:          data,
		resourcesById: resourcesById,
	}, nil

}

func createResourcesByIdMap(data *infraData, tracer trace.Tracer) map[string]Resource {
	return make(map[string]Resource)
}
