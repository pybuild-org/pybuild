package builder

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func CleanDir(path string, removeOnly bool) {
	log.Println("clean dir", path)

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		goto CREATEDIR
	}

	if err != nil {
		log.Fatalln(err)
	}

	if err := os.RemoveAll(path); err != nil {
		log.Fatalln(err)
	}

CREATEDIR:
	if removeOnly {
		return
	}

	if err := os.MkdirAll(path, 0755); err != nil {
		log.Fatalln(err)
	}
}

func MkPyFileName(version, release, arch, os string) string {
	return fmt.Sprintf(
		"cpython-%s+%s-%s-%s-install_only_stripped.tar.gz",
		version, release, arch, os,
	)
}

func MkPyDownloadUrl(version, release, arch, os string) string {
	return fmt.Sprintf(
		"https://github.com/astral-sh/python-build-standalone/releases/download/%s/%s",
		release, MkPyFileName(version, release, arch, os),
	)
}

func GetDownloadStream(url string) io.ReadCloser {
	log.Println("download", url)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	return resp.Body
}

func Decompress(r io.Reader, out string) {
	log.Println("decompress to", out)

	gzr, err := gzip.NewReader(r)
	if err != nil {
		log.Fatalln(err)
	}

	defer gzr.Close()
	tr := tar.NewReader(gzr)

	for {
		h, err := tr.Next()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalln(err)
		}

		switch p := filepath.Join(out, h.Name); h.Typeflag {

		case tar.TypeDir:
			if err := os.MkdirAll(p, h.FileInfo().Mode()); err != nil {
				log.Fatalln(err)
			}

		case tar.TypeReg:
			baseDir := filepath.Dir(p)
			if err := os.MkdirAll(baseDir, 0755); err != nil {
				log.Fatalln(err)
			}

			func() {
				outFile, err := os.OpenFile(p, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, h.FileInfo().Mode())
				if err != nil {
					log.Fatalln(err)
				}

				defer outFile.Close()
				if _, err := io.Copy(outFile, tr); err != nil {
					log.Fatalln(err)
				}
			}()

		}
	}
}

func MkPyBinPath(os, version string) string {
	if strings.Contains(os, "windows") {
		return "python.exe"
	}

	verParts := strings.SplitN(version, ".", 3)
	ver := strings.Join([]string{verParts[0], verParts[1]}, ".")

	return "bin/python" + ver
}

func RunCommand(parts ...string) {
	cmd := exec.Command(parts[0], parts[1:]...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Println("run command", cmd.String())
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
}

func CopyFile(src, dst string, d fs.DirEntry) error {
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

func CopyDir(src, dst string) {
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
			return CopyFile(path, targetPath, d)
		}

	}); err != nil {
		log.Fatalln(err)
	}
}
