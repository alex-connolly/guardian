package ast

// Node interface for storage in AST
type Node interface {
	Type() NodeType
}
