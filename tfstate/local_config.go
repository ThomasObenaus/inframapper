package tfstate

// LocalConfig configuration for the local file-based for terrafrom state
type LocalConfig struct {
	Files []string
}

func (l LocalConfig) IsRemote() bool {
	return false
}

func (l LocalConfig) RemoteConfig() RemoteConfig {
	return RemoteConfig{}
}

func (l LocalConfig) LocalConfig() LocalConfig {
	return l
}
