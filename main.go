package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/manuelarte/goslicespackagecheck/analyzer"
)

func main() {
	singlechecker.Main(analyzer.NewAnalyzer())
}
