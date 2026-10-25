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

type Option interface {
	apply(*options)
}

type options struct {
	extension          FileExtension
	dot                bool
	trim               bool
	path               FileSystemPath
	baseDir            string
	containsMarkString bool

	// Whether the operation is force.
	force bool

	// Trim specify prefix string.
	trimPrefixStr string

	// Whether blank lines are ignored.
	ignoreBlankLine bool
}

type optionFunc func(*options)

func (o optionFunc) apply(ops *options) {
	o(ops)
}

// WithSpecificFileSystemPath
// Override the original value by specifying filepath.
func WithSpecificFileSystemPath(path FileSystemPath) Option {
	return optionFunc(func(ops *options) {
		ops.path = path
	})
}

// WithSpecificBaseDir
// Override the original value by specifying base directory.
func WithSpecificBaseDir(baseDir string) Option {
	return optionFunc(func(ops *options) {
		ops.baseDir = baseDir
	})
}

// WithSpecificFileExtension
// Override the original value by specifying file extension.
func WithSpecificFileExtension(extension FileExtension) Option {
	return optionFunc(func(ops *options) {
		ops.extension = extension
	})
}

// WithSpecificFileExtensionContainsDot
// Override the original value by specifying dot.
func WithSpecificFileExtensionContainsDot(containsdot bool) Option {
	return optionFunc(func(ops *options) {
		ops.dot = containsdot
	})
}

// WithSpecificTrimOriginalFileExtension
// Override the original value by specifying trim.
func WithSpecificTrimOriginalFileExtension(trim bool) Option {
	return optionFunc(func(ops *options) {
		ops.trim = trim
	})
}

// WithSpecificContainsMarkString
// Override the original value by containsMarkString mark string.
func WithSpecificContainsMarkString(contains bool) Option {
	return optionFunc(func(ops *options) {
		ops.containsMarkString = contains
	})
}

// WithOriginalFileNameTrimPrefix WithSpecificContainsMarkString Override the original value by trimPrefixStr string.
func WithOriginalFileNameTrimPrefix(str string) Option {
	return optionFunc(func(ops *options) {
		ops.trimPrefixStr = str
	})
}

// WithIgnoreBlankLine Specify whether blank lines are ignored.
func WithIgnoreBlankLine(ignored bool) Option {
	return optionFunc(func(ops *options) {
		ops.ignoreBlankLine = ignored
	})
}
