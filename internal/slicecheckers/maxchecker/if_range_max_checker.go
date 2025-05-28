package maxchecker

import (
	"go/ast"
	"go/token"
)

// ifMaxChecker Holder to check if it's comparing max using inside a i, value := range a for loop.
type ifRangeMaxChecker struct {
	rangeKeyIdent   *ast.Ident
	rangeValueIdent *ast.Ident
	ifStmn          *ast.IfStmt
}

func (imc *ifRangeMaxChecker) apply() bool {
	maxValueIdent, ok := imc.checkMaxValueCond()
	if !ok {
		return false
	}

	return imc.checkAssignmentStmn(maxValueIdent)
}

func (imc *ifRangeMaxChecker) checkMaxValueCond() (*ast.Ident, bool) {
	binaryExpr, ok := imc.ifStmn.Cond.(*ast.BinaryExpr)
	if !ok {
		return nil, false
	}

	xIdent, isXIdent := binaryExpr.X.(*ast.Ident)
	if !isXIdent {
		return nil, false
	}

	yIdent, isYIdent := binaryExpr.Y.(*ast.Ident)
	if !isYIdent {
		return nil, false
	}

	var maxValueIdent *ast.Ident

	switch {
	case xIdent.Name == imc.rangeValueIdent.Name:
		if binaryExpr.Op != token.GEQ && binaryExpr.Op != token.GTR {
			return nil, false
		}
		maxValueIdent = yIdent
	case yIdent.Name == imc.rangeValueIdent.Name:
		if binaryExpr.Op != token.LEQ && binaryExpr.Op != token.LSS {
			return nil, false
		}
		maxValueIdent = xIdent
	default:
		return nil, false
	}
	return maxValueIdent, true
}

func (imc *ifRangeMaxChecker) checkAssignmentStmn(maxValueIdent *ast.Ident) bool {
	assignStmn, isAssignStmn := imc.ifStmn.Body.List[0].(*ast.AssignStmt)
	if !isAssignStmn {
		return false
	}

	if len(assignStmn.Lhs) == 1 {
		lhsIdent, isIdent := assignStmn.Lhs[0].(*ast.Ident)
		if !isIdent {
			return false
		}
		if lhsIdent.Name != maxValueIdent.Name {
			return false
		}
	}
	if len(assignStmn.Rhs) == 1 {
		rhsIdent, isIdent := assignStmn.Rhs[0].(*ast.Ident)
		if !isIdent {
			return false
		}
		if rhsIdent.Name != imc.rangeValueIdent.Name {
			return false
		}
	}
	return true
}
