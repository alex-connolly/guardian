package ast

import (
	"github.com/end-r/guardian/token"

	"github.com/end-r/guardian/typing"
)

type TypeDeclarationNode struct {
	Modifiers    []token.Type
	Identifier   string
	Value        Node
	ResolvedSize uint
	Resolved     typing.Type
}

// Type ...
func (n *TypeDeclarationNode) Type() NodeType { return TypeDeclaration }

type FuncDeclarationNode struct {
	Signature    *FuncTypeNode
	Body         *ScopeNode
	Modifiers    []token.Type
	ResolvedSize uint
	Resolved     typing.Type
}

func (n *FuncDeclarationNode) Type() NodeType { return FuncDeclaration }

type ClassDeclarationNode struct {
	Identifier   string
	Modifiers    []token.Type
	Supers       []*PlainTypeNode
	Interfaces   []*PlainTypeNode
	Body         *ScopeNode
	declarations map[string][]Node
	Resolved     typing.Type
}

func (n *ClassDeclarationNode) Type() NodeType { return ClassDeclaration }

type InterfaceDeclarationNode struct {
	Identifier string
	Modifiers  []token.Type
	Signatures []*FuncTypeNode
	Supers     []*PlainTypeNode
	Resolved   typing.Type
}

func (n *InterfaceDeclarationNode) Type() NodeType { return InterfaceDeclaration }

type ContractDeclarationNode struct {
	Identifier string
	Modifiers  []token.Type
	Supers     []*PlainTypeNode
	Interfaces []*PlainTypeNode
	Body       *ScopeNode
	Resolved   typing.Type
}

func (n *ContractDeclarationNode) Type() NodeType { return ContractDeclaration }

type ExplicitVarDeclarationNode struct {
	Modifiers    []token.Type
	Identifiers  []string
	DeclaredType Node
	Resolved     typing.Type
}

func (n *ExplicitVarDeclarationNode) Type() NodeType { return ExplicitVarDeclaration }

type EventDeclarationNode struct {
	Modifiers  []token.Type
	Identifier string
	Parameters []*ExplicitVarDeclarationNode
	Resolved   typing.Type
}

func (n *EventDeclarationNode) Type() NodeType { return EventDeclaration }

// LifecycleDeclarationNode ...
type LifecycleDeclarationNode struct {
	Modifiers  []token.Type
	Category   token.Type
	Parameters []*ExplicitVarDeclarationNode
	Body       *ScopeNode
}

func (n *LifecycleDeclarationNode) Type() NodeType { return LifecycleDeclaration }

type EnumDeclarationNode struct {
	Identifier string
	Modifiers  []token.Type
	Inherits   []*PlainTypeNode
	// consider whether to change this
	Enums    []string
	Resolved typing.Type
}

func (n *EnumDeclarationNode) Type() NodeType { return EnumDeclaration }

type GenericDeclarationNode struct {
	Identifier string
	Inherits   []*PlainTypeNode
	Implements []*PlainTypeNode
}
