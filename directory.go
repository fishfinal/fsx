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
	"embed"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// RuntimeMasterDir returns the root path corresponding to the current directory.
func RuntimeMasterDir() string {
	base, _ := os.Getwd()
	return base
}

func MustGetFilePathWithFileSystemPath(baseDir string, ops ...Option) string {
	options := options{
		path: GlobalDefaultDir,
	}

	for _, o := range ops {
		o.apply(&options)
	}

	path := filepath.Join(RuntimeMasterDir(), baseDir, options.path.String())
	if err := NotExistsMkdir(path); err != nil {
		panic(err)
	}

	return path
}

// GetFilePathWithFileSystemPath Get a file path with filesystem path.
func GetFilePathWithFileSystemPath(baseDir string, ops ...Option) (string, error) {
	options := options{
		path: GlobalDefaultDir,
	}

	for _, o := range ops {
		o.apply(&options)
	}

	path := filepath.Join(RuntimeMasterDir(), baseDir, options.path.String())
	if err := NotExistsMkdir(path); err != nil {
		return "", err
	}

	return path, nil
}

// NotExistsMkdir detects whether the target folder exists.
// and if it does not, create the folder.
func NotExistsMkdir(path string) error {
	if exists := Exists(path); !exists {
		if err := MkdirAll(path); err != nil {
			return err
		}
	}
	return nil
}

// IsNotExists check whether the target file or directory exists.
// Return FALSE if the file exists, TRUE otherwise.
//
// Deprecated: use Exists instead.
func IsNotExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return true
	}
	return false
}

// Exists checks whether the target file or directory exists.
// Return TRUE if the file exists, FALSE otherwise.
func Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// IsNotExistsWithEmbedFS check whether the target file or directory exists with embed.FS.
// Return FALSE if the file exists, TRUE otherwise.
//
// Deprecated: use EmbedsFileExists instead.
func IsNotExistsWithEmbedFS(embeds embed.FS, path string) bool {
	if _, err := embeds.ReadFile(path); os.IsNotExist(err) {
		return true
	}
	return false
}

// EmbedsFileExists checks whether the specified file exists in the embedded file system.
// It attempts to read the file from the embed.FS to determine its existence.
//
// Parameters:
//   - embeds: The embedded file system instance to search in
//   - path: The relative path of the file to check (relative to the embed.FS root)
//
// Returns:
//   - true if the file exists and can be read (or opened)
//   - false if the file does not exist or cannot be accessed
//
// Note: This function only checks for the existence of files, not directories.
// Use EmbedsExists() if you need to check for both files and directories.
func EmbedsFileExists(embeds embed.FS, path string) bool {
	_, err := embeds.ReadFile(path)
	return !os.IsNotExist(err)
}

// EmbedsExists checks whether the specified file or directory exists in the embedded file system.
// This is a more general function that can handle both files and directories.
//
// Parameters:
//   - embeds: The embedded file system instance to search in
//   - path: The relative path of the file or directory to check
//
// Returns:
//   - true if the file or directory exists
//   - false if it does not exist
//
// Example:
//
//	//go:embed static/*
//	var staticFS embed.FS
//
//	if EmbedsExists(staticFS, "static/css/style.css") {
//	    // File exists
//	}
//
//	if EmbedsExists(staticFS, "static/images/") {
//	    // Directory exists (note trailing slash for directories)
//	}
func EmbedsExists(embeds embed.FS, path string) bool {
	// Try to open the path (works for both files and directories)
	f, err := embeds.Open(path)
	if err != nil {
		return false
	}
	defer f.Close()
	return true
}

// EmbedsNotExists checks whether the specified file or directory does NOT exist in the embedded file system.
// This is the inverse of EmbedsExists() for semantic clarity in certain contexts.
//
// Parameters:
//   - embeds: The embedded file system instance to search in
//   - path: The relative path to check
//
// Returns:
//   - true if the file or directory does NOT exist
//   - false if it exists
//
// Usage note: Consider using !EmbedsExists() for better readability in most cases.
// This function is provided for consistency with Go's os.IsNotExist pattern.
func EmbedsNotExists(embeds embed.FS, path string) bool {
	return !EmbedsExists(embeds, path)
}

// MkdirAll creates a directory named path.
func MkdirAll(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

// NotExistsMkdirAll Check whether the directory exists and create it if it does not
func NotExistsMkdirAll(path string) error {
	var exists = true

	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		exists = false
	}

	if !exists {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}

// ReadDir Get the directory file name by recursive reading.
func ReadDir(path string, result []string) ([]string, error) {
	if result == nil {
		result = make([]string, 0)
	}

	dir, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, f := range dir {
		if f.IsDir() {
			result, err = ReadDir(filepath.Join(path, f.Name()), result)
			if err != nil {
				return nil, err
			}
		} else {
			result = append(result, filepath.Join(path, f.Name()))
		}
	}

	return result, nil
}

// WriteDir Write file to dst directory.
func WriteDir(dst string, files []string, ops ...Option) error {
	options := &options{}

	for _, o := range ops {
		o.apply(options)
	}

	if err := NotExistsMkdir(dst); err != nil {
		return err
	}

	for _, cur := range files {
		ok, err := IsFile(cur)
		if err != nil {
			return err
		}

		if ok {
			dir := filepath.Dir(cur)
			if options.trimPrefixStr != "" {
				dir = FilenameTrimPrefix(dir, options.trimPrefixStr)
			}
			dsd := filepath.Join(dst, dir)
			dsf := filepath.Join(dsd, filepath.Base(cur))

			// Read file content from original filepath.
			data, err := ioutil.ReadFile(cur)
			if err != nil {
				fmt.Printf("has errors: %+v\n", err)
				return err
			}

			// Write file content to new filepath.
			err = NotExistsMkdir(dsd)
			if err != nil {
				return err
			}

			err = ioutil.WriteFile(dsf, data, os.ModePerm)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// LastDirName Gets filename the last directory name.
func LastDirName(filename string) string {
	sli := strings.Split(filepath.Dir(filename), string(os.PathSeparator))
	if len(sli) > 1 {
		return sli[len(sli)-1]
	}
	if len(sli) == 1 {
		return sli[0]
	}
	return ""
}

// IsDir Checks whether the specified file name is a directory.
func IsDir(path string) (bool, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	if stat.IsDir() {
		return true, nil
	}
	return false, nil
}

// DirNotExists Check whether the specified file path directory does not exists.
func DirNotExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return true
	}
	return false
}

// Copy file to destination directory.
func Copy(src string, dst string, ops ...Option) error {
	options := &options{
		force: false,
	}

	for _, o := range ops {
		o.apply(options)
	}

	srcIsDir, err := IsDir(src)
	if err != nil {
		return err
	}

	if !options.force {
		if srcIsDir && DirNotExists(dst) {
			return fmt.Errorf("no such directory of dst")
		} else {
			files, err := ReadDir(src, nil)
			if err != nil {
				return err
			}

			return WriteDir(dst, files, ops...)
		}
	}

	return nil
}
