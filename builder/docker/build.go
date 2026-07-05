package docker

import (
	"fmt"
	"log"
	"path/filepath"
	"pybuild/builder"

	"github.com/google/go-containerregistry/pkg/v1/mutate"
)

func Build() {
	imageName := fmt.Sprintf(
		"%s:%s",
		builder.BuilderConfig.App,
		MetaConfig.Tag,
	)

	for _, target := range Targets {
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

		image := useImage(
			target.Image.Base,
			target.Image.OS,
			target.Image.Arch,
		)

		{
			cfg, err := image.ConfigFile()
			if err != nil {
				log.Fatalln(err)
			}

			newCfg := cfg.DeepCopy()
			newCfg.OS = target.Image.OS
			newCfg.Architecture = target.Image.Arch

			newImg, err := mutate.ConfigFile(image, newCfg)
			if err != nil {
				log.Fatalln(err)
			}

			image = newImg
		}

		cacheDir := filepath.Join(builder.TempDir, "cache")
		builder.CleanDir(cacheDir, false)

		saveImage(
			appendDir(image, baseDir, "app", cacheDir), imageName,
			filepath.Join(builder.BuilderConfig.Output, dirName+".tar"),
		)

		builder.CleanDir(cacheDir, true)
		builder.CleanDir(baseDir, true)
	}
}
