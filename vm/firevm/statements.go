package firevm

import "github.com/end-r/guardian/compiler/ast"

func (a *Arsonist) traverseSwitchStatement(n ast.SwitchStatementNode) {
	// perform switch expression
	a.Traverse(n.Target)
	// now go through the cases
	for _, c := range n.Clauses {
		a.Traverse(c)
	}
}

func (a *Arsonist) traverseCaseStatement(n ast.CaseStatementNode) {

}

func (a *Arsonist) traverseForStatement(n ast.ForStatementNode) {
	// perform init statement
	a.Traverse(n.Init)
	a.Traverse(n.Cond)
	a.Traverse(n.Block)
	a.Traverse(n.Post)
	a.VM.AddBytecode("JUMPI")

}

func (a *Arsonist) traverseReturnStatement(n ast.ReturnStatementNode) {
	for _, r := range n.Results {
		// traverse the results 1 by 1, leave on the stack
		a.Traverse(r)
	}
}

func (a *Arsonist) traverseIfStatement(n ast.IfStatementNode) {
	for _, c := range n.Conditions {
		a.Traverse(c.Condition)
		a.VM.AddBytecode("JUMPI")
		a.Traverse(c.Body)
	}
}

func (a *Arsonist) traverseAssignmentStatement(n ast.AssignmentStatementNode) {
	/*	if n.IsStorage {
			for i, r := range n.Right {
				a.Traverse(r)
				if len(n.Left) == 1 {
					a.Traverse(n.Left[0])
				} else {
					a.Traverse(n.Left[i])
				}
				a.VM.AddBytecode("STORE")
			}
		} else {
			// in memory

		}*/
}
