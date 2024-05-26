package filehelper_test

import (
	"os"
	"path/filepath"
	"testing"

	filehelper "github.com/jbterrylin/go-helper/fileHelper"
)

func TestMoveFile(t *testing.T) {
	// 创建临时目录用于目标文件测试
	dstDir, err := os.MkdirTemp("", "testdir")
	if err != nil {
		t.Fatalf("Failed to create temporary destination directory: %v", err)
	}
	defer os.RemoveAll(dstDir)

	// 目标文件路径
	dstFilePath := filepath.Join(dstDir, "movedfile")
	subDirFilePath := filepath.Join(dstDir, "subdir", "movedfile")

	tests := []struct {
		src       string
		dst       string
		expectErr bool
	}{
		{"testfile", dstFilePath, false},       // 正常移动文件
		{"testfile", "", false},                // 目标路径为空
		{"nonexistentfile", dstFilePath, true}, // 源文件不存在
		{"testfile", subDirFilePath, false},    // 目标路径目录不存在
	}

	for _, tt := range tests {
		// 为每个测试用例创建一个新的临时源文件
		srcFile, err := os.CreateTemp("", "testfile")
		if err != nil {
			t.Fatalf("Failed to create temporary source file: %v", err)
		}
		srcFilePath := srcFile.Name()
		srcFile.Close()

		// 如果源文件是"testfile"，替换为实际的临时文件路径
		if tt.src == "testfile" {
			tt.src = srcFilePath
		}

		err = filehelper.MoveFile(tt.src, tt.dst)
		if (err != nil) != tt.expectErr {
			t.Errorf("MoveFile(%v, %v) = %v; want error: %v", tt.src, tt.dst, err, tt.expectErr)
		}

		if tt.dst != "" && !tt.expectErr {
			// 检查目标文件是否存在
			if _, err := os.Stat(tt.dst); os.IsNotExist(err) {
				t.Errorf("Expected file %v to be moved, but it does not exist", tt.dst)
			}
		}

		// 检查源文件是否存在（除非目标路径为空或源文件不存在）
		if tt.dst != "" && tt.src != "nonexistentfile" {
			if _, err := os.Stat(tt.src); !os.IsNotExist(err) {
				t.Errorf("Expected source file %v to be moved, but it still exists", tt.src)
			}
		}
	}
}
