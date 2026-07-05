package docker

import (
	"log"

	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/remote"
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
