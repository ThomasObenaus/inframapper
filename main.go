package main

import (
	"log"
	"os"

	"github.com/thomas.obenaus/terrastate/aws"
	"github.com/thomas.obenaus/terrastate/mappedInfra"
	"github.com/thomas.obenaus/terrastate/terraform"
	"github.com/thomas.obenaus/terrastate/tfstate"
	"github.com/thomas.obenaus/terrastate/trace"
)

func main() {

	profile := "playground"
	region := "eu-central-1"
	tracer := trace.New(os.Stdout)

	awsInfraLoader, err := aws.NewInfraLoaderWithTracer(profile, region, tracer)
	if err != nil {
		log.Fatalf("Error creating InfraLoader for AWS: %s", err.Error())
	}

	if err := awsInfraLoader.Load(); err != nil {
		log.Fatalf("Error loading AWS infra: %s", err.Error())
	}

	awsInfra := awsInfraLoader.GetLoadedInfra()
	tracer.Trace("AWS Infra: ", awsInfra)

	tfStateLoader := tfstate.NewStateLoaderWithTracer(tracer)
	tfState, err := tfStateLoader.Load("examples/statefiles/instance.tfstate")
	if err != nil {
		log.Fatalf("Error loading terraform state: %s", err.Error())
	}

	tfInfra, err := terraform.NewInfraWithTracer(tfState, tracer)
	if err != nil {
		log.Fatalf("Error loading terraform infrastructure: %s", err.Error())
	}

	tracer.Trace("Terraform Infra: ", tfInfra)

	mapper := mappedInfra.NewMapperWithTracer(tracer)
	mappedInfra, err := mapper.Map(awsInfra, tfInfra)
	if err != nil {
		log.Fatalf("Error loading terraform infrastructure: %s", err.Error())
	}

	tracer.Trace("Mapped Infra: ", mappedInfra)

}
