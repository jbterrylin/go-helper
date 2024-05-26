package filehelper_test

import (
	"os"
	"testing"

	filehelper "github.com/jbterrylin/go-helper/fileHelper"
)

func TestCreateDir(t *testing.T) {
	// 创建临时目录进行测试
	testDirs := []string{"testdir1", "testdir2", "testdir1/subdir1"}
	defer func() {
		// 清理创建的目录
		for _, dir := range testDirs {
			os.RemoveAll(dir)
		}
	}()

	// 测试创建目录
	err := filehelper.CreateDir(testDirs...)
	if err != nil {
		t.Fatalf("CreateDir(%v) returned error: %v", testDirs, err)
	}

	// 检查目录是否存在
	for _, dir := range testDirs {
		exist, err := filehelper.DirExist(dir)
		if err != nil {
			t.Fatalf("DirExist(%v) returned error: %v", dir, err)
		}
		if !exist {
			t.Errorf("Expected directory %v to exist, but it does not", dir)
		}
	}

	// 测试重复创建目录
	err = filehelper.CreateDir(testDirs...)
	if err != nil {
		t.Fatalf("CreateDir(%v) returned error on second call: %v", testDirs, err)
	}
}
