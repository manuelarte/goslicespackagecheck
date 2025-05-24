package slicecheckers

import "go/ast"

type SliceChecker interface {
	Alternative() string
	AppliesTo(*ast.FuncDecl) bool
}
