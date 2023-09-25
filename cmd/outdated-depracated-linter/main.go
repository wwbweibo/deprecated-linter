package main

import (
	linter "github.com/wwbweibo/deprecatedlinter"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(linter.Analyzer)
}
