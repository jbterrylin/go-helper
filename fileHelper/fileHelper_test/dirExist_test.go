package filehelper_test

import (
	"os"
	"testing"

	filehelper "github.com/jbterrylin/go-helper/fileHelper"
)

func TestDirExist(t *testing.T) {
	// 创建临时目录用于测试
	tmpDir, err := os.MkdirTemp("", "testdir")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// 创建临时文件用于测试
	tmpFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	tmpFilePath := tmpFile.Name()
	tmpFile.Close()
	defer os.Remove(tmpFilePath)

	tests := []struct {
		path     string
		expected bool
		err      error
	}{
		{tmpDir, true, nil}, // 目录存在
		{tmpFilePath, false, filehelper.ErrPathIsPointToFile}, // 路径指向文件
		{"nonexistentdir", false, nil},                        // 目录不存在
	}

	for _, tt := range tests {
		result, err := filehelper.DirExist(tt.path)
		if result != tt.expected || (err != nil && err.Error() != tt.err.Error()) {
			t.Errorf("DirExist(%v) = (%v, %v); want (%v, %v)", tt.path, result, err, tt.expected, tt.err)
		}
	}
}
