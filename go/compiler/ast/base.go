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
	Parent     *ScopeNode
	ValidTypes []NodeType
	nodes      map[string][]Node
}

func (n *ScopeNode) Nodes(key string) []Node {
	return n.nodes[key]
}

func (n *ScopeNode) Declare(key string, node Node) {
	if n.nodes == nil {
		n.nodes = make(map[string][]Node)
	}
	if n.nodes[key] == nil {
		n.nodes[key] = make([]Node, 0)
	}
	n.nodes[key] = append(n.nodes[key], node)
}

func (n ScopeNode) Type() NodeType { return Scope }

func (n ScopeNode) Traverse(vm *vmgen.VM) {

}

func (n *ScopeNode) IsValid(nt NodeType) bool {
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

func (n FileNode) Traverse(vm *vmgen.VM) {

}

type PackageNode struct {
	name string
}

func (n PackageNode) Type() NodeType { return File }

func (n PackageNode) Traverse(vm *vmgen.VM) {

}

type ProgramNode struct {
}

func (n ProgramNode) Type() NodeType { return File }
