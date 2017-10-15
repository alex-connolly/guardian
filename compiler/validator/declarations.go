package validator

import "github.com/end-r/guardian/compiler/ast"

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
	}
}

func (v *Validator) validateFuncDeclaration(node ast.FuncDeclarationNode) {
	// a valid function satisfies the following properties:
	// no repeated parameter names
	// all parameter types are visible in scope
}

func (v *Validator) validateTypeDeclaration(node ast.TypeDeclarationNode) {
	// a valid function satisfies the following properties:
	// cannot be self-referential
	// must use a name without a previous definition in this scope
}

func (v *Validator) validateClassDeclaration(node ast.ClassDeclarationNode) {
	// a valid class satisfies the following properties:
	// interfaces must be valid types
	for _, inter := range node.Interfaces {
		v.requireVisibleType(inter.Names...)
	}
	// TODO: add a reporting mechanism for non-implemented interfaces
	// superclasses must be valid types
	for _, super := range node.Supers {
		v.requireVisibleType(super.Names...)
	}
}

func (v *Validator) validateContractDeclaration(node ast.ContractDeclarationNode) {
	// a valid contract satisfies the following properties:
	// superclasses must be valid types
	for _, super := range node.Supers {
		v.requireVisibleType(super.Names...)
	}

}

func (v *Validator) validateEnumDeclaration(node ast.EnumDeclarationNode) {
	// a valid enum satisfies the following properties:

}

func (v *Validator) validateEventDeclaration(node ast.EventDeclarationNode) {
	// a valid event satisfies the following properties:
	// no repeated parameter names
	// all parameter types are visible in scope
	for _, n := range node.Parameters {
		v.requireVisibleType(n.Names...)
	}
}

func (v *Validator) validateLifecycleDeclaration(node ast.LifecycleDeclarationNode) {
	// a valid constructor satisfies the following properties:
	// no repeated parameter names
	// all parameter types are visible in scope
	for _, p := range node.Parameters {
		v.validateType(p.DeclaredType)
	}
}

func (v *Validator) validateType(node ast.Node) {
	switch node.Type() {
	case ast.Reference:
		ref := node.(ast.ReferenceNode)
		v.requireVisibleType(ref.Names...)
		break
	case ast.MapType:
		ref := node.(ast.MapTypeNode)
		v.validateType(ref.Key)
		v.validateType(ref.Value)
		break
	case ast.ArrayType:
		ref := node.(ast.ArrayTypeNode)
		v.validateType(ref.Value)
		break
	case ast.FuncType:
		ref := node.(ast.FuncTypeNode)
		for _, p := range ref.Parameters {
			v.validateType(p)
		}
		for _, r := range ref.Results {
			v.validateType(r)
		}
		break
	}
}
