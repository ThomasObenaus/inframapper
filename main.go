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

	awsInfraLoader, err := aws.NewInfraLoaderWithTracer(profile, region, tracer)
	if err != nil {
		log.Fatalf("Error creating InfraLoader for AWS: %s", err.Error())
	}

	awsInfra, err := awsInfraLoader.Load()
	if err != nil {
		log.Fatalf("Error loading AWS infra: %s", err.Error())
	}
	tracer.Trace("AWS Infra: ", awsInfra)

	tfStateLoader := tfstate.NewStateLoaderWithTracer(tracer)
	_, err = tfStateLoader.Load("examples/statefiles/instance.tfstate")
	if err != nil {
		log.Fatalf("Error loading terraform state: %s", err.Error())
	}

	keys := make([]string, 2)
	keys[0] = "snapshot/base/networking/terraform.tfstate"
	keys[1] = "snapshot/base/common/terraform.tfstate"
	remoteCfg := tfstateIf.RemoteConfig{
		BucketName: "741125603121-tfstate",
		Keys:       keys,
		Profile:    "shared",
		Region:     "eu-central-1",
	}

	tfStateList, err := tfStateLoader.LoadRemoteState(remoteCfg)
	if err != nil {
		log.Fatalf("Error loading remote terraform state: %s", err.Error())
	}

	tfInfra, err := terraform.NewInfraWithTracer(tfStateList, tracerOff)
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
