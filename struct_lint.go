package depractedlinter

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

func lintStruct(t *ast.TypeSpec, s *ast.StructType, pass analysis.Pass, currentVersion string) {
	for _, f := range s.Fields.List {
		if f.Doc == nil {
			return
		}
		for _, c := range f.Doc.List {
			commonLint(f.Pos(), pass, t.Name.Name+"."+f.Names[0].Name, c.Text, currentVersion)
		}
	}
}
