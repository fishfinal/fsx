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
	"io/ioutil"
)

// CleanBlankLines clean blank lines with file.
func CleanBlankLines(data []byte) []byte {
	if len(data) == 0 {
		return nil
	}
	bu := bytes.Buffer{}
	ls := bytes.SplitAfter(data, []byte("\n"))
	for _, l := range ls {
		if len(bytes.TrimSpace(l)) == 0 {
			continue
		}
		bu.Write(l)
	}
	return bu.Bytes()
}

// CleanBlankLinesForFile clean blank lines with file.
func CleanBlankLinesForFile(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := CleanBlankLines(data)
	return lines, nil
}
