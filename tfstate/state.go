package tfstate

import (
	"github.com/thomas.obenaus/terrastate/terraform"
)

type State interface {
	FindById(id string) (terraform.Resource, error)

	//FindVpc(id string) (*Vpc, error)
	//Vpcs() []*Vpc
}

type stateImpl struct {
}
