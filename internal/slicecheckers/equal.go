package slicecheckers

import (
	"fmt"
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

var _ SliceChecker[*ast.FuncDecl] = new(EqualChecker)

// EqualChecker checks whether the *ast.FuncDecl can be replaced with slices.Equal.
type EqualChecker struct{}

func (c EqualChecker) AppliesTo(fn *ast.FuncDecl) (analysis.Diagnostic, bool) {
	if !isBoolReturned(fn.Type.Results) {
		return analysis.Diagnostic{}, false
	}

	if !areParametersArraySameType(fn.Type.Params) {
		return analysis.Diagnostic{}, false
	}

	bodyNodes := fn.Body.List

	var ifStatement *ast.IfStmt
	var rangeStatement *ast.RangeStmt
	for _, n := range bodyNodes {
		// TODO(manuelarte): steps to do:
		// check that somehow the len is compared, and if they are not equal, it returns false
		// this open the possibility to compare for !slices.Equal
		// check there is a for loop, and both arrays are comparing objects for that index
		switch casted := n.(type) {
		case *ast.IfStmt:
			if ifStatement != nil {
				return analysis.Diagnostic{}, false
			}
			ifStatement = casted
		case *ast.RangeStmt:
			if rangeStatement != nil {
				return analysis.Diagnostic{}, false
			}
			rangeStatement = casted
		case *ast.ReturnStmt:
			// do nothing
		default:
			return analysis.Diagnostic{}, false
		}
	}
	fmt.Printf("IfStatement: %+v\n", ifStatement)
	fmt.Printf("RangeStmt: %+v\n", rangeStatement)

	if ifStatement == nil || rangeStatement == nil {
		return analysis.Diagnostic{}, false
	}

	return analysis.Diagnostic{
		Pos:     fn.Pos(),
		Message: fmt.Sprintf("the function %s can be replaced by slices.Equal", fn.Name.Name),
		URL:     "", //TODO(manuelarte): add readme and then put link here
	}, true
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
	if p == nil {
		return false
	}
	if len(p.List) != 2 && (len(p.List) != 1 || len(p.List[0].Names) != 2) {
		return false
	}

	var a, b sliceField
	var okA bool
	var okB bool
	if len(p.List) == 2 {
		a, okA = newSliceField(p.List[0], 0)
		b, okB = newSliceField(p.List[1], 0)
		if !okA || !okB {
			return false
		}
	} else {
		a, okA = newSliceField(p.List[0], 0)
		b, okB = newSliceField(p.List[0], 1)
		if !okA || !okB {
			return false
		}
	}

	if a.arrType.Elt != b.arrType.Elt {
		return false
	}

	return true
}

type sliceField struct {
	*ast.Field
	arrType   *ast.ArrayType
	nameIndex int
}

func newSliceField(field *ast.Field, nameIndex int) (sliceField, bool) {
	casted, ok := field.Type.(*ast.ArrayType)
	if !ok {
		return sliceField{}, false
	}
	return sliceField{Field: field, arrType: casted, nameIndex: nameIndex}, true
}
