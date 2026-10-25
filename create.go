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
	"io/ioutil"
	"os"
	"path/filepath"
)

// CreateFile Create new file, your can specific multiple options.
func CreateFile(filename string, data []byte, ops ...Option) (string, error) {
	options := options{
		extension: FileExtensionWithDocumentOfFileText,
		baseDir:   ".",
	}

	// Exists file extension, use it.
	extension := GetFileExtension(
		filename,
		WithSpecificFileExtensionContainsDot(false),
	)
	if extension != "" {
		options.extension = FileExtension(extension)
	}

	for _, o := range ops {
		o.apply(&options)
	}

	path, err := GetFilePathWithFileSystemPath(options.baseDir, ops...)
	if err != nil {
		return "", err
	}

	filename = filepath.Join(path, CreateFilename(
		filename,
		options.extension.String(),
		ops...,
	))

	err = ioutil.WriteFile(filename, data, os.ModePerm)
	if err != nil {
		return "", err
	}

	return filename, nil
}
