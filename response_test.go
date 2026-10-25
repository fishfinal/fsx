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
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetResponseHeaderFunc(t *testing.T) {
	grids := []struct {
		filename         string
		response         *httptest.ResponseRecorder
		contentType      string
		expectedFilename string
	}{
		{
			filename:         "test/data/assets/text-files/hello.txt",
			response:         httptest.NewRecorder(),
			contentType:      "text/plain",
			expectedFilename: "UTF-8''hello.txt",
		},
		{
			filename:         "test/data/assets/images/helloshaohua.jpeg",
			response:         httptest.NewRecorder(),
			contentType:      "image/jpeg",
			expectedFilename: "UTF-8''helloshaohua.jpeg",
		},
	}

	for _, grid := range grids {
		err := DownloadFile(grid.response, grid.filename)
		assert.NoError(t, err)

		actualFilename := GetResponseHeaderFunc(grid.response, "Content-Disposition", ContentDispositionFilename())
		assert.Equal(t, grid.expectedFilename, actualFilename)
	}

}
