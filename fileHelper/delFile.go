package filehelper

import "os"

func DelFile(filePath string) error {
	return os.RemoveAll(filePath)
}
