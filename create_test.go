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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateFile(t *testing.T) {
	grids := []struct {
		baseDir                   string
		savePath                  FileSystemPath
		filename                  string
		extension                 FileExtension
		trimOriginalFileExtension bool
		useExtension              bool
		data                      []byte
		expectedContains          string
	}{
		{
			baseDir:          "test/data/create",
			filename:         "hello",
			extension:        FileExtensionWithDocumentOfFileText,
			useExtension:     true,
			data:             []byte("hello world\nhelloshaohua"),
			expectedContains: "test/data/create/hello.txt",
		},
		{
			baseDir:          "test/data/create",
			savePath:         "textfiles/notes",
			filename:         "hello",
			extension:        FileExtensionWithDocumentOfFileText,
			useExtension:     true,
			data:             []byte("hello world\nhelloshaohua"),
			expectedContains: "test/data/create/textfiles/notes/hello.txt",
		},
		{
			baseDir:          "test/data/create",
			savePath:         "markdown/posts",
			filename:         "hello",
			extension:        FileExtensionWithDocumentOfMarkdown,
			useExtension:     true,
			data:             []byte("# Hello\n## HelloWorld\n## Hello世界~\n"),
			expectedContains: "test/data/create/markdown/posts/hello.md",
		},
		{
			baseDir:      "test/data/create",
			savePath:     "images/users",
			filename:     "header",
			extension:    FileExtensionWithImageOfJPEG,
			useExtension: true,
			data: func() []byte {
				data, err := ioutil.ReadFile("test/data/assets/images/helloshaohua.jpeg")
				assert.NoError(t, err)
				return data
			}(),
			expectedContains: "test/data/create/images/users/header.jpeg",
		},
		{
			baseDir:      "test/data/create",
			savePath:     "images/users",
			filename:     "header-specific-extension.jpeg",
			extension:    FileExtensionWithImageOfJPEG,
			useExtension: false,
			data: func() []byte {
				data, err := ioutil.ReadFile("test/data/assets/images/helloshaohua.jpeg")
				assert.NoError(t, err)
				return data
			}(),
			expectedContains: "test/data/create/images/users/header-specific-extension.jpeg",
		},
		{
			baseDir:  "test/data/create",
			savePath: "images/users",
			filename: "header-specific-extension.png",
			data: func() []byte {
				data, err := ioutil.ReadFile("test/data/assets/images/helloshaohua.jpeg")
				assert.NoError(t, err)
				return data
			}(),
			expectedContains: "test/data/create/images/users/header-specific-extension.png",
		},
	}

	for _, grid := range grids {

		ops := []Option{WithSpecificBaseDir(grid.baseDir), WithSpecificFileSystemPath(grid.savePath)}
		if grid.useExtension {
			ops = append(ops, WithSpecificFileExtension(grid.extension))
		}

		actual, err := CreateFile(
			grid.filename,
			grid.data, ops...)
		assert.NoError(t, err)
		assert.Contains(t, actual, grid.expectedContains)
	}
}
