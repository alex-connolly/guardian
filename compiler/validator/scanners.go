package validator

import (
	"fmt"

	"github.com/end-r/guardian/compiler/ast"
)

func (v *Validator) scanType(destination ast.Node) Type {
	switch n := destination.(type) {
	case ast.PlainTypeNode:
		return v.scanPlainType(n)
	case ast.MapTypeNode:
		return v.scanMapType(n)
	case ast.ArrayTypeNode:
		return v.scanArrayType(n)
	case ast.FuncTypeNode:
		return v.scanFuncType(n)
	}
	return standards[Invalid]
}

func (v *Validator) scanPlainType(node ast.PlainTypeNode) Type {
	// start the scanning process for another node
	typ := v.getNamedType(node.Names...)
	if typ == standards[Unknown] {
		return v.getDeclarationNode(node.Names)
	}
	return typ
}

func (v *Validator) scanArrayType(node ast.ArrayTypeNode) Array {
	value := v.scanType(node.Value)
	return NewArray(value)
}

func (v *Validator) scanMapType(node ast.MapTypeNode) Map {
	key := v.scanType(node.Key)
	value := v.scanType(node.Value)
	return NewMap(key, value)
}

func (v *Validator) scanFuncType(node ast.FuncTypeNode) Func {
	params := v.resolveTuple(node.Parameters)
	results := v.resolveTuple(node.Results)
	return NewFunc(params, results)
}

func (v *Validator) addScanned(name string, t Type) {
	if v.scope.scanned == nil {
		v.scope.scanned = make(map[string]Type)
	}
	v.scope.scanned[name] = t
}

func (v *Validator) scanDeclaration(node ast.Node) {
	switch n := node.(type) {
	case ast.ClassDeclarationNode:
		v.scanClassDeclaration(n)
		break
	case ast.ContractDeclarationNode:
		v.scanContractDeclaration(n)
		break
	case ast.EnumDeclarationNode:
		v.scanEnumDeclaration(n)
		break
	case ast.FuncDeclarationNode:
		v.scanFuncDeclaration(n)
		break
	case ast.InterfaceDeclarationNode:
		v.scanInterfaceDeclaration(n)
		break
	case ast.ExplicitVarDeclarationNode:
		v.scanVarDeclaration(n)
		break
	case ast.EventDeclarationNode:
		v.scanEventDeclaration(n)
		break
	case ast.TypeDeclarationNode:
		v.scanTypeDeclaration(n)
		break
	case ast.LifecycleDeclarationNode:
		v.scanLifecycleDeclaration(n)
		break
	default:
		fmt.Println("?")
	}

}

func (v *Validator) getDeclarationNode(names []string) Type {
	decl := v.scope.scope.GetDeclaration(names[0])
	if decl != nil {
		v.scanDeclaration(decl)
	}
	return v.requireScannableType(names)
}

func (v *Validator) requireScannableType(names []string) Type {
	typ := v.getNamedType(names...)
	if typ == standards[Unknown] {
		v.addError(errTypeNotVisible, makeName(names))
	}
	return typ
}

func (v *Validator) scanVarDeclaration(node ast.ExplicitVarDeclarationNode) {
	for _, id := range node.Identifiers {
		v.DeclareType(id, v.scanType(node.DeclaredType))
	}
}

func (v *Validator) scanClassDeclaration(node ast.ClassDeclarationNode) {
	var supers []*Class
	for _, super := range node.Supers {
		t := v.scanPlainType(super)
		if t != standards[Unknown] {
			if c, ok := t.(Class); ok {
				supers = append(supers, &c)
			} else {
				v.addError(errTypeRequired, makeName(super.Names), "class")
			}
		}
	}

	var interfaces []*Interface
	for _, ifc := range node.Interfaces {
		t := v.scanPlainType(ifc)
		if t != standards[Unknown] {
			if c, ok := t.(Interface); ok {
				interfaces = append(interfaces, &c)
			} else {
				v.addError(errTypeRequired, makeName(ifc.Names), "interface")
			}
		}
	}

	classType := NewClass(node.Identifier, supers, interfaces, nil)
	v.DeclareType(node.Identifier, classType)
}

func (v *Validator) scanEnumDeclaration(node ast.EnumDeclarationNode) {
	var supers []*Enum
	for _, super := range node.Inherits {
		t := v.scanPlainType(super)
		if c, ok := t.(Enum); ok {
			supers = append(supers, &c)
		} else {
			v.addError(errTypeRequired, makeName(super.Names), "enum")
		}
	}

	enumType := NewEnum(node.Identifier, supers, nil)
	v.DeclareType(node.Identifier, enumType)
}

func (v *Validator) scanContractDeclaration(node ast.ContractDeclarationNode) {
	var supers []*Contract
	for _, super := range node.Supers {
		t := v.scanPlainType(super)
		if t != standards[Unknown] {
			if c, ok := t.(Contract); ok {
				supers = append(supers, &c)
			} else {
				v.addError(errTypeRequired, makeName(super.Names), "contract")
			}
		}
	}

	var interfaces []*Interface
	for _, ifc := range node.Interfaces {
		t := v.scanPlainType(ifc)
		if t != standards[Unknown] {
			if c, ok := t.(Interface); ok {
				interfaces = append(interfaces, &c)
			} else {
				v.addError(errTypeRequired, makeName(ifc.Names), "interface")
			}
		}
	}
	contractType := NewContract(node.Identifier, supers, interfaces, nil)
	v.DeclareType(node.Identifier, contractType)
}

func (v *Validator) scanInterfaceDeclaration(node ast.InterfaceDeclarationNode) {
	var supers []*Interface
	for _, super := range node.Supers {
		t := v.scanPlainType(super)
		if t != standards[Unknown] {
			if c, ok := t.(Interface); ok {
				supers = append(supers, &c)
			} else {
				v.addError(errTypeRequired, makeName(super.Names), "interface")
			}
		}
	}

	interfaceType := NewInterface(node.Identifier, supers, nil)
	v.DeclareVarOfType(node.Identifier, interfaceType)

}

func (v *Validator) scanFuncDeclaration(node ast.FuncDeclarationNode) {

	var params []Type
	for _, p := range node.Parameters {
		for _, i := range p.Identifiers {
			typ := v.scanType(p.DeclaredType)
			v.DeclareVarOfType(i, typ)
			params = append(params, typ)
		}
	}

	var results []Type
	for _, r := range node.Results {
		results = append(results, v.scanType(r))
	}

	funcType := NewFunc(NewTuple(params...), NewTuple(results...))
	v.DeclareVarOfType(node.Identifier, funcType)
}

func (v *Validator) scanEventDeclaration(node ast.EventDeclarationNode) {
	var params []Type
	for _, n := range node.Parameters {
		params = append(params, v.scanType(n.DeclaredType))
	}
	eventType := NewEvent(node.Identifier, NewTuple(params...))
	v.DeclareVarOfType(node.Identifier, eventType)
}

func (v *Validator) scanTypeDeclaration(node ast.TypeDeclarationNode) {
	v.DeclareType(node.Identifier, v.scanType(node.Value))
}

func (v *Validator) scanLifecycleDeclaration(node ast.LifecycleDeclarationNode) {
	for _, p := range node.Parameters {
		typ := v.scanType(p.DeclaredType)
		for _, i := range p.Identifiers {
			v.DeclareVarOfType(i, typ)
		}
	}
}
