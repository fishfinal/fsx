// Copyright 2021 fishfinal <contact.fishfinal@gmail.com>;
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fsx

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestAbsSuite(t *testing.T) {
	suite.Run(t, &AbsSuite{})
}

type AbsSuite struct {
	suite.Suite
	baseDir string
	genDir  string
}

func (suite *AbsSuite) BeforeTest(suiteName, testName string) {
	suite.baseDir = "."
	testDir := MustGetFilePathWithFileSystemPath(suite.baseDir, WithSpecificFileSystemPath("test"))
	suite.genDir = filepath.Join(testDir, "gen")
}

func (suite *AbsSuite) Test_AbsPath() {
	grids := []struct {
		dir            string
		create         bool
		expectedExists bool
	}{
		{
			dir:            filepath.Join(suite.genDir, "abs.path"),
			create:         true,
			expectedExists: true,
		},
		{
			dir:            filepath.Join(suite.genDir, "abs.path", "not.mkdir"),
			create:         false,
			expectedExists: false,
		},
	}

	for _, grid := range grids {
		actual, err := AbsPath(grid.dir, grid.create)
		assert.NoError(suite.T(), err)
		assert.Contains(suite.T(), actual, grid.dir)
		if grid.create {
			assert.DirExists(suite.T(), actual)
		} else {
			assert.NoDirExists(suite.T(), actual)
		}
	}
}
