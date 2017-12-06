package ast

import (
	"github.com/end-r/guardian/lexer"
	"github.com/end-r/guardian/typing"
)

type TypeDeclarationNode struct {
	Modifiers    []lexer.TokenType
	Identifier   string
	Value        Node
	ResolvedSize uint
	Resolved     typing.Type
}

// Type ...
func (n TypeDeclarationNode) Type() NodeType { return TypeDeclaration }

type FuncDeclarationNode struct {
	Identifier   string
	Parameters   []ExplicitVarDeclarationNode
	Results      []Node
	Body         *ScopeNode
	Modifiers    []lexer.TokenType
	ResolvedSize uint
	Resolved     typing.Type
}

func (n FuncDeclarationNode) Type() NodeType { return FuncDeclaration }

type ClassDeclarationNode struct {
	Identifier   string
	Modifiers    []lexer.TokenType
	Supers       []PlainTypeNode
	Interfaces   []PlainTypeNode
	Body         *ScopeNode
	declarations map[string][]Node
	Resolved     typing.Type
}

func (n ClassDeclarationNode) Type() NodeType { return ClassDeclaration }

type InterfaceDeclarationNode struct {
	Identifier string
	Modifiers  []lexer.TokenType
	Signatures []FuncTypeNode
	Supers     []PlainTypeNode
	Resolved   typing.Type
}

func (n InterfaceDeclarationNode) Type() NodeType { return InterfaceDeclaration }

type ContractDeclarationNode struct {
	Identifier string
	Modifiers  []lexer.TokenType
	Supers     []PlainTypeNode
	Interfaces []PlainTypeNode
	Body       *ScopeNode
	Resolved   typing.Type
}

func (n ContractDeclarationNode) Type() NodeType { return ContractDeclaration }

type ExplicitVarDeclarationNode struct {
	Modifiers    []lexer.TokenType
	Identifiers  []string
	DeclaredType Node
	Resolved     typing.Type
}

func (n ExplicitVarDeclarationNode) Type() NodeType { return ExplicitVarDeclaration }

type EventDeclarationNode struct {
	Modifiers  []lexer.TokenType
	Identifier string
	Parameters []ExplicitVarDeclarationNode
	Resolved   typing.Type
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
	Inherits   []PlainTypeNode
	// consider whether to change this
	Enums    []string
	Resolved typing.Type
}

func (n EnumDeclarationNode) Type() NodeType { return EnumDeclaration }
