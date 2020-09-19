package main

import (
	"archive/zip"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	src, err := filepath.Abs("./3.5 files/task.zip")
	if err != nil {
		panic(err)
	}

	dest, err := filepath.Abs("./3.5 files/task_temp")
	if err != nil {
		panic(err)
	}

	if err := Unzip(src, dest); err != nil {
		panic(err)
	}
	defer os.RemoveAll(dest)

	var walkFunc = func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		r := csv.NewReader(file)
		data, err := r.ReadAll()
		if err != nil {
			return nil
		}

		if len(data) != 10 {
			return nil
		}

		fmt.Println(data[4][2])
		return nil
	}

	if err := filepath.Walk(dest, walkFunc); err != nil {
		panic(err)
	}
}

// https://stackoverflow.com/questions/20357223/easy-way-to-unzip-file-with-golang
func Unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	os.MkdirAll(dest, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		// Check for ZipSlip (Directory traversal)
		if !strings.HasPrefix(path, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", path)
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}
