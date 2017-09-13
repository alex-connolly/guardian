package evm

import (
    "github.com/end-r/guardian/compiler/ast"
    "github.com/end-r/guardian/compiler/lexer"
)

func (e *EVMTraverser) traverseArrayLiteral(n ast.ArrayLiteralNode) {
	for _, expr := range n.Data {
		e.Traverse(expr)
	}
}

func (e *EVMTraverser) traverseSliceExpression(n ast.SliceExpressionNode) {

}

func (e *EVMTraverser) traverseCompositeLiteral(n ast.CompositeLiteralNode) {

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

func (e *EVMTraverser) traverseBinaryExpr(n ast.BinaryExpressionNode) {
	e.Traverse(n.Left)
	e.Traverse(n.Right)
	// operation
	e.VM.AddBytecode(binaryOps[n.Operator])
}

var unaryOps = map[lexer.TokenType]string{
	lexer.TknNot: "NOT",
}

func (e *EVMTraverser) traverseUnaryExpr(n ast.UnaryExpressionNode) {
	e.VM.AddBytecode(unaryOps[n.Operator])
	e.Traverse(n.Operand)
}

func (e *EVMTraverser) traverseCallExpr(n ast.CallExpressionNode) {
	for _, arg := range n.Arguments {
		e.Traverse(arg)
	}
	// parameters are at the top of the stack
	// jump to the top of the function
}

func (e *EVMTraverser) traverseLiteral(n ast.LiteralNode) {
	// Literal Nodes are directly converted to push instructions
	var parameters []byte
	bytes := n.GetBytes()
	parameters = append(parameters, byte(len(bytes)))
	parameters = append(parameters, bytes...)
	e.VM.AddBytecode("PUSH", parameters...)
}

func (e *EVMTraverser) traverseIndex(n ast.IndexExpressionNode) {
	// evaluate the index
	e.Traverse(n.Index)
	// then MLOAD it at the index offset
	e.Traverse(n.Expression)
	e.VM.AddBytecode("GET")
}

func (e *EVMTraverser) traverseMapLiteral(n ast.MapLiteralNode) {
	for k, v := range n.Data {
		e.Traverse(k)
		e.Traverse(v)
	}
	// push the size of the map
	e.VM.AddBytecode("PUSH", byte(1), byte(len(n.Data)))
	e.VM.AddBytecode("MAP")
}

func (e *EVMTraverser) traverseReference(n ast.ReferenceNode) {

	// reference e.g. dog.tail.wag()
	// get the object
	/*if n.InStorage {
		// if in storage
		// only the top level name is accessible in storage
		// everything else is accessed
		e.VM.AddBytecode("PUSH", len(n.Names[0]), n.Names[0])
		e.VM.AddBytecode("LOAD")

		// now get the sub-references
		// e.VM.AddBytecode("", params)
	} else {
		e.VM.AddBytecode("GET")
	}*/
}
