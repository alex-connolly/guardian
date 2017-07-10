package ast

type TypeDeclarationNode struct {
	identifier string
}

func (n *TypeDeclarationNode) Type() { return TypeDeclaration }

type FuncDeclarationNode struct {
	identifier string
	isAbstract bool
}

func (n *FuncDeclarationNode) Type() { return FuncDeclaration }

type ClassDeclarationNode struct {
	identifier   string
	isAbstract   bool
	declarations map[string][]Node
}

func (n *ClassDeclarationNode) Type() { return ClassDeclaration }

func (n *ClassDeclarationNode) Validate(t NodeType) {
	switch t {
	case ClassDeclaration, InterfaceDeclaration, ContractDeclaration:
		return true
	}
	return false
}

type InterfaceDeclarationNode struct {
	identifier   string
	isAbstract   bool
	declarations map[string][]Node
}

func (n *InterfaceDeclarationNode) Type() { return InterfaceDeclaration }

type ContractDeclarationNode struct {
	identifier   string
	isAbstract   bool
	declarations map[string][]Node
}

func (n *ContractDeclarationNode) Type() { return ContractDeclaration }
