package container

import (
	"fmt"
	"path/filepath"
	"pybuild/builder"

	v1 "github.com/google/go-containerregistry/pkg/v1"
)

func Build() {
	imageName := fmt.Sprintf(
		"%s:%s",
		builder.BuilderConfig.App,
		MetaConfig.Tag,
	)

	for _, target := range Targets {
		dirName := target.Python.Arch + "-" + target.Python.OS + "-" + "container"
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

		var image v1.Image

		{
			baseImg := useImage(
				target.Image.Base,
				target.Image.Arch,
				target.Image.OS,
			)

			image = appendDirLayer(baseImg, baseDir, "/app")
		}

	}
}
