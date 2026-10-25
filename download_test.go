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
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestDownloadSuite(t *testing.T) {
	suite.Run(t, new(DownloadTestSuite))
}

type DownloadTestSuite struct {
	suite.Suite
	tempDir     string
	testFile    string
	testContent []byte
}

func (suite *DownloadTestSuite) SetupTest() {
	// 创建临时目录
	var err error
	suite.tempDir, err = os.MkdirTemp("", "download_test")
	suite.Require().NoError(err, "创建临时目录失败")

	// 创建测试文件
	suite.testFile = filepath.Join(suite.tempDir, "test.txt")
	suite.testContent = []byte("测试内容，Hello 世界!")
	err = os.WriteFile(suite.testFile, suite.testContent, 0644)
	suite.Require().NoError(err, "创建测试文件失败")
}

func (suite *DownloadTestSuite) TearDownTest() {
	// 清理临时目录
	if suite.tempDir != "" {
		suite.NoError(os.RemoveAll(suite.tempDir), "清理临时目录失败")
	}
}

func (suite *DownloadTestSuite) TestDownloadFile_Success() {
	// 准备测试
	rr := httptest.NewRecorder()

	// 执行测试
	err := DownloadFile(rr, suite.testFile)

	// 验证结果
	suite.NoError(err, "DownloadFile 不应返回错误")
	suite.Equal(http.StatusOK, rr.Code, "HTTP 状态码不正确")

	// 验证响应头
	expectedDisposition := `attachment; filename*=UTF-8''test.txt`
	suite.Equal(expectedDisposition, rr.Header().Get("Content-Disposition"), "Content-Disposition 不正确")

	expectedContentType := "application/octet-stream"
	suite.Equal(expectedContentType, rr.Header().Get("Content-Type"), "Content-Type 不正确")
	expectedContentLength := "28"
	suite.Equal(expectedContentLength, rr.Header().Get("Content-Length"), "Content-Length 不正确")

	// 验证响应体
	suite.Equal(suite.testContent, rr.Body.Bytes(), "响应体内容不正确")
}

func (suite *DownloadTestSuite) TestDownloadFile_NonExistentFile() {
	// 准备测试
	rr := httptest.NewRecorder()
	nonExistentFile := filepath.Join(suite.tempDir, "nonexistent.txt")

	// 执行测试
	err := DownloadFile(rr, nonExistentFile)

	// 验证结果
	suite.Error(err, "对于不存在的文件应返回错误")
	suite.Equal(http.StatusNotFound, rr.Code, "当文件不存在时不应设置 HTTP 状态码")
}

func (suite *DownloadTestSuite) TestDownloadFile_WithSpecialCharacters() {
	// 准备特殊文件名测试文件
	specialFile := filepath.Join(suite.tempDir, "测试 文件@123.txt")
	err := os.WriteFile(specialFile, suite.testContent, 0644)
	suite.Require().NoError(err, "创建特殊文件名测试文件失败")

	// 准备测试
	rr := httptest.NewRecorder()

	// 执行测试
	err = DownloadFile(rr, specialFile)

	// 验证结果
	suite.NoError(err, "DownloadFile 不应返回错误")

	// 验证文件名编码
	expectedDisposition := `attachment; filename*=UTF-8''%E6%B5%8B%E8%AF%95%20%E6%96%87%E4%BB%B6%40123.txt`
	suite.Equal(expectedDisposition, rr.Header().Get("Content-Disposition"), "特殊文件名编码不正确")
}

func (suite *DownloadTestSuite) TestDownloadFile_WithDifferentMimeTypes() {
	// 测试不同 MIME 类型的文件
	testCases := []struct {
		filename     string
		content      []byte
		expectedMime string
	}{
		{"image.png", []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}, "application/octet-stream"},
		{"test.json", []byte(`{"key": "value"}`), "application/octet-stream"},
		{"binary.bin", []byte{0x00, 0x01, 0x02, 0x03}, "application/octet-stream"},
	}

	for _, tc := range testCases {
		filePath := filepath.Join(suite.tempDir, tc.filename)
		err := os.WriteFile(filePath, tc.content, 0644)
		suite.Require().NoError(err, "创建测试文件失败: %s", tc.filename)

		rr := httptest.NewRecorder()
		err = DownloadFile(rr, filePath)

		suite.NoError(err, "DownloadFile 不应返回错误: %s", tc.filename)
		suite.Equal(tc.expectedMime, rr.Header().Get("Content-Type"), "MIME 类型不匹配: %s", tc.filename)
	}
}

func (suite *DownloadTestSuite) TestDownloadFileByIOReader() {

	// 准备测试
	rr := httptest.NewRecorder()

	// 执行测试
	filename := "测试 文件@123.txt"
	err := DownloadFileByIOReader(rr, bytes.NewReader(suite.testContent), filename, int64(len(suite.testContent)))

	// 验证结果
	suite.NoError(err, "DownloadFile 不应返回错误")

	// 验证文件名编码
	expectedDisposition := `attachment; filename*=UTF-8''%E6%B5%8B%E8%AF%95%20%E6%96%87%E4%BB%B6%40123.txt`
	suite.Equal(expectedDisposition, rr.Header().Get("Content-Disposition"), "特殊文件名编码不正确")
}
