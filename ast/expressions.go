package ast

import (
	"github.com/end-r/guardian/lexer"
	"github.com/end-r/guardian/util"
)

// BinaryExpressionNode ...
type BinaryExpressionNode struct {
	Left, Right ExpressionNode
	Operator    lexer.TokenType
	Resolved    util.Type
}

// Type ...
func (n BinaryExpressionNode) Type() NodeType { return BinaryExpression }

// UnaryExpressionNode ...
type UnaryExpressionNode struct {
	Operator lexer.TokenType
	Operand  ExpressionNode
	Resolved util.Type
}

func (n UnaryExpressionNode) Type() NodeType { return UnaryExpression }

type LiteralNode struct {
	Data        string
	LiteralType lexer.TokenType
	Resolved    util.Type
}

func (n LiteralNode) Type() NodeType { return Literal }

func (n LiteralNode) GetBytes() []byte {
	return nil
}

type CompositeLiteralNode struct {
	TypeName string
	Fields   map[string]ExpressionNode
	Resolved util.Type
}

func (n CompositeLiteralNode) Type() NodeType { return CompositeLiteral }

type IndexExpressionNode struct {
	Expression ExpressionNode
	Index      ExpressionNode
	Resolved   util.Type
}

func (n IndexExpressionNode) Type() NodeType { return IndexExpression }

type SliceExpressionNode struct {
	Expression ExpressionNode
	Low, High  ExpressionNode
	Max        ExpressionNode
	Resolved   util.Type
}

func (n SliceExpressionNode) Type() NodeType { return SliceExpression }

type CallExpressionNode struct {
	Call      ExpressionNode
	Arguments []ExpressionNode
	Resolved  util.Type
}

func (n CallExpressionNode) Type() NodeType { return CallExpression }

type ArrayLiteralNode struct {
	Signature ArrayTypeNode
	Data      []ExpressionNode
	Resolved  util.Type
}

func (n ArrayLiteralNode) Type() NodeType { return ArrayLiteral }

type MapLiteralNode struct {
	Signature MapTypeNode
	Data      map[ExpressionNode]ExpressionNode
	Resolved  util.Type
}

func (n MapLiteralNode) Type() NodeType { return MapLiteral }

type FuncLiteralNode struct {
	Parameters []ExplicitVarDeclarationNode
	Results    []Node
	Scope      *ScopeNode
	Resolved   util.Type
}

// Type ...
func (n FuncLiteralNode) Type() NodeType { return FuncLiteral }

type IdentifierNode struct {
	Name     string
	Resolved util.Type
}

func (n IdentifierNode) Type() NodeType { return Identifier }

type ReferenceNode struct {
	Parent    ExpressionNode
	Reference ExpressionNode
	Resolved  util.Type
}

func (n ReferenceNode) Type() NodeType { return Reference }
