package analyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	"github.com/manuelarte/goslicespackagecheck/internal/slicecheckers/concatchecker"
	"github.com/manuelarte/goslicespackagecheck/internal/slicecheckers/equalchecker"
	"github.com/manuelarte/goslicespackagecheck/internal/slicecheckers/maxchecker"
)

const (
	ConcatCheckName = "concat"
	EqualCheckName  = "equal"
	MaxCheckName    = "max"
)

func NewAnalyzer() *analysis.Analyzer {
	cfg := config{}
	f := goslicespackagecheck{cfg: &cfg}

	a := &analysis.Analyzer{
		Name:     "goslicespackagecheck",
		Doc:      "linter that checks when the new slices package can be used.",
		URL:      "https://github.com/manuelarte/goslicespackagecheck",
		Run:      f.run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}

	a.Flags.BoolVar(&cfg.concat, ConcatCheckName, true,
		"Check whether some loops can be replaced by slices.Concat")

	a.Flags.BoolVar(&cfg.equal, EqualCheckName, true,
		"Check whether some functions can be replaced by slices.Equal")

	a.Flags.BoolVar(&cfg.max, MaxCheckName, true,
		"Check whether some loops can be replaced by slices.Max")

	return a
}

type goslicespackagecheck struct {
	cfg *config
}

//nolint:gocognit // if for config
func (g *goslicespackagecheck) run(pass *analysis.Pass) (any, error) {
	insp, found := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !found {
		//nolint:nilnil // impossible case.
		return nil, nil
	}

	nodeFilter := []ast.Node{
		(*ast.ForStmt)(nil),
		(*ast.FuncDecl)(nil),
		(*ast.RangeStmt)(nil),
	}

	insp.Preorder(nodeFilter, func(n ast.Node) {
		switch node := n.(type) {
		case *ast.FuncDecl:
			if g.cfg.equal {
				ec := equalchecker.EqualChecker{}
				if diag, ok := ec.AppliesTo(node); ok {
					pass.Report(diag)
				}
			}

		case *ast.RangeStmt:
			if g.cfg.max {
				mc := maxchecker.MaxRangeChecker{}
				if diag, ok := mc.AppliesTo(node); ok {
					pass.Report(diag)
				}
			}
			if g.cfg.concat {
				mc := concatchecker.ConcatRangeChecker{}
				if diag, ok := mc.AppliesTo(node); ok {
					pass.Report(diag)
				}
			}
		case *ast.ForStmt:
			if g.cfg.max {
				mc := maxchecker.MaxForChecker{}
				if diag, ok := mc.AppliesTo(node); ok {
					pass.Report(diag)
				}
			}
		}
	})

	//nolint:nilnil //any, error
	return nil, nil
}

type config struct {
	concat bool
	equal  bool
	max    bool
}
