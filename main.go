package main

import (
	"log"

	terraform "github.com/hashicorp/terraform/terraform"
)

func main() {
	log.Println("Hello world")
	state := terraform.State{}

	log.Printf("State %s", state)
}
