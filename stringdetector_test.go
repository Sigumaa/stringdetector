package stringdetector_test

import (
	"testing"

	"github.com/Sigumaa/stringdetector"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	defer stringdetector.ExportSet("Url,Id,Utc")()
	analysistest.Run(t, testdata, stringdetector.Analyzer, "a")
}
