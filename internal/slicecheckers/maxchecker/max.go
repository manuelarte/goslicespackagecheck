package maxchecker

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"

	"github.com/manuelarte/goslicespackagecheck/internal/slicecheckers"
)

var (
	_ slicecheckers.SliceChecker[*ast.RangeStmt] = new(MaxRangeChecker)
	_ slicecheckers.SliceChecker[*ast.ForStmt]   = new(MaxForChecker)
)

type MaxRangeChecker struct{}

func (m *MaxRangeChecker) AppliesTo(r *ast.RangeStmt) (analysis.Diagnostic, bool) {
	if len(r.Body.List) != 1 {
		return analysis.Diagnostic{}, false
	}

	ifStmn, ok := r.Body.List[0].(*ast.IfStmt)
	if !ok {
		return analysis.Diagnostic{}, false
	}

	_, isArrayIndexValueIdent := r.X.(*ast.Ident)
	if !isArrayIndexValueIdent {
		return analysis.Diagnostic{}, false
	}

	rangeKeyIdent, isKeyIdent := r.Key.(*ast.Ident)
	if !isKeyIdent {
		return analysis.Diagnostic{}, false
	}

	rangeValueIdent, isValueIdent := r.Value.(*ast.Ident)
	if !isValueIdent {
		return analysis.Diagnostic{}, false
	}

	irmc := ifRangeMaxChecker{
		rangeKeyIdent:   rangeKeyIdent,
		rangeValueIdent: rangeValueIdent,
		ifStmn:          ifStmn,
	}

	if !irmc.apply() {
		return analysis.Diagnostic{}, false
	}

	return analysis.Diagnostic{
		Pos:     r.Pos(),
		Message: "this for loop can be replaced by slices.Max",
		URL:     "https://github.com/manuelarte/goslicespackagecheck/tree/main?tab=readme-ov-file#slicesmax",
	}, true
}

type MaxForChecker struct{}

func (m *MaxForChecker) AppliesTo(f *ast.ForStmt) (analysis.Diagnostic, bool) {
	if len(f.Body.List) != 1 {
		return analysis.Diagnostic{}, false
	}

	ifStmn, ok := f.Body.List[0].(*ast.IfStmt)
	if !ok {
		return analysis.Diagnostic{}, false
	}

	iIdent, isInitOk := m.checkForInit(f)
	if !isInitOk {
		return analysis.Diagnostic{}, false
	}

	arrayIdent, isCondOk := m.checkForCond(f)
	if !isCondOk {
		return analysis.Diagnostic{}, false
	}

	// check that ifStmn is checking for value >=
	iffrc := ifForMaxChecker{iIdent: iIdent, arrayIdent: arrayIdent, ifStmn: ifStmn}
	if !iffrc.apply() {
		return analysis.Diagnostic{}, false
	}

	return analysis.Diagnostic{
		Pos:     f.Pos(),
		Message: "this for loop can be replaced by slices.Max",
		URL:     "", // TODO(manuelarte): add readme and then put link here
	}, true
}

func (m *MaxForChecker) checkForInit(f *ast.ForStmt) (*ast.Ident, bool) {
	iAssignStmn, isIAssignStmn := f.Init.(*ast.AssignStmt)
	if !isIAssignStmn {
		return nil, false
	}

	if len(iAssignStmn.Lhs) != 1 {
		return nil, false
	}

	lhsIdent, isLHSIdent := iAssignStmn.Lhs[0].(*ast.Ident)
	if !isLHSIdent {
		return nil, false
	}
	rhsBasicLit, isRHSBasicLit := iAssignStmn.Rhs[0].(*ast.BasicLit)
	if !isRHSBasicLit {
		return nil, false
	}
	if rhsBasicLit.Value != "0" {
		return nil, false
	}

	return lhsIdent, true
}

// checking i < len(a) -> returns a, true.
func (m *MaxForChecker) checkForCond(f *ast.ForStmt) (*ast.Ident, bool) {
	condExpr, isBinaryExpr := f.Cond.(*ast.BinaryExpr)
	if !isBinaryExpr {
		return nil, false
	}
	if condExpr.Op != token.LSS {
		return nil, false
	}
	// check x is the same as init
	yCallExpr, isCallExpr := condExpr.Y.(*ast.CallExpr)
	if !isCallExpr {
		return nil, false
	}
	lenFun, isIdentFun := yCallExpr.Fun.(*ast.Ident)
	if !isIdentFun {
		return nil, false
	}
	if lenFun.Name != "len" || len(yCallExpr.Args) != 1 {
		return nil, false
	}

	arrayIdent, isArrayIdent := yCallExpr.Args[0].(*ast.Ident)
	if !isArrayIdent {
		return nil, false
	}

	return arrayIdent, true
}
