// Copyright 2021 fishfinal <contact.fishfinal@gmail.com>;
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fsx

import (
	"archive/zip"
	"compress/flate"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const (
	ZIPExtension = "zip"
)

// CreateZIPFileWithFilename CreateZIPFile Create ZIP file, return the storage path and error.
// source parameter specifies a file or directory.
func CreateZIPFileWithFilename(source string) (string, error) {
	filename, dir, err := GetFileOrDirectoryName(source)
	if err != nil {
		return "", err
	}

	name := strings.Join([]string{filename, ZIPExtension}, ".")
	path := filepath.Join(filepath.Dir(source), name)

	archive, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer archive.Close()

	writer := zip.NewWriter(archive)
	writer.RegisterCompressor(zip.Deflate, func(out io.Writer) (io.WriteCloser, error) {
		return flate.NewWriter(out, flate.BestCompression)
	})

	if dir {
		dir, err := ReadDir(source, nil)
		if err != nil {
			return "", err
		}

		for _, f := range dir {
			err := write(writer, f, filename)
			if err != nil {
				return "", err
			}
		}

		if err = writer.Close(); err != nil {
			return "", err
		}

		return path, nil
	}

	err = write(writer, source, filename)
	if err != nil {
		return "", err
	}

	if err = writer.Close(); err != nil {
		return "", err
	}

	return path, nil
}

func write(writer *zip.Writer, origin string, latest string) error {
	file, err := os.OpenFile(origin, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}

	stat, err := file.Stat()
	if err != nil {
		return err
	}

	fw, err := writer.CreateHeader(&zip.FileHeader{
		Name:     trim(file.Name(), latest),
		Modified: stat.ModTime(),
	})

	if _, err := io.Copy(fw, file); err != nil {
		return err
	}

	return file.Close()
}

func trim(current, filename string) string {
	old := current[0:strings.LastIndex(current, filename)]
	return strings.Replace(current, old, "", 1)
}
