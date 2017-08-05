package ast

import "axia/guardian/go/compiler/lexer"

type ExpressionNode interface {
	Type() NodeType
}

// BinaryExpressionNode ...
type BinaryExpressionNode struct {
	Left, Right ExpressionNode
	Operator    lexer.TokenType
}

// Type ...
func (n BinaryExpressionNode) Type() NodeType { return BinaryExpression }

type UnaryExpressionNode struct {
	Operator lexer.TokenType
	Operand  ExpressionNode
}

func (n UnaryExpressionNode) Type() NodeType { return UnaryExpression }

type LiteralNode struct {
	Data        string
	LiteralType lexer.TokenType
}

func (n LiteralNode) Type() NodeType { return Literal }

type CompositeLiteralNode struct {
	Reference ExpressionNode
}

func (n CompositeLiteralNode) Type() NodeType { return CompositeLiteral }

type IndexExpressionNode struct {
	Expression ExpressionNode
	Index      ExpressionNode
}

func (n IndexExpressionNode) Type() NodeType { return IndexExpression }

type GenericExpressionNode struct {
	Expression Node
}

func (n GenericExpressionNode) Type() NodeType { return GenericExpression }

type SliceExpressionNode struct {
	Expression ExpressionNode
	Low, High  ExpressionNode
	Max        ExpressionNode
}

func (n SliceExpressionNode) Type() NodeType { return SliceExpression }

type CallExpressionNode struct {
	Call      ExpressionNode
	Arguments []ExpressionNode
}

func (n CallExpressionNode) Type() NodeType { return CallExpression }

type ArrayLiteralNode struct {
	Key  ReferenceNode
	Size ExpressionNode
	Data []ExpressionNode
}

func (n ArrayLiteralNode) Type() NodeType { return ArrayLiteral }

type MapLiteralNode struct {
	Key   ReferenceNode
	Value ReferenceNode
	Data  map[ExpressionNode]ExpressionNode
}

func (n MapLiteralNode) Type() NodeType { return MapLiteral }

type ReferenceNode struct {
	Names []string
}

func (n ReferenceNode) Type() NodeType { return Reference }
