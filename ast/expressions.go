package ast

import (
	"github.com/end-r/guardian/token"

	"github.com/end-r/guardian/typing"
)

// BinaryExpressionNode ...
type BinaryExpressionNode struct {
	Left, Right ExpressionNode
	Operator    token.Type
	Resolved    typing.Type
}

// Type ...
func (n *BinaryExpressionNode) Type() NodeType { return BinaryExpression }

// UnaryExpressionNode ...
type UnaryExpressionNode struct {
	Operator token.Type
	Operand  ExpressionNode
	Resolved typing.Type
}

func (n *UnaryExpressionNode) Type() NodeType { return UnaryExpression }

type LiteralNode struct {
	Data        string
	LiteralType token.Type
	Resolved    typing.Type
}

func (n *LiteralNode) Type() NodeType { return Literal }

func (n *LiteralNode) GetBytes() []byte {
	return []byte(n.Data)
}

type CompositeLiteralNode struct {
	TypeName string
	Fields   map[string]ExpressionNode
	Resolved typing.Type
}

func (n *CompositeLiteralNode) Type() NodeType { return CompositeLiteral }

type IndexExpressionNode struct {
	Expression ExpressionNode
	Index      ExpressionNode
	Resolved   typing.Type
}

func (n *IndexExpressionNode) Type() NodeType { return IndexExpression }

type SliceExpressionNode struct {
	Expression ExpressionNode
	Low, High  ExpressionNode
	Max        ExpressionNode
	Resolved   typing.Type
}

func (n *SliceExpressionNode) Type() NodeType { return SliceExpression }

type CallExpressionNode struct {
	Call      ExpressionNode
	Arguments []ExpressionNode
	Resolved  typing.Type
}

func (n *CallExpressionNode) Type() NodeType { return CallExpression }

type ArrayLiteralNode struct {
	Signature *ArrayTypeNode
	Data      []ExpressionNode
	Resolved  typing.Type
}

func (n *ArrayLiteralNode) Type() NodeType { return ArrayLiteral }

type MapLiteralNode struct {
	Signature *MapTypeNode
	Data      map[ExpressionNode]ExpressionNode
	Resolved  typing.Type
}

func (n *MapLiteralNode) Type() NodeType { return MapLiteral }

type FuncLiteralNode struct {
	Parameters []*ExplicitVarDeclarationNode
	Results    []Node
	Scope      *ScopeNode
	Resolved   typing.Type
}

// Type ...
func (n *FuncLiteralNode) Type() NodeType { return FuncLiteral }

type IdentifierNode struct {
	Name     string
	Resolved typing.Type
}

func (n *IdentifierNode) Type() NodeType { return Identifier }

type ReferenceNode struct {
	Parent    ExpressionNode
	Reference ExpressionNode
	Resolved  typing.Type
}

func (n *ReferenceNode) Type() NodeType { return Reference }
