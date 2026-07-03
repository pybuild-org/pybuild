package main

import (
	"log"
	"net/http"
	"os"
	"pybuild/funcjob"
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

		name := n.Attrs["name"]
		netMode := false

		if strings.HasPrefix(name, "url:") {
			netMode = true
			name = strings.TrimPrefix(name, "url:")

		} else if !strings.HasSuffix(name, ".xml") {
			name = name + ".xml"
		}

		if netMode {
			r, err := http.Get(name)
			if err != nil {
				log.Fatalln(err)
			}

			defer r.Body.Close()
			i.Run(r.Body)

		} else {
			f, err := os.Open(name)
			if err != nil {
				log.Fatalln(err)
			}

			defer f.Close()
			i.Run(f)
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

	case "run":
		i.PopStack()

		name := n.Attrs["job"]
		funcjob.Run(name)

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
		parts := strings.Fields(n.Value)
		util.RunCommand(parts, os.Environ())
		i.PopStack()

	case "config", "prop":
		i.PopStack()

	}
}
