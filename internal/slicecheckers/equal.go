package slicecheckers

import (
	"go/ast"
)

var _ SliceChecker = new(EqualChecker)

// EqualChecker checks whether the *ast.FuncDecl can be replaced with slices.Equal.
type EqualChecker struct{}

func (c EqualChecker) Alternative() string {
	return "slices.Equal"
}

func (c EqualChecker) AppliesTo(fn *ast.FuncDecl) bool {
	// TODO(manuelarte): steps to do:
	// check both two input parameters and both are arrays of the same type
	// check output parameter is a bool
	// check that somehow the len is compared, and if they are not equal, it returns false
	// this open the possibility to compare for !slices.Equal
	// check there is a for loop, and both arrays are comparing objects for that index

	if !isBoolReturned(fn.Type.Results) {
		return false
	}

	if !areParametersArraySameType(fn.Type.Params) {
		return false
	}

	return true
}

func isBoolReturned(r *ast.FieldList) bool {
	if r == nil || len(r.List) != 1 {
		return false
	}
	// Must return bool
	if ident, ok := r.List[0].Type.(*ast.Ident); !ok || ident.Name != "bool" {
		return false
	}

	return true
}

func areParametersArraySameType(p *ast.FieldList) bool {
	if p == nil || len(p.List) != 2 {
		return false
	}

	a, b := p.List[0], p.List[1]

	// Both must be slices
	sliceA, okA := a.Type.(*ast.ArrayType)
	sliceB, okB := b.Type.(*ast.ArrayType)

	if !okA || !okB || (sliceA.Elt != sliceB.Elt) {
		return false
	}

	// check if they are

	return true
}
