package awsresource

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func newAWSSession(profile string, region string) (*session.Session, error) {
	verboseCredErrors := true

	cfg := aws.Config{Region: aws.String(region), CredentialsChainVerboseErrors: &verboseCredErrors}
	sessionOpts := session.Options{Profile: profile, Config: cfg}

	return session.NewSessionWithOptions(sessionOpts)
}
