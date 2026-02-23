package analyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var LogLintAnalyzer = &analysis.Analyzer{
	Name: "loglint",
	Doc:  "check log messages",
	Run:  run,
}

var loggers = map[string]bool{
	"log":             true,
	"log/slog":        true,
	"go.uber.org/zap": true,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			sel, ok := call.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}

			obj := pass.TypesInfo.ObjectOf(sel.Sel)
			if obj == nil || obj.Pkg() == nil || !loggers[obj.Pkg().Path()] {
				return true
			}

			pass.Reportf(call.Pos(), "logging found")

			return true
		})
	}

	return nil, nil
}
