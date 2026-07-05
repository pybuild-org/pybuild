package docker

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
	"github.com/google/go-containerregistry/pkg/v1/stream"
)

func useImage(image, os, arch string) v1.Image {
	log.Println("use image", image, os, arch)

	repo, err := name.ParseReference(image)
	if err != nil {
		log.Fatalln(err)
	}

	options := []remote.Option{
		remote.WithPlatform(v1.Platform{
			OS:           os,
			Architecture: arch,
		}),
	}

	img, err := remote.Image(repo, options...)
	if err != nil {
		log.Fatalln(err)
	}

	return img
}

func appendDir(img v1.Image, src, dst string) v1.Image {
	log.Println("copy", src, "to", dst)

	pr, pw := io.Pipe()

	go func() {
		tw := tar.NewWriter(pw)

		err := filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
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

			header, err := tar.FileInfoHeader(info, "")
			if err != nil {
				return err
			}

			targetPath := dst + "/" + filepath.ToSlash(relPath)
			header.Name = targetPath
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
		})

		if err != nil {
			tw.Close()
			pw.CloseWithError(err)
			return
		}

		if err := tw.Close(); err != nil {
			pw.CloseWithError(err)
			return
		}

		pw.Close()
	}()

	layer := stream.NewLayer(pr)
	newImg, err := mutate.AppendLayers(img, layer)
	if err != nil {
		log.Fatalln(err)
	}

	return newImg
}
