package ast

import (
	"github.com/end-r/guardian/lexer"
	"github.com/end-r/guardian/typing"
)

// BinaryExpressionNode ...
type BinaryExpressionNode struct {
	Left, Right ExpressionNode
	Operator    lexer.TokenType
	Resolved    typing.Type
}

// Type ...
func (n BinaryExpressionNode) Type() NodeType { return BinaryExpression }

// UnaryExpressionNode ...
type UnaryExpressionNode struct {
	Operator lexer.TokenType
	Operand  ExpressionNode
	Resolved core.Type
}

func (n UnaryExpressionNode) Type() NodeType { return UnaryExpression }

type LiteralNode struct {
	Data        string
	LiteralType lexer.TokenType
	Resolved    core.Type
}

func (n LiteralNode) Type() NodeType { return Literal }

func (n LiteralNode) GetBytes() []byte {
	return nil
}

type CompositeLiteralNode struct {
	TypeName string
	Fields   map[string]ExpressionNode
	Resolved core.Type
}

func (n CompositeLiteralNode) Type() NodeType { return CompositeLiteral }

type IndexExpressionNode struct {
	Expression ExpressionNode
	Index      ExpressionNode
	Resolved   core.Type
}

func (n IndexExpressionNode) Type() NodeType { return IndexExpression }

type SliceExpressionNode struct {
	Expression ExpressionNode
	Low, High  ExpressionNode
	Max        ExpressionNode
	Resolved   core.Type
}

func (n SliceExpressionNode) Type() NodeType { return SliceExpression }

type CallExpressionNode struct {
	Call      ExpressionNode
	Arguments []ExpressionNode
	Resolved  core.Type
}

func (n CallExpressionNode) Type() NodeType { return CallExpression }

type ArrayLiteralNode struct {
	Signature ArrayTypeNode
	Data      []ExpressionNode
	Resolved  core.Type
}

func (n ArrayLiteralNode) Type() NodeType { return ArrayLiteral }

type MapLiteralNode struct {
	Signature MapTypeNode
	Data      map[ExpressionNode]ExpressionNode
	Resolved  core.Type
}

func (n MapLiteralNode) Type() NodeType { return MapLiteral }

type FuncLiteralNode struct {
	Parameters []ExplicitVarDeclarationNode
	Results    []Node
	Scope      *ScopeNode
	Resolved   core.Type
}

// Type ...
func (n FuncLiteralNode) Type() NodeType { return FuncLiteral }

type IdentifierNode struct {
	Name     string
	Resolved core.Type
}

func (n IdentifierNode) Type() NodeType { return Identifier }

type ReferenceNode struct {
	Parent    ExpressionNode
	Reference ExpressionNode
	Resolved  core.Type
}

func (n ReferenceNode) Type() NodeType { return Reference }
