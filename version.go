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
	"fmt"
	"strings"

	"github.com/fishfinal/expression"
)

// VersionGithub Github version number format.
type VersionGithub struct {
	Major string `json:"major"`
	Minor string `json:"minor"`
	Patch string `json:"patch"`
}

// Join version section, default use dot.
func (v *VersionGithub) Join(element ...string) string {
	return strings.Join(element, ".")
}

// GetFullVersion Get a full version string.
// Such as 1.16.6, will return 1.16.6
func (v *VersionGithub) GetFullVersion() string {
	return strings.Join([]string{v.Major, v.Minor, v.Patch}, ".")
}

// GetMinorVersion Get a minor version string.
// Such as 1.16.6, will return 1.16
func (v *VersionGithub) GetMinorVersion() string {
	return strings.Join([]string{v.Major, v.Minor}, ".")
}

// GetVersion Get Github version information in version number format.
func GetVersion(version string) (*VersionGithub, error) {
	replaced, err := expression.RegularReplaceMatch(`go`, version, "")
	if err != nil {
		return nil, err
	}

	slice := strings.Split(replaced, ".")
	if len(slice) != 3 {
		return nil, fmt.Errorf("go version invalid")
	}

	v := &VersionGithub{
		Major: slice[0],
		Minor: slice[1],
		Patch: slice[2],
	}

	return v, nil
}
