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
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestDirectorySuite(t *testing.T) {
	suite.Run(t, &DirectorySuite{})
}

type DirectorySuite struct {
	suite.Suite
}

func (suite *DirectorySuite) TestIsNotExists() {
	grids := []struct {
		filename string
		expected bool
	}{
		{
			filename: "create.go",
			expected: false,
		},
		{
			filename: "none.go",
			expected: true,
		},
	}

	for _, grid := range grids {
		actual := IsNotExists(grid.filename)
		assert.Equal(suite.T(), grid.expected, actual)
	}
}

func (suite *DirectorySuite) TestExists() {
	grids := []struct {
		filename string
		expected bool
	}{
		{
			filename: "create.go",
			expected: true,
		},
		{
			filename: "none.go",
			expected: false,
		},
	}

	for _, grid := range grids {
		actual := Exists(grid.filename)
		assert.Equal(suite.T(), grid.expected, actual)
	}
}

func (suite *DirectorySuite) TestIsNotExistsWithEmbedFS() {
	grids := []struct {
		filename string
		expected bool
	}{
		{
			filename: "test/data/embed/a.txt",
			expected: false,
		},
		{
			filename: "test/data/embed/b.txt",
			expected: false,
		},
		{
			filename: "test/data/embed/c.txt",
			expected: false,
		},
		{
			filename: "test/data/embed/VERSION",
			expected: false,
		},
		{
			filename: "a.txt",
			expected: true,
		},
		{
			filename: "/hello/world/not-exists-file.txt",
			expected: true,
		},
	}

	for _, grid := range grids {
		actual := IsNotExistsWithEmbedFS(efs, grid.filename)
		assert.Equal(suite.T(), grid.expected, actual)
	}
}

func (suite *DirectorySuite) TestEmbedsFileExists() {
	grids := []struct {
		filename string
		expected bool
	}{
		{
			filename: "test/data/embed/a.txt",
			expected: true,
		},
		{
			filename: "test/data/embed/b.txt",
			expected: true,
		},
		{
			filename: "test/data/embed/c.txt",
			expected: true,
		},
		{
			filename: "test/data/embed/VERSION",
			expected: true,
		},
		{
			filename: "a.txt",
			expected: false,
		},
		{
			filename: "/hello/world/not-exists-file.txt",
			expected: false,
		},
	}

	for _, grid := range grids {
		actual := EmbedsFileExists(efs, grid.filename)
		assert.Equal(suite.T(), grid.expected, actual)
	}
}

func (suite *DirectorySuite) TestMustNotExistsMkdir() {
	grids := []struct {
		path  string
		error error
	}{
		{
			path:  "test/data/tem/images/post",
			error: nil,
		},
		{
			path:  "test/data/tem/docs/images",
			error: nil,
		},
		{
			path:  "test/data/tem/article/comment",
			error: nil,
		},
	}

	for _, grid := range grids {
		err := NotExistsMkdir(grid.path)
		assert.NoError(suite.T(), err)
	}
}

func (suite *DirectorySuite) TestNotExistsMkdir() {
	grids := []struct {
		path  string
		error error
	}{
		{
			path:  "test/data/create/dir/a/helloshaohua",
			error: nil,
		},
		{
			path:  "test/data/create/dir/b/helloshaohua",
			error: nil,
		},
		{
			path:  "test/data/create/dir/c/helloshaohua",
			error: nil,
		},
	}

	for _, grid := range grids {
		err := NotExistsMkdir(grid.path)
		assert.NoError(suite.T(), err)
	}
}

func (suite *DirectorySuite) TestLastDirName() {
	grids := []struct {
		filename string
		expected string
	}{
		{
			filename: "hello/world/helloshaohua/a.txt",
			expected: "helloshaohua",
		},
		{
			filename: "hello/world/a.txt",
			expected: "world",
		},
		{
			filename: "helloshaohua/a.txt",
			expected: "helloshaohua",
		},
		{
			filename: "a.txt",
			expected: ".",
		},
	}

	for _, grid := range grids {
		actual := LastDirName(grid.filename)
		assert.Equal(suite.T(), grid.expected, actual)
	}
}

func (suite *DirectorySuite) TestIsDir() {
	grids := []struct {
		filename string
		error    error
		expected bool
	}{
		{
			filename: "test",
			error:    nil,
			expected: true,
		},
		{
			filename: "none.go",
			error:    errors.New("stat none.go: no such file or directory"),
			expected: false,
		},
		{
			filename: "create.go",
			error:    nil,
			expected: false,
		},
	}

	for _, grid := range grids {
		actual, err := IsDir(grid.filename)
		if err != nil {
			assert.Equal(suite.T(), grid.error.Error(), err.Error())
		} else {
			assert.Equal(suite.T(), grid.expected, actual)
		}
	}
}

func (suite *DirectorySuite) TestDirNotExists() {
	grids := []struct {
		path     string
		expected bool
	}{
		{
			path:     "test",
			expected: false,
		},
		{
			path:     "hello",
			expected: true,
		},
		{
			path:     "world",
			expected: true,
		},
		{
			path:     "hello/world/hello-world",
			expected: true,
		},
	}

	for _, grid := range grids {
		actual := DirNotExists(grid.path)
		assert.Equal(suite.T(), grid.expected, actual,
			"want %t but got %t, path: %s\n", grid.expected, actual, grid.path)
	}
}

func (suite *DirectorySuite) TestNotExistsMkdirAll() {
	grids := []struct {
		dir      string
		expected error
	}{
		{
			dir: "test/not.exists.mkdir.all/hello",
		},
		{
			dir: "test/not.exists.mkdir.all/world",
		},
		{
			dir: "test/not.exists.mkdir.all/helloshaohua",
		},
		{
			dir: "test/not.exists.mkdir.all/wu.shaohua@foxmail.com",
		},
	}

	for _, grid := range grids {
		err := NotExistsMkdirAll(grid.dir)
		assert.NoError(suite.T(), err)
	}
}

func (suite *DirectorySuite) TestReadDir() {
	grids := []struct {
		path  string
		error error
	}{
		{
			path:  "test",
			error: nil,
		},
		{
			path:  ".",
			error: nil,
		},
	}

	for _, grid := range grids {
		actual, err := ReadDir(grid.path, nil)
		assert.NoError(suite.T(), err)
		assert.Greater(suite.T(), len(actual), 0)
	}
}

func (suite *DirectorySuite) TestWriteDir() {
	grids := []struct {
		src string
		dst string
		ops []Option
	}{
		// Contains the hierarchy of copy directories.
		{
			src: "test/data",
			dst: "test/gen/filesystem/helloworld",
			ops: nil,
		},

		// Does not include the hierarchy of copied directories.
		{
			src: "test/data",
			dst: "test/gen/filesystem/shaohua",
			ops: []Option{
				WithOriginalFileNameTrimPrefix("test/data"),
			},
		},
	}

	for _, grid := range grids {
		actual, err := ReadDir(grid.src, nil)
		assert.NoError(suite.T(), err)
		assert.NotEmpty(suite.T(), actual)

		err = WriteDir(grid.dst, actual, grid.ops...)
		assert.NoError(suite.T(), err)
	}
}

func (suite *DirectorySuite) TestCopy() {
	grids := []struct {
		src string
		dst string
		ops []Option
	}{
		// Contains the hierarchy of copy directories.
		{
			src: "test/data",
			dst: "test/gen/filesystem/helloworld",
			ops: nil,
		},

		// Does not include the hierarchy of copied directories.
		{
			src: "test/data",
			dst: "test/gen/filesystem/shaohua",
			ops: []Option{
				WithOriginalFileNameTrimPrefix("test/data"),
			},
		},
	}

	for _, grid := range grids {
		actual, err := ReadDir(grid.src, nil)
		assert.NoError(suite.T(), err)
		assert.NotEmpty(suite.T(), actual)

		err = WriteDir(grid.dst, actual, grid.ops...)
		assert.NoError(suite.T(), err)
	}
}
