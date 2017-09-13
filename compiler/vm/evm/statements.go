package evm

import "github.com/end-r/guardian/compiler/ast"

func (e *EVMTraverser) traverseSwitchStatement(n ast.SwitchStatementNode) {

}

func (e *EVMTraverser) traverseCaseStatement(n ast.CaseStatementNode) {

}

func (e *EVMTraverser) traverseForStatement(n ast.ForStatementNode) {
	e.Traverse(n.Init)
	// if condition, do Block
	e.Traverse(n.Cond)
	e.Traverse(n.Block)
	e.Traverse(n.Post)
	// jump back to the start of the loop
}

func (e *EVMTraverser) traverseReturnStatement(n ast.ReturnStatementNode) {
	for _, r := range n.Results {
		e.Traverse(r)
	}
}

func (e *EVMTraverser) traverseIfStatement(n ast.IfStatementNode) {
	for _, c := range n.Conditions {
		e.Traverse(c.Condition)
		e.Traverse(c.Body)
	}
}

func (e *EVMTraverser) traverseAssignmentStatement(n ast.AssignmentStatementNode) {

}
