package ast

type TypeDeclarationNode struct {
	identifier string
}

func (n *TypeDeclarationNode) Type() { return TypeDeclaration }

type FuncDeclarationNode struct {
	identifier string
}

func (n *FuncDeclarationNode) Type() { return FuncDeclaration }

type ClassDeclarationNode struct {
	identifier string
}

func (n *ClassDeclarationNode) Type() { return ClassDeclaration }

type InterfaceDeclarationNode struct {
	identifier string
}

func (n *InterfaceDeclarationNode) Type() { return InterfaceDeclaration }

type ContractDeclarationNode struct {
	identifier string
}

func (n *ContractDeclarationNode) Type() { return ContractDeclaration }
