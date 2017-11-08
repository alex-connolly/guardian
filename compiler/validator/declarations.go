package validator

import (
	"fmt"

	"github.com/end-r/guardian/compiler/ast"
)

func (v *Validator) validateDeclaration(node ast.Node) {
	switch node.Type() {
	case ast.FuncDeclaration:
		v.validateFuncDeclaration(node.(ast.FuncDeclarationNode))
		break
	case ast.TypeDeclaration:
		v.validateTypeDeclaration(node.(ast.TypeDeclarationNode))
		break
	case ast.ClassDeclaration:
		v.validateClassDeclaration(node.(ast.ClassDeclarationNode))
		break
	case ast.ContractDeclaration:
		v.validateContractDeclaration(node.(ast.ContractDeclarationNode))
		break
	case ast.EnumDeclaration:
		v.validateEnumDeclaration(node.(ast.EnumDeclarationNode))
		break
	case ast.EventDeclaration:
		v.validateEventDeclaration(node.(ast.EventDeclarationNode))
		break
	case ast.LifecycleDeclaration:
		v.validateLifecycleDeclaration(node.(ast.LifecycleDeclarationNode))
		break
	case ast.InterfaceDeclaration:
		v.validateInterfaceDeclaration(node.(ast.InterfaceDeclarationNode))
		break
	case ast.ExplicitVarDeclaration:
		v.validateExplicitVarDeclaration(node.(ast.ExplicitVarDeclarationNode))
		break
	}
}

func (v *Validator) validateExplicitVarDeclaration(node ast.ExplicitVarDeclarationNode) {
	for _, id := range node.Identifiers {
		typ := v.validateType(node.DeclaredType)
		fmt.Printf("Explicitly Declaring %s of type %s\n", id, typ)
		v.DeclareVarOfType(id, typ)
	}
}

func (v *Validator) validateInterfaceDeclaration(node ast.InterfaceDeclarationNode) {

	funcs := map[string]Func{}
	for _, function := range node.Signatures {
		f := v.validateType(function).(Func)
		funcs[function.Identifier] = f
	}

	t := v.getNamedType(node.Identifier)

	if cl, ok := t.(Interface); ok {
		cl.Funcs = funcs
	}

}

func (v *Validator) validateFuncDeclaration(node ast.FuncDeclarationNode) {
	v.validateScope(node.Body)
}

func (v *Validator) validateTypeDeclaration(node ast.TypeDeclarationNode) {

}

func (v *Validator) validateClassDeclaration(node ast.ClassDeclarationNode) {

	properties := v.validateScope(node.Body)

	t := v.getNamedType(node.Identifier)

	if cl, ok := t.(Class); ok {
		cl.Properties = properties
	}

}

func (v *Validator) validateContractDeclaration(node ast.ContractDeclarationNode) {

	properties := v.validateScope(node.Body)

	t := v.getNamedType(node.Identifier)

	if cl, ok := t.(Contract); ok {
		cl.Properties = properties
	}

}

func (v *Validator) validateEnumDeclaration(node ast.EnumDeclarationNode) {

	list := map[string]bool{}
	for _, s := range node.Enums {
		list[s] = true
	}

	t := v.getNamedType(node.Identifier)

	if cl, ok := t.(Enum); ok {
		cl.Items = list
	}
}

func (v *Validator) validateEventDeclaration(node ast.EventDeclarationNode) {

}

func (v *Validator) validateLifecycleDeclaration(node ast.LifecycleDeclarationNode) {
	// a valid constructor satisfies the following properties:
	// no repeated parameter names
	// all parameter types are visible in scope

	v.validateScope(node.Body)
}
