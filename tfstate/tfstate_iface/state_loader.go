package tfstate_iface

import (
	"github.com/hashicorp/terraform/terraform"
)

type StateLoader interface {
	// Load loads a terraform state file
	Load(filename string) (*terraform.State, error)

	// LoadRemoteState loads state from an aws S3 bucket
	LoadRemoteState(remoteCfg RemoteConfig) ([]*terraform.State, error)
}
