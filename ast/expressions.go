package ast

import (
	"github.com/end-r/guardian/token"

	"github.com/end-r/guardian/typing"
)

// BinaryExpressionNode ...
type BinaryExpressionNode struct {
	Begin, Final uint
	Left, Right  ExpressionNode
	Operator     token.Type
	Resolved     typing.Type
}

func (n *BinaryExpressionNode) Start() uint               { return n.Begin }
func (n *BinaryExpressionNode) End() uint                 { return n.Final }
func (n *BinaryExpressionNode) Type() NodeType            { return BinaryExpression }
func (n *BinaryExpressionNode) ResolvedType() typing.Type { return n.Resolved }

// UnaryExpressionNode ...
type UnaryExpressionNode struct {
	Begin, Final uint
	Operator     token.Type
	Operand      ExpressionNode
	Resolved     typing.Type
}

func (n *UnaryExpressionNode) Start() uint               { return n.Begin }
func (n *UnaryExpressionNode) End() uint                 { return n.Final }
func (n *UnaryExpressionNode) Type() NodeType            { return UnaryExpression }
func (n *UnaryExpressionNode) ResolvedType() typing.Type { return n.Resolved }

type LiteralNode struct {
	Begin, Final uint
	Data         string
	LiteralType  token.Type
	Resolved     typing.Type
}

func (n *LiteralNode) Start() uint               { return n.Begin }
func (n *LiteralNode) End() uint                 { return n.Final }
func (n *LiteralNode) Type() NodeType            { return Literal }
func (n *LiteralNode) ResolvedType() typing.Type { return n.Resolved }

func (n *LiteralNode) GetBytes() []byte {
	return []byte(n.Data)
}

type CompositeLiteralNode struct {
	Begin, Final uint
	TypeName     *PlainTypeNode
	Fields       map[string]ExpressionNode
	Resolved     typing.Type
}

func (n *CompositeLiteralNode) Start() uint               { return n.Begin }
func (n *CompositeLiteralNode) End() uint                 { return n.Final }
func (n *CompositeLiteralNode) Type() NodeType            { return CompositeLiteral }
func (n *CompositeLiteralNode) ResolvedType() typing.Type { return n.Resolved }

type IndexExpressionNode struct {
	Begin, Final uint
	Expression   ExpressionNode
	Index        ExpressionNode
	Resolved     typing.Type
}

func (n *IndexExpressionNode) Start() uint               { return n.Begin }
func (n *IndexExpressionNode) End() uint                 { return n.Final }
func (n *IndexExpressionNode) Type() NodeType            { return IndexExpression }
func (n *IndexExpressionNode) ResolvedType() typing.Type { return n.Resolved }

type SliceExpressionNode struct {
	Begin, Final uint
	Expression   ExpressionNode
	Low, High    ExpressionNode
	Max          ExpressionNode
	Resolved     typing.Type
}

func (n *SliceExpressionNode) Start() uint               { return n.Begin }
func (n *SliceExpressionNode) End() uint                 { return n.Final }
func (n *SliceExpressionNode) Type() NodeType            { return SliceExpression }
func (n *SliceExpressionNode) ResolvedType() typing.Type { return n.Resolved }

type CallExpressionNode struct {
	Begin, Final uint
	Call         ExpressionNode
	Arguments    []ExpressionNode
	Resolved     typing.Type
}

func (n *CallExpressionNode) Start() uint               { return n.Begin }
func (n *CallExpressionNode) End() uint                 { return n.Final }
func (n *CallExpressionNode) Type() NodeType            { return CallExpression }
func (n *CallExpressionNode) ResolvedType() typing.Type { return n.Resolved }

type ArrayLiteralNode struct {
	Begin, Final uint
	Signature    *ArrayTypeNode
	Data         []ExpressionNode
	Resolved     typing.Type
}

func (n *ArrayLiteralNode) Start() uint               { return n.Begin }
func (n *ArrayLiteralNode) End() uint                 { return n.Final }
func (n *ArrayLiteralNode) Type() NodeType            { return ArrayLiteral }
func (n *ArrayLiteralNode) ResolvedType() typing.Type { return n.Resolved }

type MapLiteralNode struct {
	Begin, Final uint
	Signature    *MapTypeNode
	Data         map[ExpressionNode]ExpressionNode
	Resolved     typing.Type
}

func (n *MapLiteralNode) Start() uint               { return n.Begin }
func (n *MapLiteralNode) End() uint                 { return n.Final }
func (n *MapLiteralNode) Type() NodeType            { return MapLiteral }
func (n *MapLiteralNode) ResolvedType() typing.Type { return n.Resolved }

type FuncLiteralNode struct {
	Begin, Final uint
	Parameters   []*ExplicitVarDeclarationNode
	Results      []Node
	Scope        *ScopeNode
	Resolved     typing.Type
}

func (n *FuncLiteralNode) Start() uint               { return n.Begin }
func (n *FuncLiteralNode) End() uint                 { return n.Final }
func (n *FuncLiteralNode) Type() NodeType            { return FuncLiteral }
func (n *FuncLiteralNode) ResolvedType() typing.Type { return n.Resolved }

type IdentifierNode struct {
	Begin, Final uint
	Name         string
	Parameters   []Node
	Resolved     typing.Type
}

func (n *IdentifierNode) Start() uint               { return n.Begin }
func (n *IdentifierNode) End() uint                 { return n.Final }
func (n *IdentifierNode) Type() NodeType            { return Identifier }
func (n *IdentifierNode) ResolvedType() typing.Type { return n.Resolved }

type ReferenceNode struct {
	Begin, Final uint
	Parent       ExpressionNode
	Reference    ExpressionNode
	Resolved     typing.Type
}

func (n *ReferenceNode) Start() uint               { return n.Begin }
func (n *ReferenceNode) End() uint                 { return n.Final }
func (n *ReferenceNode) Type() NodeType            { return Reference }
func (n *ReferenceNode) ResolvedType() typing.Type { return n.Resolved }

type KeywordNode struct {
	Begin, Final uint
	Resolved     typing.Type
	Keyword      token.Type
	TypeNode     Node
	Arguments    []ExpressionNode
}

func (n *KeywordNode) Start() uint               { return n.Begin }
func (n *KeywordNode) End() uint                 { return n.Final }
func (n *KeywordNode) Type() NodeType            { return Keyword }
func (n *KeywordNode) ResolvedType() typing.Type { return n.Resolved }
