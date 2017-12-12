package evm

import (
	"github.com/end-r/vmgen"

	"github.com/end-r/guardian/ast"
	"github.com/end-r/guardian/lexer"
)

func (e *GuardianEVM) traverseSwitchStatement(n *ast.SwitchStatementNode) (code vmgen.Bytecode) {
	// always traverse the target
	e.traverse(n.Target)
	// switch statements are implicitly converted to if statements
	// may be a better way to do this
	// Solidity doesn't have a switch so shrug
	return code
}

func (e *GuardianEVM) traverseCaseStatement(n *ast.CaseStatementNode) (code vmgen.Bytecode) {
	return code
}

func (e *GuardianEVM) traverseForStatement(n *ast.ForStatementNode) (code vmgen.Bytecode) {

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

func increment(varName string) (code vmgen.Bytecode) {

	code.Concat(pusher([]byte(name)...)
	code.Add("DUP1")
	code.Add("MLOAD")
	code.Concat(pusher(1))
	code.Add("ADD")
	code.Add("MSTORE")
}

func (e *GuardianEVM) traverseForEachStatement(n *ast.ForEachStatementNode) (code vmgen.Bytecode) {
	// starting from index 0
	// same for
	// allocate memory for the index
	// NOTE:
	// can potentially support dmaps by encoding a backing array as well
	// would be more expensive - add a keyword?

	switch n.ResolvedType {
	case typing.Array:
		// TODO: index size must be large enough for any vars
		name := n.Variables[0])
		e.allocateMemory(name, 10)
		memloc := e.lookupMemory(name).offset
		increment := increment(name)
		block := e.traverse(n.Block)
		code.Add(pusher(0))

		code.Add(pusher(memloc))
		code.Add("MSTORE")
		code.Add("JUMPDEST")
		code.Add(pusher(memloc))
		code.Add("MLOAD")
		code.Add("LT")
		code.AddMarked(pusher(1 + len(increment) + len(body)))
		code.Add("JUMPI")
		
		code.Concat(block)
		code.Concat(increment)

		code.Add("JUMPDEST")
		break
	case typing.Map:
		break
	}
	if n.Variables

}

func (e *GuardianEVM) traverseReturnStatement(n *ast.ReturnStatementNode) (code vmgen.Bytecode) {
	for _, r := range n.Results {
		// leave each of them on the stack in turn
		e.traverse(r)
	}
	// jump back to somewhere
	return code
}

func (e *GuardianEVM) traverseIfStatement(n *ast.IfStatementNode) (code vmgen.Bytecode) {
	for _, c := range n.Conditions {
		cond := e.traverse(c.Condition)
		block = e.traverse(c.Body)
		code.Concat(cond)
		code.Concat("ISZERO")
		code.AddMarked("PUSH", code.MarkOffset(len(block)+1))
		code.Concat("JUMPI")
		code.Concat(block)
	}

	code.Concat(e.traverse(n.Else))

	return code
}

func (e *GuardianEVM) traverseAssignmentStatement(n *ast.AssignmentStatementNode) (code vmgen.Bytecode) {
	// consider mismatched lengths
	if len(n.Left) > 1 && len(n.Right) == 1 {
		for _, l := range n.Left {
			r := n.Right[0]
			code.Concat(e.assign(l, r, true))
		}
	} else {
		for i, l := range n.Left {
			r := n.Right[i]
			code.Concat(e.assign(l, r, true))
		}
	}
	return code
}

func (e *GuardianEVM) assign(l, r ast.ExpressionNode, inStorage bool) (code vmgen.Bytecode) {
	e.currentlyAssigning = l
	e.traverse(l)
	e.traverse(r)

	// assignments are either in memory or storage depending on the context
	if inStorage {
		e.allocateStorage(name, size)
		code.Add("SSTORE")
	} else {
		e.allocateMemory(name, size)
		code.Add("MSTORE")
	}
	return code
}

func hasModifier(mods []lexer.TokenType, modifier lexer.TokenType) bool {
	for _, m := range mods {
		if m == modifier {
			return true
		}
	}
	return false
}
