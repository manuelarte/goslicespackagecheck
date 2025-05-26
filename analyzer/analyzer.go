package analyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	"github.com/manuelarte/goslicespackagecheck/internal/slicecheckers/equalchecker"
)

const (
	EqualCheckName = "equal"
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

	a.Flags.BoolVar(&cfg.equal, EqualCheckName, true,
		"Checks that constructors are placed after the structure declaration.")

	return a
}

type goslicespackagecheck struct {
	cfg *config
}

func (g *goslicespackagecheck) run(pass *analysis.Pass) (any, error) {
	insp, found := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !found {
		//nolint:nilnil // impossible case.
		return nil, nil
	}

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}

	insp.Preorder(nodeFilter, func(n ast.Node) {
		//nolint:gocritic // more cases to be added
		switch node := n.(type) {
		case *ast.FuncDecl:
			ec := equalchecker.EqualChecker{}
			if diag, ok := ec.AppliesTo(node); ok {
				pass.Report(diag)
			}
		}
	})

	//nolint:nilnil //any, error
	return nil, nil
}

type config struct {
	equal bool
}
