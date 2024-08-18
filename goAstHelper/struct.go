package goasthelper

import "go/ast"

// path and file path may different if user give folder instead of file path
type AstFile struct {
	Path     string // from user
	FilePath string // path for file
	AstNode  *ast.File
}
