package tplhelper

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	filehelper "github.com/jbterrylin/go-helper/fileHelper"
	ophelper "github.com/jbterrylin/go-helper/opHelper"
)

type Tpl struct {
	TplDir              string
	TempDir             string
	ReplaceFileName     string
	fileNamePlaceHolder string
}

func NewTpl(tplDir, tempDir, replaceFileName string) Tpl {
	return Tpl{
		TplDir:              tplDir,
		TempDir:             ophelper.Or(tempDir, ".\\"+defaultTempDir),
		ReplaceFileName:     replaceFileName,
		fileNamePlaceHolder: fileNamePlaceHolder,
	}
}

func (tpl *Tpl) SetFileNamePlaceHolder(fileNamePlaceHolder string) {
	tpl.fileNamePlaceHolder = fileNamePlaceHolder
}

func (tpl *Tpl) PreviewTpl(data map[string]interface{}) (resultMap map[string]string, err error) {
	resultMap = make(map[string]string)
	dataList, _, needMkdir, err := tpl.getNeedList("")
	if err != nil {
		return
	}

	if err = filehelper.CreateDir(needMkdir...); err != nil {
		return
	}

	for _, value := range dataList {
		var f *os.File
		f, err = os.OpenFile(value.TempPath, os.O_CREATE|os.O_WRONLY, 0o755)
		if err != nil {
			return
		}
		if err = value.Template.Execute(f, data); err != nil {
			return
		}
		_ = f.Close()

		f, err = os.OpenFile(value.TempPath, os.O_CREATE|os.O_RDONLY, 0o755)
		if err != nil {
			return
		}

		builder := strings.Builder{}
		var data []byte
		data, err = io.ReadAll(f)
		if err != nil {
			return
		}
		builder.Write(data)

		resultMap[value.AutoMoveFilePath[1:]] = builder.String()
		_ = f.Close()
	}

	defer func() { // 移除中间文件
		if err := os.RemoveAll(tpl.TempDir); err != nil {
			return
		}
	}()

	return
}

// zipFilePath: path with file name e.x: "./code.zip"
func (tpl *Tpl) CreateTpl(data map[string]interface{}, autoMoveBasePath, zipFilePath string) (err error) {
	dataList, fileList, needMkdir, err := tpl.getNeedList(autoMoveBasePath)
	if err != nil {
		return err
	}

	if err = filehelper.CreateDir(needMkdir...); err != nil {
		return err
	}

	for _, value := range dataList {
		f, err := os.OpenFile(value.TempPath, os.O_CREATE|os.O_WRONLY, 0o755)
		if err != nil {
			return err
		}
		if err = value.Template.Execute(f, data); err != nil {
			return err
		}
		_ = f.Close()
	}

	defer func() { // 移除中间文件
		if err := os.RemoveAll(tpl.TempDir); err != nil {
			return
		}
	}()

	if zipFilePath != "" {
		if err = filehelper.ZipFiles(zipFilePath, fileList, ".", "."); err != nil {
			return err
		}
	}
	if autoMoveBasePath != "" {
		for _, value := range dataList {
			fileExist, err := filehelper.FileExist(value.AutoMoveFilePath)
			if err != nil {
				return err
			}
			if fileExist {
				return ErrFileExist
			}
		}
		for _, value := range dataList { // 移动文件
			if err := filehelper.MoveFile(value.TempPath, value.AutoMoveFilePath); err != nil {
				return err
			}
		}
	}
	return
}

func (tpl *Tpl) getNeedList(autoMoveBasePath string) (dataList []tplData, fileList []string, needMkdir []string, err error) {
	tplFileList, err := getAllTplFileByPath(tpl.TplDir, nil)
	if err != nil {
		return
	}

	dataList = make([]tplData, 0, len(tplFileList))
	fileList = make([]string, 0, len(tplFileList))
	needMkdir = make([]string, 0, len(tplFileList))

	for _, tplFilePath := range tplFileList {
		// Parse the template file
		var tplFile *template.Template
		tplFile, err = template.ParseFiles(tplFilePath)
		if err != nil {
			return
		}

		// Replace placeholders in the template file path and names
		replacedTplFilePath := strings.ReplaceAll(tplFilePath, tpl.fileNamePlaceHolder, tpl.ReplaceFileName)
		fileName := strings.TrimPrefix(tplFilePath, tpl.TplDir+string(os.PathSeparator))
		targetFileName := strings.TrimSuffix(fileName, ".tpl")
		replacedFileName := strings.ReplaceAll(targetFileName, tpl.fileNamePlaceHolder, tpl.ReplaceFileName)

		// Generate temporary path and auto move file path
		tempPath := filepath.Join(tpl.TempDir, replacedFileName)
		autoMoveFilePath := strings.ReplaceAll(strings.ReplaceAll(replacedTplFilePath, tpl.TplDir, autoMoveBasePath), ".tpl", "")

		// Collect directories that need to be created
		if lastSeparator := strings.LastIndex(tempPath, string(os.PathSeparator)); lastSeparator != -1 {
			needMkdir = append(needMkdir, tempPath[:lastSeparator])
		}

		// Add to file list and data list
		fileList = append(fileList, tempPath)
		dataList = append(dataList, tplData{
			LocationPath:     tplFilePath,
			Template:         tplFile,
			TempPath:         tempPath,
			AutoMoveFilePath: autoMoveFilePath,
		})
	}

	return
}

// TODO: inject code
