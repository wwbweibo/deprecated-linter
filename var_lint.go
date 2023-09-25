package depractedlinter

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

func lintVar(gen *ast.GenDecl, pass analysis.Pass, currentVersion string) {
	if !gen.Lparen.IsValid() {
		doc := gen.Doc
		if doc != nil {
			commonLint(gen.Pos(), pass, gen.Tok.String()+" "+gen.Specs[0].(*ast.ValueSpec).Names[0].Name, doc.List[0].Text, currentVersion)
		}
		return
	}
	for _, v := range gen.Specs {
		vv := v.(*ast.ValueSpec)
		for _, c := range vv.Doc.List {
			commonLint(vv.Pos(), pass, vv.Names[0].Name, c.Text, currentVersion)
		}
	}
}
