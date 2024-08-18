package goasthelper

import (
	"go/ast"
)

type GoAstHelper[Resp any] struct {
	Path   []string                                              // 2 type, first is point to file directly, for example a/b/c/d.go, second is point to folder, only need to show last folder, for example a/b/c
	Action func(path string, n ast.Node) (resp *Resp, next bool) // true mean skip
}

func GetNewGoAstHelper[Resp any](path []string, action func(path string, n ast.Node) (resp *Resp, next bool)) (goAstHelper GoAstHelper[Resp]) {
	goAstHelper = GoAstHelper[Resp]{
		Path:   path,
		Action: action,
	}
	return
}

func (a *GoAstHelper[Resp]) Run() (resps []Resp, err error) {
	astFiles, err := ParseGoFiles(a.Path)
	if err != nil {
		return
	}

	for _, astFile := range astFiles {
		ast.Inspect(astFile.AstNode, func(n ast.Node) bool {
			resp, next := a.Action(astFile.Path, n)
			if resp != nil {
				resps = append(resps, *resp)
			}
			return next
		})
	}

	return
}

func (a *GoAstHelper[Resp]) RunAndReturnByPath() (resps map[string][]Resp, err error) {
	resps = map[string][]Resp{}

	astFiles, err := ParseGoFiles(a.Path)
	if err != nil {
		return
	}

	for _, astFile := range astFiles {
		ast.Inspect(astFile.AstNode, func(n ast.Node) bool {
			resp, next := a.Action(astFile.Path, n)
			if resp != nil {
				resps[astFile.Path] = append(resps[astFile.Path], *resp)
			}
			return next
		})
	}

	return
}

func (a *GoAstHelper[Resp]) RunAndReturnByFilePath() (resps map[string]Resp, err error) {
	resps = map[string]Resp{}

	astFiles, err := ParseGoFiles(a.Path)
	if err != nil {
		return
	}

	for _, astFile := range astFiles {
		ast.Inspect(astFile.AstNode, func(n ast.Node) bool {
			resp, next := a.Action(astFile.Path, n)
			if resp != nil {
				resps[astFile.Path] = *resp
			}
			return next
		})
	}

	return
}
