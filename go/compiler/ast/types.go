package ast

// NodeType denotes the type of node
type NodeType int

const (
	ContractDeclaration = iota
	ClassDeclaration
	FuncDeclaration

	Literal
	CompositeLiteral
	BinaryExpression
	UnaryExpression
	GenericExpression
	SliceExpression
	IndexExpression

	AssignmentStatement
	ReturnStatement
	BranchStatement
	IfStatement
	SwitchStatement
	CaseClause
	BlockStatement
	ForStatement
)
