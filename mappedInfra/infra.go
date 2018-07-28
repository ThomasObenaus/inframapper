package mappedInfra

import (
	"fmt"

	"github.com/thomas.obenaus/terrastate/aws"
	"github.com/thomas.obenaus/terrastate/terraform"
)

type Infra interface {
}

type infraImpl struct {
}

func DoMapping(aws *aws.Infra, tf *terraform.Infra) (Infra, error) {

	return nil, fmt.Errorf("N/A")
}
