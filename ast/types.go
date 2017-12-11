package ast

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
	Variable bool
	Key      Node
	Value    Node
}

func (n *MapTypeNode) Type() NodeType { return MapType }

type ArrayTypeNode struct {
	Variable bool
	Length   int
	Value    Node
}

func (n *ArrayTypeNode) Type() NodeType { return ArrayType }

type PlainTypeNode struct {
	Variable bool
	Names    []string
}

func (n *PlainTypeNode) Type() NodeType { return PlainType }

type FuncTypeNode struct {
	Variable   bool
	Identifier string
	Parameters []Node
	Results    []Node
}

func (n *FuncTypeNode) Type() NodeType { return FuncType }

func (t NodeType) isExpression() bool {
	switch t {
	case UnaryExpression, BinaryExpression, SliceExpression, CallExpression, IndexExpression,
		GenericExpression, MapLiteral, ArrayLiteral, CompositeLiteral, Literal, Reference:
		return true
	}
	return false
}
