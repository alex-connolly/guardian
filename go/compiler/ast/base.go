package ast

import "github.com/end-r/vmgen"

// Node interface for storage in AST
type Node interface {
	Type() NodeType
	Validate(NodeType) bool
	Traverse(*vmgen.VM)
}

type ExpressionNode interface {
	Node
}

type DeclarationNode interface {
}

type ScopeNode interface {
	Declare(string, Node)
}

type StatementNode interface {
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

func (n FileNode) Traverse(vm *vmgen.VM) {

}

type PackageNode struct {
	name string
}

func (n PackageNode) Type() NodeType { return File }

func (n PackageNode) Validate(t NodeType) bool {
	return true
}
func (n PackageNode) Declare(key string, node Node) {

}

func (n PackageNode) Traverse(vm *vmgen.VM) {

}

type ProgramNode struct {
}

func (n ProgramNode) Type() NodeType { return File }

func (n ProgramNode) Validate(t NodeType) bool {
	return true
}
func (n ProgramNode) Declare(key string, node Node) {

}
