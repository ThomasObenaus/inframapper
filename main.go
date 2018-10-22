package main

import (
	"log"
	"os"

	"github.com/thomasobenaus/inframapper/mappedInfra"
	"github.com/thomasobenaus/inframapper/tfstate"
	"github.com/thomasobenaus/inframapper/trace"
)

func main() {

	profile := "develop"
	region := "eu-central-1"
	tracer := trace.New(os.Stdout)

	keys := make([]string, 2)
	keys[0] = "snapshot/base/networking/terraform.tfstate"
	keys[1] = "snapshot/base/common/terraform.tfstate"
	remoteCfg := tfstate.RemoteConfig{
		BucketName: "741125603121-tfstate",
		Keys:       keys,
		Profile:    "shared",
		Region:     "eu-central-1",
	}

	mappedInfra, err := mappedInfra.LoadAndMap(profile, region, remoteCfg, tracer)
	if err != nil {
		log.Fatalf("Error loading/ mapping infrastructure: %s", err.Error())
	}

	var mappedResStr string
	var unMappedAwsResStr string
	for _, res := range mappedInfra.MappedResources() {
		mappedResStr += "\t" + res.String() + "\n"
	}
	for _, res := range mappedInfra.UnMappedAwsResources() {
		unMappedAwsResStr += "\t" + res.String() + "\n"
	}

	tracer.Trace("Mapped Infra: ", mappedInfra)
	tracer.Trace("Mapped Resources [", len(mappedInfra.MappedResources()), "]:")
	tracer.Trace(mappedResStr)
	tracer.Trace("UnMapped AWS Resources [", len(mappedInfra.UnMappedAwsResources()), "]:")
	tracer.Trace(unMappedAwsResStr)

}
