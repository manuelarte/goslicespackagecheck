package maxchecker

import "go/ast"

// ifRange Holder to check if it's comparing max using inside a i, value := range a for loop
type ifRange struct {
	valueIdent *ast.Ident
}
