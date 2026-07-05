package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
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

	var targets []Target
	if err := json.NewDecoder(f).Decode(&targets); err != nil {
		log.Fatalln(err)
	}

	var wg sync.WaitGroup
	limit := make(chan struct{}, 10)

	for _, target := range targets {
		wg.Add(1)
		limit <- struct{}{}

		go func(target Target) {
			defer func() {
				<-limit
				wg.Done()
			}()

			log.Println("build", fmt.Sprintf("%s/%s", target.GOOS, target.GOARCH))

			ext := ""
			if target.GOOS == "windows" {
				ext = ".exe"
			}

			name := fmt.Sprintf("pybuild-%s-%s%s", target.GOOS, target.GOARCH, ext)
			file := filepath.Join(".github/build/dist", name)

			run(
				[]string{"go", "build", "-o", file, "."},
				[]string{
					"CGO_ENABLED=0",
					fmt.Sprintf("GOOS=%s", target.GOOS),
					fmt.Sprintf("GOARCH=%s", target.GOARCH),
				}
			)
		}(target)
	}
}

func run(parts, env []string) {
	cmd := exec.Command(parts[0], parts[1:]...)

	cmd.Env = env
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()
}
