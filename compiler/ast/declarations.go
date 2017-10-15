package ast

import "github.com/end-r/guardian/compiler/lexer"

type TypeDeclarationNode struct {
	Identifier string
	Value      Node
}

// Type ...
func (n TypeDeclarationNode) Type() NodeType { return TypeDeclaration }

type FuncDeclarationNode struct {
	Identifier string
	Parameters []ExplicitVarDeclarationNode
	Results    []Node
	Body       *ScopeNode
	IsAbstract bool
}

func (n FuncDeclarationNode) Type() NodeType { return FuncDeclaration }

type ClassDeclarationNode struct {
	Identifier   string
	IsAbstract   bool
	Supers       []ReferenceNode
	Interfaces   []ReferenceNode
	Body         *ScopeNode
	declarations map[string][]Node
}

func (n ClassDeclarationNode) Type() NodeType { return ClassDeclaration }

type InterfaceDeclarationNode struct {
	Identifier string
	IsAbstract bool
	Body       *ScopeNode
	Supers     []ReferenceNode
}

func (n InterfaceDeclarationNode) Type() NodeType { return InterfaceDeclaration }

type ContractDeclarationNode struct {
	Identifier string
	IsAbstract bool
	Supers     []ReferenceNode
	Interfaces []ReferenceNode
	Body       *ScopeNode
}

func (n ContractDeclarationNode) Type() NodeType { return ContractDeclaration }

type ExplicitVarDeclarationNode struct {
	Identifiers  []string
	DeclaredType Node
}

func (n ExplicitVarDeclarationNode) Type() NodeType { return ExplicitVarDeclaration }

type EventDeclarationNode struct {
	Identifier string
	Parameters []ReferenceNode
}

func (n EventDeclarationNode) Type() NodeType { return EventDeclaration }

// LifecycleDeclarationNode ...
type LifecycleDeclarationNode struct {
	Category   lexer.TokenType
	Parameters []ExplicitVarDeclarationNode
	Body       *ScopeNode
}

func (n LifecycleDeclarationNode) Type() NodeType { return LifecycleDeclaration }

type EnumDeclarationNode struct {
	Identifier string
	IsAbstract bool
	Inherits   []ReferenceNode
	Body       *ScopeNode
}

func (n EnumDeclarationNode) Type() NodeType { return EnumDeclaration }
