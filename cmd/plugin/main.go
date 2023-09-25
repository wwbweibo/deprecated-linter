package main

import (
	linter "deprecatedlinter"
	a "golang.org/x/tools/go/analysis"
)

//func main() {
//	singlechecker.Main(linter.Analyzer)
//}
//
//type analyzerPlugin struct{}
//
//func (*analyzerPlugin) GetAnalyzers() []*a.Analyzer {
//	return []*a.Analyzer{
//		linter.Analyzer,
//	}
//}
//
//var AnalyzerPlugin analyzerPlugin

func New(conf any) ([]*a.Analyzer, error) {
	return []*a.Analyzer{linter.Analyzer}, nil
}
