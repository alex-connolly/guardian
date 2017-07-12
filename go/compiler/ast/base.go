package ast

// Node interface for storage in AST
type Node interface {
	Type() NodeType
	Validate(NodeType) bool
	Declare(string, Node)
	Traverse()
}

type FileNode struct {
	name string
}

func (n FileNode) Type() NodeType { return File }

func (n FileNode) Validate(t NodeType) bool {
	return true
}
func (n FileNode) Declare(key string, node Node) {

}

func (n FileNode) Traverse() {

}

type PackageNode struct {
	name string
}

type ProgramNode struct {
}
