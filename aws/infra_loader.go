package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/thomasobenaus/inframapper/aws/iface"
	"github.com/thomasobenaus/inframapper/trace"
)

// InfraLoader is a simple interface for loading AWS resources
type InfraLoader interface {
	Load() (Infra, error)
}

type infraLoaderImpl struct {
	tracer     trace.Tracer
	awsProfile string
	awsRegion  string

	ec2IF iface.EC2IF
}

func (sl *infraLoaderImpl) Load() (Infra, error) {
	// VPC - section
	sl.tracer.Info("Load vpcs ...")
	vpcs, err := LoadVpcs(sl.ec2IF, sl.tracer)
	if err != nil {
		sl.tracer.Error("Error loading vpcs: ", err.Error())
	}
	sl.tracer.Info("Load vpcs (", len(vpcs), ") ... done")

	// put the data together
	infraData := &infraData{
		profile: sl.awsProfile,
		region:  sl.awsRegion,
		vpcs:    vpcs,
	}

	// create the infra
	return newInfraWithTracer(infraData, sl.tracer)
}

// NewInfraLoader creates a InfraLoader instance
func NewInfraLoader(awsProfile string, awsRegion string) (InfraLoader, error) {
	return NewInfraLoaderWithTracer(awsProfile, awsRegion, nil)
}

// NewInfraLoaderWithTracer creates a InfraLoader instance using the given Tracer for logging.
// With this InfraLoader data from an aws account can be loaded. The account + needed access credentials are
// defined by awsProfile and awsRegion.
func NewInfraLoaderWithTracer(awsProfile string, awsRegion string, tracer trace.Tracer) (InfraLoader, error) {

	if len(awsProfile) == 0 {
		return nil, fmt.Errorf("AWS profile is empty")
	}

	if len(awsRegion) == 0 {
		return nil, fmt.Errorf("AWS region is empty")
	}

	sess, err := newAWSSession(awsProfile, awsRegion)
	if err != nil {
		return nil, fmt.Errorf("Unable to create loader: %s", err.Error())
	}

	if sess == nil {
		return nil, fmt.Errorf("Unable to create loader: session is nil")
	}

	if tracer == nil {
		tracer = trace.Off()
	}

	resourdeLoader := &infraLoaderImpl{
		tracer:     tracer,
		awsProfile: awsProfile,
		awsRegion:  awsRegion,

		ec2IF: ec2.New(sess),
	}

	return resourdeLoader, nil
}
