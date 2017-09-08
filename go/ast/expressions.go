package ast

import (
	"github.com/end-r/vmgen"

	"github.com/end-r/guardian/go/lexer"
)

// BinaryExpressionNode ...
type BinaryExpressionNode struct {
	Left, Right ExpressionNode
	Operator    lexer.TokenType
}

// Type ...
func (n BinaryExpressionNode) Type() NodeType { return BinaryExpression }

func (n BinaryExpressionNode) Traverse(vm *vmgen.VM) {
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

func (n UnaryExpressionNode) Traverse(vm *vmgen.VM) {
	switch n.Operator {
	case lexer.TknNot:
		// push data
		// add not instruction
		vm.AddBytecode("NOT")
		break
	}
}

type LiteralNode struct {
	Data        string
	LiteralType lexer.TokenType
}

func (n LiteralNode) Type() NodeType { return Literal }

func (n LiteralNode) Traverse(vm *vmgen.VM) {
	// Literal Nodes are directly converted to push instructions
	parameters := make([]byte, 0)
	bytes := n.GetBytes()
	parameters = append(parameters, byte(len(bytes)))
	parameters = append(parameters, bytes...)
	vm.AddBytecode("PUSH", parameters...)
}

func (n LiteralNode) GetBytes() []byte {
	return nil
}

type CompositeLiteralNode struct {
	Reference ExpressionNode
	Fields    map[string]ExpressionNode
}

func (n CompositeLiteralNode) Type() NodeType { return CompositeLiteral }

func (n CompositeLiteralNode) Traverse(vm *vmgen.VM) {

}

type IndexExpressionNode struct {
	Expression ExpressionNode
	Index      ExpressionNode
}

func (n IndexExpressionNode) Type() NodeType { return IndexExpression }

func (n IndexExpressionNode) Traverse(vm *vmgen.VM) {
	// evaluate the index
	n.Index.Traverse(vm)
	// then MLOAD it at the index offset
	n.Expression.Traverse(vm)
	vm.AddBytecode("GET")
}

type SliceExpressionNode struct {
	Expression ExpressionNode
	Low, High  ExpressionNode
	Max        ExpressionNode
}

func (n SliceExpressionNode) Type() NodeType { return SliceExpression }

func (n SliceExpressionNode) Traverse(vm *vmgen.VM) {

}

type CallExpressionNode struct {
	Call      ExpressionNode
	Arguments []ExpressionNode
}

func (n CallExpressionNode) Type() NodeType { return CallExpression }

func (n CallExpressionNode) Traverse(vm *vmgen.VM) {
	// push arguments onto the stack
	for _, a := range n.Arguments {
		a.Type()
		n.Traverse(vm)
		// return values should be on the stack right now
	}

}

type ArrayLiteralNode struct {
	Key  ReferenceNode
	Size ExpressionNode
	Data []ExpressionNode
}

func (n ArrayLiteralNode) Type() NodeType { return ArrayLiteral }

func (n ArrayLiteralNode) Traverse(vm *vmgen.VM) {
	for _, en := range n.Data {
		en.Traverse(vm)
	}
	// push the size of the array
	vm.AddBytecode("PUSH", byte(1), byte(len(n.Data)))
	vm.AddBytecode("ARRAY")
}

type MapLiteralNode struct {
	Key   ReferenceNode
	Value ReferenceNode
	Data  map[ExpressionNode]ExpressionNode
}

func (n MapLiteralNode) Type() NodeType { return MapLiteral }

func (n MapLiteralNode) Traverse(vm *vmgen.VM) {
	for k, v := range n.Data {
		k.Traverse(vm)
		v.Traverse(vm)
	}
	// push the size of the map
	vm.AddBytecode("PUSH", byte(1), byte(len(n.Data)))
	vm.AddBytecode("MAP")
}

type ReferenceNode struct {
	Names []string
}

func (n ReferenceNode) Type() NodeType { return Reference }

func (n ReferenceNode) Traverse(vm *vmgen.VM) {
	// reference e.g. dog.tail.wag()
	// get the object
	// if in storage
	vm.AddBytecode("LOAD")
}
