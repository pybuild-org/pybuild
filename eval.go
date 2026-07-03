package main

import (
	"log"
	"net/http"
	"os"
	"pybuild/strprop"
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

	case "config":
		configType, ok := n.Attrs["type"]
		if ok && configType == "group" {
			strprop.Next(n.Attrs["name"])
		}

	case "prop":
		p := i.ParentNode()
		if p == nil || p.Name != "config" {
			return
		}

		pp := i.Node(-3)
		if pp == nil || pp.Name != "config" {
			strprop.Update(p.Attrs["name"], "", n.Attrs["name"], n.Attrs["value"])
			return
		}

		configType, ok := pp.Attrs["type"]
		if ok && configType == "group" {
			strprop.Update(pp.Attrs["name"], p.Attrs["name"], n.Attrs["name"], n.Attrs["value"])
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

	case "config", "prop":
		i.PopStack()

	}
}
