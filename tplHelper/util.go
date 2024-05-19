package tplhelper

import (
	"os"
	"strings"
)

func getAllTplFileByPath(pathName string, fileList []string) ([]string, error) {
	files, err := os.ReadDir(pathName)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if file.IsDir() {
			fileList, err = getAllTplFileByPath(pathName+string(os.PathSeparator)+file.Name(), fileList)
			if err != nil {
				return nil, err
			}
		} else {
			if strings.HasSuffix(file.Name(), ".tpl") {
				fileList = append(fileList, pathName+string(os.PathSeparator)+file.Name())
			}
		}
	}
	return fileList, err
}
