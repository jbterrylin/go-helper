package filehelper

import (
	"os"
	"path/filepath"
)

// MoveFile 将文件从 src 移动到 dst
func MoveFile(src, dst string) error {
	// 如果目标路径为空，返回 nil
	if dst == "" {
		return nil
	}

	// 获取 src 的绝对路径
	srcAbs, err := filepath.Abs(src)
	if err != nil {
		return err
	}

	// 获取 dst 的绝对路径
	dstAbs, err := filepath.Abs(dst)
	if err != nil {
		return err
	}

	// 获取 dst 目录的路径
	dstDir := filepath.Dir(dstAbs)

	// 检查 dst 目录是否存在，不存在则创建
	if err := ensureDirExists(dstDir); err != nil {
		return err
	}

	// 移动文件
	return os.Rename(srcAbs, dstAbs)
}

// ensureDirExists 检查目录是否存在，不存在则创建
func ensureDirExists(dir string) error {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		// 创建目录及其所有父目录
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return err
		}
		return nil
	}
	return err
}
