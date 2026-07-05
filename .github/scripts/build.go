package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Target struct {
	GOOS   string `json:"GOOS"`
	GOARCH string `json:"GOARCH"`
}

var Targets []Target

func main() {
	platforms, err := os.Open(".github/platforms.json")
	if err != nil {
		log.Fatalln(err)
	}

	if err := json.NewDecoder(platforms).Decode(&Targets); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(Targets)
}
