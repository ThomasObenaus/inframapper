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
	m.tracer.Info("Map (", len(vpcs), ") vpcs")
	for _, awsVpc := range vpcs {
		if awsVpc == nil {
			m.tracer.Warn("Ignore vpc which is nil")
			continue
		}

		tfResource := tf.FindByID(awsVpc.ID())
		mapFrom := awsVpc.ID()
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
// this two information are then mapped in order to get a joined view about the linkage between code an really deployed resources.
// The AWS account containing the resources is defined through the given awsProfile and awsRegion.
// The terraform state to be read in is specified by the stateBackend which can be a remote or a local one.
func LoadAndMap(awsProfile string, awsRegion string, stateBackend tfstate.StateBackend, tracer trace.Tracer) (Infra, error) {

	if tracer == nil {
		tracer = trace.Off()
	}

	// Step 1: Load the AWS infra
	tracer.Info("Step 1: Loading AWS Infra ...")
	awsInfraLoader, err := aws.NewInfraLoaderWithTracer(awsProfile, awsRegion, tracer)
	if err != nil {
		return nil, fmt.Errorf("Error loading AWS infra: %s", err.Error())
	}

	awsInfra, err := awsInfraLoader.Load()
	if err != nil {
		return nil, fmt.Errorf("Error loading AWS infra: %s", err.Error())
	}
	tracer.Info("Step 1: Loading AWS Infra ... done\n")

	// Step 2a: Load the terraform state for the infra
	tracer.Info("Step 2a: Loading Terraform state ...")
	tfStateLoader := tfstate.NewStateLoaderWithTracer(tracer)

	var tfStateList []*hc_terraform.State
	if stateBackend.IsRemote() {
		tfStateList, err = tfStateLoader.LoadRemoteState(stateBackend.RemoteConfig())
		if err != nil {
			return nil, fmt.Errorf("Error loading remote terraform state: %s", err.Error())
		}
	} else {
		tfStateList, err = tfStateLoader.LoadFiles(stateBackend.LocalConfig().Files)
		if err != nil {
			return nil, fmt.Errorf("Error loading local terraform state: %s", err.Error())
		}
	}

	tracer.Info("Step 2a: Loading Terraform state ... done\n")

	// Step 2b: Create terraform infra representation based on the loaded state.
	tracer.Info("Step 2a: Step 2b: Create terraform infra representation based on the loaded state ...")
	tfInfra, err := terraform.NewInfraWithTracer(tfStateList, tracer)
	if err != nil {
		return nil, fmt.Errorf("Error creatin terraform infra: %s", err.Error())
	}
	tracer.Info("Step 2a: Step 2b: Create terraform infra representation based on the loaded state ... done\n")

	// Step 3: Map terraform state and aws infra
	tracer.Info("Step 3: Map terraform state and aws infra ...")
	mapper := NewMapperWithTracer(tracer)
	mappedInfra, err := mapper.Map(awsInfra, tfInfra)
	if err != nil {
		return nil, fmt.Errorf("Error mapping infra: %s", err.Error())
	}
	tracer.Info("Step 3: Map terraform state and aws infra ... done\n")

	return mappedInfra, nil
}
