package main

import (
	"fmt"
	"go/token"
	"log"
	"os"

	// "go/ast"
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
)

func main() {
	root := "../azure-sdk-for-go/services/preview/network/mgmt/2015-05-01-preview/network"
	// fset := token.NewFileSet()

	// if err := filepath.Walk(root, walk); err != nil {
	// 	log.Fatal(err)
	// }

	// pkgs, err := parser.ParseDir(fset, root, nil, 0)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, pkg := range pkgs {
	// 	for name := range pkg.Files {
	// 		fmt.Println(name)
	// 	}
	// }
	path := fmt.Sprintf("%s/models.go", root)
	fset := token.NewFileSet()

	f, err := decorator.ParseFile(fset, path, nil, 0)
	if err != nil {
		log.Fatal(err)
	}

	filterDecl(f)

	// structs := getTypeSpecs(fset, f)
	// _ = structs
	// for _, val := range structs {
	// 	fmt.Println(val.Name)
	// 	// os.Exit(1)
	// }

	// var buf bytes.Buffer
	// if err := format.Node(&buf, f, nil); err != nil {
	// 	log.Fatal(err)
	// }
	// decorator.Print(f)
	// if err := format.Node(&buf, fset, f); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(buf.String())
	return
}

func walk(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	// if !info.IsDir() {
	// 	if err := parseFile(path); err != nil {
	// 		return err
	// 	}
	// }

	return nil
}

func getTypeSpecs(fset *token.FileSet, f *dst.File) []*dst.TypeSpec {
	typeSpecs := make([]*dst.TypeSpec, 0)
	for _, val := range f.Decls {
		decl, ok := val.(*dst.GenDecl)
		if !ok {
			continue
		}

		if decl.Tok != token.TYPE {
			continue
		}

		typeSpec, ok := decl.Specs[0].(*dst.TypeSpec)
		if !ok {
			continue
		}

		typeSpecs = append(typeSpecs, typeSpec)
	}
	return typeSpecs
}

func filterDecl(f *dst.File) {
	n := 0
	for _, decl := range f.Decls {
		if keepDecl(decl) {
			f.Decls[n] = decl
			n++
		}
	}
	f.Decls = f.Decls[:n]
}

func keepDecl(decl dst.Decl) bool {
	if gen, ok := decl.(*dst.GenDecl); ok && gen.Tok == token.TYPE {
		t, ok := gen.Specs[0].(*dst.TypeSpec)
		if ok {
			fmt.Println(t.Name.Name)
		}
		_ = t
		return true
	}
	return false
}

// func keepType(name string) bool {
// 	if strings.HasSuffix(name, "Future") || strings.HasSuffix(name, "Result") || strings.HasSuffix(name, "Iterator") || strings.HasSuffix(name, "Page")
// }
