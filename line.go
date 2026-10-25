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

// Lines Count the number of data rows
func Lines(data []byte) int {
	if len(data) == 0 {
		return 0
	}
	return len(bytes.Split(data, []byte("\n")))
}

// LinesForFile Count the number of data rows from file.
func LinesForFile(filename string) (int, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	return Lines(data), nil
}

// Characters Count the number of data characters
func Characters(data []byte) int {
	if len(data) == 0 {
		return 0
	}
	data = CleanBlankLines(data)

	bu := bytes.Buffer{}
	defer bu.Reset()

	ls := bytes.SplitAfter(data, []byte("\n"))
	for _, l := range ls {
		bu.Write(bytes.TrimSpace(l))
	}

	return len(bytes.Runes(bu.Bytes()))
}

// CharactersForFile Count the number of data characters
func CharactersForFile(filename string) (int, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	return Characters(data), nil
}
