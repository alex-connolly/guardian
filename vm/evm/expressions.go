package evm

import (
	"github.com/end-r/guardian/compiler/ast"
	"github.com/end-r/guardian/compiler/lexer"
)

func (e *Traverser) traverseArrayLiteral(n ast.ArrayLiteralNode) bytecode {

	b := make(bytecode, 0)
	// create array
	for _, expr := range n.Data {
		b = append(b, e.Traverse(expr)...)
	}
	return b
}

func (e *Traverser) traverseSliceExpression(n ast.SliceExpressionNode) bytecode {

}

func (e *Traverser) traverseCompositeLiteral(n ast.CompositeLiteralNode) bytecode {

	if e.inStorage() {

	} else {
		// the object exists in memory
		// things will have a byte offset

	}

}

var binaryOps = map[lexer.TokenType]string{
	lexer.TknAdd: "ADD",
	lexer.TknSub: "SUB",
	lexer.TknMul: "MUL",
	lexer.TknDiv: "DIV",
	lexer.TknMod: "MOD",
	lexer.TknShl: "SHL",
	lexer.TknShr: "SHR",
	lexer.TknAnd: "AND",
	lexer.TknOr:  "OR",
	lexer.TknXor: "XOR",
}

func (e *Traverser) traverseBinaryExpr(n ast.BinaryExpressionNode) {
	/* alter stack:

	| Operand 1 |
	| Operand 2 |
	| Operator  |

	Note that these operands may contain further expressions of arbitrary depth.
	*/
	e.Traverse(n.Left)
	e.Traverse(n.Right)
	// operation
	e.AddBytecode(binaryOps[n.Operator])
}

var unaryOps = map[lexer.TokenType]string{
	lexer.TknNot: "NOT",
}

func (e *Traverser) traverseUnaryExpr(n ast.UnaryExpressionNode) {
	/* alter stack:

	| Expression 1 |
	| Operand      |

	Note that these expressions may contain further expressions of arbitrary depth.
	*/
	e.Traverse(n.Operand)
	e.AddBytecode(unaryOps[n.Operator])
}

func (e *Traverser) traverseCallExpr(n ast.CallExpressionNode) {
	for _, arg := range n.Arguments {
		e.Traverse(arg)
	}
	// parameters are at the top of the stack
	// jump to the top of the function

}

func (e *Traverser) traverseLiteral(n ast.LiteralNode) {
	// Literal Nodes are directly converted to push instructions
	// these nodes must be divided into blocks of 16 bytes
	// in order to maintain

	bytes := n.GetBytes()
	const maxLength = 16

	for remaining := len(bytes); remaining > maxLength; remaining -= maxLength {
		base := len(bytes) - remaining
		e.AddBytecode("PUSH", bytes[base:base+maxLength]...)
	}

}

func (e *Traverser) traverseIndex(n ast.IndexExpressionNode) {

	e.Traverse(n.Index)
	e.Traverse(n.Expression)

	if e.inStorage() {
		e.AddBytecode("SLOAD")
	} else {
		e.AddBytecode("MLOAD")
	}

}

func (e *Traverser) traverseMapLiteral(n ast.MapLiteralNode) {
	// the evm doesn't support maps in the same way firevm does
	// Solidity converts things to a mapping
	// all keys to all values etc
	// precludes iteration
	// TODO: can we do it better?
	e.AddBytecode("")
}

func (e *Traverser) traverseFuncLiteral(n ast.FuncLiteralNode) {
	// create a hook

}

func (e *Traverser) traverseReference(n ast.ReferenceNode) {

	e.AddBytecode("PUSH")

	if e.inStorage() {
		e.AddBytecode("SLOAD")
	} else {
		e.AddBytecode("MLOAD")
	}

	// reference e.g. dog.tail.wag()
	// get the object
	/*if n.InStorage {
		// if in storage
		// only the top level name is accessible in storage
		// everything else is accessed
		e.AddBytecode("PUSH", len(n.Names[0]), n.Names[0])
		e.AddBytecode("LOAD")

		// now get the sub-references
		// e.AddBytecode("", params)
	} else {
		e.AddBytecode("GET")
	}*/
}
