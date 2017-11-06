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
	case ast.ExplicitVarDeclaration:
		v.validateExplicitVarDeclaration(node.(ast.ExplicitVarDeclarationNode))
		break
	}
}

func (v *Validator) validateExplicitVarDeclaration(node ast.ExplicitVarDeclarationNode) {
	for _, id := range node.Identifiers {
		v.DeclareVarOfType(id, v.resolveType(node.DeclaredType))
	}
}

func (v *Validator) validateInterfaceDeclaration(node ast.InterfaceDeclarationNode) {
	var supers []*Interface
	for _, super := range node.Supers {
		t := v.requireVisibleType(super.Names...)
		if c, ok := t.(Interface); ok {
			supers = append(supers, &c)
		} else {
			v.addError(errTypeRequired, makeName(super.Names), "interface")
		}
	}

	funcs := map[string]Func{}
	for _, function := range node.Signatures {
		f := v.validateType(function).(Func)
		funcs[function.Identifier] = f
	}

	interfaceType := NewInterface(node.Identifier, funcs, supers)
	v.DeclareVarOfType(node.Identifier, interfaceType)
}

func (v *Validator) validateFuncDeclaration(node ast.FuncDeclarationNode) {
	// a valid function satisfies the following properties:
	// no repeated parameter names
	// all parameter types are visible in scope

	var params []Type
	for _, p := range node.Parameters {
		for _, i := range p.Identifiers {
			typ := v.validateType(p.DeclaredType)
			v.DeclareVarOfType(i, typ)
			params = append(params, typ)
		}
	}

	var results []Type
	for _, r := range node.Results {
		results = append(results, v.validateType(r))
	}

	funcType := NewFunc(NewTuple(params...), NewTuple(results...))
	v.DeclareVarOfType(node.Identifier, funcType)

	v.validateScope(node.Body)
}

func (v *Validator) validateTypeDeclaration(node ast.TypeDeclarationNode) {
	// a valid function satisfies the following properties:
	// cannot be self-referential
	// must use a name without a previous definition in this scope
	typ := v.validateType(node.Value)
	v.DeclareVarOfType(node.Identifier, typ)
}

func (v *Validator) validateClassDeclaration(node ast.ClassDeclarationNode) {
	// a valid class satisfies the following properties:
	// TODO: add a reporting mechanism for non-implemented interfaces
	// superclasses must be valid types
	var supers []*Class
	for _, super := range node.Supers {
		t := v.requireVisibleType(super.Names...)
		if c, ok := t.(Class); ok {
			supers = append(supers, &c)
		} else {
			v.addError(errTypeRequired, makeName(super.Names), "class")
		}
	}

	var interfaces []*Interface
	for _, ifc := range node.Interfaces {
		t := v.requireVisibleType(ifc.Names...)
		if c, ok := t.(Interface); ok {
			interfaces = append(interfaces, &c)
		} else {
			v.addError(errTypeRequired, makeName(ifc.Names), "interface")
		}
	}

	properties := v.validateScope(node.Body)

	classType := NewClass(node.Identifier, supers, interfaces, properties)
	v.DeclareType(node.Identifier, classType)

}

func (v *Validator) validateContractDeclaration(node ast.ContractDeclarationNode) {
	// a valid contract satisfies the following properties:
	// superclasses must be valid types
	var supers []*Contract
	for _, super := range node.Supers {
		t := v.findReference(super.Names...)
		if t == standards[Invalid] {
			v.addError(errTypeNotVisible)
		} else {
			if c, ok := t.(Contract); ok {
				supers = append(supers, &c)
			} else {
				v.addError(errTypeRequired, makeName(super.Names), "contract")
			}
		}
	}

	var interfaces []*Interface
	for _, ifc := range node.Interfaces {
		t := v.requireVisibleType(ifc.Names...)
		if c, ok := t.(Interface); ok {
			interfaces = append(interfaces, &c)
		} else {
			v.addError(errTypeRequired, makeName(ifc.Names), "interface")
		}
	}

	properties := v.validateScope(node.Body)

	contractType := NewContract(node.Identifier, supers, interfaces, properties)
	v.DeclareType(node.Identifier, contractType)

}

func (v *Validator) validateEnumDeclaration(node ast.EnumDeclarationNode) {
	// a valid enum satisfies the following properties:
	// superclasses must be valid types
	var supers []*Enum
	for _, super := range node.Inherits {
		t := v.requireVisibleType(super.Names...)
		if c, ok := t.(Enum); ok {
			supers = append(supers, &c)
		} else {
			v.addError(errTypeRequired, makeName(super.Names), "enum")
		}
	}

	list := map[string]bool{}
	for _, s := range node.Enums {
		list[s] = true
	}

	enumType := NewEnum(node.Identifier, supers, list)
	v.DeclareType(node.Identifier, enumType)
}

func (v *Validator) validateEventDeclaration(node ast.EventDeclarationNode) {
	var params []Type
	for _, n := range node.Parameters {
		params = append(params, v.validateType(n.DeclaredType))
	}
	eventType := NewEvent(node.Identifier, NewTuple(params...))
	v.DeclareVarOfType(node.Identifier, eventType)
}

func (v *Validator) validateLifecycleDeclaration(node ast.LifecycleDeclarationNode) {
	// a valid constructor satisfies the following properties:
	// no repeated parameter names
	// all parameter types are visible in scope
	for _, p := range node.Parameters {
		typ := v.validateType(p.DeclaredType)
		for _, i := range p.Identifiers {
			v.DeclareVarOfType(i, typ)
		}
	}

	v.validateScope(node.Body)
}
