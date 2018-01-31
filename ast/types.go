package ast

import (
	"github.com/end-r/guardian/typing"
	"github.com/end-r/guardian/util"
)

// NodeType denotes the type of node
type NodeType int

const (
	ContractDeclaration NodeType = iota
	ClassDeclaration
	FuncDeclaration
	LifecycleDeclaration
	EnumDeclaration
	InterfaceDeclaration
	TypeDeclaration
	EventDeclaration
	ExplicitVarDeclaration
	GenericDeclaration
	ArrayType
	MapType
	FuncType
	PlainType
	Identifier
	Reference
	Literal
	CompositeLiteral
	MapLiteral
	ArrayLiteral
	FuncLiteral
	Keyword
	BinaryExpression
	UnaryExpression
	SliceExpression
	IndexExpression
	CallExpression
	AssignmentStatement
	ReturnStatement
	BranchStatement
	IfStatement
	Condition
	SwitchStatement
	CaseStatement
	BlockStatement
	ForStatement
	ForEachStatement
	FlowStatement
	ImportStatement
	PackageStatement
	File
	Package
	Scope
)

var (
	AllDeclarations = []NodeType{
		ClassDeclaration, EnumDeclaration, InterfaceDeclaration,
		FuncDeclaration, LifecycleDeclaration, EventDeclaration,
		ContractDeclaration, TypeDeclaration, ExplicitVarDeclaration,
	}

	AllStatements = []NodeType{
		ForStatement, IfStatement, PackageStatement, ImportStatement,
		SwitchStatement, ForEachStatement,
	}

	AllExpressions = []NodeType{
		UnaryExpression, BinaryExpression, SliceExpression, CallExpression, IndexExpression,
		MapLiteral, ArrayLiteral, CompositeLiteral, Literal, Reference,
	}
)

type MapTypeNode struct {
	Begin, Final util.Location
	Variable     bool
	Key          Node
	Value        Node
}

func (n *MapTypeNode) Start() util.Location      { return n.Begin }
func (n *MapTypeNode) End() util.Location        { return n.Final }
func (n *MapTypeNode) Type() NodeType            { return MapType }
func (n *MapTypeNode) ResolvedType() typing.Type { return typing.Unknown() }

type ArrayTypeNode struct {
	Begin, Final util.Location
	Variable     bool
	Length       int
	Value        Node
}

func (n *ArrayTypeNode) Start() util.Location      { return n.Begin }
func (n *ArrayTypeNode) End() util.Location        { return n.Final }
func (n *ArrayTypeNode) Type() NodeType            { return ArrayType }
func (n *ArrayTypeNode) ResolvedType() typing.Type { return typing.Unknown() }

type PlainTypeNode struct {
	Begin, Final util.Location
	Variable     bool
	Parameters   []Node
	Names        []string
}

func (n *PlainTypeNode) Start() util.Location      { return n.Begin }
func (n *PlainTypeNode) End() util.Location        { return n.Final }
func (n *PlainTypeNode) Type() NodeType            { return PlainType }
func (n *PlainTypeNode) ResolvedType() typing.Type { return typing.Unknown() }

type FuncTypeNode struct {
	// mods for interfaces
	Mods         typing.Modifiers
	Generics     []*GenericDeclarationNode
	Begin, Final util.Location
	Variable     bool
	Identifier   string
	Parameters   []Node
	Results      []Node
}

func (n *FuncTypeNode) Start() util.Location      { return n.Begin }
func (n *FuncTypeNode) End() util.Location        { return n.Final }
func (n *FuncTypeNode) ResolvedType() typing.Type { return typing.Unknown() }
func (n *FuncTypeNode) Type() NodeType            { return FuncType }

func (t NodeType) isExpression() bool {
	switch t {
	case UnaryExpression, BinaryExpression, SliceExpression, CallExpression, IndexExpression,
		MapLiteral, ArrayLiteral, CompositeLiteral, Literal, Reference:
		return true
	}
	return false
}
