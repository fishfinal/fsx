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
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// DownloadFile will write download file into the HTTP response body.
func DownloadFile(resp http.ResponseWriter, path string) error {
	// Stat to the file
	fileInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			resp.WriteHeader(http.StatusNotFound)
		} else {
			resp.WriteHeader(http.StatusInternalServerError)
		}
		return fmt.Errorf("failed to stat file, %w", err)
	}

	file, err := os.Open(path)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return err
	}
	defer file.Close()

	var filetype = "application/octet-stream"
	var filename = filepath.Base(path)

	// 安全编码文件名（RFC 5987）
	escaped := url.PathEscape(filename)
	escaped = strings.ReplaceAll(escaped, "@", "%40")

	// Setting http response header.
	resp.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename*=UTF-8''%s`, escaped))
	resp.Header().Set("Content-Type", filetype)
	resp.Header().Set("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))
	resp.Header().Set("X-Content-Type-Options", "nosniff")

	// 流式传输文件内容，使用 32KB 缓冲区
	if _, err = io.CopyBuffer(resp, file, make([]byte, 32*1024)); err != nil {
		return fmt.Errorf("transmission interrupt: %w", err)
	}

	return nil
}

// DownloadFileByIOReader streams file content from an io.Reader to the HTTP response as a downloadable file.
// filesize: File size in bytes. Pass -1 to let the function attempt to determine the size automatically
// (only works when reader implements io.ReadSeeker).
func DownloadFileByIOReader(resp http.ResponseWriter, reader io.Reader, filename string, filesize int64) error {
	filetype := "application/octet-stream"

	// RFC 5987 filename encoding for safe HTTP header transmission
	escaped := url.PathEscape(filename)
	escaped = strings.ReplaceAll(escaped, "@", "%40")

	// Auto-detect file size when caller passes -1
	if filesize < 0 {
		// Attempt to determine size if reader supports seeking
		if seeker, ok := reader.(io.ReadSeeker); ok {
			// Save current position
			currentPos, err := seeker.Seek(0, io.SeekCurrent)
			if err != nil {
				return fmt.Errorf("failed to get current position: %w", err)
			}

			// Determine total file size by seeking to end
			size, err := seeker.Seek(0, io.SeekEnd)
			if err != nil {
				return fmt.Errorf("failed to seek to end: %w", err)
			}

			filesize = size
			// Reset to original position for subsequent reading
			_, err = seeker.Seek(currentPos, io.SeekStart)
			if err != nil {
				return fmt.Errorf("failed to reset position: %w", err)
			}
		}
	}

	// Set HTTP response headers
	resp.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename*=UTF-8''%s`, escaped))
	resp.Header().Set("Content-Type", filetype)

	// Set Content-Length only when known (enables progress indicators and better client handling)
	if filesize >= 0 {
		resp.Header().Set("Content-Length", strconv.FormatInt(filesize, 10))
	}

	resp.Header().Set("X-Content-Type-Options", "nosniff")

	// Stream file content using a 32KB buffer for optimal performance
	buffer := make([]byte, 32*1024)
	if _, err := io.CopyBuffer(resp, reader, buffer); err != nil {
		return fmt.Errorf("transmission interrupt: %w", err)
	}

	return nil
}
