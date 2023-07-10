package main

import (
	"github.com/Sigumaa/stringdetector"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(stringdetector.Analyzer) }
