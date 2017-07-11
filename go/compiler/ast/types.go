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
)
