package main

import (
	"log"
	"os"
	"pybuild/xmlinterp"
)

var i *xmlinterp.Interpreter

func main() {
	var scriptPath string

	args := os.Args
	if len(args) != 2 {
		if _, err := os.Stat("target.xml"); os.IsNotExist(err) {
			log.Fatalln("usage: pybuild target.xml")
		}

		scriptPath = "target.xml"
	} else {
		scriptPath = args[1]
	}

	i = xmlinterp.New(onTagOpen, onTagClose)
	f, err := os.Open(scriptPath)
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()
	i.Run(f)
}
