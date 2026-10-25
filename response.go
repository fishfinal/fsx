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
	"strings"
)

// GetResponseHeader Gets response header field value.
func GetResponseHeader(resp http.ResponseWriter, field string) string {
	return resp.Header().Get(field)
}

// GetResponseHeaderFunc GetResponseHeader Gets response header field value,
// processing of response field values by attaching custom functions.
func GetResponseHeaderFunc(resp http.ResponseWriter, field string, process func(val string) string) string {
	return process(GetResponseHeader(resp, field))
}

// ContentDispositionFilename Gets filename for response header Content-Disposition.
// Header such as Content-Disposition: attachment; filename="helloshaohua.txt"
func ContentDispositionFilename() func(val string) string {
	return func(val string) string {
		sli := strings.SplitAfter(val, ";")
		for _, item := range sli {
			if space := strings.TrimSpace(item); strings.Contains(space, "filename") {
				fsl := strings.SplitAfter(space, "=")
				return strings.Trim(fsl[len(fsl)-1], "\"")
			}
		}
		return val
	}
}
