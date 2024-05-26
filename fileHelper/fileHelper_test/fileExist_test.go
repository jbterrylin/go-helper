package filehelper_test

import (
	"os"
	"testing"

	filehelper "github.com/jbterrylin/go-helper/fileHelper"
)

func TestFileExist(t *testing.T) {
	// 创建临时文件用于测试
	tmpFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	tmpFilePath := tmpFile.Name()
	tmpFile.Close()
	defer os.Remove(tmpFilePath)

	// 创建临时目录用于测试
	tmpDir, err := os.MkdirTemp("", "testdir")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	tests := []struct {
		path     string
		expected bool
		err      error
	}{
		{tmpFilePath, true, nil},                        // 文件存在
		{tmpDir, false, filehelper.ErrPathIsPointToDir}, // 路径指向目录
		{"nonexistentfile", false, nil},                 // 文件不存在
	}

	for _, tt := range tests {
		result, err := filehelper.FileExist(tt.path)
		if result != tt.expected || (err != nil && err.Error() != tt.err.Error()) {
			t.Errorf("FileExist(%v) = (%v, %v); want (%v, %v)", tt.path, result, err, tt.expected, tt.err)
		}
	}
}
