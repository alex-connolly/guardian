package evm

import (
	"github.com/end-r/guardian/typing"

	"github.com/end-r/vmgen"

	"github.com/end-r/guardian/ast"
)

func (e *GuardianEVM) traverseSwitchStatement(n *ast.SwitchStatementNode) (code vmgen.Bytecode) {
	// always traverse the target
	var generatedExpressions []vmgen.Bytecode
	if n.Target != nil {
		e.traverse(n.Target)
		for _, c := range n.Cases.Sequence {
			switch cas := c.(type) {
			case *ast.CaseStatementNode:
				for _, exp := range cas.Expressions {
					generatedExpressions = append(generatedExpressions, e.traverseExpression(exp))
				}
				break
			}
		}
		sum := 0
		for _, exp := range generatedExpressions {
			sum += 1 // for duplicating the top of the stack
			sum += exp.Length()
			sum += 3 // for EQ, PUSH, JUMP
		}
		for _, exp := range generatedExpressions {
			code.Concat(exp)
			code.Concat(pushMarker(sum))
			code.Add("EQ")
			code.Add("JUMPI")
		}
	} else {
		for _, c := range n.Cases.Sequence {
			switch cas := c.(type) {
			case *ast.CaseStatementNode:
				for _, exp := range cas.Expressions {
					generatedExpressions = append(generatedExpressions, e.traverseExpression(exp))
				}
				break
			}
		}
		sum := 0
		for _, exp := range generatedExpressions {
			sum += 1 // for duplicating the top of the stack
			sum += exp.Length()
			sum += 3 // for EQ, PUSH, JUMP
		}
		for _, exp := range generatedExpressions {
			code.Concat(exp)
			code.Concat(pushMarker(sum))
			code.Add("JUMPI")
		}
	}

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
	// jump back to the top of the loop
	// jumpdest
	// continue after the loop

	init := e.traverse(n.Init)

	cond := e.traverseExpression(n.Cond)

	block := e.traverseScope(n.Block)

	post := e.traverse(n.Post)

	code.Concat(init)
	code.Add("JUMPDEST")
	top := code.Length()
	code.Concat(cond)
	code.Concat(pushMarker(block.Length() + post.Length() + 1))
	code.Add("JUMPI")
	code.Concat(block)
	code.Concat(post)
	bottom := code.Length()
	code.Concat(pushMarker(bottom - top))
	code.Add("JUMP")
	code.Add("JUMPDEST")

	return code
}

func (evm *GuardianEVM) increment(varName string) (code vmgen.Bytecode) {
	code.Concat(push([]byte(varName)))
	code.Add("DUP1")
	code.Add("MLOAD")
	code.Concat(push(uintAsBytes(1)))
	code.Add("ADD")
	code.Add("MSTORE")
	return code
}

func (evm *GuardianEVM) traverseForEachStatement(n *ast.ForEachStatementNode) (code vmgen.Bytecode) {
	// starting from index 0
	// same for
	// allocate memory for the index
	// NOTE:
	// can potentially support dmaps by encoding a backing array as well
	// would be more expensive - add a keyword?

	switch n.ResolvedType.(type) {
	case *typing.Array:
		// TODO: index size must be large enough for any vars
		name := n.Variables[0]
		e.allocateMemory(name, 10)
		memloc := e.lookupMemory(name).offset
		increment := e.increment(name)
		block := e.traverseScope(n.Block)
		code.Concat(push(encodeUint(0)))

		code.Concat(push(encodeUint(memloc)))
		code.Add("MSTORE")
		code.Add("JUMPDEST")
		code.Concat(push(encodeUint(memloc)))
		code.Add("MLOAD")
		code.Add("LT")
		code.Concat(pushMarker(1 + increment.Length() + block.Length()))
		code.Add("JUMPI")

		code.Concat(block)
		code.Concat(increment)

		code.Add("JUMPDEST")
		break
	case *typing.Map:
		break
	}
	return code
}

func (e *GuardianEVM) traverseReturnStatement(n *ast.ReturnStatementNode) (code vmgen.Bytecode) {
	for _, r := range n.Results {
		// leave each of them on the stack in turn
		code.Concat(e.traverse(r))
	}
	// jump back to somewhere
	// top of stack should now be return address
	code.Add("JUMP")
	return code
}

func (e *GuardianEVM) traverseControlFlowStatement(n *ast.FlowStatementNode) (code vmgen.Bytecode) {
	return code
}

func (e *GuardianEVM) traverseIfStatement(n *ast.IfStatementNode) (code vmgen.Bytecode) {
	conds := make([]vmgen.Bytecode, len(n.Conditions))
	blocks := make([]vmgen.Bytecode, len(n.Conditions))
	end := 0

	for _, c := range n.Conditions {
		cond := e.traverse(c.Condition)
		conds = append(conds, cond)
		body := e.traverseScope(c.Body)
		blocks = append(blocks, body)
		end += cond.Length() + body.Length() + 3 + 1
	}

	for i := range n.Conditions {
		code.Concat(conds[i])
		code.Add("ISZERO")
		code.Concat(pushMarker(blocks[i].Length() + 1))
		code.Add("JUMPI")
		code.Concat(blocks[i])
		code.Concat(pushMarker(end))
	}

	code.Concat(e.traverseScope(n.Else))

	return code
}

func (e *GuardianEVM) traverseAssignmentStatement(n *ast.AssignmentStatementNode) (code vmgen.Bytecode) {
	for i, l := range n.Left {
		r := n.Right[i]
		code.Concat(e.assign(l, r, e.inStorage))
	}
	return code
}

func (e *GuardianEVM) assign(l, r ast.ExpressionNode, inStorage bool) (code vmgen.Bytecode) {
	// get the location
	//code.Concat(e.traverseExpression(l))
	// do the calculation
	code.Concat(e.traverseExpression(r))
	if inStorage {
		code.Add("SSTORE")
	} else {
		code.Add("MSTORE")
	}
	return code
}
