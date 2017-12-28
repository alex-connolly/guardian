package ast

import (
	"github.com/end-r/guardian/token"

	"github.com/end-r/guardian/typing"
)

type TypeDeclarationNode struct {
	Modifiers    Modifiers
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
	Modifiers    Modifiers
	Generics     []*GenericDeclarationNode
	ResolvedSize uint
	Resolved     typing.Type
}

func (n *FuncDeclarationNode) Type() NodeType { return FuncDeclaration }

type ClassDeclarationNode struct {
	Identifier   string
	Modifiers    Modifiers
	Supers       []*PlainTypeNode
	Interfaces   []*PlainTypeNode
	Body         *ScopeNode
	declarations map[string][]Node
	Generics     []*GenericDeclarationNode
	Resolved     typing.Type
}

func (n *ClassDeclarationNode) Type() NodeType { return ClassDeclaration }

type InterfaceDeclarationNode struct {
	Identifier string
	Modifiers  Modifiers
	Signatures []*FuncTypeNode
	Supers     []*PlainTypeNode
	Generics   []*GenericDeclarationNode
	Resolved   typing.Type
}

func (n *InterfaceDeclarationNode) Type() NodeType { return InterfaceDeclaration }

type ContractDeclarationNode struct {
	Identifier string
	Modifiers  Modifiers
	Supers     []*PlainTypeNode
	Interfaces []*PlainTypeNode
	Generics   []*GenericDeclarationNode
	Body       *ScopeNode
	Resolved   typing.Type
}

func (n *ContractDeclarationNode) Type() NodeType { return ContractDeclaration }

type Annotation struct {
	Name       string
	Parameters []ExpressionNode
	Required   []Node
}

type Modifiers struct {
	Annotations []*Annotation
	Modifiers   []string
}

type ExplicitVarDeclarationNode struct {
	Modifiers    Modifiers
	Identifiers  []string
	DeclaredType Node
	Resolved     typing.Type
	IsConstant   bool
	Value        ExpressionNode
}

func (n *ExplicitVarDeclarationNode) Type() NodeType { return ExplicitVarDeclaration }

type EventDeclarationNode struct {
	Modifiers  Modifiers
	Identifier string
	Generics   []*GenericDeclarationNode
	Parameters []*ExplicitVarDeclarationNode
	Resolved   typing.Type
}

func (n *EventDeclarationNode) Type() NodeType { return EventDeclaration }

// LifecycleDeclarationNode ...
type LifecycleDeclarationNode struct {
	Modifiers  Modifiers
	Category   token.Type
	Parameters []*ExplicitVarDeclarationNode
	Body       *ScopeNode
}

func (n *LifecycleDeclarationNode) Type() NodeType { return LifecycleDeclaration }

type EnumDeclarationNode struct {
	Identifier string
	Modifiers  Modifiers
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
