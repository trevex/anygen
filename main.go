package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

func main() {
	b, err := os.ReadFile(os.Args[2])
	if err != nil {
		panic(err)
	}
	src := string(b)

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	for _, decl := range f.Decls {
		switch d := decl.(type) {
		case *ast.FuncDecl:
			if d.Doc != nil {
				for _, comment := range d.Doc.List {
					if strings.HasPrefix(comment.Text, "//anygen:") {
						ast.Print(fset, d)
					}
				}
			}
		}
	}

	// ast.Print(fset, f)
}
