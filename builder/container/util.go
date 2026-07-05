package container

import (
	"archive/tar"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/layout"
	"github.com/google/go-containerregistry/pkg/v1/mutate"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/google/go-containerregistry/pkg/v1/tarball"
)

func useImage(base, arch, os string) v1.Image {
	log.Println("use iamge", base, arch, os)

	ref, err := name.ParseReference(base)
	if err != nil {
		log.Fatalln(err)
	}

	img, err := remote.Image(ref, remote.WithPlatform(v1.Platform{
		Architecture: arch,
		OS:           os,
	}))

	if err != nil {
		log.Fatalln(err)
	}

	return img
}

func streamDirToTar(src, dst string, tw *tar.Writer) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		tarPath := filepath.ToSlash(filepath.Join(dst, relPath))
		if tarPath == dst && info.IsDir() {
			return nil
		}

		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}

		header.Name = tarPath
		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		if !info.Mode().IsRegular() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}

		defer file.Close()
		_, err = io.Copy(tw, file)
		return err
	})
}

func appendDirLayer(base v1.Image, src, target string) v1.Image {
	log.Println("copy", src, "to", target)

	newLayer, err := tarball.LayerFromOpener(func() (io.ReadCloser, error) {
		pr, pw := io.Pipe()

		go func() {
			tw := tar.NewWriter(pw)
			err := streamDirToTar(src, target, tw)
			tw.Close()
			pw.CloseWithError(err)
		}()

		return pr, nil
	})

	if err != nil {
		log.Fatalln(err)
	}

	newImg, err := mutate.AppendLayers(base, newLayer)
	if err != nil {
		log.Fatalln(err)
	}

	return newImg
}

func saveDockerImage(img v1.Image, imageName, outFile string) {
	log.Println("save docker image to", outFile)

	ref, err := name.ParseReference(imageName)
	if err != nil {
		log.Fatalln(err)
	}

	file, err := os.Create(outFile)
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()
	if err := tarball.Write(ref, img, file); err != nil {
		log.Fatalln(err)
	}
}

func saveOciImage(img v1.Image, imageName, outFile string) {
	log.Println("save oci image to", outFile)

	tmpDir, err := os.MkdirTemp("", "oci-layout-*")
	if err != nil {
		log.Fatalln(err)
	}

	defer os.RemoveAll(tmpDir)
	path, err := layout.FromPath(tmpDir)
	if err != nil {
		log.Fatalln(err)
	}

	parsedTag, err := name.NewTag(imageName, name.StrictValidation)
	if err != nil {
		log.Fatalln(err)
	}

	err = path.AppendImage(img, layout.WithAnnotations(map[string]string{
		"org.opencontainers.image.ref.name": parsedTag.TagStr(),
	}))

	if err != nil {
		log.Fatalln(err)
	}

	tarFile, err := os.Create(outFile)
	if err != nil {
		log.Fatalln(err)
	}

	defer tarFile.Close()
	tw := tar.NewWriter(tarFile)
	defer tw.Close()

	if err := filepath.Walk(tmpDir, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		rel, err := filepath.Rel(tmpDir, p)
		if err != nil {
			return err
		}
		if rel == "." {
			return nil
		}

		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}

		header.Name = filepath.ToSlash(rel)
		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		if !info.Mode().IsRegular() {
			return nil
		}

		f, err := os.Open(p)
		if err != nil {
			return err
		}

		defer f.Close()
		_, err = io.Copy(tw, f)
		return err

	}); err != nil {
		log.Fatalln(err)
	}
}
