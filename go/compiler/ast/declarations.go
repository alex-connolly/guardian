package ast

import "github.com/end-r/vmgen"

type TypeDeclarationNode struct {
	Identifier string
	Value      ReferenceNode
}

// Type ...
func (n TypeDeclarationNode) Type() NodeType { return TypeDeclaration }

func (n TypeDeclarationNode) Traverse(vm *vmgen.VM) {

}

type FuncDeclarationNode struct {
	Identifier string
	Parameters []ExplicitVarDeclarationNode
	Results    []ReferenceNode
	Body       ScopeNode
	IsAbstract bool
}

func (n FuncDeclarationNode) Type() NodeType { return FuncDeclaration }

func (n FuncDeclarationNode) Traverse(vm *vmgen.VM) {

}

type ClassDeclarationNode struct {
	Identifier   string
	IsAbstract   bool
	Supers       []ReferenceNode
	Interfaces   []ReferenceNode
	Body         ScopeNode
	declarations map[string][]Node
}

func (n ClassDeclarationNode) Type() NodeType { return ClassDeclaration }

func (n ClassDeclarationNode) Traverse(vm *vmgen.VM) {

}

type InterfaceDeclarationNode struct {
	Identifier string
	IsAbstract bool
	Body       ScopeNode
	Supers     []ReferenceNode
}

func (n InterfaceDeclarationNode) Type() NodeType { return InterfaceDeclaration }

func (n InterfaceDeclarationNode) Traverse(vm *vmgen.VM) {

}

type ContractDeclarationNode struct {
	Identifier string
	IsAbstract bool
	Supers     []ReferenceNode
	Interfaces []ReferenceNode
	Body       ScopeNode
}

func (n ContractDeclarationNode) Type() NodeType { return ContractDeclaration }

func (n ContractDeclarationNode) Traverse(vm *vmgen.VM) {
	// find our variable declarations
}

type MapTypeNode struct {
	Key   Node
	Value Node
}

func (n MapTypeNode) Type() NodeType { return MapType }

func (n MapTypeNode) Traverse(vm *vmgen.VM) {

}

type ArrayTypeNode struct {
	Value Node
}

func (n ArrayTypeNode) Type() NodeType { return ArrayType }

func (n ArrayTypeNode) Traverse(vm *vmgen.VM) {

}

type FuncTypeNode struct {
	Value Node
}

func (n FuncTypeNode) Type() NodeType { return FuncType }

func (n FuncTypeNode) Traverse(vm *vmgen.VM) {

}

type ExplicitVarDeclarationNode struct {
	Identifiers  []string
	DeclaredType ReferenceNode
}

func (n ExplicitVarDeclarationNode) Type() NodeType { return ExplicitVarDeclaration }

func (n ExplicitVarDeclarationNode) Traverse(vm *vmgen.VM) {

}

type EventDeclarationNode struct {
	Identifier string
	Parameters []ReferenceNode
}

func (n EventDeclarationNode) Type() NodeType { return EventDeclaration }

func (n EventDeclarationNode) Traverse(vm *vmgen.VM) {
	vm.AddBytecode("LOG")
}

// constructors are functions
// have their own type to distinguish them from functions
// declared as construtor(params) {
//
// }
type ConstructorDeclarationNode struct {
	Parameters []ExplicitVarDeclarationNode
}

func (n ConstructorDeclarationNode) Type() NodeType { return ConstructorDeclaration }

func (n ConstructorDeclarationNode) Traverse(vm *vmgen.VM) {

}

type EnumDeclarationNode struct {
	Cases []ReferenceNode
}

func (n EnumDeclarationNode) Type() NodeType { return EnumDeclaration }

func (n EnumDeclarationNode) Traverse(vm *vmgen.VM) {

}
