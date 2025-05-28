package maxchecker

import (
	"go/ast"
	"go/token"
)

// ifMaxChecker Holder to check if it's comparing max using inside a i, value := range a for loop.
type ifForMaxChecker struct {
	arrayIdent *ast.Ident
	iIdent     *ast.Ident
	ifStmn     *ast.IfStmt
}

func (imc *ifForMaxChecker) apply() bool {
	maxValueIdent, ok := imc.checkMaxValueCond()
	if !ok {
		return false
	}

	return imc.checkAssignmentStmn(maxValueIdent)
}

func (imc *ifForMaxChecker) checkMaxValueCond() (*ast.Ident, bool) {
	binaryExpr, ok := imc.ifStmn.Cond.(*ast.BinaryExpr)
	if !ok {
		return nil, false
	}

	var maxValueIdent *ast.Ident
	var indexExpr *ast.IndexExpr

	xIdent, isXIdent := binaryExpr.X.(*ast.Ident)
	yIdent, isYIdent := binaryExpr.Y.(*ast.Ident)
	xIndexExpr, isXIndexExpr := binaryExpr.X.(*ast.IndexExpr)
	yIndexExpr, isYIndexExpr := binaryExpr.Y.(*ast.IndexExpr)
	if isXIdent && isYIndexExpr {
		maxValueIdent = xIdent
		indexExpr = yIndexExpr
		if binaryExpr.Op != token.LEQ && binaryExpr.Op != token.LSS {
			return nil, false
		}
	} else if isXIndexExpr && isYIdent {
		indexExpr = xIndexExpr
		maxValueIdent = yIdent
		if binaryExpr.Op != token.GEQ && binaryExpr.Op != token.GTR {
			return nil, false
		}
	} else {
		return nil, false
	}

	if indexExpr.Index.(*ast.Ident).Name != imc.iIdent.Name {
		return nil, false
	}

	return maxValueIdent, true
}

func (imc *ifForMaxChecker) checkAssignmentStmn(maxValueIdent *ast.Ident) bool {
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
		rhsIdent, isIndexExpr := assignStmn.Rhs[0].(*ast.IndexExpr)
		if !isIndexExpr {
			return false
		}
		rhsIndexIdent, isIdent := rhsIdent.Index.(*ast.Ident)
		if !isIdent {
			return false
		}
		if rhsIndexIdent.Name != imc.iIdent.Name {
			return false
		}
	}
	return true
}
