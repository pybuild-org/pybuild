package container

import (
	"archive/tar"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

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

func copyDirToImageLayer(base v1.Image, src, target string) v1.Image {
	log.Println("copy", src, "to", target)

	tmpTar, err := os.CreateTemp("", "image-layer-*.tar")
	if err != nil {
		log.Println(err)
	}

	defer os.Remove(tmpTar.Name())
	defer tmpTar.Close()

	tw := tar.NewWriter(tmpTar)
	if err = filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		if relPath == "." {
			return nil
		}

		header, err := tar.FileInfoHeader(info, info.Name())
		if err != nil {
			return err
		}

		tarPath := filepath.ToSlash(filepath.Join(target, relPath))
		header.Name = strings.TrimPrefix(tarPath, "/")
		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		if info.Mode().IsRegular() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}

			defer file.Close()
			if _, err := io.Copy(tw, file); err != nil {
				return err
			}
		}

		return nil

	}); err != nil {
		log.Println(err)
	}

	tw.Close()

	layer, err := tarball.LayerFromOpener(func() (io.ReadCloser, error) {
		return os.Open(tmpTar.Name())
	})

	if err != nil {
		log.Println(err)
	}

	newImg, err := mutate.AppendLayers(base, layer)
	if err != nil {
		log.Println(err)
	}

	return newImg
}
