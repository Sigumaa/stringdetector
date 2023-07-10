package stringdetector

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var flag string

const doc = "stringdetector is a static analysis tool that detects the given string."

var Analyzer = &analysis.Analyzer{
	Name: "stringdetector",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func init() {
	Analyzer.Flags.StringVar(&flag, "flag", "", "string to detect")
}

func run(pass *analysis.Pass) (any, error) {
	if flag == "" {
		return nil, nil
	}

	ds := strings.Split(flag, ",")

	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.Ident)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.Ident:
			for _, d := range ds {
				if n.Name == d {
					pass.Reportf(n.Pos(), "%s is detected", d)
				}
			}
		}
	})

	return nil, nil
}
