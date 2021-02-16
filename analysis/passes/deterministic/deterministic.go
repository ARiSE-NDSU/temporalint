package deterministic

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "time",
	Doc:      "reports time related violations in temporal workflows",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	// A workflow body is in-fact a function declaration that has a specific first parameter of context.Workflow
	// Therfore, filter out only the nodes that are of the type 'FuncDec' from all the nodes
	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}

	// Preorder will parse all the nodes in the filter and execute function parameter on each node.
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		// Based on our filter n can be only of FuncDecl type
		functionDecl := n.(*ast.FuncDecl)

		// If we dont have sufficient arguments, then exit, we need to have at least 1 parameter which is workflow.Context
		if len(functionDecl.Type.Params.List) < 1 {
			return
		}

		currentExpr := functionDecl.Type.Params.List[0].Type.(*ast.SelectorExpr)
		currentClass := currentExpr.X.(*ast.Ident)

		// If the first argument isn't, by specification, what we expect to be a workflows first argument, then exit
		if !(currentClass.Name == "workflow" && currentExpr.Sel.Name == "Context") {
			return
		}

		// If everything else passed, perform an ast analysis on the function body
		ast.Inspect(functionDecl.Body, func(node ast.Node) bool {
			functionCall, isFunctionCall := node.(*ast.CallExpr)
			if isFunctionCall {
				selector, isSelector := functionCall.Fun.(*ast.SelectorExpr)
				if isSelector {
					identifier, isIdentifier := selector.X.(*ast.Ident)
					if isIdentifier {
						if identifier.Name == "time" && (selector.Sel.Name == "Now" || selector.Sel.Name == "Sleep") {
							pass.Reportf(functionCall.Fun.Pos(), "Deterministic constraint violation, please consider using temporal sdk functions for managing time")
						}
					}
				}
			}
			return false
		})
	})
	return nil, nil
}
