package tfstate

// RemoteConfig configuratipon for the remote storage (S3) for terrafrom state
type RemoteConfig struct {
	BucketName string
	Keys       []string
	Profile    string
	Region     string
}
