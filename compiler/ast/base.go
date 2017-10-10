package ast

// Node interface for storage in AST
type Node interface {
	Type() NodeType
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
	Parent       *ScopeNode
	ValidTypes   []NodeType
	Declarations map[string]Node
	Sequence     []Node
}

func (n *ScopeNode) AddSequential(node Node) {
	if n.Sequence == nil {
		n.Sequence = make([]Node, 0)
	}
	n.Sequence = append(n.Sequence, node)
}

func (n *ScopeNode) AddDeclaration(key string, node Node) {
	// declarations is a map to shortcut lookups
	// could change value to array for overloaded methods etc
	// don't think supporting overloading is a good idea at this stage
	if n.Declarations == nil {
		n.Declarations = make(map[string]Node)
	}
	n.Declarations[key] = node
}

func (n ScopeNode) Type() NodeType { return Scope }

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

type PackageNode struct {
	name string
}

func (n PackageNode) Type() NodeType { return File }

type ProgramNode struct {
}

func (n ProgramNode) Type() NodeType { return File }
