package maxchecker

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"

	"github.com/manuelarte/goslicespackagecheck/internal/slicecheckers"
)

var _ slicecheckers.SliceChecker[*ast.RangeStmt] = new(MaxChecker)

type MaxChecker struct {
}

func (m *MaxChecker) AppliesTo(r *ast.RangeStmt) (analysis.Diagnostic, bool) {
	if len(r.Body.List) != 1 {
		return analysis.Diagnostic{}, false
	}

	ifStmn, ok := r.Body.List[0].(*ast.IfStmt)
	if !ok {
		return analysis.Diagnostic{}, false
	}

	// TODO(manuelarte): check that ifStmn, check if the RangeStmt is using i, value := range a, or i := 0; i < len(a)
	// and act accordingly, for the 1st case, I need to get the Ident name (in the example value)
	// for the typical for loop with i:0, I need to check that they are checking a[i]
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

	bmc := ifMaxChecker{
		rangeKeyIdent:   rangeKeyIdent,
		rangeValueIdent: rangeValueIdent,
		ifStmn:          ifStmn,
	}

	if !bmc.apply() {
		return analysis.Diagnostic{}, false
	}
	return analysis.Diagnostic{
		Pos:     r.Pos(),
		Message: "this for loop can be replaced by slices.Max",
		URL:     "", // TODO(manuelarte): add readme and then put link here
	}, true
}
