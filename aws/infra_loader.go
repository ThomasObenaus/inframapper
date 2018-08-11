package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/thomasobenaus/inframapper/trace"
)

type InfraLoader interface {
	Load() (Infra, error)
}

type infraLoaderImpl struct {
	session    *session.Session
	tracer     trace.Tracer
	awsProfile string
	awsRegion  string
	infra      Infra
}

func (sl *infraLoaderImpl) Load() (Infra, error) {

	if err := sl.Validate(); err != nil {
		return nil, err
	}

	// VPC - section
	sl.tracer.Trace("Load vpcs ...")
	vpcs, err := sl.loadVpcs()
	if err != nil {
		return nil, err
	}
	sl.tracer.Trace("Load vpcs (", len(vpcs), ")...done")

	// put the data together
	infraData := &infraData{
		profile: sl.awsProfile,
		region:  sl.awsRegion,
		vpcs:    vpcs,
	}

	// create the infra
	return newInfraWithTracer(infraData, sl.tracer)
}

// Validate if the preconditions to load the infrastructure are met
func (sl *infraLoaderImpl) Validate() error {
	if sl.session == nil {
		return fmt.Errorf("Session is nil")
	}

	if sl.tracer == nil {
		return fmt.Errorf("Tracer is nil")
	}
	return nil
}

// NewInfraLoader creates a InfraLoader instance
func NewInfraLoader(awsProfile string, awsRegion string) (InfraLoader, error) {
	return NewInfraLoaderWithTracer(awsProfile, awsRegion, nil)
}

// NewInfraLoaderWithTracer creates a InfraLoader instance using the given Tracer for logging
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
		session:    sess,
		tracer:     tracer,
		awsProfile: awsProfile,
		awsRegion:  awsRegion,
	}

	return resourdeLoader, nil
}
