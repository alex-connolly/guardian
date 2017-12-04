package evm

import (
	"github.com/end-r/vmgen"

	"github.com/end-r/guardian/ast"
	"github.com/end-r/guardian/lexer"
)

func (e *GuardianEVM) traverseSwitchStatement(n ast.SwitchStatementNode) (code vmgen.Bytecode) {
	// always traverse the target
	e.traverse(n.Target)
	// switch statements are implicitly converted to if statements
	// may be a better way to do this
	// Solidity doesn't have a switch so shrug
	return code
}

func (e *GuardianEVM) traverseCaseStatement(n ast.CaseStatementNode) (code vmgen.Bytecode) {
	return code
}

func (e *GuardianEVM) traverseForStatement(n ast.ForStatementNode) (code vmgen.Bytecode) {

	// init statement
	// jumpdest
	// condition
	// jump to end
	// regular loop processes would occur here
	// post statement
	// jump back to them top of the loop
	// jumpdest
	// continue after the loop

	init := e.traverse(n.Init)

	cond := e.traverseExpression(n.Cond)

	block := e.traverse(n.Block)

	post := e.traverse(n.Post)

	code.Concat(init)
	code.Add("JUMPDEST")
	code.Concat(cond)
	code.Add("JUMPI")
	code.Concat(block)
	code.Concat(post)
	code.Add("JUMP")

	return code
}

func (e *GuardianEVM) traverseReturnStatement(n ast.ReturnStatementNode) (code vmgen.Bytecode) {
	for _, r := range n.Results {
		e.traverse(r)
	}
	return code
}

func (e *GuardianEVM) traverseIfStatement(n ast.IfStatementNode) (code vmgen.Bytecode) {
	for _, c := range n.Conditions {
		e.traverse(c.Condition)
		e.traverse(c.Body)
	}
	return code
}

func (e *GuardianEVM) traverseAssignmentStatement(n ast.AssignmentStatementNode) (code vmgen.Bytecode) {
	// consider mismatched lengths
	if len(n.Left) > 1 && len(n.Right) == 1 {
		for _, l := range n.Left {
			e.traverse(l)
			e.traverse(n.Right[0])
			// assignments are either in memory or storage depending on the context
			if e.inStorage() || hasModifier(n, lexer.TknStorage) {
				code.Add("")
			}
		}
	} else {
		for i, l := range n.Left {
			e.traverse(l)
			e.traverse(n.Right[i])
			// assignments are either in memory or storage depending on the context
			if e.inStorage() || hasModifier(n, lexer.TknStorage) {
				code.Add("")
			}
		}
	}
	return code
}

func hasModifier(n ast.Node, modifier lexer.TokenType) bool {
	return false
}
