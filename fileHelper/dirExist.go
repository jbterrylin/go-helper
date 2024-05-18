package filehelper

import "os"

func DirExist(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	if !fi.IsDir() {
		return false, ErrPathIsPointToFile
	}
	return true, nil
}
