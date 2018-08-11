package tfstate

import (
	"fmt"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	terraform "github.com/hashicorp/terraform/terraform"
	"github.com/thomasobenaus/inframapper/tfstate/iface"
	"github.com/thomasobenaus/inframapper/trace"
)

type tfStateLoader struct {
	tracer trace.Tracer
}

func (sl *tfStateLoader) Validate() error {
	if sl.tracer == nil {
		return fmt.Errorf("Tracer is nil")
	}
	return nil
}

func (sl *tfStateLoader) loadRemoteStateImpl(remoteCfg iface.RemoteConfig, downloader iface.S3DownloaderAPI) ([]*terraform.State, error) {

	tfStateList := make([]*terraform.State, 0)

	for _, key := range remoteCfg.Keys {
		var buffer []byte
		wBuffer := aws.NewWriteAtBuffer(buffer)
		stateFile := remoteCfg.BucketName + "/" + key
		sl.tracer.Trace("Loading ", stateFile, "...")

		// Write the contents of S3 Object to the buffer
		nBytes, err := downloader.Download(wBuffer, &s3.GetObjectInput{
			Bucket: aws.String(remoteCfg.BucketName),
			Key:    aws.String(key),
		})

		if err != nil {
			return nil, fmt.Errorf("Failed to download state-file '%s', %v", stateFile, err)
		}
		sl.tracer.Trace("Loading ", stateFile, " ", nBytes, " B...done")

		tfState, err := Parse(wBuffer.Bytes())
		if err != nil {
			return nil, err
		}
		tfStateList = append(tfStateList, tfState)
	}

	return tfStateList, nil

}
func (sl *tfStateLoader) LoadRemoteState(remoteCfg iface.RemoteConfig) ([]*terraform.State, error) {

	emptyList := make([]*terraform.State, 0)

	// create session
	verboseCredErrors := true
	cfg := aws.Config{Region: aws.String(remoteCfg.Region), CredentialsChainVerboseErrors: &verboseCredErrors}
	sessionOpts := session.Options{Profile: remoteCfg.Profile, Config: cfg}
	session, err := session.NewSessionWithOptions(sessionOpts)
	if err != nil {
		return emptyList, err
	}

	// Create a downloader with the session and default options
	downloader := s3manager.NewDownloader(session)

	return sl.loadRemoteStateImpl(remoteCfg, downloader)
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
func NewStateLoader() iface.StateLoader {
	return NewStateLoaderWithTracer(nil)
}

// NewStateLoaderWithTracer creates a new instance of a StateLoader with tracing
func NewStateLoaderWithTracer(tracer trace.Tracer) iface.StateLoader {
	if tracer == nil {
		tracer = trace.Off()
	}
	return &tfStateLoader{
		tracer: tracer,
	}
}
