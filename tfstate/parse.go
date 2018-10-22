package tfstate

import (
	"encoding/json"
	"fmt"

	terraform "github.com/hashicorp/terraform/terraform"
)

// Parse unmarshals the given binary array into terraform state structure.
func Parse(data []byte) (*terraform.State, error) {

	state := &terraform.State{}
	err := json.Unmarshal(data, state)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse given data: %s", err.Error())
	}

	return state, nil
}
