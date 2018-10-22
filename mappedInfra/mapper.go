package mappedInfra

import (
	"fmt"

	hc_terraform "github.com/hashicorp/terraform/terraform"
	"github.com/thomasobenaus/inframapper/terraform"
	"github.com/thomasobenaus/inframapper/trace"

	"github.com/thomasobenaus/inframapper/aws"
	"github.com/thomasobenaus/inframapper/tfstate"
)

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

// LoadAndMap loads all resources of an AWS account and the terraform state that represents the code for this resources.
// this two infromation are then mapped in order to get a joined view about the linkage between code an realy deployed resources.
// The AWS account containing the resources is defined through the given awsProfile and awsRegion.
// The terraform state to be read in is specified by the stateBackend which can be a remote or a local one.
func LoadAndMap(awsProfile string, awsRegion string, stateBackend tfstate.StateBackend, tracer trace.Tracer) (Infra, error) {

	if tracer == nil {
		tracer = trace.Off()
	}

	// Step 1: Load the AWS infra
	tracer.Trace("Step 1: Loading AWS Infra ...")
	awsInfraLoader, err := aws.NewInfraLoaderWithTracer(awsProfile, awsRegion, tracer)
	if err != nil {
		return nil, fmt.Errorf("Error loading AWS infra: %s", err.Error())
	}

	awsInfra, err := awsInfraLoader.Load()
	if err != nil {
		return nil, fmt.Errorf("Error loading AWS infra: %s", err.Error())
	}
	tracer.Trace("Step 1: Loading AWS Infra ... done")

	// Step 2a: Load the terraform state for the infra
	tracer.Trace("\n\nStep 2a: Loading Terraform state ...")
	tfStateLoader := tfstate.NewStateLoaderWithTracer(tracer)

	var tfStateList []*hc_terraform.State
	if stateBackend.IsRemote() {
		tfStateList, err = tfStateLoader.LoadRemoteState(stateBackend.RemoteConfig())
		if err != nil {
			return nil, fmt.Errorf("Error loading remote terraform state: %s", err.Error())
		}
	} else {
		return nil, fmt.Errorf("Error loading local state is not implemented yet")
	}

	tracer.Trace("Step 2a: Loading Terraform state ... done")

	// Step 2b: Create terraform infra representation based on the loaded state.
	tracer.Trace("\n\nStep 2a: Step 2b: Create terraform infra representation based on the loaded state ...")
	tfInfra, err := terraform.NewInfraWithTracer(tfStateList, tracer)
	if err != nil {
		return nil, fmt.Errorf("Error creatin terraform infra: %s", err.Error())
	}
	tracer.Trace("\n\nStep 2a: Step 2b: Create terraform infra representation based on the loaded state ... done")

	// Step 3: Map terraform state and aws infra
	tracer.Trace("\n\nStep 3: Map terraform state and aws infra ...")
	mapper := NewMapperWithTracer(tracer)
	mappedInfra, err := mapper.Map(awsInfra, tfInfra)
	if err != nil {
		return nil, fmt.Errorf("Error mapping infra: %s", err.Error())
	}
	tracer.Trace("\n\nStep 3: Map terraform state and aws infra ... done")

	return mappedInfra, nil
}
