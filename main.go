package main

import (
	"log"
	"os"
	"pybuild/builder"
	"pybuild/builder/standalone"
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

	strprop.Bind("standalone targets", &standalone.Targets)
	funcjob.Register("standalone build", standalone.Build)

	defer f.Close()
	defer builder.Clean()
	i.Run(f)
}
