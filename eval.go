package main

import (
	"log"
	"net/http"
	"os"
	"pybuild/util"
	"strings"
)

func onTagOpen() {
	n := i.CurrentNode()
	if n == nil {
		return
	}

	switch n.Name {

	case "xml":
		i.PopStack()

	case "use":
		i.PopStack()

		if file, ok := n.Attrs["file"]; file != "" && ok {
			f, err := os.Open(file)
			if err != nil {
				log.Fatalln(err)
			}

			defer f.Close()
			i.Run(f)

		} else if url, ok := n.Attrs["url"]; url != "" && ok {
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
		i.PopStack()

	case "exec":
		util.ExecCommand(strings.Fields(n.Value), os.Environ())
		i.PopStack()

	}
}
