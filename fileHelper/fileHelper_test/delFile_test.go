package filehelper_test

import (
	"os"
	"testing"

	filehelper "github.com/jbterrylin/go-helper/fileHelper"
)

func TestDelFile(t *testing.T) {
	// 创建临时文件进行测试
	tmpFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	tmpFilePath := tmpFile.Name()
	tmpFile.Close()

	// 检查临时文件是否存在
	if _, err := os.Stat(tmpFilePath); os.IsNotExist(err) {
		t.Fatalf("Temporary file does not exist: %v", tmpFilePath)
	}

	// 测试删除临时文件
	err = filehelper.DelFile(tmpFilePath)
	if err != nil {
		t.Fatalf("DelFile(%v) returned error: %v", tmpFilePath, err)
	}

	// 检查临时文件是否被删除
	if _, err := os.Stat(tmpFilePath); !os.IsNotExist(err) {
		t.Errorf("Expected file %v to be deleted, but it still exists", tmpFilePath)
	}

	// 创建临时目录进行测试
	tmpDir, err := os.MkdirTemp("", "testdir")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}

	// 检查临时目录是否存在
	if _, err := os.Stat(tmpDir); os.IsNotExist(err) {
		t.Fatalf("Temporary directory does not exist: %v", tmpDir)
	}

	// 测试删除临时目录
	err = filehelper.DelFile(tmpDir)
	if err != nil {
		t.Fatalf("DelFile(%v) returned error: %v", tmpDir, err)
	}

	// 检查临时目录是否被删除
	if _, err := os.Stat(tmpDir); !os.IsNotExist(err) {
		t.Errorf("Expected directory %v to be deleted, but it still exists", tmpDir)
	}
}
