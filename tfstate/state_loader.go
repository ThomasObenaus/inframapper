// Package tfstate contains code for loading and parsing terraform state.
// TfState can be loaded from a local file or from remote (S3).
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

// StateLoader is a interface providing the possibility to load terraform state
// from remote (S3) or a local file.
type StateLoader interface {
	// Load loads a terraform state file
	Load(filename string) (*terraform.State, error)

	// LoadsFiles loads a list of terraform state files
	LoadFiles(files []string) ([]*terraform.State, error)

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

func (sl *tfStateLoader) loadRemoteStateImpl(remoteCfg RemoteConfig, downloader iface.S3DownloaderAPI) ([]*terraform.State, error) {

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
func (sl *tfStateLoader) LoadRemoteState(remoteCfg RemoteConfig) ([]*terraform.State, error) {

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

func (sl *tfStateLoader) LoadFiles(files []string) ([]*terraform.State, error) {
	stateList := make([]*terraform.State, 0)

	for _, file := range files {
		state, err := sl.Load(file)
		if err != nil {
			sl.tracer.Trace("Error loading state file '", file, "': ", err.Error())
			continue
		}
		stateList = append(stateList, state)
	}

	return stateList, nil
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
