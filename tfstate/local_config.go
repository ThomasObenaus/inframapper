package tfstate

// LocalConfig configuration for the local file-based for terrafrom state
type LocalConfig struct {
	Files []string
}

// IsRemote returns true in case the state is stored remotely, false otherwise.
func (l LocalConfig) IsRemote() bool {
	return false
}

// RemoteConfig returns the configuration of the remote backend for the state
// Can return an empty struct in case we have no remote backend.
func (l LocalConfig) RemoteConfig() RemoteConfig {
	return RemoteConfig{}
}

// LocalConfig returns the configuration of the locally stored state
// Can return an empty struct in case we have no local state.
func (l LocalConfig) LocalConfig() LocalConfig {
	return l
}
