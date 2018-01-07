package ast

import (
	"go/ast"

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

func (n *BinaryExpressionNode) ResolvedType() typing.Type { return n.Resolved }

// UnaryExpressionNode ...
type UnaryExpressionNode struct {
	Operator token.Type
	Operand  ExpressionNode
	Resolved typing.Type
}

func (n *UnaryExpressionNode) Type() NodeType { return UnaryExpression }

func (n *UnaryExpressionNode) ResolvedType() typing.Type { return n.Resolved }

type LiteralNode struct {
	Data        string
	LiteralType token.Type
	Resolved    typing.Type
}

func (n *LiteralNode) Type() NodeType { return Literal }

func (n *LiteralNode) ResolvedType() typing.Type { return n.Resolved }

func (n *LiteralNode) GetBytes() []byte {
	return []byte(n.Data)
}

type CompositeLiteralNode struct {
	TypeName string
	Fields   map[string]ExpressionNode
	Resolved typing.Type
}

func (n *CompositeLiteralNode) Type() NodeType { return CompositeLiteral }

func (n *CompositeLiteralNode) ResolvedType() typing.Type { return n.Resolved }

type IndexExpressionNode struct {
	Expression ExpressionNode
	Index      ExpressionNode
	Resolved   typing.Type
}

func (n *IndexExpressionNode) Type() NodeType { return IndexExpression }

func (n *IndexExpressionNode) ResolvedType() typing.Type { return n.Resolved }

type SliceExpressionNode struct {
	Expression ExpressionNode
	Low, High  ExpressionNode
	Max        ExpressionNode
	Resolved   typing.Type
}

func (n *SliceExpressionNode) Type() NodeType { return SliceExpression }

func (n *SliceExpressionNode) ResolvedType() typing.Type { return n.Resolved }

type CallExpressionNode struct {
	Call      ExpressionNode
	Arguments []ExpressionNode
	Resolved  typing.Type
}

func (n *CallExpressionNode) Type() NodeType { return CallExpression }

func (n *CallExpressionNode) ResolvedType() typing.Type { return n.Resolved }

type ArrayLiteralNode struct {
	Signature *ArrayTypeNode
	Data      []ExpressionNode
	Resolved  typing.Type
}

func (n *ArrayLiteralNode) Type() NodeType { return ArrayLiteral }

func (n *ArrayLiteralNode) ResolvedType() typing.Type { return n.Resolved }

type MapLiteralNode struct {
	Signature *MapTypeNode
	Data      map[ExpressionNode]ExpressionNode
	Resolved  typing.Type
}

func (n *MapLiteralNode) Type() NodeType { return MapLiteral }

func (n *MapLiteralNode) ResolvedType() typing.Type { return n.Resolved }

type FuncLiteralNode struct {
	Parameters []*ExplicitVarDeclarationNode
	Results    []Node
	Scope      *ScopeNode
	Resolved   typing.Type
}

// Type ...
func (n *FuncLiteralNode) Type() NodeType { return FuncLiteral }

func (n *FuncLiteralNode) ResolvedType() typing.Type { return n.Resolved }

type IdentifierNode struct {
	Name       string
	Parameters []Node
	Resolved   typing.Type
}

func (n *IdentifierNode) Type() NodeType { return Identifier }

func (n *IdentifierNode) ResolvedType() typing.Type { return n.Resolved }

type ReferenceNode struct {
	Parent    ExpressionNode
	Reference ExpressionNode
	Resolved  typing.Type
}

func (n *ReferenceNode) Type() NodeType { return Reference }

func (n *ReferenceNode) ResolvedType() typing.Type { return n.Resolved }

type KeywordNode struct {
	Resolved typing.Type
	Keyword  token.Type
	TypeNode ast.Node
}

func (n *KeywordNode) Type() NodeType { return Keyword }

func (n *KeywordNode) ResolvedType() typing.Type { return n.Resolved }
