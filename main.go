package main

import (
	"log"
	"os"
	"pybuild/builder"
	"pybuild/builder/docker"
	"pybuild/builder/standalone"
	"pybuild/funcjob"
	"pybuild/strprop"
	"pybuild/strtpl"
	"pybuild/xmlinterp"
	"strings"
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
		if !strings.HasSuffix(scriptPath, ".xml") {
			scriptPath += ".xml"
		}
	}

	i = xmlinterp.New(onTagOpen, onTagClose)
	f, err := os.Open(scriptPath)
	if err != nil {
		log.Fatalln(err)
	}

	strprop.Bind("template", &strtpl.TemplateConfig)

	strprop.Bind("builder", &builder.BuilderConfig)
	funcjob.Register("setup builder", builder.SetupBuilder)

	strprop.Bind("python", &builder.PythonConfig)
	funcjob.Register("setup python", builder.SetupPython)

	strprop.Bind("standalone targets", &standalone.Targets)
	funcjob.Register("build standalone", standalone.Build)

	strprop.Bind("docker image meta", &docker.MetaConfig)
	strprop.Bind("docker image targets", &docker.Targets)
	funcjob.Register("docker build", docker.Build)

	defer f.Close()
	defer builder.Clean()
	i.Run(f)
}
