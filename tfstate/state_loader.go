package tfstate

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	terraform "github.com/hashicorp/terraform/terraform"
	"github.com/thomas.obenaus/inframapper/trace"
)

type StateLoader interface {
	// Load loads a terraform state file
	Load(filename string) (*terraform.State, error)

	// LoadRemoteState loads state from an aws S3 bucket
	LoadRemoteState(remoteCfg RemoteConfig) ([]*terraform.State, error)
}

type tfStateLoader struct {
	tracer trace.Tracer
}

func (sl *tfStateLoader) Validate() error {
	if sl.tracer == nil {
		return fmt.Errorf("Tracer is nil")
	}
	return nil
}

func (sl *tfStateLoader) LoadRemoteState(remoteCfg RemoteConfig) ([]*terraform.State, error) {

	tfStateList := make([]*terraform.State, 0)

	// create session
	verboseCredErrors := true
	cfg := aws.Config{Region: aws.String(remoteCfg.Region), CredentialsChainVerboseErrors: &verboseCredErrors}
	sessionOpts := session.Options{Profile: remoteCfg.Profile, Config: cfg}
	session, err := session.NewSessionWithOptions(sessionOpts)
	if err != nil {
		return tfStateList, err
	}

	// Create a downloader with the session and default options
	downloader := s3manager.NewDownloader(session)
	var buffer []byte
	wBuffer := aws.NewWriteAtBuffer(buffer)

	for _, key := range remoteCfg.Keys {
		log.Printf("bucket: %s", remoteCfg.BucketName)
		log.Printf("key: %s", key)

		// Write the contents of S3 Object to the file
		n, err := downloader.Download(wBuffer, &s3.GetObjectInput{
			Bucket: aws.String(remoteCfg.BucketName),
			Key:    aws.String(remoteCfg.Keys[0]),
		})

		if err != nil {
			return nil, fmt.Errorf("Failed to download state-file, %v", err)
		}

		fmt.Printf("file downloaded, %d bytes\n", n)
		tfState, err := Parse(wBuffer.Bytes())
		if err != nil {
			return nil, err
		}

		tfStateList = append(tfStateList, tfState)
	}

	return tfStateList, nil
}

func (sl *tfStateLoader) Load(filename string) (*terraform.State, error) {

	if err := sl.Validate(); err != nil {
		return nil, err
	}

	sl.tracer.Trace("Loading tfstate from '", filename, "'...")
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Error reading file %s: %s", filename, err.Error())
	}

	sl.tracer.Trace("Parse state...")
	tfstate, err := Parse(data)
	sl.tracer.Trace("Parse state...done")

	sl.tracer.Trace("Loading tfstate from '", filename, "'...done")
	return tfstate, err
}

// NewStateLoader creates a new instance of a StateLoader without tracing
func NewStateLoader() StateLoader {
	return NewStateLoaderWithTracer(nil)
}

// NewStateLoaderWithTracer creates a new instance of a StateLoader with tracing
func NewStateLoaderWithTracer(tracer trace.Tracer) StateLoader {
	if tracer == nil {
		tracer = trace.Off()
	}
	return &tfStateLoader{
		tracer: tracer,
	}
}
