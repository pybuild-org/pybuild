package container

import (
	"log"

	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/remote"
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
