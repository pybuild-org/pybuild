package main

import (
	"log"
	"os"
	"pybuild/builder"
	"pybuild/funcjob"
	"pybuild/strprop"
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

	strprop.Bind("builder", &builder.BuilderConfig)
	funcjob.Register("setup builder", builder.SetupBuilder)

	strprop.Bind("python", &builder.PythonConfig)
	funcjob.Register("setup python", builder.SetupPython)

	defer f.Close()
	i.Run(f)
}
