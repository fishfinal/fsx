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
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLines(t *testing.T) {
	grids := []struct {
		file     string
		expected int
	}{
		{
			file:     "./test/data/assets/text-files/hello.txt",
			expected: 0,
		},
		{
			file:     "./test/data/assets/text-files/helloworld.txt",
			expected: 12,
		},
		{
			file:     "./test/data/assets/text-files/helloshaohua.txt",
			expected: 22,
		},
		{
			file:     "./test/data/assets/text-files/tom.txt",
			expected: 2,
		},
	}

	for _, grid := range grids {
		data, err := ioutil.ReadFile(grid.file)
		assert.NoError(t, err)

		actual := Lines(data)
		assert.Equal(t, grid.expected, actual)
	}
}

func TestLinesForFile(t *testing.T) {
	grids := []struct {
		file     string
		expected int
		err      error
	}{
		{
			file:     "./test/data/assets/text-files/hello.txt",
			expected: 0,
			err:      nil,
		},
		{
			file:     "./test/data/assets/text-files/helloworld.txt",
			expected: 12,
			err:      nil,
		},
		{
			file:     "./test/data/assets/text-files/helloshaohua.txt",
			expected: 22,
			err:      nil,
		},
		{
			file:     "./test/data/assets/text-files/helloshaohua-not-exists.txt",
			expected: 0,
			err:      errors.New("no such file or directory"),
		},
	}

	for _, grid := range grids {
		actual, err := LinesForFile(grid.file)
		if err != nil {
			assert.Contains(t, err.Error(), grid.err.Error())
		}
		assert.Equal(t, grid.expected, actual)
	}
}

func TestCharacters(t *testing.T) {
	grids := []struct {
		file     string
		expected int
	}{
		{
			file:     "./test/data/assets/text-files/hellochina.txt",
			expected: 110,
		},
		{
			file:     "./test/data/assets/text-files/helloshaohua.txt",
			expected: 99,
		},
		{
			file:     "./test/data/assets/text-files/helloworld.txt",
			expected: 31,
		},
		{
			file:     "./test/data/assets/text-files/hello.txt",
			expected: 0,
		},
	}

	for _, grid := range grids {
		file, err := ioutil.ReadFile(grid.file)
		assert.NoError(t, err)

		actual := Characters(file)
		assert.Equal(t, grid.expected, actual)
	}
}

func TestCharactersForFile(t *testing.T) {
	grids := []struct {
		file     string
		expected int
		err      error
	}{
		{
			file:     "./test/data/assets/text-files/hellochina.txt",
			expected: 110,
			err:      nil,
		},
		{
			file:     "./test/data/assets/text-files/helloshaohua.txt",
			expected: 99,
			err:      nil,
		},
		{
			file:     "./test/data/assets/text-files/helloworld.txt",
			expected: 31,
			err:      nil,
		},
		{
			file:     "./test/data/assets/text-files/hello.txt",
			expected: 0,
			err:      nil,
		},
		{
			file:     "./test/data/assets/text-files/hello-not-exists.txt",
			expected: 0,
			err:      errors.New("no such file or directory"),
		},
	}

	for _, grid := range grids {
		actual, err := CharactersForFile(grid.file)
		if err != nil {
			assert.Contains(t, err.Error(), grid.err.Error())
		}

		assert.Equal(t, grid.expected, actual)
	}
}
