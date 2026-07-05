package container

import (
	"archive/tar"
	"io"
	"log"
	"os"
	"path/filepath"
)

func copyDir(tw *tar.Writer, src, dst string) {
	if err := filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		header, err := tar.FileInfoHeader(info, info.Name())
		if err != nil {
			return err
		}

		imagePath := filepath.Join(dst, relPath)
		header.Name = imagePath
		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		if info.Mode().IsRegular() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}

			defer file.Close()
			_, err = io.Copy(tw, file)
			if err != nil {
				return err
			}
		}

		return nil

	}); err != nil {
		log.Fatalln(err)
	}
}
