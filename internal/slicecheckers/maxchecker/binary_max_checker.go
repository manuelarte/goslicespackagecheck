package maxchecker

import (
	"go/ast"
	"go/token"
)

// ifMaxChecker Holder to check if it's comparing max using inside a i, value := range a for loop.
type ifMaxChecker struct {
	rangeKeyIdent   *ast.Ident
	rangeValueIdent *ast.Ident
	ifStmn          *ast.IfStmt
}

func (bmc *ifMaxChecker) apply() bool {
	maxValueIdent, ok := bmc.checkMaxValueCond()
	if !ok {
		return false
	}

	assignStmn, isAssignStmn := bmc.ifStmn.Body.List[0].(*ast.AssignStmt)
	if !isAssignStmn {
		return false
	}

	return bmc.checkAssignmentStmn(assignStmn, maxValueIdent)
}

func (bmc *ifMaxChecker) checkMaxValueCond() (*ast.Ident, bool) {
	binaryExpr, ok := bmc.ifStmn.Cond.(*ast.BinaryExpr)
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
	case xIdent.Name == bmc.rangeValueIdent.Name:
		if binaryExpr.Op != token.GEQ && binaryExpr.Op != token.GTR {
			return nil, false
		}
		maxValueIdent = yIdent
	case yIdent.Name == bmc.rangeValueIdent.Name:
		if binaryExpr.Op != token.LEQ && binaryExpr.Op != token.LSS {
			return nil, false
		}
		maxValueIdent = xIdent
	default:
		return nil, false
	}
	return maxValueIdent, true
}

func (bmc *ifMaxChecker) checkAssignmentStmn(assignStmn *ast.AssignStmt, maxValueIdent *ast.Ident) bool {
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
		if rhsIdent.Name != bmc.rangeValueIdent.Name {
			return false
		}
	}
	return true
}
