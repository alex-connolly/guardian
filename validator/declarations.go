package validator

import (
	"fmt"

	"github.com/end-r/guardian/typing"

	"github.com/end-r/guardian/ast"
)

func (v *Validator) validateType(destination ast.Node) typing.Type {
	switch n := destination.(type) {
	case ast.PlainTypeNode:
		return v.validatePlainType(n)
	case ast.MapTypeNode:
		return v.validateMapType(n)
	case ast.ArrayTypeNode:
		return v.validateArrayType(n)
	case ast.FuncTypeNode:
		return v.validateFuncType(n)
	}
	return typing.Invalid()
}

func (v *Validator) validatePlainType(node ast.PlainTypeNode) typing.Type {
	// start the validating process for another node
	typ := v.getNamedType(node.Names...)
	if typ == typing.Unknown() {
		return v.getDeclarationNode(node.Names)
	}
	return typ
}

func (v *Validator) validateArrayType(node ast.ArrayTypeNode) typing.Array {
	value := v.validateType(node.Value)
	return typing.Array{Value: value, Length: node.Length, Variable: node.Variable}
}

func (v *Validator) validateMapType(node ast.MapTypeNode) typing.Map {
	key := v.validateType(node.Key)
	value := v.validateType(node.Value)
	return typing.Map{Key: key, Value: value}
}

func (v *Validator) validateFuncType(node ast.FuncTypeNode) typing.Func {
	var params []typing.Type
	for _, p := range node.Parameters {
		params = append(params, v.validateType(p))
	}
	var results []typing.Type
	for _, r := range node.Results {
		results = append(results, v.validateType(r))
	}
	return typing.Func{Params: typing.NewTuple(params...), Results: typing.NewTuple(results...)}
}

func (v *Validator) validateDeclaration(node ast.Node) {
	switch n := node.(type) {
	case ast.ClassDeclarationNode:
		v.validateClassDeclaration(n)
		break
	case ast.ContractDeclarationNode:
		v.validateContractDeclaration(n)
		break
	case ast.EnumDeclarationNode:
		v.validateEnumDeclaration(n)
		break
	case ast.FuncDeclarationNode:
		v.validateFuncDeclaration(n)
		break
	case ast.InterfaceDeclarationNode:
		v.validateInterfaceDeclaration(n)
		break
	case ast.ExplicitVarDeclarationNode:
		v.validateVarDeclaration(n)
		break
	case ast.EventDeclarationNode:
		v.validateEventDeclaration(n)
		break
	case ast.TypeDeclarationNode:
		v.validateTypeDeclaration(n)
		break
	case ast.LifecycleDeclarationNode:
		v.validateLifecycleDeclaration(n)
		break
	default:
		fmt.Println("?")
	}

}

func (v *Validator) getDeclarationNode(names []string) typing.Type {
	if v.isParsingBuiltins {
		if v.builtinScope != nil {
			decl := v.builtinScope.GetDeclaration(names[0])
			if decl != nil {
				v.validateDeclaration(decl)
			}
		}
	} else {
		for scope := v.scope; scope != nil; scope = scope.parent {
			decl := scope.scope.GetDeclaration(names[0])
			if decl != nil {
				v.validateDeclaration(decl)
				break
			}
		}
	}
	return v.requireValidType(names)
}

func (v *Validator) requireValidType(names []string) typing.Type {
	typ := v.getNamedType(names...)
	if typ == typing.Unknown() {
		v.addError(errTypeNotVisible, makeName(names))
	}
	return typ
}

func (v *Validator) validateVarDeclaration(node ast.ExplicitVarDeclarationNode) {
	for _, id := range node.Identifiers {
		typ := v.validateType(node.DeclaredType)
		v.declareContextualVar(id, typ)
		//fmt.Printf("Declared: %s as %s\n", id, WriteType(typ))
	}
}

func (v *Validator) validateClassDeclaration(node ast.ClassDeclarationNode) {
	var supers []*Class
	for _, super := range node.Supers {
		t := v.validatePlainType(super)
		if t != standards[unknown] {
			if c, ok := t.(Class); ok {
				supers = append(supers, &c)
			} else {
				v.addError(errTypeRequired, makeName(super.Names), "class")
			}
		}
	}

	var interfaces []*Interface
	for _, ifc := range node.Interfaces {
		t := v.validatePlainType(ifc)
		if t != standards[unknown] {
			if c, ok := t.(Interface); ok {
				interfaces = append(interfaces, &c)
			} else {
				v.addError(errTypeRequired, makeName(ifc.Names), "interface")
			}
		}
	}

	types, properties, lifecycles := v.validateScope(node.Body)

	classType := NewClass(node.Identifier, supers, interfaces, types, properties, lifecycles)

	v.declareContextualType(node.Identifier, classType)
}

func (v *Validator) validateEnumDeclaration(node ast.EnumDeclarationNode) {
	var supers []*Enum
	for _, super := range node.Inherits {
		t := v.validatePlainType(super)
		if c, ok := t.(Enum); ok {
			supers = append(supers, &c)
		} else {
			v.addError(errTypeRequired, makeName(super.Names), "enum")
		}
	}

	list := node.Enums

	enumType := NewEnum(node.Identifier, supers, list)

	v.declareContextualType(node.Identifier, enumType)

}

func (v *Validator) validateContractDeclaration(node ast.ContractDeclarationNode) {
	var supers []*Contract
	for _, super := range node.Supers {
		t := v.validatePlainType(super)
		if t != standards[unknown] {
			if c, ok := t.(Contract); ok {
				supers = append(supers, &c)
			} else {
				v.addError(errTypeRequired, makeName(super.Names), "contract")
			}
		}
	}

	var interfaces []*Interface
	for _, ifc := range node.Interfaces {
		t := v.validatePlainType(ifc)
		if t != standards[unknown] {
			if c, ok := t.(Interface); ok {
				interfaces = append(interfaces, &c)
			} else {
				v.addError(errTypeRequired, makeName(ifc.Names), "interface")
			}
		}
	}

	types, properties, lifecycles := v.validateScope(node.Body)

	contractType := NewContract(node.Identifier, supers, interfaces, types, properties, lifecycles)

	v.declareContextualType(node.Identifier, contractType)
}

func (v *Validator) validateInterfaceDeclaration(node ast.InterfaceDeclarationNode) {
	var supers []*Interface
	for _, super := range node.Supers {
		t := v.validatePlainType(super)
		if t != standards[unknown] {
			if c, ok := t.(Interface); ok {
				supers = append(supers, &c)
			} else {
				v.addError(errTypeRequired, makeName(super.Names), "interface")
			}
		}
	}

	funcs := map[string]Func{}
	for _, function := range node.Signatures {
		f := v.validateContextualType(function).(Func)
		funcs[function.Identifier] = f
	}

	interfaceType := NewInterface(node.Identifier, supers, funcs)

	v.declareContextualType(node.Identifier, interfaceType)

}

func (v *Validator) validateContextualType(node ast.Node) typing.Type {
	return v.validateType(node)
}

func (v *Validator) declareContextualVar(name string, typ typing.Type) {
	if v.isParsingBuiltins {
		v.DeclareBuiltinOfType(name, typ)
	} else {
		v.DeclareVarOfType(name, typ)
	}
}

func (v *Validator) validateFuncDeclaration(node ast.FuncDeclarationNode) {

	var params []Type
	for _, p := range node.Parameters {
		for _ = range p.Identifiers {
			typ := v.validateContextualType(p.DeclaredType)
			//TODO: declare this later
			//v.DeclareVarOfType(i, typ)
			params = append(params, typ)
		}
	}

	var results []Type
	for _, r := range node.Results {
		results = append(results, v.validateType(r))
	}

	funcType := NewFunc(NewTuple(params...), NewTuple(results...))

	v.declareContextualVar(node.Identifier, funcType)

	v.validateScope(node.Body)

}

func (v *Validator) validateEventDeclaration(node ast.EventDeclarationNode) {
	var params []Type
	for _, n := range node.Parameters {
		params = append(params, v.validateType(n.DeclaredType))
	}
	eventType := NewEvent(node.Identifier, NewTuple(params...))

	v.declareContextualVar(node.Identifier, eventType)
}

func (v *Validator) declareContextualType(name string, typ typing.Type) {
	if v.isParsingBuiltins {
		v.DeclareBuiltinType(name, typ)
	} else {
		v.DeclareType(name, typ)
	}
}

func (v *Validator) validateTypeDeclaration(node ast.TypeDeclarationNode) {
	typ := v.validateType(node.Value)
	v.declareContextualType(node.Identifier, typ)
}

func (v *Validator) validateLifecycleDeclaration(node ast.LifecycleDeclarationNode) {
	// TODO: enforce location
	types := make([]Type, 0)
	for _, p := range node.Parameters {
		typ := v.validateType(p.DeclaredType)
		for _, i := range p.Identifiers {
			v.declareContextualVar(i, typ)
			types = append(types, typ)
		}
	}
	v.validateScope(node.Body)
	l := NewLifecycle(node.Category, types)
	v.declareLifecycle(node.Category, l)
}
