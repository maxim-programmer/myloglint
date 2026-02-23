package analyzer

import (
	"go/ast"
	"go/token"
	"strings"
	"unicode"

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

			if len(call.Args) > 0 {
				if lit, ok := call.Args[0].(*ast.BasicLit); ok && lit.Kind == token.STRING {
					s := strings.Trim(lit.Value, `"`)
					if len(s) > 0 {
						if unicode.IsUpper(rune(s[0])) {
							pass.Reportf(call.Pos(), "log messages must begin with a lowercase letter")
						}

						if !isAllEnglishLetters(s) {
							pass.Reportf(call.Pos(), "log messages must be in english only")
						}

						if hasSpecialChar(s) {
							pass.Reportf(call.Pos(), "log messages must not contain special characters or emojis")
						}
					}
				}
			}

			return true
		})
	}

	return nil, nil
}

func isAllEnglishLetters(s string) bool {
	s = strings.ToLower(s)
	for _, r := range s {
		if !unicode.IsLetter(r) {
			continue
		}
		if r < 'a' || r > 'z' {
			return false
		}
	}

	return true
}

func hasSpecialChar(s string) bool {
	for _, r := range s {
		if r == ' ' {
			continue
		}
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return true
		}
	}
	return false
}
