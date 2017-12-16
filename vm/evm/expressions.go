package evm

import (
	"fmt"

	"github.com/end-r/guardian/token"

	"github.com/end-r/vmgen"

	"github.com/end-r/guardian/ast"
)

func (e *GuardianEVM) traverseExpression(n ast.ExpressionNode) (code vmgen.Bytecode) {
	fmt.Println("expr")
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

var binaryOps = map[token.Type]string{
	token.Add: "ADD",
	token.Sub: "SUB",
	token.Mul: "MUL",
	token.Div: "DIV",
	token.Mod: "MOD",
	token.Shl: "SHL",
	token.Shr: "SHR",
	token.And: "AND",
	token.Or:  "OR",
	token.Xor: "XOR",
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

var unaryOps = map[token.Type]string{
	token.Not: "NOT",
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

	fmt.Println("Literal")
	// Literal Nodes are directly converted to push instructions
	// these nodes must be divided into blocks of 16 bytes
	// in order to maintain

	// maximum number size is 256 bits (32 bytes)
	switch n.LiteralType {
	case token.Integer, token.Float:
		if len(n.Data) > 32 {
			// error
		} else {
			code.Add(fmt.Sprintf("PUSH%d", len(n.Data)), []byte(n.Data)...)
		}
		break
	case token.String:
		bytes := []byte(n.Data)
		max := 32
		size := 0
		for size = len(bytes); size > max; size -= max {
			code.Add("PUSH32", bytes[len(bytes)-size:len(bytes)-size+max]...)
		}
		op := fmt.Sprintf("PUSH%d", size)
		code.Add(op, bytes[size:len(bytes)]...)
		break
	}
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

	/*fakeKey := e.currentlyAssigning

	i := 0
	 TODO: deterministic iteration
	for _, v := range n.Data {
		// each storage slot must be 32 bytes regardless of contents
		slot := EncodeName(fakeKey + "key")
		code.Concat(push(slot))
		code.Add("SSTORE")
		i++
	}*/
	return code
}

func (e *GuardianEVM) traverseFuncLiteral(n *ast.FuncLiteralNode) (code vmgen.Bytecode) {
	// create an internal hook

	// parameters should have been pushed onto the stack by the caller
	// take them off and put them in memory
	for _, p := range n.Parameters {
		for _, i := range p.Identifiers {
			e.allocateMemory(i, p.Resolved.Size())
			code.Add("MSTORE")
		}
	}

	code.Concat(e.traverse(n.Scope))

	for _, p := range n.Parameters {
		for _, i := range p.Identifiers {
			e.freeMemory(i)
		}
	}

	return code
}

func (e *GuardianEVM) traverseIdentifier(n *ast.IdentifierNode) (code vmgen.Bytecode) {
	if e.inStorage() {
		s := e.lookupStorage(n.Name)
		if s != nil {
			return s.retrieve()
		}
	} else {
		m := e.lookupMemory(n.Name)
		if m != nil {
			return m.retrieve()
		}
		e.allocateMemory(n.Name, n.ResolvedType().Size())
		m = e.lookupMemory(n.Name)
		if m != nil {
			return m.retrieve()
		}
	}
	return code
}

func (e *GuardianEVM) traverseReference(n *ast.ReferenceNode) (code vmgen.Bytecode) {
	code.Concat(e.traverse(n.Parent))

	//resolved := n.Parent.ResolvedType()

	//code.Concat(e.traverseContextual(resolved, n.Reference))

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
