package builder

import (
	"log"
	"os"
)

func cleanDir(path string, removeOnly bool) {
	log.Println("clean dir", path)

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		goto CREATEDIR
	}

	if err != nil {
		log.Fatalln(err)
	}

	if err := os.RemoveAll(path); err != nil {
		log.Fatalln(err)
	}

CREATEDIR:
	if removeOnly {
		return
	}

	if err := os.MkdirAll(path, 0755); err != nil {
		log.Fatalln(err)
	}
}
