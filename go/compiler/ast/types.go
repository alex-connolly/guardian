package ast

// NodeType denotes the type of node
type NodeType int

const (
	ContractDeclaration NodeType = iota
	ClassDeclaration
	FuncDeclaration
	InterfaceDeclaration
	TypeDeclaration
	ArrayType
	MapType

	Reference
	Literal
	CompositeLiteral
	MapLiteral
	ArrayLiteral
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
	SwitchStatement
	CaseStatement
	BlockStatement
	ForStatement

	File
)

func (t *NodeType) isExpression() bool {
	switch t {
	case UnaryExpression, BinaryExpression, SliceExpression, CallExpression, IndexExpression,
		SliceExpression, GenericExpression, MapLiteral, ArrayLiteral, CompositeLiteral, Literal, Reference:
		return true
	}
	return false
}
