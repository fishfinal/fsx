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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInSlice(t *testing.T) {
	grids := []struct {
		source   []interface{}
		find     interface{}
		expected bool
	}{
		{
			source:   []interface{}{"hello", "world"},
			find:     "hello",
			expected: true,
		},
		{
			source:   []interface{}{"hello", "world"},
			find:     "not_exists",
			expected: false,
		},
		{
			source:   []interface{}{"hello", "world"},
			find:     5,
			expected: false,
		},
	}

	for _, grid := range grids {
		actual := InSlice(grid.source, grid.find)
		assert.Equal(t, actual, grid.expected)
	}
}
