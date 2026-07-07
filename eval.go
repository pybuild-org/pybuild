package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"pybuild/funcjob"
	"pybuild/strprop"
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

		src := n.Attrs["src"]
		if !strings.HasSuffix(src, ".xml") {
			src += ".xml"
		}

		if strings.HasPrefix(src, "http://") || strings.HasPrefix(src, "https://") {
			resp, err := http.Get(src)
			if err != nil {
				log.Fatalln(err)
			}

			defer resp.Body.Close()
			i.Run(resp.Body)

		} else {
			f, err := os.Open(src)
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

		job, isJob := n.Attrs["job"]
		command, isCommand := n.Attrs["command"]

		if isJob {
			funcjob.Run(job)

		} else if isCommand {
			parts := strings.Fields(command)
			cmd := exec.Command(parts[0], parts[1:]...)

			cmd.Env = os.Environ()
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			log.Println("run command", cmd.String())
			if err := cmd.Run(); err != nil {
				log.Fatalln(err)
			}
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

	case "config", "prop":
		i.PopStack()

	}
}
