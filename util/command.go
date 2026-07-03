package util

import (
	"log"
	"os"
	"os/exec"
)

func RunCommand(parts, env []string) {
	cmd := exec.Command(parts[0], parts[1:]...)

	cmd.Env = env
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Println("run command:", cmd.String())
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
}
