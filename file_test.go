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
	"bytes"
	"embed"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateFilename(t *testing.T) {
	grids := []struct {
		filename  string
		extension string
		trim      bool
		expected  string
	}{
		{
			filename:  "hello/world/helloshaohua/abc.txt",
			extension: "png",
			trim:      true,
			expected:  "hello/world/helloshaohua/abc.png",
		},
		{
			filename:  "hello/world/abc.jpeg",
			extension: "txt",
			trim:      false,
			expected:  "hello/world/abc.jpeg.txt",
		},
	}

	for _, grid := range grids {
		actual := CreateFilename(grid.filename, grid.extension,
			WithSpecificTrimOriginalFileExtension(grid.trim))
		assert.Equal(t, grid.expected, actual)
	}
}

func TestGetFileExtension(t *testing.T) {
	grids := []struct {
		filename string
		dot      bool
		expected string
	}{
		{
			filename: "hello/world/helloshaohua/abc.txt",
			dot:      true,
			expected: ".txt",
		},
		{
			filename: "hello/world/abc.jpeg",
			dot:      false,
			expected: "jpeg",
		},
		{
			filename: "helloshaohua/abc.png",
			dot:      false,
			expected: "png",
		},
		{
			filename: "abc",
			dot:      true,
			expected: "",
		},
	}

	for _, grid := range grids {
		actual := GetFileExtension(grid.filename,
			WithSpecificFileExtensionContainsDot(grid.dot))
		assert.Equal(t, grid.expected, actual)
	}
}

//go:embed test/data/embed
var efs embed.FS

func TestGetFileContentWithEmbedFS(t *testing.T) {
	grids := []struct {
		filename        string
		expectedError   error
		expectedContext []byte
	}{
		{
			filename:        "a.txt",
			expectedError:   fmt.Errorf(`open a.txt: file does not exist`),
			expectedContext: []byte(`a.txt file.`),
		},
		{
			filename:        "test/data/embed/a.txt",
			expectedError:   nil,
			expectedContext: []byte(`a.txt file.`),
		},
		{
			filename:        "test/data/embed/b.txt",
			expectedError:   nil,
			expectedContext: []byte(`b.txt file.`),
		},
		{
			filename:        "test/data/embed/c.txt",
			expectedError:   nil,
			expectedContext: []byte(`c.txt file.`),
		},
		{
			filename:        "test/data/embed/VERSION",
			expectedError:   nil,
			expectedContext: []byte(`v1.2.1`),
		},
		{
			filename:      "test/data/embed/readme.md",
			expectedError: nil,
			expectedContext: []byte(`# hello world

I'm helloshaohua.`),
		},
	}

	for _, grid := range grids {
		actual, err := GetFileContentBytesWithEmbedFS(efs, grid.filename)
		if err != nil {
			assert.Equal(t, grid.expectedError.Error(), err.Error())
		} else {
			assert.Equal(t, grid.expectedContext, actual)
		}
	}
}

func TestGetFileContentStringSliceWithEmbedFS(t *testing.T) {
	grids := []struct {
		filename        string
		option          Option
		expectedError   error
		expectedContext []string
	}{
		{
			filename:        "a.txt",
			expectedError:   fmt.Errorf(`open a.txt: file does not exist`),
			expectedContext: []string{`a.txt file.`},
		},
		{
			filename:        "test/data/embed/a.txt",
			expectedError:   nil,
			expectedContext: []string{`a.txt file.`},
		},
		{
			filename:        "test/data/embed/b.txt",
			expectedError:   nil,
			expectedContext: []string{`b.txt file.`},
		},
		{
			filename:        "test/data/embed/c.txt",
			expectedError:   nil,
			expectedContext: []string{`c.txt file.`},
		},
		{
			filename:        "test/data/embed/VERSION",
			expectedError:   nil,
			expectedContext: []string{"v1.2.1"},
		},
		{
			filename:        "test/data/embed/readme.md",
			expectedError:   nil,
			option:          WithIgnoreBlankLine(true),
			expectedContext: []string{"# hello world", "I'm helloshaohua."},
		},
		{
			filename:        "test/data/embed/readme.md",
			expectedError:   nil,
			option:          WithIgnoreBlankLine(false),
			expectedContext: []string{"# hello world", "", "I'm helloshaohua."},
		},
	}

	for _, grid := range grids {
		options := make([]Option, 0)
		if grid.option != nil {
			options = append(options, grid.option)
		}

		actual, err := GetFileContentStringSliceWithEmbedFS(efs, grid.filename, options...)
		if err != nil {
			assert.Equal(t, grid.expectedError.Error(), err.Error())
		} else {
			assert.Equal(t, grid.expectedContext, actual)
		}
	}
}

func TestGetFileContentStringSliceWithBuffer(t *testing.T) {
	file := `hello world
helloshaohua

hello@shaohua.com`

	// by default
	buffer, err := GetFileContentStringSliceWithBuffer(
		bytes.NewBufferString(file),
	)
	assert.NoError(t, err)
	assert.Len(t, buffer, 4)

	// ignore blank line
	buffer, err = GetFileContentStringSliceWithBuffer(
		bytes.NewBufferString(file),
		WithIgnoreBlankLine(true),
	)
	assert.NoError(t, err)
	assert.Len(t, buffer, 3)

}

func TestTouch(t *testing.T) {
	grids := []struct {
		filename string
	}{
		{
			filename: "test/data/touch/.history",
		},
		{
			filename: "test/data/touch/helloshaohua/.keep",
		},
		{
			filename: "test/data/touch/tom/.history",
		},
		{
			filename: "test/data/touch/lily/.history",
		},
		{
			filename: "test/data/touch/kitty/hello/world/.history",
		},
	}

	for _, grid := range grids {
		err := Touch(grid.filename)
		assert.NoError(t, err)
	}
}

func TestTruncate(t *testing.T) {
	grids := []struct {
		filename string
		content  string
		filesize int64
	}{
		{
			filename: "test/data/truncate/helloshaohua.txt",
			content:  `wu.shaohua@foxmail.com`,
			filesize: 0,
		},
		{
			filename: "test/data/truncate/users.txt",
			content: `张三
李四
王五`,
			filesize: 0,
		},
	}

	for _, grid := range grids {
		err := os.WriteFile(grid.filename, []byte(grid.content), os.ModePerm)
		assert.NoError(t, err)

		// Read file size on the truncate file after.
		info, err := os.Stat(grid.filename)
		assert.NoError(t, err)
		assert.NotEmpty(t, info.Size())

		err = Truncate(grid.filename, 0)
		assert.NoError(t, err)

		// Read file size on the truncate file after.
		info, err = os.Stat(grid.filename)
		assert.NoError(t, err)
		assert.NotNil(t, info)
		assert.Equal(t, grid.filesize, info.Size())
	}
}

func TestGetFileOrDirectoryName(t *testing.T) {
	grids := []struct {
		path     string
		dir      bool
		filename string
		err      error
	}{
		{
			path:     "test/data/assets/images/helloshaohua.jpeg",
			dir:      false,
			filename: "helloshaohua",
			err:      nil,
		},
		{
			path:     "test/data/assets/images/",
			dir:      true,
			filename: "images",
			err:      nil,
		},
	}

	for _, grid := range grids {
		filename, dir, err := GetFileOrDirectoryName(grid.path)
		assert.Equal(t, grid.filename, filename)
		assert.Equal(t, grid.dir, dir)
		assert.Equal(t, grid.err, err)
	}
}

func TestIsFile(t *testing.T) {
	grids := []struct {
		filename string
		error    error
		expected bool
	}{
		{
			filename: "test",
			error:    nil,
			expected: false,
		},
		{
			filename: "none.go",
			error:    errors.New("stat none.go: no such file or directory"),
			expected: false,
		},
		{
			filename: "create.go",
			error:    nil,
			expected: true,
		},
	}

	for _, grid := range grids {
		actual, err := IsFile(grid.filename)
		if err != nil {
			assert.Equal(t, grid.error.Error(), err.Error())
		} else {
			assert.Equal(t, grid.expected, actual)
		}
	}
}

func TestFilenameTrimPrefix(t *testing.T) {
	grids := []struct {
		filename string
		prefix   string
		expected string
	}{
		{
			filename: "D:/Users/hello/world/hello.txt",
			prefix:   "D:/Users/hello/world",
			expected: "/hello.txt",
		},
		{
			filename: "/Users/hello/world/hello.txt",
			prefix:   "/Users/hello/",
			expected: "world/hello.txt",
		},
		{
			filename: "/Users/hello/world/hello.txt",
			prefix:   "/Users/",
			expected: "hello/world/hello.txt",
		},
	}

	for _, grid := range grids {
		actual := FilenameTrimPrefix(grid.filename, grid.prefix)
		assert.Equal(t, grid.expected, actual)
	}
}

func TestGetFileContentWithStringSlice(t *testing.T) {
	grids := []struct {
		filename string
		ops      []Option
		expected int
	}{
		{
			filename: "./test/data/git/.gitignore",
			ops:      nil,
			expected: 24,
		},
		{
			filename: "./test/data/git/.gitignore",
			ops: []Option{
				WithIgnoreBlankLine(true),
			},
			expected: 19,
		},
	}

	for _, grid := range grids {
		actual, err := GetFileContentStringSliceWithFilename(grid.filename, grid.ops...)
		assert.NoError(t, err)
		assert.Len(t, actual, grid.expected)
	}
}

func TestGetContentWithRegularExpression(t *testing.T) {
	grids := []struct {
		filename   string
		expression string
		expected   int
	}{
		{
			filename:   "test/data/regular/app.log",
			expression: `\{"connnect"\s?:\s?.*,\s?"file_name"\s?:\s?".*"\}`,
			expected:   1,
		},
	}

	for _, grid := range grids {
		data, err := ioutil.ReadFile(grid.filename)
		assert.NoError(t, err)

		actual, err := GetContentWithRegularExpression(string(data), grid.expression)
		assert.NoError(t, err)
		assert.Len(t, actual, grid.expected)
	}
}
