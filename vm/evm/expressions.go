package evm

import (
	"fmt"

	"github.com/end-r/guardian/typing"

	"github.com/end-r/guardian/token"

	"github.com/end-r/vmgen"

	"github.com/end-r/guardian/ast"
)

func (e *GuardianEVM) traverseValue(n ast.ExpressionNode) (code vmgen.Bytecode) {
	return code
}

func (e *GuardianEVM) traverseLocation(n ast.ExpressionNode) (code vmgen.Bytecode) {
	return code
}

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

var binaryOps = map[token.Type]BinaryOperator{
	token.Add:        additionOrConcatenation,
	token.Sub:        simpleOperator("SUB"),
	token.Mul:        simpleOperator("MUL"),
	token.Div:        signedOperator("DIV", "SDIV"),
	token.Mod:        signedOperator("MOD", "SMOD"),
	token.Shl:        simpleOperator("SHL"),
	token.Shr:        simpleOperator("SHR"),
	token.And:        simpleOperator("AND"),
	token.Or:         simpleOperator("OR"),
	token.Xor:        simpleOperator("XOR"),
	token.As:         ignoredOperator(),
	token.Gtr:        signedOperator("GT", "SGT"),
	token.Lss:        signedOperator("LT", "SLT"),
	token.Eql:        simpleOperator("EQL"),
	token.Neq:        reversedOperator("EQL"),
	token.Geq:        reversedSignedOperator("LT", "SLT"),
	token.Leq:        reversedSignedOperator("GT", "SGT"),
	token.LogicalAnd: ignoredOperator(),
	token.LogicalOr:  ignoredOperator(),
}

type BinaryOperator func(n *ast.BinaryExpressionNode) vmgen.Bytecode

func reversedSignedOperator(unsigned, signed string) BinaryOperator {
	return func(n *ast.BinaryExpressionNode) (code vmgen.Bytecode) {
		left, lok := typing.ResolveUnderlying(n.Left.ResolvedType()).(*typing.NumericType)
		right, rok := typing.ResolveUnderlying(n.Left.ResolvedType()).(*typing.NumericType)
		if lok && rok {
			if left.Signed || right.Signed {
				code.Add(signed)
			} else {
				code.Add(unsigned)
			}
			code.Add("NOT")
		}
		return code
	}
}

func reversedOperator(mnemonic string) BinaryOperator {
	return func(n *ast.BinaryExpressionNode) (code vmgen.Bytecode) {
		code.Add(mnemonic)
		code.Add("NOT")
		return code
	}
}

func ignoredOperator() BinaryOperator {
	return func(n *ast.BinaryExpressionNode) (code vmgen.Bytecode) {
		return code
	}
}

func simpleOperator(mnemonic string) BinaryOperator {
	return func(n *ast.BinaryExpressionNode) (code vmgen.Bytecode) {
		code.Add(mnemonic)
		return code
	}
}

func signedOperator(unsigned, signed string) BinaryOperator {
	return func(n *ast.BinaryExpressionNode) (code vmgen.Bytecode) {
		//fmt.Println(typing.WriteType(n.Left.ResolvedType()))
		left, lok := typing.ResolveUnderlying(n.Left.ResolvedType()).(*typing.NumericType)
		right, rok := typing.ResolveUnderlying(n.Right.ResolvedType()).(*typing.NumericType)
		if lok && rok {
			if left.Signed || right.Signed {
				code.Add(signed)
			} else {
				code.Add(unsigned)
			}
		}
		return code
	}
}

func additionOrConcatenation(n *ast.BinaryExpressionNode) (code vmgen.Bytecode) {
	switch n.Resolved.(type) {
	case *typing.NumericType:
		code.Add("ADD")
		return code
	default:
		// otherwise must be a string
		return code

	}
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

	code.Concat(binaryOps[n.Operator](n))

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

func (e *GuardianEVM) traverseContractCall(n *ast.CallExpressionNode) (code vmgen.Bytecode) {
	// calls another contract, The CREATE opcode takes three values:
	// value (ie. initial amount of ether),
	code.Add("GAS")
	// memory start
	code.Add("PUSH")
	// memory length
	code.Add("PUSH")
	code.Add("CREATE")
	return code
}

func (e *GuardianEVM) traverseClassCall(n *ast.CallExpressionNode) (code vmgen.Bytecode) {
	return code
}

func (e *GuardianEVM) traverseFunctionCall(n *ast.CallExpressionNode) (code vmgen.Bytecode) {
	for _, arg := range n.Arguments {
		code.Concat(e.traverseExpression(arg))
	}

	// traverse the call expression
	// should leave the function address on top of the stack

	// need to get annotations through this process

	if n.Call.Type() == ast.Identifier {
		i := n.Call.(*ast.IdentifierNode)
		if b, ok := builtins[i.Name]; ok {
			code.Concat(b(e))
			return code
		}
	}

	call := e.traverse(n.Call)

	code.Concat(call)

	// parameters are at the top of the stack
	// jump to the top of the function
	return code
}

func (e *GuardianEVM) traverseCallExpr(n *ast.CallExpressionNode) (code vmgen.Bytecode) {
	e.expression = n

	switch typing.ResolveUnderlying(n.Call.ResolvedType()).(type) {
	case *typing.Func:
		return e.traverseFunctionCall(n)
	case *typing.Contract:
		return e.traverseContractCall(n)
	case *typing.Class:
		return e.traverseClassCall(n)
	}
	return code
}

func (e *GuardianEVM) traverseLiteral(n *ast.LiteralNode) (code vmgen.Bytecode) {

	// Literal Nodes are directly converted to push instructions
	// these nodes must be divided into blocks of 16 bytes
	// in order to maintain

	// maximum number size is 256 bits (32 bytes)
	switch n.LiteralType {
	case token.Integer, token.Float:
		if len(n.Data) > 32 {
			// error
		} else {
			fmt.Println("HERE")
			// TODO: type size or data size
			code.Add(fmt.Sprintf("PUSH%d", bytesRequired(int(n.Resolved.Size()))), []byte(n.Data)...)
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

	// TODO: bounds checking?

	// load the data
	code.Concat(e.traverseExpression(n.Expression))

	typ := n.Expression.ResolvedType()

	// calculate offset
	// evaluate index
	code.Concat(e.traverseExpression(n.Index))
	// get size of type
	code.Concat(push(encodeUint(typ.Size())))
	// offset = size of type * index
	code.Add("MUL")

	code.Add("ADD")

	return code
}

const (
	mapLiteralReserved   = "gevm_map_literal_%d"
	arrayLiteralReserved = "gevm_array_literal_%d"
)

// reserve name for map literals: gevm_map_literal_{count}
// if you try to name things it won't let you
func (evm *GuardianEVM) traverseMapLiteral(n *ast.MapLiteralNode) (code vmgen.Bytecode) {

	fakeKey := fmt.Sprintf(mapLiteralReserved, evm.mapLiteralCount)

	// must be deterministic iteration here
	for _, v := range n.Data {
		// each storage slot must be 32 bytes regardless of contents
		slot := EncodeName(fakeKey + evm.traverse(n))
		code.Concat(push(slot))
		code.Add("SSTORE")
	}

	evm.mapLiteralCount++

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

	code.Concat(e.traverseScope(n.Scope))

	for _, p := range n.Parameters {
		for _, i := range p.Identifiers {
			e.freeMemory(i)
		}
	}

	return code
}

func (e *GuardianEVM) traverseIdentifier(n *ast.IdentifierNode) (code vmgen.Bytecode) {

	if e.inStorage {
		s := e.lookupStorage(n.Name)
		if s != nil {
			return s.retrieve()
		}
		e.allocateStorage(n.Name, n.Resolved.Size())
		s = e.lookupStorage(n.Name)
		if s != nil {
			return s.retrieve()
		}
	} else {
		m := e.lookupMemory(n.Name)
		if m != nil {
			return m.retrieve()
		}
		e.allocateMemory(n.Name, n.Resolved.Size())
		m = e.lookupMemory(n.Name)
		if m != nil {
			return m.retrieve()
		}
	}
	return code
}

func (e *GuardianEVM) traverseContextual(t typing.Type, expr ast.ExpressionNode) (code vmgen.Bytecode) {
	switch expr.(type) {
	case *ast.IdentifierNode:
		switch t.(type) {
		case *typing.Class:

			break
		case *typing.Interface:
			break
		}
		break
	case *ast.CallExpressionNode:
		break
	}
	return code
}

func (e *GuardianEVM) traverseReference(n *ast.ReferenceNode) (code vmgen.Bytecode) {
	code.Concat(e.traverse(n.Parent))

	resolved := n.Parent.ResolvedType()

	ctx := e.traverseContextual(resolved, n.Reference)

	code.Concat(ctx)

	if e.inStorage {
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
