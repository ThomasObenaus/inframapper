// Package aws contains code needed to query aws infrastructure. This means reading all relevant data
// from a given AWS account that is needed to reflec the current state of the resources.
// No resources on the account will be removed or created.
package aws

import (
	"fmt"
	"strconv"

	"github.com/thomasobenaus/inframapper/trace"
)

// Infra represents an interface reflecting all resources in the loaded
// AWS infrastructure.
// To obtain an Infra object you have to use the InfraLoader in order to
// read in the resources from an AWS account.
type Infra interface {

	// FindByID returns the AWS resource that matches the given id.
	FindByID(id string) Resource

	// FindByID returns the AWS VPC that matches the given id.
	FindVpc(id string) *Vpc

	// Vpcs returns all vpc's
	Vpcs() []*Vpc

	// NumResources returns the number of available resources.
	NumResources() int

	// Returns the AWS region for which the resources where read in.
	Region() string

	String() string
}

type infraData struct {
	profile string
	region  string

	vpcs []*Vpc
}

type infraImpl struct {
	tracer trace.Tracer

	data          *infraData
	resourcesByID map[string]Resource
}

func (infra *infraImpl) Region() string {
	if infra.data == nil || len(infra.data.region) == 0 {
		return "UNKNOWN"
	}
	return infra.data.region
}

func (infra *infraImpl) String() string {
	if infra.data == nil {
		return "INVALID"
	}

	return "[" + infra.data.profile + "] " + infra.Region() + ", #resources=" + strconv.Itoa(infra.NumResources())
}

func (infra *infraImpl) FindByID(id string) Resource {
	return infra.resourcesByID[id]
}

func (infra *infraImpl) NumResources() int {
	return len(infra.resourcesByID)
}

func (infra *infraImpl) FindVpc(id string) *Vpc {
	if infra.Vpcs() == nil {
		return nil
	}

	for _, vpc := range infra.Vpcs() {
		if vpc != nil && vpc.ID() == id {
			return vpc
		}
	}
	return nil
}

func (infra *infraImpl) Vpcs() []*Vpc {
	if infra.data == nil {
		infra.tracer.Error("Error: infra.data is nil, return nil.")
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

	resourcesByID, err := createResourcesByIDMap(data, tracer)
	if err != nil {
		return nil, err
	}

	return &infraImpl{
		tracer:        tracer,
		data:          data,
		resourcesByID: resourcesByID,
	}, nil

}

func newInfra(data *infraData) (Infra, error) {
	return newInfraWithTracer(data, nil)
}

func createResourcesByIDMap(data *infraData, tracer trace.Tracer) (map[string]Resource, error) {

	var empty = make(map[string]Resource)

	if data == nil {
		return empty, fmt.Errorf("Data is nil")
	}

	var result = make(map[string]Resource)

	for _, vpc := range data.vpcs {
		if vpc == nil {
			continue
		}
		result[vpc.ID()] = vpc
	}

	// TODO add mapping for more resources here
	return result, nil
}
