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
	"github.com/stretchr/testify/suite"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlankLineSuite(t *testing.T) {
	suite.Run(t, &BlankLineSuite{})
}

type BlankLineSuite struct {
	suite.Suite
	baseDir string
	dataDir string
	genDir  string
}

func (suite *BlankLineSuite) BeforeTest(suiteName, testName string) {
	suite.baseDir = "."
	testDir := MustGetFilePathWithFileSystemPath(suite.baseDir, WithSpecificFileSystemPath("test"))
	suite.dataDir = filepath.Join(testDir, "data")
	suite.genDir = filepath.Join(testDir, "gen")
}

func (suite *BlankLineSuite) TestCleanBlankLines() {
	grids := []struct {
		file     string
		expected int
		err      error
	}{
		{
			file:     filepath.Join(suite.dataDir, "assets", "text-files", "hellochina.txt"),
			expected: 16,
			err:      nil,
		},
		{
			file:     filepath.Join(suite.dataDir, "assets", "text-files", "helloshaohua.txt"),
			expected: 11,
			err:      nil,
		},
		{
			file:     filepath.Join(suite.dataDir, "assets", "text-files", "helloworld.txt"),
			expected: 6,
			err:      nil,
		},
		{
			file:     filepath.Join(suite.dataDir, "assets", "text-files", "hello.txt"),
			expected: 0,
			err:      nil,
		},
		{
			file:     filepath.Join(suite.dataDir, "assets", "text-files", "hello-not-exists.txt"),
			expected: 0,
			err:      errors.New("no such file or directory"),
		},
	}

	for _, grid := range grids {
		data, err := os.ReadFile(grid.file)
		if err != nil {
			assert.Contains(suite.T(), err.Error(), grid.err.Error())
		}

		actual := Lines(CleanBlankLines(data))
		assert.Equal(suite.T(), grid.expected, actual)
	}
}

func (suite *BlankLineSuite) TestCleanBlankLinesForFile() {
	grids := []struct {
		file     string
		expected int
		err      error
	}{
		{
			file:     filepath.Join(suite.dataDir, "assets", "text-files", "hellochina.txt"),
			expected: 16,
			err:      nil,
		},
		{
			file:     filepath.Join(suite.dataDir, "assets", "text-files", "helloshaohua.txt"),
			expected: 11,
			err:      nil,
		},
		{
			file:     filepath.Join(suite.dataDir, "assets", "text-files", "helloworld.txt"),
			expected: 6,
			err:      nil,
		},
		{
			file:     filepath.Join(suite.dataDir, "assets", "text-files", "hello.txt"),
			expected: 0,
			err:      nil,
		},
		{
			file:     filepath.Join(suite.dataDir, "assets", "text-files", "hello-not-exists.txt"),
			expected: 0,
			err:      errors.New("no such file or directory"),
		},
	}

	for _, grid := range grids {

		data, err := CleanBlankLinesForFile(grid.file)
		if err != nil {
			assert.Contains(suite.T(), err.Error(), grid.err.Error())
		}

		assert.Equal(suite.T(), grid.expected, Lines(data))
	}
}
