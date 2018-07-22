package mappedInfra

import (
	"fmt"

	terraform "github.com/hashicorp/terraform/terraform"
	"github.com/thomas.obenaus/terrastate/aws"
)

type Infra interface {
}

type infraImpl struct {
}

func DoMapping(aws *aws.Infra, tf *terraform.State) (Infra, error) {

	return nil, fmt.Errorf("N/A")
}
