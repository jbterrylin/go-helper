package filehelper_test

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	filehelper "github.com/jbterrylin/go-helper/fileHelper"
)

func TestZipFiles(t *testing.T) {
	// 创建临时目录用于测试
	tmpDir, err := os.MkdirTemp("", "testzip")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// 创建临时文件用于测试
	files := []string{}
	for i := 1; i <= 3; i++ {
		tmpFile, err := os.CreateTemp(tmpDir, "testfile")
		if err != nil {
			t.Fatalf("Failed to create temporary file: %v", err)
		}
		tmpFilePath := tmpFile.Name()
		files = append(files, tmpFilePath)
		tmpFile.Close()

		// 写入一些数据到临时文件
		if err := os.WriteFile(tmpFilePath, []byte("test content"), 0644); err != nil {
			t.Fatalf("Failed to write to temporary file: %v", err)
		}
	}

	// 创建临时ZIP文件用于测试
	tmpZipFile, err := os.CreateTemp(tmpDir, "testzipfile.zip")
	if err != nil {
		t.Fatalf("Failed to create temporary zip file: %v", err)
	}
	tmpZipFilePath := tmpZipFile.Name()
	tmpZipFile.Close()

	// 测试压缩文件
	err = filehelper.ZipFiles(tmpZipFilePath, files, tmpDir, "newdir")
	if err != nil {
		t.Fatalf("ZipFiles(%v, %v, %v, %v) returned error: %v", tmpZipFilePath, files, tmpDir, "newdir", err)
	}

	// 检查ZIP文件内容
	zipReader, err := zip.OpenReader(tmpZipFilePath)
	if err != nil {
		t.Fatalf("Failed to open zip file: %v", err)
	}
	defer zipReader.Close()

	for _, f := range zipReader.File {
		rc, err := f.Open()
		if err != nil {
			t.Fatalf("Failed to open file in zip: %v", err)
		}
		content, err := io.ReadAll(rc)
		rc.Close()
		if err != nil {
			t.Fatalf("Failed to read file content in zip: %v", err)
		}

		expectedContent := []byte("test content")
		if !bytes.Equal(content, expectedContent) {
			t.Errorf("File content in zip is %v; want %v", string(content), string(expectedContent))
		}

		expectedName := strings.Replace(f.Name, "newdir", tmpDir, -1)
		if !strings.HasPrefix(expectedName, tmpDir) {
			t.Errorf("File path in zip is %v; want prefix %v", f.Name, "newdir")
		}
	}
}
