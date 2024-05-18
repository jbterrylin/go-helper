package filehelper

import (
	"os"
)

func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := DirExist(v)
		if err != nil {
			return err
		}
		if !exist {
			if err := os.MkdirAll(v, os.ModePerm); err != nil {
				return err
			}
		}
	}
	return err
}
