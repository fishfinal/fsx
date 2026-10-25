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
)

// MarkContentPrefixClear Mark content prefix clear.
func MarkContentPrefixClear(data []byte, mark string, ops ...Option) []byte {
	options := options{
		containsMarkString: true,
	}

	for _, o := range ops {
		o.apply(&options)
	}

	index := bytes.Index(data, []byte(mark))
	if index > 0 {
		if !options.containsMarkString {
			return data[index+(len([]byte(mark))):]
		}
		return data[index:]
	}

	return data
}

// MarkContentSuffixClear Mark content suffix clear.
func MarkContentSuffixClear(data []byte, mark string, ops ...Option) []byte {
	options := options{
		containsMarkString: false,
	}

	for _, o := range ops {
		o.apply(&options)
	}

	index := bytes.Index(data, []byte(mark))
	if index > 0 {
		if options.containsMarkString {
			return data[:index+(len([]byte(mark)))]
		}
		return data[:index]
	}
	return data
}
