package tfstate

// RemoteConfig configuratipon for the remote storage (S3) for terrafrom state
type RemoteConfig struct {
	BucketName string
	Keys       []string
	Profile    string
	Region     string
}

// IsRemote returns true in case the state is stored remotely, false otherwise.
func (l RemoteConfig) IsRemote() bool {
	return true
}

// RemoteConfig returns the configuration of the remote backend for the state
// Can return an empty struct in case we have no remote backend.
func (l RemoteConfig) RemoteConfig() RemoteConfig {
	return l
}

// LocalConfig returns the configuration of the locally stored state
// Can return an empty struct in case we have no local state.
func (l RemoteConfig) LocalConfig() LocalConfig {
	return LocalConfig{}
}
