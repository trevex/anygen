package anygen

import (
	"go/ast"
	"go/parser"
	"go/token"
)

type Anygen struct {
	fset *token.FileSet
	pkgs map[string]*ast.Package
	// pkgfs map[string][]*ast.File
}

func New(dir string) (*Anygen, error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, nil, parser.ParseComments)
	if err != nil {
		return nil, err // TODO: wrap
	}

	// // We need the map of files to directly pass it in for type-checks later
	// pkgfs := make(map[string][]*ast.File, len(pkgs))
	// for name, pkg := range pkgs {
	// 	fs := make([]*ast.File, 0, len(pkg.Files))
	// 	for _, f := range pkg.Files {
	// 		fs = append(fs, f)
	// 	}
	// 	pkgfs[name] = fs

	// }

	return &Anygen{
		fset,
		pkgs,
	}, nil
}

func (a *Anygen) AddTemplate(name, src string) error {
	return nil
}
