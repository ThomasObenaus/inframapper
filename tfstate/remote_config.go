package tfstate

// RemoteConfig configuratipon for the remote storage (S3) for terrafrom state
type RemoteConfig struct {
	BucketName string
	Keys       []string
	Profile    string
	Region     string
}

func (l RemoteConfig) IsRemote() bool {
	return true
}

func (l RemoteConfig) RemoteConfig() RemoteConfig {
	return l
}

func (l RemoteConfig) LocalConfig() LocalConfig {
	return LocalConfig{}
}
