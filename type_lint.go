package depractedlinter

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

func lintType(gen *ast.GenDecl, pass analysis.Pass, currentVersion string) {
	// 首先获取struct上的deprecated注释
	doc := gen.Doc
	if doc != nil {
		for _, comment := range doc.List {
			commonLint(gen.Pos(), pass, gen.Tok.String()+" "+gen.Specs[0].(*ast.TypeSpec).Name.Name, comment.Text, currentVersion)
		}
	}
	for _, spec := range gen.Specs {
		typeSpec := spec.(*ast.TypeSpec)
		if s, ok := typeSpec.Type.(*ast.StructType); ok {
			lintStruct(typeSpec, s, pass, currentVersion)
		}
		if i, ok := typeSpec.Type.(*ast.InterfaceType); ok {
			lintInterface(typeSpec, i, pass, currentVersion)
		}
	}
}
