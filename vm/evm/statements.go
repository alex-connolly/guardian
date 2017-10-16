package evm

import "github.com/end-r/guardian/compiler/ast"

func (e *Traverser) traverseSwitchStatement(n ast.SwitchStatementNode) {
	// always traverse the target
	e.Traverse(n.Target)

}

func (e *Traverser) traverseCaseStatement(n ast.CaseStatementNode) {

}

func (e *Traverser) traverseForStatement(n ast.ForStatementNode) {
	e.Traverse(n.Init)
	// if condition, do Block
	e.Traverse(n.Cond)
	e.Traverse(n.Block)
	e.Traverse(n.Post)
	// jump back to the start of the loop
}

func (e *Traverser) traverseReturnStatement(n ast.ReturnStatementNode) {
	for _, r := range n.Results {
		e.Traverse(r)
	}
}

func (e *Traverser) traverseIfStatement(n ast.IfStatementNode) {
	for _, c := range n.Conditions {
		e.Traverse(c.Condition)
		e.Traverse(c.Body)
	}
}

func (e *Traverser) traverseAssignmentStatement(n ast.AssignmentStatementNode) {
	// assignments are either in memory or storage depending on the context

}
