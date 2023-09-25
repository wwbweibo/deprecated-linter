package depractedlinter

import (
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/analysis"
	"os"
	"os/exec"
	"regexp"
)

var Analyzer = &analysis.Analyzer{
	Name: "deprecatedlinter",
	Doc:  "check if the code contains deprecated comments is outdated",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	currentVersion := getCurrentVersion()
	inspect := func(node ast.Node) bool {
		file, ok := node.(*ast.File)
		if !ok {
			return true
		}
		for _, decl := range file.Decls {
			if function, ok := decl.(*ast.FuncDecl); ok {
				lintFunction(function, *pass, currentVersion)
			}
			if gen, ok := decl.(*ast.GenDecl); ok {
				if gen.Tok == token.TYPE {
					lintType(gen, *pass, currentVersion)
				}
				if gen.Tok == token.VAR || gen.Tok == token.CONST {
					lintVar(gen, *pass, currentVersion)
				}
			}
		}
		return true
	}

	for _, f := range pass.Files {
		ast.Inspect(f, inspect)
	}
	return nil, nil
}

func getCurrentVersion() string {
	version := os.Getenv("VERSION")
	if version != "" {
		return version
	}
	// 尝试从gitlab ci变量中获取
	version = os.Getenv("CI_COMMIT_TAG")
	if version != "" {
		exp, _ := regexp.Compile("^(test)|(dev)|(release)|(prerelease)-")
		return exp.ReplaceAllString(version, "")
	}
	version = os.Getenv("CI_COMMIT_BRANCH")
	if version != "" {
		exp, _ := regexp.Compile("^(test)|(dev)|(release)|(prerelease)/")
		return exp.ReplaceAllString(version, "")
	}
	// 尝试调用git命令行获取
	cmd := exec.Command("git", "status")
	out, _ := cmd.Output()
	exp, _ := regexp.Compile("v[\\d]+.[\\d]+.[\\d]+")
	return exp.FindString(string(out))
}
