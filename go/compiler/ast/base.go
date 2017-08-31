package ast

import "github.com/end-r/vmgen"

// Node interface for storage in AST
type Node interface {
	Type() NodeType
	Traverse(*vmgen.VM)
}

type ExpressionNode interface {
	Node
}

type DeclarationNode interface {
	Node
}

type StatementNode interface {
	Node
}

type ScopeNode struct {
	ValidTypes []NodeType
	Nodes      []Node
}

func (n ScopeNode) Add(new Node) {
	if n.Nodes == nil {
		n.Nodes = make([]Node, 0)
	}
	n.Nodes = append(n.Nodes, new)
}

func (n ScopeNode) Type() NodeType { return Scope }

func (n ScopeNode) Traverse(vm *vmgen.VM) {
	for _, n := range n.Nodes {
		n.Traverse(vm)
	}
}

func (n ScopeNode) IsValid(nt NodeType) bool {
	for _, t := range n.ValidTypes {
		if t == nt {
			return true
		}
	}
	return false
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
