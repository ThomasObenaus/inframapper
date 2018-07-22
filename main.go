package main

import (
	"log"
	"os"

	"github.com/thomas.obenaus/terrastate/aws"
	"github.com/thomas.obenaus/terrastate/trace"
)

func main() {

	profile := "playground"
	region := "eu-central-1"
	tracer := trace.New(os.Stdout)

	aws, err := aws.NewInfraLoaderWithTracer(profile, region, tracer)
	if err != nil {
		log.Fatalf("Error creating InfraLoader for AWS: %s", err.Error())
	}

	if err := aws.Load(); err != nil {
		log.Fatalf("Error loading AWS infra: %s", err.Error())

	}

}
