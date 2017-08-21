package ast

import (
	"encoding/binary"

	"github.com/end-r/firevm"
	"github.com/end-r/guardian/go/compiler/lexer"
)

type ExpressionNode interface {
	Type() NodeType
	Traverse(firevm.VM)
}

// BinaryExpressionNode ...
type BinaryExpressionNode struct {
	Left, Right ExpressionNode
	Operator    lexer.TokenType
}

// Type ...
func (n BinaryExpressionNode) Type() NodeType { return BinaryExpression }

func (n BinaryExpressionNode) Traverse(vm firevm.VM) {
	n.Left.Traverse(vm)
	n.Right.Traverse(vm)
	//bytecode.TraverseOperator(vm, n.Operator)
}

// UnaryExpressionNode ...
type UnaryExpressionNode struct {
	Operator lexer.TokenType
	Operand  ExpressionNode
}

func (n UnaryExpressionNode) Type() NodeType { return UnaryExpression }

func (n UnaryExpressionNode) Traverse(vm firevm.VM) {
	switch n.Operator {
	case lexer.TknNot:
		// push data
		// add not instruction
		vm.AddInstruction("NOT")
		break
	}
}

type LiteralNode struct {
	Data        string
	LiteralType lexer.TokenType
}

func (n LiteralNode) Type() NodeType { return Literal }

func (n LiteralNode) Traverse(vm firevm.VM) {
	// Literal Nodes are directly converted to push instructions
	parameters := make([]byte, 0)
	bytes := n.GetBytes()
	parameters = append(parameters, len(bytes))
	parameters = append(parameters, bytes...)
	vm.AddInstruction("PUSH", parameters...)
}

func (n LiteralNode) GetBytes() []byte {
	switch n.LiteralType {
	case lexer.String:
		return []byte(n.Data)
	case lexer.Int:
		bs := make([]byte, 4)
		binary.LittleEndian.PutUint32(bs, 31415926)
		return []byte()
	case lexer.Float:

	}
}

type CompositeLiteralNode struct {
	Reference ExpressionNode
	Fields    map[string]ExpressionNode
}

func (n CompositeLiteralNode) Type() NodeType { return CompositeLiteral }

func (n CompositeLiteralNode) Traverse(vm firevm.VM) {
	for k, v := range n.Fields {

	}
}

type IndexExpressionNode struct {
	Expression ExpressionNode
	Index      ExpressionNode
}

func (n IndexExpressionNode) Type() NodeType { return IndexExpression }

func (n IndexExpressionNode) Traverse(vm firevm.VM) {
	// evaluate the index
	n.Index.Traverse(vm)
	// then MLOAD it at the index offset
	n.Expression.Traverse(vm)
	vm.AddInstruction("GET")
}

type SliceExpressionNode struct {
	Expression ExpressionNode
	Low, High  ExpressionNode
	Max        ExpressionNode
}

func (n SliceExpressionNode) Type() NodeType { return SliceExpression }

func (n SliceExpressionNode) Traverse(vm firevm.VM) {

}

type CallExpressionNode struct {
	Call      ExpressionNode
	Arguments []ExpressionNode
}

func (n CallExpressionNode) Type() NodeType { return CallExpression }

func (n CallExpressionNode) Traverse(vm firevm.VM) {
	// push arguments onto the stack
	for _, a := range n.Arguments {
		n.Traverse()
		// return values should be on the stack right now
	}

}

type ArrayLiteralNode struct {
	Key  ReferenceNode
	Size ExpressionNode
	Data []ExpressionNode
}

func (n ArrayLiteralNode) Type() NodeType { return ArrayLiteral }

func (n ArrayLiteralNode) Traverse(vm firevm.VM) {
	for _, en := range n.Data {
		en.Traverse(vm)
	}
	// push the size of the array
	vm.AddInstruction("PUSH", byte(1), byte(len(n.Data)))
	vm.AddInstruction("ARRAY")
}

type MapLiteralNode struct {
	Key   ReferenceNode
	Value ReferenceNode
	Data  map[ExpressionNode]ExpressionNode
}

func (n MapLiteralNode) Type() NodeType { return MapLiteral }

func (n MapLiteralNode) Traverse(vm firevm.VM) {
	for k, v := range n.Data {
		k.Traverse(vm)
		v.Traverse(vm)
	}
	// push the size of the map
	vm.AddInstruction("PUSH", byte(1), byte(len(n.Data)))
	vn.AddInstruction("MAP")
}

type ReferenceNode struct {
	Names []string
}

func (n ReferenceNode) Type() NodeType { return Reference }

func (n ReferenceNode) Traverse(vm firevm.VM) {

}
