package standalone

import (
	"archive/zip"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func copyFile(src, dst string, d fs.DirEntry) error {
	sf, err := os.Open(src)
	if err != nil {
		return err
	}

	defer sf.Close()
	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return err
	}

	info, err := d.Info()
	if err != nil {
		return err
	}

	df, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, info.Mode())
	if err != nil {
		return err
	}

	defer df.Close()
	if _, err := io.Copy(df, sf); err != nil {
		return err
	}

	return nil
}

func copyDir(src, dst string) {
	log.Println("copy", src, "to", dst)

	if err := filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		if d.IsDir() && d.Name() == "__pycache__" {
			return filepath.SkipDir
		}

		targetPath := filepath.Join(dst, relPath)

		if d.IsDir() {
			info, err := d.Info()
			if err != nil {
				return err
			}

			return os.MkdirAll(targetPath, info.Mode())

		} else {
			return copyFile(path, targetPath, d)
		}

	}); err != nil {
		log.Fatalln(err)
	}
}

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
