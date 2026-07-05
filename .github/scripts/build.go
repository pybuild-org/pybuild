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

func main() {
	f, err := os.Open(".github/build/targets.json")
	if err != nil {
		log.Fatalln(err)
	}

	var Targets []Target
	if err := json.NewDecoder(f).Decode(&Targets); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(Targets)
}
