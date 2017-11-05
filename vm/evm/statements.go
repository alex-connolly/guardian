package evm

import (
	"github.com/end-r/guardian/compiler/ast"
	"github.com/end-r/guardian/compiler/lexer"
)

func (e *Traverser) traverseSwitchStatement(n ast.SwitchStatementNode) {
	// always traverse the target
	e.Traverse(n.Target)
	// switch statements are implicitly converted to if statements
	// may be a better way to do this
	// Solidity doesn't have a switch so shrug

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
	// consider mismatched lengths
	if len(n.Left) > 1 && len(n.Right) == 1 {
		for _, l := range n.Left {
			e.Traverse(l)
			e.Traverse(n.Right[0])
			// assignments are either in memory or storage depending on the context
			if e.inStorage() || hasModifier(n, lexer.TknStorage) {
				e.AddBytecode("")
			}
		}
	} else {
		for i, l := range n.Left {
			e.Traverse(l)
			e.Traverse(n.Right[i])
			// assignments are either in memory or storage depending on the context
			if e.inStorage() || hasModifier(n, lexer.TknStorage) {
				e.AddBytecode("")
			}
		}
	}

}

func hasModifier(n ast.Node, modifier lexer.TokenType) bool {
	return false
}
