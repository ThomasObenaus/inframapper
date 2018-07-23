package aws

import (
	"fmt"

	"github.com/thomas.obenaus/terrastate/trace"
)

type Infra interface {
	FindById(id string) Resource

	FindVpc(id string) *Vpc
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

func (infra *infraImpl) FindById(id string) Resource {
	return infra.resourcesById[id]
}

func (infra *infraImpl) FindVpc(id string) *Vpc {
	if infra.Vpcs() == nil {
		return nil
	}

	for _, vpc := range infra.Vpcs() {
		if vpc != nil && vpc.Id() == id {
			return vpc
		}
	}
	return nil
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

	resourcesById, err := createResourcesByIdMap(data, tracer)
	if err != nil {
		return nil, err
	}

	return &infraImpl{
		tracer:        tracer,
		data:          data,
		resourcesById: resourcesById,
	}, nil

}

func newInfra(data *infraData) (Infra, error) {
	return newInfraWithTracer(data, nil)
}

func createResourcesByIdMap(data *infraData, tracer trace.Tracer) (map[string]Resource, error) {

	var empty = make(map[string]Resource)

	if data == nil {
		return empty, fmt.Errorf("Data is nil")
	}

	var result = make(map[string]Resource)

	for _, vpc := range data.vpcs {
		if vpc == nil {
			continue
		}
		result[vpc.Id()] = vpc
	}

	// TODO add mapping for more resources here
	return result, nil
}
