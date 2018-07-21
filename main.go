package main

import (
	"log"

	"github.com/thomas.obenaus/terrastate/awsstate"
)

func main() {

	profile := "shared"
	region := "us-east-1"

	awsSl, err := awsstate.NewStateLoader(profile, region)
	if err != nil {
		log.Fatalf("Error creating StateLoader for AWS: %s", err.Error())
	}

	if err := awsSl.Load(); err != nil {
		log.Fatalf("Error loading AWS state: %s", err.Error())

	}

}
