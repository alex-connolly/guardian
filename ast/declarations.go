package ast

import (
	"github.com/end-r/guardian/token"
	"github.com/end-r/guardian/typing"
	"github.com/end-r/guardian/util"
)

type TypeDeclarationNode struct {
	Begin, Final util.Location
	Modifiers    typing.Modifiers
	Identifier   string
	Value        Node
	ResolvedSize uint
	Resolved     typing.Type
}

// Type ...
func (n *TypeDeclarationNode) Type() NodeType       { return TypeDeclaration }
func (n *TypeDeclarationNode) Start() util.Location { return n.Begin }
func (n *TypeDeclarationNode) End() util.Location   { return n.Final }

type FuncDeclarationNode struct {
	Begin, Final util.Location
	Signature    *FuncTypeNode
	Body         *ScopeNode
	Modifiers    typing.Modifiers
	Generics     []*GenericDeclarationNode
	ResolvedSize uint
	Resolved     typing.Type
}

func (n *FuncDeclarationNode) Type() NodeType       { return FuncDeclaration }
func (n *FuncDeclarationNode) Start() util.Location { return n.Begin }
func (n *FuncDeclarationNode) End() util.Location   { return n.Final }

type ClassDeclarationNode struct {
	Begin, Final util.Location
	Identifier   string
	Modifiers    typing.Modifiers
	Supers       []*PlainTypeNode
	Interfaces   []*PlainTypeNode
	Body         *ScopeNode
	declarations map[string][]Node
	Generics     []*GenericDeclarationNode
	Resolved     typing.Type
}

func (n *ClassDeclarationNode) Type() NodeType       { return ClassDeclaration }
func (n *ClassDeclarationNode) Start() util.Location { return n.Begin }
func (n *ClassDeclarationNode) End() util.Location   { return n.Final }

type InterfaceDeclarationNode struct {
	Begin, Final util.Location
	Identifier   string
	Modifiers    typing.Modifiers
	Signatures   []*FuncTypeNode
	Supers       []*PlainTypeNode
	Generics     []*GenericDeclarationNode
	Resolved     typing.Type
}

func (n *InterfaceDeclarationNode) Type() NodeType       { return InterfaceDeclaration }
func (n *InterfaceDeclarationNode) Start() util.Location { return n.Begin }
func (n *InterfaceDeclarationNode) End() util.Location   { return n.Final }

type ContractDeclarationNode struct {
	Begin, Final util.Location
	Identifier   string
	Modifiers    typing.Modifiers
	Supers       []*PlainTypeNode
	Interfaces   []*PlainTypeNode
	Generics     []*GenericDeclarationNode
	Body         *ScopeNode
	Resolved     typing.Type
}

func (n *ContractDeclarationNode) Type() NodeType       { return ContractDeclaration }
func (n *ContractDeclarationNode) Start() util.Location { return n.Begin }
func (n *ContractDeclarationNode) End() util.Location   { return n.Final }

type ExplicitVarDeclarationNode struct {
	Begin, Final util.Location
	Modifiers    typing.Modifiers
	Identifiers  []string
	DeclaredType Node
	Resolved     typing.Type
	IsConstant   bool
	Value        ExpressionNode
}

func (n *ExplicitVarDeclarationNode) Type() NodeType       { return ExplicitVarDeclaration }
func (n *ExplicitVarDeclarationNode) Start() util.Location { return n.Begin }
func (n *ExplicitVarDeclarationNode) End() util.Location   { return n.Final }

type EventDeclarationNode struct {
	Begin, Final util.Location
	Modifiers    typing.Modifiers
	Identifier   string
	Generics     []*GenericDeclarationNode
	Parameters   []*ExplicitVarDeclarationNode
	Resolved     typing.Type
}

func (n *EventDeclarationNode) Type() NodeType       { return EventDeclaration }
func (n *EventDeclarationNode) Start() util.Location { return n.Begin }
func (n *EventDeclarationNode) End() util.Location   { return n.Final }

// LifecycleDeclarationNode ...
type LifecycleDeclarationNode struct {
	Begin, Final util.Location
	Modifiers    typing.Modifiers
	Category     token.Type
	Parameters   []*ExplicitVarDeclarationNode
	Body         *ScopeNode
}

func (n *LifecycleDeclarationNode) Type() NodeType       { return LifecycleDeclaration }
func (n *LifecycleDeclarationNode) Start() util.Location { return n.Begin }
func (n *LifecycleDeclarationNode) End() util.Location   { return n.Final }

type EnumDeclarationNode struct {
	Begin, Final util.Location
	Identifier   string
	Modifiers    typing.Modifiers
	Inherits     []*PlainTypeNode
	// consider whether to change this
	Enums    []string
	Resolved typing.Type
}

func (n *EnumDeclarationNode) Type() NodeType       { return EnumDeclaration }
func (n *EnumDeclarationNode) Start() util.Location { return n.Begin }
func (n *EnumDeclarationNode) End() util.Location   { return n.Final }

type GenericDeclarationNode struct {
	Begin, Final util.Location
	Identifier   string
	Inherits     []*PlainTypeNode
	Implements   []*PlainTypeNode
}

func (n *GenericDeclarationNode) Type() NodeType       { return GenericDeclaration }
func (n *GenericDeclarationNode) Start() util.Location { return n.Begin }
func (n *GenericDeclarationNode) End() util.Location   { return n.Final }
