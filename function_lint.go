package depractedlinter

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

func lintFunction(function *ast.FuncDecl, pass analysis.Pass, currentVersion string) {
	if function.Doc == nil {
		return
	}
	for _, c := range function.Doc.List {
		commonLint(function.Pos(), pass, function.Name.Name, c.Text, currentVersion)
	}
}
