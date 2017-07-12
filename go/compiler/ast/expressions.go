package ast

import "axia/guardian/go/compiler/lexer"

// BinaryExpressionNode ...
type BinaryExpressionNode struct {
	Left, Right Node
	Operator    lexer.TokenType
}

// Type ...
func (n BinaryExpressionNode) Type() NodeType { return BinaryExpression }

func (n BinaryExpressionNode) Validate(t NodeType) bool {
	return true
}

func (n BinaryExpressionNode) Declare(key string, node Node) {

}

type UnaryExpressionNode struct {
	Operator lexer.TokenType
	Operand  Node
}

func (n UnaryExpressionNode) Type() NodeType { return UnaryExpression }

func (n UnaryExpressionNode) Validate(t NodeType) bool {
	return true
}

func (n UnaryExpressionNode) Declare(key string, node Node) {}

type LiteralNode struct {
	LiteralType lexer.TokenType
}

func (n LiteralNode) Type() NodeType { return Literal }

func (n LiteralNode) Validate(t NodeType) bool {
	return true
}

func (n LiteralNode) Declare(key string, node Node) {

}

type CompositeLiteralNode struct {
}

func (n CompositeLiteralNode) Type() NodeType { return CompositeLiteral }

func (n CompositeLiteralNode) Validate(t NodeType) bool {
	return true
}

func (n CompositeLiteralNode) Declare(key string, node Node) {

}

type IndexExpressionNode struct {
	Expression Node
	Index      Node
}

func (n IndexExpressionNode) Type() NodeType { return IndexExpression }

func (n IndexExpressionNode) Validate(t NodeType) bool {
	return true
}

func (n IndexExpressionNode) Declare(key string, node Node) {

}

type GenericExpressionNode struct {
	Expression Node
}

func (n GenericExpressionNode) Type() NodeType { return GenericExpression }

func (n GenericExpressionNode) Validate(t NodeType) bool {
	return true
}

func (n GenericExpressionNode) Declare(key string, node Node) {

}

type SliceExpressionNode struct {
	Expression Node
	Low, High  Node
	Max        Node
}

func (n SliceExpressionNode) Type() NodeType { return SliceExpression }

func (n SliceExpressionNode) Validate(t NodeType) bool {
	return true
}

func (n SliceExpressionNode) Declare(key string, node Node) {

}

type CallExpressionNode struct {
	Name      string
	Arguments []Node
}

func (n CallExpressionNode) Type() NodeType { return CallExpression }

func (n CallExpressionNode) Validate(t NodeType) bool {
	return true
}

func (n CallExpressionNode) Declare(key string, node Node) {

}

type ArrayLiteralNode struct {
	Key  Node
	Data []Node
}

func (n ArrayLiteralNode) Type() NodeType { return ArrayLiteral }

func (n ArrayLiteralNode) Validate(t NodeType) bool {
	return true
}

func (n ArrayLiteralNode) Declare(key string, node Node) {

}

type MapLiteralNode struct {
	Key   Node
	Value Node
	Data  map[string]Node
}

func (n MapLiteralNode) Type() NodeType { return MapLiteral }

func (n MapLiteralNode) Validate(t NodeType) bool {
	return true
}

func (n MapLiteralNode) Declare(key string, node Node) {

}

type ReferenceNode struct {
	Name string
}

func (n ReferenceNode) Type() NodeType { return Reference }

func (n ReferenceNode) Validate(t NodeType) bool {
	return false
}

func (n ReferenceNode) Declare(key string, node Node) {

}
