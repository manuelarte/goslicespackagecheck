package slicecheckers

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

type SliceChecker[N ast.Node] interface {
	AppliesTo(N) (analysis.Diagnostic, bool)
}
