package standalone

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
)

func compress(src, dst string) {
	log.Println("compress", src, "to", dst)

	f, err := os.Create(dst)
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()
	archive := zip.NewWriter(f)
	defer archive.Close()

	absSrcDir, err := filepath.Abs(src)
	if err != nil {
		log.Fatalln(err)
	}

	if err := filepath.WalkDir(absSrcDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(absSrcDir, path)
		if err != nil {
			return err
		}

		if relPath == "." {
			return nil
		}

		fi, err := d.Info()
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(fi)
		if err != nil {
			return err
		}

		header.Name = filepath.ToSlash(relPath)

		if d.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		fileToZip, err := os.Open(path)
		if err != nil {
			return err
		}

		defer fileToZip.Close()
		_, err = io.Copy(writer, fileToZip)
		if err != nil {
			return err
		}

		return nil

	}); err != nil {
		log.Fatalln(err)
	}
}
