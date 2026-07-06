package docker

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"pybuild/builder"
	"strings"
	"text/template"

	"github.com/google/go-containerregistry/pkg/v1/mutate"
)

func Build() {
	imageName := fmt.Sprintf(
		"%s:%s",
		builder.BuilderConfig.App,
		MetaConfig.Tag,
	)

	for _, target := range Targets {
		isWindows := strings.Contains(target.Python.OS, "windows")
		dirName := target.Python.Arch + "-" + target.Python.OS + "-docker-image"
		baseDir := filepath.Join(builder.TempDir, dirName)
		builder.CleanDir(baseDir, false)

		builder.CopyDir(builder.BuilderConfig.Source, baseDir)

		{
			url := builder.MkPyDownloadUrl(
				builder.BuilderConfig.Version,
				builder.BuilderConfig.Release,
				target.Python.Arch,
				target.Python.OS,
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
			content := builder.ShLauncher
			if isWindows {
				content = builder.CmdLauncher
			}

			tpl, err := template.New("launcher").Parse(content)
			if err != nil {
				log.Fatalln(err)
			}

			ext := ".sh"
			if isWindows {
				ext = ".cmd"
			}

			p := filepath.Join(baseDir, "launcher"+ext)
			f, err := os.Create(p)
			if err != nil {
				log.Fatalln(err)
			}

			data := map[string]string{
				"RUN": target.Launcher.Run,
				"PYTHON": builder.MkPyBinPath(
					target.Python.OS,
					builder.BuilderConfig.Version,
				),
			}

			defer f.Close()
			if err := tpl.Execute(f, data); err != nil {
				log.Fatalln(err)
			}
		}()

		cacheDir := filepath.Join(builder.TempDir, "cache")
		builder.CleanDir(cacheDir, false)

		image := appendDir(useImage(
			target.Image.Base,
			target.Image.OS,
			target.Image.Arch,
		), baseDir, "app", cacheDir)

		{
			cfg, err := image.ConfigFile()
			if err != nil {
				log.Fatalln(err)
			}

			newCfg := cfg.DeepCopy()
			newCfg.OS = target.Image.OS
			newCfg.Architecture = target.Image.Arch
			newCfg.Config.Entrypoint = []string{"/app/launcher.sh"}
			if isWindows {
				newCfg.Config.Entrypoint = []string{"C:\\app\\launcher.cmd"}
			}

			newImg, err := mutate.ConfigFile(image, newCfg)
			if err != nil {
				log.Fatalln(err)
			}

			image = newImg
		}

		saveImage(
			image, imageName,
			filepath.Join(builder.BuilderConfig.Output, dirName+".tar"),
		)

		builder.CleanDir(cacheDir, true)
		builder.CleanDir(baseDir, true)
	}
}
