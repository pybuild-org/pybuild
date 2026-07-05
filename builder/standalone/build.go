package standalone

import (
	_ "embed"
	"log"
	"os"
	"path/filepath"
	"pybuild/builder"
	"strings"
	"text/template"
)

//go:embed template.sh
var ShLauncher string

//go:embed template.cmd
var CmdLauncher string

func Build() {
	for _, target := range Targets {
		dirName := target.Python.Arch + "-" + target.Python.OS
		baseDir := filepath.Join(builder.TempDir, dirName)
		builder.CleanDir(baseDir, false)

		CopyDir(builder.BuilderConfig.Source, baseDir)

		{
			url := builder.MkPyDownloadUrl(
				builder.BuilderConfig.Version,
				builder.BuilderConfig.Release,
				builder.PythonConfig.Arch,
				builder.PythonConfig.OS,
			)

			func() {
				s := builder.GetDownloadStream(url)
				defer s.Close()
				builder.Decompress(s, baseDir)
			}()
		}

		{
			if len(target.Pip.Downloads) > 0 {
				pipDir := filepath.Join(baseDir, "__pip_install__")
				builder.CleanDir(pipDir, false)

				builder.RunCommand(append([]string{
					builder.Python, "-m", "pip", "download",
					"--index-url=https://pypi.tuna.tsinghua.edu.cn/simple",
					"--only-binary=:all:",
					"--platform=" + target.Pip.Platform,
					"--implementation=cp",
					"--no-cache-dir",
					"--dest=" + pipDir,
				}, target.Pip.Downloads...)...)
			}
		}

		func() {
			isWindows := strings.Contains(target.Python.OS, "windows")

			content := ShLauncher
			if isWindows {
				content = CmdLauncher
			}

			tpl, err := template.New("launcher").Parse(content)
			if err != nil {
				log.Fatalln(err)
			}

			ext := ""
			if isWindows {
				ext = ".cmd"
			}

			p := filepath.Join(baseDir, builder.BuilderConfig.App+ext)
			f, err := os.Create(p)
			if err != nil {
				log.Fatalln(err)
			}

			data := map[string]string{
				"RUN": target.Launcher.Run,
				"PYTHON": builder.MkPyBinPath(
					builder.PythonConfig.OS,
					builder.BuilderConfig.Version,
				),
			}

			defer f.Close()
			if err := tpl.Execute(f, data); err != nil {
				log.Fatalln(err)
			}
		}()

		compress(
			baseDir,
			filepath.Join(builder.BuilderConfig.Output, dirName+".zip"),
		)
	}
}
