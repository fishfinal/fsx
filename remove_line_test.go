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

func TestCleanSpecificContentLines(t *testing.T) {
	grids := []struct {
		file     string
		fineStr  string
		expected int
	}{
		{
			file:     "test/data/assets/text-files/helloworld.txt",
			fineStr:  "山东",
			expected: 5,
		},
		{
			file:     "test/data/assets/text-files/helloworld.txt",
			fineStr:  "中国",
			expected: 5,
		},
		{
			file:     "test/data/assets/text-files/helloworld.txt",
			fineStr:  "hello",
			expected: 3,
		},
		{
			file:     "test/data/assets/text-files/helloworld.txt",
			fineStr:  "world",
			expected: 5,
		},
	}

	for _, grid := range grids {
		data, err := ioutil.ReadFile(grid.file)
		assert.NoError(t, err)

		// Clean blank lines.
		data = CleanBlankLines(data)

		// Clean specific content lines
		data = CleanSpecificContentLines(data, grid.fineStr)
		assert.Equal(t, grid.expected, Lines(data))
	}
}

func TestCleanSpecificContentLinesForFile(t *testing.T) {
	grids := []struct {
		file     string
		fineStr  string
		expected int
	}{
		{
			file:     "test/data/assets/text-files/helloworld.txt",
			fineStr:  "山东",
			expected: 11,
		},
		{
			file:     "test/data/assets/text-files/helloworld.txt",
			fineStr:  "中国",
			expected: 11,
		},
		{
			file:     "test/data/assets/text-files/helloworld.txt",
			fineStr:  "hello",
			expected: 9,
		},
		{
			file:     "test/data/assets/text-files/helloworld.txt",
			fineStr:  "world",
			expected: 11,
		},
	}

	for _, grid := range grids {
		data, err := CleanSpecificContentLinesForFile(grid.file, grid.fineStr)
		assert.NoError(t, err)
		assert.Equal(t, grid.expected, Lines(data))
	}
}
