package validator

import (
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
	}
}

func (v *Validator) validateInterfaceDeclaration(node ast.InterfaceDeclarationNode) {
	var supers []Interface
	for _, super := range node.Supers {
		t := v.requireVisibleType(super.Names...)
		if c, ok := t.(Interface); ok {
			supers = append(supers, c)
		} else {
			v.addError(errTypeRequired, makeName(super.Names), "interface")
		}
	}
	interfaceType := NewInterface(node.Identifier, nil, supers)
	v.DeclareType(node.Identifier, interfaceType)
}

func (v *Validator) validateFuncDeclaration(node ast.FuncDeclarationNode) {
	// a valid function satisfies the following properties:
	// no repeated parameter names
	// all parameter types are visible in scope

	var params []Type
	for _, p := range node.Parameters {
		for _, i := range p.Identifiers {
			v.addDeclaration(i, p.DeclaredType)
			params = append(params, v.validateType(p.DeclaredType))
		}
	}

	var results []Type
	for _, r := range node.Results {
		results = append(results, v.validateType(r))
	}

	funcType := NewFunc(NewTuple(params...), NewTuple(results...))
	v.DeclareType(node.Identifier, funcType)
}

func (v *Validator) validateTypeDeclaration(node ast.TypeDeclarationNode) {
	// a valid function satisfies the following properties:
	// cannot be self-referential
	// must use a name without a previous definition in this scope
	typ := v.validateType(node.Value)
	v.DeclareType(node.Identifier, typ)
}

func (v *Validator) validateClassDeclaration(node ast.ClassDeclarationNode) {
	// a valid class satisfies the following properties:
	// interfaces must be valid types
	var interfaces []Interface
	for _, ifc := range node.Interfaces {
		t := v.requireVisibleType(ifc.Names...)
		if c, ok := t.(Interface); ok {
			interfaces = append(interfaces, c)
		} else {
			v.addError(errTypeRequired, makeName(ifc.Names), "interface")
		}
	}
	// TODO: add a reporting mechanism for non-implemented interfaces
	// superclasses must be valid types
	var supers []Class
	for _, super := range node.Supers {
		t := v.requireVisibleType(super.Names...)
		if c, ok := t.(Class); ok {
			supers = append(supers, c)
		} else {
			v.addError(errTypeRequired, makeName(super.Names), "class")
		}
	}
	classType := NewClass(node.Identifier, nil, nil, supers)
	v.DeclareType(node.Identifier, classType)
}

func (v *Validator) validateContractDeclaration(node ast.ContractDeclarationNode) {
	// a valid contract satisfies the following properties:
	// superclasses must be valid types
	var supers []Contract
	for _, super := range node.Supers {
		t := v.findReference(super.Names...)
		if t == standards[Invalid] {
			v.addError("not found")
		} else {
			if c, ok := t.(Contract); ok {
				supers = append(supers, c)
			} else {
				v.addError(errTypeRequired, makeName(super.Names), "contract")
			}
		}
	}
	contractType := NewContract(node.Identifier, supers, nil)
	v.DeclareType(node.Identifier, contractType)
}

func (v *Validator) validateEnumDeclaration(node ast.EnumDeclarationNode) {
	// a valid enum satisfies the following properties:
	// superclasses must be valid types
	var supers []Enum
	for _, super := range node.Inherits {
		t := v.requireVisibleType(super.Names...)
		if c, ok := t.(Enum); ok {
			supers = append(supers, c)
		} else {
			v.addError(errTypeRequired, makeName(super.Names), "enum")
		}
	}
	enumType := NewEnum(node.Identifier, supers)
	v.DeclareType(node.Identifier, enumType)
}

func (v *Validator) validateEventDeclaration(node ast.EventDeclarationNode) {
	var params []Type
	for _, n := range node.Parameters {
		params = append(params, v.validateType(n.DeclaredType))
	}
	eventType := NewEvent(node.Identifier, NewTuple(params...))
	v.DeclareType(node.Identifier, eventType)
}

func (v *Validator) validateLifecycleDeclaration(node ast.LifecycleDeclarationNode) {
	// a valid constructor satisfies the following properties:
	// no repeated parameter names
	// all parameter types are visible in scope
	for _, p := range node.Parameters {
		v.validateType(p.DeclaredType)
		for _, i := range p.Identifiers {
			v.addDeclaration(i, p.DeclaredType)
		}
	}
}
