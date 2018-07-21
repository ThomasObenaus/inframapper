package awsstate

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
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

	ss3Ep := s3.New(sl.session)
	buckets, err := ss3Ep.ListBuckets(nil)
	if err != nil {
		return err
	}

	log.Printf("Buckerts: %s", buckets)

	return nil
}

func NewStateLoader(awsProfile string, awsRegion string) (StateLoader, error) {

	sess, err := newAWSSession(awsProfile, awsRegion)
	if err != nil {
		return nil, fmt.Errorf("Unable to create loader: %s", err.Error())
	}

	if sess == nil {
		return nil, fmt.Errorf("Unable to create loader: session is nil")
	}

	stateLoader := &stateLoaderImpl{
		session:    sess,
		tracer:     trace.Off(),
		awsProfile: awsProfile,
		awsRegion:  awsRegion,
	}

	return stateLoader, nil
}
