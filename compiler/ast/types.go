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

	File
	Scope
)

type MapTypeNode struct {
	Key   Node
	Value Node
}

func (n MapTypeNode) Type() NodeType { return MapType }

type ArrayTypeNode struct {
	Value Node
}

func (n ArrayTypeNode) Type() NodeType { return ArrayType }

type FuncTypeNode struct {
	Parameters []Node
	Results    []Node
}

func (n FuncTypeNode) Type() NodeType { return FuncType }

func (t NodeType) isExpression() bool {
	switch t {
	case UnaryExpression, BinaryExpression, SliceExpression, CallExpression, IndexExpression,
		GenericExpression, MapLiteral, ArrayLiteral, CompositeLiteral, Literal, Reference:
		return true
	}
	return false
}
