package evm

import (
	"fmt"

	"github.com/end-r/vmgen"

	"github.com/end-r/guardian/ast"
	"github.com/end-r/guardian/lexer"
)

func (e *GuardianEVM) traverseExpression(n ast.ExpressionNode) (code vmgen.Bytecode) {
	switch node := n.(type) {
	case *ast.ArrayLiteralNode:
		return e.traverseArrayLiteral(node)
	case *ast.FuncLiteralNode:
		return e.traverseFuncLiteral(node)
	case *ast.MapLiteralNode:
		return e.traverseMapLiteral(node)
	case *ast.CompositeLiteralNode:
		return e.traverseCompositeLiteral(node)
	case *ast.UnaryExpressionNode:
		return e.traverseUnaryExpr(node)
	case *ast.BinaryExpressionNode:
		return e.traverseBinaryExpr(node)
	case *ast.CallExpressionNode:
		return e.traverseCallExpr(node)
	case *ast.IndexExpressionNode:
		return e.traverseIndex(node)
	case *ast.SliceExpressionNode:
		return e.traverseSliceExpression(node)
	case *ast.IdentifierNode:
		return e.traverseIdentifier(node)
	case *ast.ReferenceNode:
		return e.traverseReference(node)
	case *ast.LiteralNode:
		return e.traverseLiteral(node)
	}
	return code
}

func (e *GuardianEVM) traverseArrayLiteral(n *ast.ArrayLiteralNode) (code vmgen.Bytecode) {
	/*
		// encode the size first
		code.Add(uintAsBytes(len(n.Data))...)

		for _, expr := range n.Data {
			code.Concat(e.traverseExpression(expr))
		}*/

	return code
}

func (e *GuardianEVM) traverseSliceExpression(n *ast.SliceExpressionNode) (code vmgen.Bytecode) {
	// evaluate the original expression first

	// get the data
	code.Concat(e.traverseExpression(n.Expression))

	// ignore the first (item size) * lower

	// ignore the ones after

	return code
}

func (e *GuardianEVM) traverseCompositeLiteral(n *ast.CompositeLiteralNode) (code vmgen.Bytecode) {
	/*
		var ty validator.Class
		for _, f := range ty.Fields {
			if n.Fields[f] != nil {

			} else {
				n.Fields[f].Size()
			}
		}

		for _, field := range n.Fields {
			// evaluate each field
			code.Concat(e.traverseExpression(field))
		}*/
	return code
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

func (e *GuardianEVM) traverseBinaryExpr(n *ast.BinaryExpressionNode) (code vmgen.Bytecode) {
	/* alter stack:

	| Operand 1 |
	| Operand 2 |
	| Operator  |

	Note that these operands may contain further expressions of arbitrary depth.
	*/
	code.Concat(e.traverseExpression(n.Left))
	code.Concat(e.traverseExpression(n.Right))

	//op := binaryOps[n.Operator]

	return code
}

var unaryOps = map[lexer.TokenType]string{
	lexer.TknNot: "NOT",
}

func (e *GuardianEVM) traverseUnaryExpr(n *ast.UnaryExpressionNode) (code vmgen.Bytecode) {
	/* alter stack:

	| Expression 1 |
	| Operand      |

	Note that these expressions may contain further expressions of arbitrary depth.
	*/
	code.Concat(e.traverseExpression(n.Operand))
	code.Add(unaryOps[n.Operator])
	return code
}

func (e *GuardianEVM) traverseCallExpr(n *ast.CallExpressionNode) (code vmgen.Bytecode) {

	for _, arg := range n.Arguments {
		code.Concat(e.traverseExpression(arg))
	}

	// traverse the call expression
	// should leave the function address on top of the stack

	call := e.traverse(n.Call)

	// check to see whether we need to replace the callhash
	// with the builtin code
	if res, ok := checkBuiltin(call); ok {
		call = res
	}

	code.Concat(call)

	// parameters are at the top of the stack
	// jump to the top of the function
	return code
}

func checkBuiltin(code vmgen.Bytecode) (res vmgen.Bytecode, isBuiltin bool) {
	/*for name, b := range builtins {
		if code.CompareBytes(EncodeName(name)) {
			return b(), true
		}
	}*/
	return code, false
}

var builtins = map[string]Builtin{
	// arithmetic
	"addmod":  simpleInstruction("ADDMOD"),
	"mulmod":  simpleInstruction("MULMOD"),
	"balance": simpleInstruction("BALANCE"),
	// transactional
	"transfer":     nil,
	"send":         nil,
	"delegateCall": simpleInstruction("DELEGATECALL"),
	// cryptographic
	"sha3":      simpleInstruction("SHA3"),
	"keccak256": nil,
	"sha356":    nil,
	"ecrecover": nil,
	"ripemd160": nil,
}

type Builtin func() vmgen.Bytecode

// returns an anon func to handle simplest cases
func simpleInstruction(mnemonic string) Builtin {
	return func() (code vmgen.Bytecode) {
		code.Add(mnemonic)
		return code
	}
}

func (e *GuardianEVM) traverseLiteral(n *ast.LiteralNode) (code vmgen.Bytecode) {
	// Literal Nodes are directly converted to push instructions
	// these nodes must be divided into blocks of 16 bytes
	// in order to maintain

	// maximum number size is 256 bits (32 bytes)

	bytes := []byte(n.Data)
	max := 32
	size := 0
	for size = len(bytes); size > max; size -= max {
		code.Add("PUSH32", bytes[len(bytes)-size:len(bytes)-size+max]...)
	}
	op := fmt.Sprintf("PUSH%d", size)
	code.Add(op, bytes[size:len(bytes)]...)
	return code
}

func (e *GuardianEVM) traverseIndex(n *ast.IndexExpressionNode) (code vmgen.Bytecode) {

	// load the data
	code.Concat(e.traverseExpression(n.Expression))

	// calculate offset, get bytes
	code.Concat(e.traverseExpression(n.Index))

	return code
}

func (e *GuardianEVM) traverseMapLiteral(n *ast.MapLiteralNode) (code vmgen.Bytecode) {
	/* use indirection to do this
	// map literals will use an extra 20k gas
	//fakeKey := e.generateNextIndirect()
	// keccak256(bytes32(key) + bytes32(position))

	i := 0
	// TODO: deterministic iteration
	for k, v := range n.Data {
		slot := EncodeName(fakeKey + bytes32(i))
		code.Add()
		code.Add()
		i++
	}
	code.Add("PUSH", fakeKey)*/
	return code
}

func (e *GuardianEVM) traverseFuncLiteral(n *ast.FuncLiteralNode) (code vmgen.Bytecode) {
	// create an internal hook

	return code
}

func isStorage(name string) bool {
	return false
}

func (e *GuardianEVM) traverseIdentifier(n *ast.IdentifierNode) (code vmgen.Bytecode) {
	if isStorage(n.Name) {
		return e.lookupStorage(n.Name).retrive()
	} else {
		return code
		//return e.lookupMemory(n.Name).retrieve()
	}
}

func (e *GuardianEVM) traverseReference(n *ast.ReferenceNode) (code vmgen.Bytecode) {

	code.Concat(e.traverse(n.Parent))

	if e.inStorage() {
		code.Add("SLOAD")
	} else {
		code.Add("MLOAD")
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
	return code
}
