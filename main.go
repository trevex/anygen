package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"os"
	"strings"
)

func main() {
	// PARSE
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, os.Args[2], nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	fs := []*ast.File{} // TODO pre-alloc
	for _, f := range pkgs["main"].Files {
		fs = append(fs, f)
	}

	// TYPECHECK
	info := types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
		Defs:  make(map[*ast.Ident]types.Object),
		Uses:  make(map[*ast.Ident]types.Object),
	}
	conf := types.Config{
		Importer: importer.Default(),
	}
	// pkg, err := conf.Check("", fset, fs, &info)
	_, err = conf.Check("", fset, fs, &info)
	if err != nil {
		log.Fatal(err)
	}

	// FIND ANYGEN COMMENTS
	for _, decl := range fs[0].Decls {
		switch d := decl.(type) {
		case *ast.FuncDecl:
			if d.Doc != nil {
				for _, comment := range d.Doc.List {
					if strings.HasPrefix(comment.Text, "//anygen:") {
						// ast.Print(fset, d)
						fmt.Println(comment.Text)
						fmt.Println(d.Name.Obj.Kind.String() + " " + d.Name.Name)
						def, ok := info.Defs[d.Name]
						if !ok {
							panic("expected def")
						}
						fdef, ok := def.(*types.Func)
						if !ok {
							panic("expected func")
						}
						fmt.Println(fdef.String())
						sig, ok := fdef.Type().(*types.Signature)
						if !ok {
							panic("expected signature")
						}
						params := sig.Params()
						for i := 0; i < params.Len(); i++ {
							param := params.At(i)
							fmt.Println(param.Name() + " " + param.Type().String())
							fmt.Printf("%v\n", param.Type())
							pt, ok := param.Type().(*types.Named)
							if ok {
								fmt.Println(pt.String())
								put, ok := pt.Underlying().(*types.Struct)
								if ok {
									for j := 0; j < put.NumFields(); j++ {
										fmt.Println("   " + put.Field(j).String())
									} // also put.Tag(j)!
								}
							}
						}
						fmt.Println(sig.Params().String())
					}
				}
			}
		}
	}

	// ast.Print(fset, f)

	// // Print package-level variables in initialization order.
	// fmt.Printf("InitOrder: %v\n\n", info.InitOrder)

	// // For each named object, print the line and
	// // column of its definition and each of its uses.
	// fmt.Println("Defs and Uses of each named object:")
	// usesByObj := make(map[types.Object][]string)
	// for id, obj := range info.Uses {
	// 	posn := fset.Position(id.Pos())
	// 	lineCol := fmt.Sprintf("%d:%d", posn.Line, posn.Column)
	// 	usesByObj[obj] = append(usesByObj[obj], lineCol)
	// }
	// var items []string
	// for obj, uses := range usesByObj {
	// 	sort.Strings(uses)
	// 	item := fmt.Sprintf("%s:\n  defined at %s\n  used at %s",
	// 		types.ObjectString(obj, types.RelativeTo(pkg)),
	// 		fset.Position(obj.Pos()),
	// 		strings.Join(uses, ", "))
	// 	items = append(items, item)
	// }
	// sort.Strings(items) // sort by line:col, in effect
	// fmt.Println(strings.Join(items, "\n"))
	// fmt.Println()

	// fmt.Println("Types and Values of each expression:")
	// items = nil
	// for expr, tv := range info.Types {
	// 	var buf strings.Builder
	// 	posn := fset.Position(expr.Pos())
	// 	tvstr := tv.Type.String()
	// 	if tv.Value != nil {
	// 		tvstr += " = " + tv.Value.String()
	// 	}
	// 	// line:col | expr | mode : type = value
	// 	fmt.Fprintf(&buf, "%1d:%2d | %-19s | %-7s : %s",
	// 		posn.Line, posn.Column, exprString(fset, expr),
	// 		mode(tv), tvstr)
	// 	items = append(items, buf.String())
	// }
	// sort.Strings(items)
	// fmt.Println(strings.Join(items, "\n"))
}

// func mode(tv types.TypeAndValue) string {
// 	switch {
// 	case tv.IsVoid():
// 		return "void"
// 	case tv.IsType():
// 		return "type"
// 	case tv.IsBuiltin():
// 		return "builtin"
// 	case tv.IsNil():
// 		return "nil"
// 	case tv.Assignable():
// 		if tv.Addressable() {
// 			return "var"
// 		}
// 		return "mapindex"
// 	case tv.IsValue():
// 		return "value"
// 	default:
// 		return "unknown"
// 	}
// }

// func exprString(fset *token.FileSet, expr ast.Expr) string {
// 	var buf strings.Builder
// 	format.Node(&buf, fset, expr)
// 	return buf.String()
// }
