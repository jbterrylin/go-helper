package filehelper

import "os"

func FileExist(path string) (bool, error) {
	fi, err := os.Lstat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	if fi.IsDir() {
		return false, ErrPathIsPointToDir
	}
	return true, nil
}
