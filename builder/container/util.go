package container

import (
	"archive/tar"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
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
