package awsstate

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/thomas.obenaus/terrastate/trace"
)

type ResourceLoader interface {
	Load() error
	Validate() error
}

type resourceLoaderImpl struct {
	session    *session.Session
	tracer     trace.Tracer
	awsProfile string
	awsRegion  string
}

func (sl *resourceLoaderImpl) Load() error {

	_, err := sl.loadVpc()
	if err != nil {
		return err
	}

	return nil
}

func (sl *resourceLoaderImpl) Validate() error {
	if sl.session == nil {
		return fmt.Errorf("Session is nil")
	}

	if sl.tracer == nil {
		return fmt.Errorf("Tracer is nil")
	}
	return nil
}

// NewResourceLoader creates a ResourceLoader instance
func NewResourceLoader(awsProfile string, awsRegion string) (ResourceLoader, error) {
	return NewResourceLoaderWithTracer(awsProfile, awsRegion, nil)
}

// NewResourceLoaderWithTracer creates a ResourceLoader instance using the given Tracer for logging
func NewResourceLoaderWithTracer(awsProfile string, awsRegion string, tracer trace.Tracer) (ResourceLoader, error) {

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

	resourdeLoader := &resourceLoaderImpl{
		session:    sess,
		tracer:     tracer,
		awsProfile: awsProfile,
		awsRegion:  awsRegion,
	}

	return resourdeLoader, nil
}
