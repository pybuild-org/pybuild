package main

import (
	"log"
	"net/http"
	"os"
)

func onTagOpen() {
	n := i.CurrentNode()
	if n == nil {
		return
	}

	switch n.Name {

	case "xml":
		i.PushStack()

	case "use":
		i.PushStack()

		if file := n.Attrs["file"]; file != "" {
			f, err := os.Open(file)
			if err != nil {
				log.Fatalln(err)
			}

			defer f.Close()
			i.Run(f)

		} else if url := n.Attrs["url"]; url != "" {
			r, err := http.Get(url)
			if err != nil {
				log.Fatalln(err)
			}

			defer r.Body.Close()
			i.Run(r.Body)
		}

	}
}

func onTagClose() {
	n := i.CurrentNode()
	if n == nil {
		return
	}

	switch n.Name {

	case "log":
		log.Println(n.Value)
		i.PushStack()

	}
}
