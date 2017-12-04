package ast

import (
	"github.com/end-r/guardian/lexer"
	"github.com/end-r/guardian/util"
)

type TypeDeclarationNode struct {
	Modifiers    []lexer.TokenType
	Identifier   string
	Value        Node
	ResolvedType util.Type
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
	Supers       []PlainTypeNode
	Interfaces   []PlainTypeNode
	Body         *ScopeNode
	declarations map[string][]Node
	ResolvedType util.Type
}

func (n ClassDeclarationNode) Type() NodeType { return ClassDeclaration }

type InterfaceDeclarationNode struct {
	Identifier   string
	Modifiers    []lexer.TokenType
	Signatures   []FuncTypeNode
	Supers       []PlainTypeNode
	ResolvedType util.Type
}

func (n InterfaceDeclarationNode) Type() NodeType { return InterfaceDeclaration }

type ContractDeclarationNode struct {
	Identifier   string
	Modifiers    []lexer.TokenType
	Supers       []PlainTypeNode
	Interfaces   []PlainTypeNode
	Body         *ScopeNode
	ResolvedType util.Type
}

func (n ContractDeclarationNode) Type() NodeType { return ContractDeclaration }

type ExplicitVarDeclarationNode struct {
	Modifiers    []lexer.TokenType
	Identifiers  []string
	DeclaredType Node
	ResolvedType util.Type
}

func (n ExplicitVarDeclarationNode) Type() NodeType { return ExplicitVarDeclaration }

type EventDeclarationNode struct {
	Modifiers    []lexer.TokenType
	Identifier   string
	Parameters   []ExplicitVarDeclarationNode
	ResolvedType util.Type
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
	Enums        []string
	ResolvedType util.Type
}

func (n EnumDeclarationNode) Type() NodeType { return EnumDeclaration }
