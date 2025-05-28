//nolint:cyclop // refactor later
package concatchecker

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"

	"github.com/manuelarte/goslicespackagecheck/internal/slicecheckers"
)

var _ slicecheckers.SliceChecker[*ast.RangeStmt] = new(ConcatRangeChecker)

// ConcatRangeChecker checks whether the RangeStmt can be replaced with slices.Concat.
type ConcatRangeChecker struct {
}

func (c *ConcatRangeChecker) AppliesTo(r *ast.RangeStmt) (analysis.Diagnostic, bool) {
	if len(r.Body.List) != 1 {
		return analysis.Diagnostic{}, false
	}

	xIdent, isArrayIndexValueIdent := r.X.(*ast.Ident)
	if !isArrayIndexValueIdent {
		return analysis.Diagnostic{}, false
	}

	// check xIdent is an array
	if xIdent.Obj.Decl == nil {
		return analysis.Diagnostic{}, false
	}
	xAssignStmn, isAssignStmn := xIdent.Obj.Decl.(*ast.AssignStmt)
	if !isAssignStmn {
		return analysis.Diagnostic{}, false
	}
	if len(xAssignStmn.Rhs) != 1 {
		return analysis.Diagnostic{}, false
	}
	compositeLit, isRHSIdent := xAssignStmn.Rhs[0].(*ast.CompositeLit)
	if !isRHSIdent {
		return analysis.Diagnostic{}, false
	}
	if compositeLit.Type == nil {
		return analysis.Diagnostic{}, false
	}

	_, isArrayType := compositeLit.Type.(*ast.ArrayType)
	if !isArrayType {
		return analysis.Diagnostic{}, false
	}

	rangeValueIdent, isValueIdent := r.Value.(*ast.Ident)
	if !isValueIdent {
		return analysis.Diagnostic{}, false
	}

	assignStmn, isAssignStmn := r.Body.List[0].(*ast.AssignStmt)
	if !isAssignStmn {
		return analysis.Diagnostic{}, false
	}

	if len(assignStmn.Lhs) != 1 {
		return analysis.Diagnostic{}, false
	}
	if len(assignStmn.Rhs) != 1 {
		return analysis.Diagnostic{}, false
	}
	arrayIdent, isLHSIdent := assignStmn.Lhs[0].(*ast.Ident)
	appendCallExpr, isRHSCallExpr := assignStmn.Rhs[0].(*ast.CallExpr)
	if !isLHSIdent || !isRHSCallExpr {
		return analysis.Diagnostic{}, false
	}
	appendFuncIdent, isAppendFuncIdent := appendCallExpr.Fun.(*ast.Ident)
	if !isAppendFuncIdent {
		return analysis.Diagnostic{}, false
	}

	if appendFuncIdent.Name != "append" {
		return analysis.Diagnostic{}, false
	}

	//nolint:mnd // append has two parameters
	if len(appendCallExpr.Args) != 2 {
		return analysis.Diagnostic{}, false
	}
	arg0, isArg0Ident := appendCallExpr.Args[0].(*ast.Ident)
	if !isArg0Ident {
		return analysis.Diagnostic{}, false
	}
	if arg0.Name != arrayIdent.Name {
		return analysis.Diagnostic{}, false
	}

	arg1, isArg1Ident := appendCallExpr.Args[1].(*ast.Ident)
	if !isArg1Ident {
		return analysis.Diagnostic{}, false
	}
	if arg1.Name != rangeValueIdent.Name {
		return analysis.Diagnostic{}, false
	}

	return analysis.Diagnostic{
		Pos:     r.Pos(),
		Message: "this for loop can be replaced by slices.Concat",
		URL:     "https://github.com/manuelarte/goslicespackagecheck/tree/main?tab=readme-ov-file#slicesconcat",
	}, true
}
