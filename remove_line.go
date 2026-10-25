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
	"regexp"
)

// CleanSpecificContentLines clean specific content lines.
func CleanSpecificContentLines(data []byte, findStr string) []byte {
	buffer := bytes.Buffer{}

	lines := bytes.SplitAfter(data, []byte("\n"))
	for _, line := range lines {
		if len(regexp.MustCompile(findStr).Find(line)) != 0 {
			continue
		}

		buffer.Write(line)
	}

	return buffer.Bytes()
}

// CleanSpecificContentLinesForFile clean specific content lines for file.
func CleanSpecificContentLinesForFile(filename string, findStr string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	buffer := bytes.Buffer{}
	lines := bytes.SplitAfter(data, []byte("\n"))
	for _, line := range lines {
		if len(regexp.MustCompile(findStr).Find(line)) != 0 {
			continue
		}

		buffer.Write(line)
	}

	return buffer.Bytes(), nil
}
