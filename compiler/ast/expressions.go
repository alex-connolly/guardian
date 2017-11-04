package ast

import (
	"github.com/end-r/guardian/compiler/lexer"
)

// BinaryExpressionNode ...
type BinaryExpressionNode struct {
	Left, Right ExpressionNode
	Operator    lexer.TokenType
}

// Type ...
func (n BinaryExpressionNode) Type() NodeType { return BinaryExpression }

// UnaryExpressionNode ...
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

func (n LiteralNode) GetBytes() []byte {
	return nil
}

type CompositeLiteralNode struct {
	Reference ExpressionNode
	Fields    map[string]ExpressionNode
}

func (n CompositeLiteralNode) Type() NodeType { return CompositeLiteral }

type IndexExpressionNode struct {
	Expression ExpressionNode
	Index      ExpressionNode
}

func (n IndexExpressionNode) Type() NodeType { return IndexExpression }

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
	Signature ArrayTypeNode
	Data      []ExpressionNode
}

func (n ArrayLiteralNode) Type() NodeType { return ArrayLiteral }

type MapLiteralNode struct {
	Signature MapTypeNode
	Data      map[ExpressionNode]ExpressionNode
}

func (n MapLiteralNode) Type() NodeType { return MapLiteral }

type FuncLiteralNode struct {
	Signature FuncTypeNode
	Scope     *ScopeNode
}

// Type ...
func (n FuncLiteralNode) Type() NodeType { return FuncLiteral }

type IdentifierNode struct {
	Name string
}

func (n IdentifierNode) Type() NodeType { return Identifier }

type ReferenceNode struct {
	Parent    ExpressionNode
	Reference ExpressionNode
}

func (n ReferenceNode) Type() NodeType { return Reference }
