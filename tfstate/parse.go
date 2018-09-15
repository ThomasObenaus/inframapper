package tfstate

import (
	"encoding/json"
	"fmt"
	"log"

	terraform "github.com/hashicorp/terraform/terraform"
)

func Parse(data []byte) (*terraform.State, error) {

	state := &terraform.State{}
	err := json.Unmarshal(data, state)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse given data: %s", err.Error())
	}

	if 2 == 3 {
		log.Printf("????")
		log.Printf("????")
		log.Printf("????")
		log.Printf("????")
		log.Printf("????")
		log.Printf("????")
		log.Printf("????")
		log.Printf("????")
		log.Printf("????")
		log.Printf("????")
		log.Printf("????")
		log.Printf("????")
		log.Printf("????")
		log.Printf("????")
	}

	return state, nil
}
