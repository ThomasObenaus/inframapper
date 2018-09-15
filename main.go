package main

import (
	"log"
	"os"

	"github.com/thomasobenaus/inframapper/aws"
	"github.com/thomasobenaus/inframapper/mappedInfra"
	"github.com/thomasobenaus/inframapper/terraform"
	"github.com/thomasobenaus/inframapper/tfstate"
	tfstateIf "github.com/thomasobenaus/inframapper/tfstate/iface"
	"github.com/thomasobenaus/inframapper/trace"
)

func main() {

	profile := "develop"
	region := "eu-central-1"
	tracer := trace.New(os.Stdout)
	tracerOff := trace.Off()

	// load the aws infra
	tracer.Trace("Loading AWS Infra ...")
	awsInfraLoader, err := aws.NewInfraLoaderWithTracer(profile, region, tracer)
	if err != nil {
		log.Fatalf("Error creating InfraLoader for AWS: %s", err.Error())
	}

	awsInfra, err := awsInfraLoader.Load()
	if err != nil {
		log.Fatalf("Error loading AWS infra: %s", err.Error())
	}
	tracer.Trace("Loading AWS Infra ... done")
	tracer.Trace("AWS Infra: ", awsInfra)

	// Load the terraform state for the infra
	tracer.Trace("\n\nLoading Terraform state ...")
	keys := make([]string, 2)
	keys[0] = "snapshot/base/networking/terraform.tfstate"
	keys[1] = "snapshot/base/common/terraform.tfstate"
	remoteCfg := tfstateIf.RemoteConfig{
		BucketName: "741125603121-tfstate",
		Keys:       keys,
		Profile:    "shared",
		Region:     "eu-central-1",
	}

	tfStateLoader := tfstate.NewStateLoaderWithTracer(tracer)
	tfStateList, err := tfStateLoader.LoadRemoteState(remoteCfg)
	if err != nil {
		log.Fatalf("Error loading remote terraform state: %s", err.Error())
	}

	tfInfra, err := terraform.NewInfraWithTracer(tfStateList, tracerOff)
	if err != nil {
		log.Fatalf("Error loading terraform infrastructure: %s", err.Error())
	}
	tracer.Trace("Loading Terraform state ... done")
	tracer.Trace("Terraform Infra: ", tfInfra)

	// Mapping tf-state and aws infra
	tracer.Trace("\n\nMapping tf-state and aws infra ...")
	mapper := mappedInfra.NewMapperWithTracer(tracer)
	mappedInfra, err := mapper.Map(awsInfra, tfInfra)
	if err != nil {
		log.Fatalf("Error loading terraform infrastructure: %s", err.Error())
	}
	tracer.Trace("Mapping tf-state and aws infra ... done")

	var mappedInfraStr string
	var unMappedInfraStr string
	var unMappedTfStateStr string
	for _, res := range mappedInfra.Resources() {

		resStr := "\t" + res.String() + "\n"
		if res.IsMapped() {
			mappedInfraStr += resStr
		} else if res.HasAws() {
			unMappedInfraStr += resStr
		} else {
			unMappedTfStateStr += resStr
		}
	}
	tracer.Trace("Mapped Infra [", mappedInfra.NumResources(), "]:")
	tracer.Trace("Mapped:")
	tracer.Trace(mappedInfraStr)
	tracer.Trace("Unmapped (aws only):")
	tracer.Trace(unMappedInfraStr)
	tracer.Trace("Unmapped (terraform only):")
	tracer.Trace(unMappedTfStateStr)

}
