package ast

// Node interface for storage in AST
type Node interface {
	Type() NodeType
	Validate(string, NodeType)
	Declare(string, Node)
}

type FileNode struct {
	name string
}

type PackageNode struct {
	name string
}

type ProgramNode struct {
}
