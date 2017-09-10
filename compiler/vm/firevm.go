package vm

import (
	"github.com/end-r/firevm"
	"github.com/end-r/guardian/compiler/ast"
	"github.com/end-r/vmgen"
)

// An Arsonist burns down trees
type Arsonist struct {
	VM *vmgen.VM
}

// Traverse ...
func (a *Arsonist) Traverse(node ast.Node) {
	// initialise the vm
	if a.VM == nil {
		a.VM = firevm.NewVM()
	}
	switch node.Type() {
	case ast.ClassDeclaration:
		a.traverseClass(node.(ast.ClassDeclarationNode))
		break
	case ast.InterfaceDeclaration:
		a.traverseInterface(node.(ast.InterfaceDeclarationNode))
		break
	case ast.EnumDeclaration:
		a.traverseEnum(node.(ast.EnumDeclarationNode))
		break
	case ast.EventDeclaration:
		a.traverseEvent(node.(ast.EventDeclarationNode))
		break
	case ast.ContractDeclaration:
		a.traverseContract(node.(ast.ContractDeclarationNode))
		break
	}
}

func (a *Arsonist) traverseClass(n ast.ClassDeclarationNode) {

}

func (a *Arsonist) traverseInterface(n ast.InterfaceDeclarationNode) {

}

func (a *Arsonist) traverseEnum(n ast.EnumDeclarationNode) {

}

func (a *Arsonist) traverseContract(n ast.ContractDeclarationNode) {

}

func (a *Arsonist) traverseEvent(n ast.EventDeclarationNode) {

}
