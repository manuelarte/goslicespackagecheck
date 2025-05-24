package analyzer

import (
	"go/ast"

	"github.com/manuelarte/goslicespackagecheck/internal"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

func NewAnalyzer() *analysis.Analyzer {
	f := goslicespackagecheck{}

	a := &analysis.Analyzer{
		Name:     "goslicespackagecheck",
		Doc:      "linter that checks when the new slices package can be used.",
		URL:      "https://github.com/manuelarte/goslicespackagecheck",
		Run:      f.run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}

	return a
}

type goslicespackagecheck struct{}

func (g *goslicespackagecheck) run(pass *analysis.Pass) (any, error) {
	insp, found := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !found {
		//nolint:nilnil // impossible case.
		return nil, nil
	}

	fp := internal.NewFileProcessor()

	nodeFilter := []ast.Node{
		(*ast.File)(nil),
		(*ast.FuncDecl)(nil),
		(*ast.TypeSpec)(nil),
	}

	insp.Preorder(nodeFilter, func(n ast.Node) {
		switch node := n.(type) {
		case *ast.File:

		case *ast.FuncDecl:

		case *ast.TypeSpec:

		}
	})

	//nolint:nilnil //any, error
	return nil, nil
}
