package awsstate

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/thomas.obenaus/terrastate/trace"
)

type StateLoader interface {
	Load() error
}

type stateLoaderImpl struct {
	session    *session.Session
	tracer     trace.Tracer
	awsProfile string
	awsRegion  string
}

func (sl *stateLoaderImpl) Load() error {

	_, err := sl.loadVpc()
	if err != nil {
		return err
	}

	return nil
}

func (sl *stateLoaderImpl) Validate() error {
	if sl.session == nil {
		return fmt.Errorf("Session is nil")
	}

	if sl.tracer == nil {
		return fmt.Errorf("Tracer is nil")
	}
	return nil
}

func NewStateLoader(awsProfile string, awsRegion string) (StateLoader, error) {
	return NewStateLoaderWithTracer(awsProfile, awsRegion, nil)
}

func NewStateLoaderWithTracer(awsProfile string, awsRegion string, tracer trace.Tracer) (StateLoader, error) {

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

	stateLoader := &stateLoaderImpl{
		session:    sess,
		tracer:     tracer,
		awsProfile: awsProfile,
		awsRegion:  awsRegion,
	}

	return stateLoader, nil
}
