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
	"io"
	"mime/multipart"
	"os"
)

// UploadFileForMultipartFields Mock form upload file.
func UploadFileForMultipartFields(values map[string]interface{}) (contentType string, buffer bytes.Buffer, err error) {
	w := multipart.NewWriter(&buffer)
	defer w.Close()
	for key, r := range values {
		switch val := r.(type) {
		case *os.File:
			var fw io.Writer
			if fw, err = w.CreateFormFile(key, val.Name()); err != nil {
				return
			}
			if _, err = io.Copy(fw, val); err != nil {
				return
			}
		case string:
			if err = w.WriteField(key, val); err != nil {
				return
			}
		}
	}
	contentType = w.FormDataContentType()

	return
}
