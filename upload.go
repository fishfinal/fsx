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
	"net/http"

	"github.com/gabriel-vasile/mimetype"
)

func UploadIsAllowMIMEWithFormFile(request *http.Request, key string, allows []string) (bool, error) {
	// Gets the file for upload form.
	file, _, err := request.FormFile(key)
	if err != nil {
		return false, err
	}
	defer file.Close()

	// Gets the file content mime.
	mime, err := mimetype.DetectReader(file)
	if err != nil {
		return false, err
	}

	// Match allow mime.
	mimeStr := mime.String()
	allowed := false
	for _, allow := range allows {
		if allow == mimeStr {
			allowed = true
		}
	}

	return allowed, nil
}
