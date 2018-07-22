package aws

import (
	"fmt"

	"github.com/thomas.obenaus/terrastate/trace"
)

type Infra interface {
	FindById(id string) (Resource, error)

	FindVpc(id string) (*Vpc, error)
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
	// FIXME implement this one
	return nil, fmt.Errorf("N/A")
}

func (infra *infraImpl) FindVpc(id string) (*Vpc, error) {
	// FIXME implement this one
	return nil, fmt.Errorf("N/A")
}

func (infra *infraImpl) Vpcs() []*Vpc {
	if infra.data == nil {
		infra.tracer.Trace("Error: infra.data is nil, return nil.")
		return nil
	}

	return infra.data.vpcs
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

func newInfra(data *infraData) (Infra, error) {
	return newInfraWithTracer(data, nil)
}

func createResourcesByIdMap(data *infraData, tracer trace.Tracer) map[string]Resource {
	// FIXME implement this one
	return make(map[string]Resource)
}
