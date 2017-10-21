package ast

import "github.com/end-r/guardian/compiler/lexer"

type TypeDeclarationNode struct {
	Modifiers  []lexer.TokenType
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
	Modifiers  []lexer.TokenType
}

func (n FuncDeclarationNode) Type() NodeType { return FuncDeclaration }

type ClassDeclarationNode struct {
	Identifier   string
	Modifiers    []lexer.TokenType
	Supers       []ReferenceNode
	Interfaces   []ReferenceNode
	Body         *ScopeNode
	declarations map[string][]Node
}

func (n ClassDeclarationNode) Type() NodeType { return ClassDeclaration }

type InterfaceDeclarationNode struct {
	Identifier string
	Modifiers  []lexer.TokenType
	Signatures []FuncTypeNode
	Supers     []ReferenceNode
}

func (n InterfaceDeclarationNode) Type() NodeType { return InterfaceDeclaration }

type ContractDeclarationNode struct {
	Identifier string
	Modifiers  []lexer.TokenType
	Supers     []ReferenceNode
	Interfaces []ReferenceNode
	Body       *ScopeNode
}

func (n ContractDeclarationNode) Type() NodeType { return ContractDeclaration }

type ExplicitVarDeclarationNode struct {
	Modifiers    []lexer.TokenType
	Identifiers  []string
	DeclaredType Node
}

func (n ExplicitVarDeclarationNode) Type() NodeType { return ExplicitVarDeclaration }

type EventDeclarationNode struct {
	Modifiers  []lexer.TokenType
	Identifier string
	Parameters []ReferenceNode
}

func (n EventDeclarationNode) Type() NodeType { return EventDeclaration }

// LifecycleDeclarationNode ...
type LifecycleDeclarationNode struct {
	Modifiers  []lexer.TokenType
	Category   lexer.TokenType
	Parameters []ExplicitVarDeclarationNode
	Body       *ScopeNode
}

func (n LifecycleDeclarationNode) Type() NodeType { return LifecycleDeclaration }

type EnumDeclarationNode struct {
	Identifier string
	Modifiers  []lexer.TokenType
	Inherits   []ReferenceNode
	// consider whether to change this
	Enums []string
}

func (n EnumDeclarationNode) Type() NodeType { return EnumDeclaration }
