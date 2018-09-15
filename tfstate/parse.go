package tfstate

import (
	"encoding/json"
	"fmt"

	terraform "github.com/hashicorp/terraform/terraform"
)

func Parse(data []byte) (*terraform.State, error) {

	state := &terraform.State{}
	err := json.Unmarshal(data, state)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse given data: %s", err.Error())
	}

	return state, nil
}
