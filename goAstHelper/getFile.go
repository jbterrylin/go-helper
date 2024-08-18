package goasthelper

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

func ParseGoFiles(paths []string) (astFiles []AstFile, err error) {
	for _, path := range paths {
		var files []string
		if !strings.HasSuffix(path, ".go") {
			// 获取所有匹配的 Go 文件路径
			files, err = filepath.Glob(path + "/*.go")
			if err != nil {
				return
			}
		} else {
			files = append(files, path)
		}

		// 遍历所有文件
		for _, file := range files {
			// 读取文件内容
			var content []byte
			content, err = os.ReadFile(file)
			if err != nil {
				return
			}

			// 解析文件的AST
			fset := token.NewFileSet()
			var node *ast.File
			node, err = parser.ParseFile(fset, file, content, parser.ParseComments)
			if err != nil {
				return
			}

			astFiles = append(astFiles, AstFile{
				Path:     path,
				FilePath: file,
				AstNode:  node,
			})
		}
	}

	return
}
