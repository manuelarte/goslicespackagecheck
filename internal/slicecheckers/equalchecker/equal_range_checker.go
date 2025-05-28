package equalchecker

import (
	"go/ast"
	"go/token"
)

type equalRangeChecker struct {
	aName string
	bName string
	*ast.RangeStmt
}

//nolint:gocognit,nestif // refactor later
func (c equalRangeChecker) applies() bool {
	if len(c.Body.List) == 1 {
		if ifStmn, ok := c.Body.List[0].(*ast.IfStmt); ok {
			if ifStmn.Body == nil || len(ifStmn.Body.List) != 1 {
				return false
			}
			if _, isReturnStmn := ifStmn.Body.List[0].(*ast.ReturnStmt); !isReturnStmn {
				return false
			}
			if binaryStmn, isBinaryStmn := ifStmn.Cond.(*ast.BinaryExpr); isBinaryStmn {
				if binaryStmn.Op == token.NEQ && binaryStmn.X != nil && binaryStmn.Y != nil {
					if xIndexExpr, isXIndexExpr := binaryStmn.X.(*ast.IndexExpr); isXIndexExpr {
						if yIndexExpr, isYIndexExpr := binaryStmn.Y.(*ast.IndexExpr); isYIndexExpr {
							return compareIndexExpr(xIndexExpr, yIndexExpr) && c.areXAndYParams(xIndexExpr, yIndexExpr)
						}
					}
				}
			}
		}
	}
	return false
}

func (c equalRangeChecker) areXAndYParams(x, y *ast.IndexExpr) bool {
	if xIdent, isXIdent := x.X.(*ast.Ident); isXIdent {
		if yIdent, isYIdent := y.X.(*ast.Ident); isYIdent {
			switch xIdent.Name {
			case c.aName:
				return yIdent.Name == c.bName
			case c.bName:
				return yIdent.Name == c.aName
			default:
				return false
			}
		}
	}
	return false
}

func compareIndexExpr(x, y *ast.IndexExpr) bool {
	if xIdent, isXIdent := x.Index.(*ast.Ident); isXIdent {
		if yIdent, isYIdent := y.Index.(*ast.Ident); isYIdent {
			return xIdent.Name == yIdent.Name
		}
	}
	return false
}
