package tfstate

// StateBackend is an interface that specifies how the state is stored either locally (filebased)
// or remotely (in S3).
type StateBackend interface {

	// IsRemote returns true in case the state is stored remotely, false otherwise.
	IsRemote() bool

	// RemoteConfig returns the configuration of the remote backend for the state
	// Can return an empty struct in case we have no remote backend.
	RemoteConfig() RemoteConfig

	// LocalConfig returns the configuration of the locally stored state
	// Can return an empty struct in case we have no local state.
	LocalConfig() LocalConfig
}
