package analyzer_test

import (
	"myloglint/analyzer"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()

	analysistest.Run(t, testdata, analyzer.LogLintAnalyzer, "test/valid", "test/invalid")
}
