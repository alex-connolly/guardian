package ast

import "github.com/end-r/guardian/typing"

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
	GenericExpression
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
		GenericExpression, MapLiteral, ArrayLiteral, CompositeLiteral, Literal, Reference,
	}
)

type MapTypeNode struct {
	Begin, Final uint
	Variable     bool
	Key          Node
	Value        Node
}

func (n *MapTypeNode) Start() uint               { return n.Begin }
func (n *MapTypeNode) End() uint                 { return n.Final }
func (n *MapTypeNode) Type() NodeType            { return MapType }
func (n *MapTypeNode) ResolvedType() typing.Type { return typing.Unknown() }

type ArrayTypeNode struct {
	Begin, Final uint
	Variable     bool
	Length       int
	Value        Node
}

func (n *ArrayTypeNode) Start() uint               { return n.Begin }
func (n *ArrayTypeNode) End() uint                 { return n.Final }
func (n *ArrayTypeNode) Type() NodeType            { return ArrayType }
func (n *ArrayTypeNode) ResolvedType() typing.Type { return typing.Unknown() }

type PlainTypeNode struct {
	Begin, Final uint
	Variable     bool
	Parameters   []Node
	Names        []string
}

func (n *PlainTypeNode) Start() uint               { return n.Begin }
func (n *PlainTypeNode) End() uint                 { return n.Final }
func (n *PlainTypeNode) Type() NodeType            { return PlainType }
func (n *PlainTypeNode) ResolvedType() typing.Type { return typing.Unknown() }

type FuncTypeNode struct {
	Begin, Final uint
	Variable     bool
	Identifier   string
	Parameters   []Node
	Results      []Node
}

func (n *FuncTypeNode) Start() uint               { return n.Begin }
func (n *FuncTypeNode) End() uint                 { return n.Final }
func (n *FuncTypeNode) ResolvedType() typing.Type { return typing.Unknown() }
func (n *FuncTypeNode) Type() NodeType            { return FuncType }

func (t NodeType) isExpression() bool {
	switch t {
	case UnaryExpression, BinaryExpression, SliceExpression, CallExpression, IndexExpression,
		GenericExpression, MapLiteral, ArrayLiteral, CompositeLiteral, Literal, Reference:
		return true
	}
	return false
}
