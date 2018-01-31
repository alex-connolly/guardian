package ast

import (
	"github.com/end-r/guardian/token"
	"github.com/end-r/guardian/typing"
	"github.com/end-r/guardian/util"
)

// BinaryExpressionNode ...
type BinaryExpressionNode struct {
	Begin, Final util.Location
	Left, Right  ExpressionNode
	Operator     token.Type
	Resolved     typing.Type
}

func (n *BinaryExpressionNode) Start() util.Location      { return n.Begin }
func (n *BinaryExpressionNode) End() util.Location        { return n.Final }
func (n *BinaryExpressionNode) Type() NodeType            { return BinaryExpression }
func (n *BinaryExpressionNode) ResolvedType() typing.Type { return n.Resolved }

// UnaryExpressionNode ...
type UnaryExpressionNode struct {
	Begin, Final util.Location
	Operator     token.Type
	Operand      ExpressionNode
	Resolved     typing.Type
}

func (n *UnaryExpressionNode) Start() util.Location      { return n.Begin }
func (n *UnaryExpressionNode) End() util.Location        { return n.Final }
func (n *UnaryExpressionNode) Type() NodeType            { return UnaryExpression }
func (n *UnaryExpressionNode) ResolvedType() typing.Type { return n.Resolved }

type LiteralNode struct {
	Begin, Final util.Location
	Data         string
	LiteralType  token.Type
	Resolved     typing.Type
}

func (n *LiteralNode) Start() util.Location      { return n.Begin }
func (n *LiteralNode) End() util.Location        { return n.Final }
func (n *LiteralNode) Type() NodeType            { return Literal }
func (n *LiteralNode) ResolvedType() typing.Type { return n.Resolved }

func (n *LiteralNode) GetBytes() []byte {
	return []byte(n.Data)
}

type CompositeLiteralNode struct {
	Begin, Final util.Location
	TypeName     *PlainTypeNode
	Fields       map[string]ExpressionNode
	Resolved     typing.Type
}

func (n *CompositeLiteralNode) Start() util.Location      { return n.Begin }
func (n *CompositeLiteralNode) End() util.Location        { return n.Final }
func (n *CompositeLiteralNode) Type() NodeType            { return CompositeLiteral }
func (n *CompositeLiteralNode) ResolvedType() typing.Type { return n.Resolved }

type IndexExpressionNode struct {
	Begin, Final util.Location
	Expression   ExpressionNode
	Index        ExpressionNode
	Resolved     typing.Type
}

func (n *IndexExpressionNode) Start() util.Location      { return n.Begin }
func (n *IndexExpressionNode) End() util.Location        { return n.Final }
func (n *IndexExpressionNode) Type() NodeType            { return IndexExpression }
func (n *IndexExpressionNode) ResolvedType() typing.Type { return n.Resolved }

type SliceExpressionNode struct {
	Begin, Final util.Location
	Expression   ExpressionNode
	Low, High    ExpressionNode
	Max          ExpressionNode
	Resolved     typing.Type
}

func (n *SliceExpressionNode) Start() util.Location      { return n.Begin }
func (n *SliceExpressionNode) End() util.Location        { return n.Final }
func (n *SliceExpressionNode) Type() NodeType            { return SliceExpression }
func (n *SliceExpressionNode) ResolvedType() typing.Type { return n.Resolved }

type CallExpressionNode struct {
	Begin, Final util.Location
	Call         ExpressionNode
	Arguments    []ExpressionNode
	Resolved     typing.Type
}

func (n *CallExpressionNode) Start() util.Location      { return n.Begin }
func (n *CallExpressionNode) End() util.Location        { return n.Final }
func (n *CallExpressionNode) Type() NodeType            { return CallExpression }
func (n *CallExpressionNode) ResolvedType() typing.Type { return n.Resolved }

type ArrayLiteralNode struct {
	Begin, Final util.Location
	Signature    *ArrayTypeNode
	Data         []ExpressionNode
	Resolved     typing.Type
}

func (n *ArrayLiteralNode) Start() util.Location      { return n.Begin }
func (n *ArrayLiteralNode) End() util.Location        { return n.Final }
func (n *ArrayLiteralNode) Type() NodeType            { return ArrayLiteral }
func (n *ArrayLiteralNode) ResolvedType() typing.Type { return n.Resolved }

type MapLiteralNode struct {
	Begin, Final util.Location
	Signature    *MapTypeNode
	Data         map[ExpressionNode]ExpressionNode
	Resolved     typing.Type
}

func (n *MapLiteralNode) Start() util.Location      { return n.Begin }
func (n *MapLiteralNode) End() util.Location        { return n.Final }
func (n *MapLiteralNode) Type() NodeType            { return MapLiteral }
func (n *MapLiteralNode) ResolvedType() typing.Type { return n.Resolved }

type FuncLiteralNode struct {
	Generics     []*GenericDeclarationNode
	Begin, Final util.Location
	Parameters   []*ExplicitVarDeclarationNode
	Results      []Node
	Scope        *ScopeNode
	Resolved     typing.Type
}

func (n *FuncLiteralNode) Start() util.Location      { return n.Begin }
func (n *FuncLiteralNode) End() util.Location        { return n.Final }
func (n *FuncLiteralNode) Type() NodeType            { return FuncLiteral }
func (n *FuncLiteralNode) ResolvedType() typing.Type { return n.Resolved }

type IdentifierNode struct {
	Begin, Final util.Location
	Name         string
	Parameters   []Node
	Resolved     typing.Type
}

func (n *IdentifierNode) Start() util.Location      { return n.Begin }
func (n *IdentifierNode) End() util.Location        { return n.Final }
func (n *IdentifierNode) Type() NodeType            { return Identifier }
func (n *IdentifierNode) ResolvedType() typing.Type { return n.Resolved }

type ReferenceNode struct {
	Begin, Final util.Location
	Parent       ExpressionNode
	Reference    ExpressionNode
	Resolved     typing.Type
}

func (n *ReferenceNode) Start() util.Location      { return n.Begin }
func (n *ReferenceNode) End() util.Location        { return n.Final }
func (n *ReferenceNode) Type() NodeType            { return Reference }
func (n *ReferenceNode) ResolvedType() typing.Type { return n.Resolved }

type KeywordNode struct {
	Begin, Final util.Location
	Resolved     typing.Type
	Keyword      token.Type
	TypeNode     Node
	Arguments    []ExpressionNode
}

func (n *KeywordNode) Start() util.Location      { return n.Begin }
func (n *KeywordNode) End() util.Location        { return n.Final }
func (n *KeywordNode) Type() NodeType            { return Keyword }
func (n *KeywordNode) ResolvedType() typing.Type { return n.Resolved }
