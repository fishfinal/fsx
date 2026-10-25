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
	"runtime"
	"strings"
	"testing"

	"github.com/fishfinal/expression"
	"github.com/stretchr/testify/assert"
)

func TestGetVersion(t *testing.T) {
	replaced, err := expression.RegularReplaceMatch(`go`, runtime.Version(), "")
	assert.NoError(t, err)
	assert.NotEmpty(t, replaced)

	version, err := GetVersion(replaced)
	assert.NoError(t, err)
	assert.NotNil(t, version)
}

func TestGetFullVersion(t *testing.T) {
	version, err := GetVersion(runtime.Version())
	assert.NoError(t, err)
	assert.NotNil(t, version)

	fullVersion := version.GetFullVersion()
	assert.Len(t, strings.Split(fullVersion, "."), 3)
}

func TestGetMinorVersion(t *testing.T) {
	version, err := GetVersion(runtime.Version())
	assert.NoError(t, err)
	assert.NotNil(t, version)

	minorVersion := version.GetMinorVersion()
	assert.Len(t, strings.Split(minorVersion, "."), 2)
}
