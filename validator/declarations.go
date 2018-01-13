package validator

import (
	"fmt"
	"reflect"

	"github.com/end-r/guardian/token"

	"github.com/end-r/guardian/typing"

	"github.com/end-r/guardian/ast"
)

func (v *Validator) validateType(destination ast.Node) typing.Type {
	switch n := destination.(type) {
	case *ast.PlainTypeNode:
		return v.validatePlainType(n)
	case *ast.MapTypeNode:
		return v.validateMapType(n)
	case *ast.ArrayTypeNode:
		return v.validateArrayType(n)
	case *ast.FuncTypeNode:
		return v.validateFuncType(n)
	}
	return typing.Invalid()
}

func (v *Validator) validatePlainType(node *ast.PlainTypeNode) typing.Type {
	// start the validating process for another node
	typ := v.getNamedType(node.Names...)
	if typ == typing.Unknown() {
		typ = v.getDeclarationNode(node.Names)
	}
	// validate parameters if necessary
	switch a := typ.(type) {
	case *typing.Class:
		if len(node.Parameters) != len(a.Generics) {
			v.addError(errWrongParameterLength)
		}
		for i, p := range node.Parameters {
			if !a.Generics[i].Accepts(v.validateType(p)) {
				v.addError(errInvalidParameter)
			}
		}
		break
	case *typing.Interface:
		if len(node.Parameters) != len(a.Generics) {
			v.addError(errWrongParameterLength)
		}
		for i, p := range node.Parameters {
			if !a.Generics[i].Accepts(v.validateType(p)) {
				v.addError(errInvalidParameter)
			}
		}
		break
	case *typing.Contract:
		if len(node.Parameters) != len(a.Generics) {
			v.addError(errWrongParameterLength)
		}
		for i, p := range node.Parameters {
			if !a.Generics[i].Accepts(v.validateType(p)) {
				v.addError(errInvalidParameter)
			}
		}
		break
	default:
		if len(node.Parameters) > 0 {
			v.addError(errCannotParametrizeType)
		}
		break
	}

	return typ
}

func (v *Validator) validateArrayType(node *ast.ArrayTypeNode) *typing.Array {
	value := v.validateType(node.Value)
	return &typing.Array{
		Value:    value,
		Length:   node.Length,
		Variable: node.Variable,
	}
}

func (v *Validator) validateMapType(node *ast.MapTypeNode) *typing.Map {
	key := v.validateType(node.Key)
	value := v.validateType(node.Value)
	return &typing.Map{
		Key:   key,
		Value: value,
	}
}

func (v *Validator) validateFuncType(node *ast.FuncTypeNode) typing.Type {
	var params []typing.Type
	if node == nil {
		return typing.Invalid()
	}
	if node.Parameters != nil {
		for _, p := range node.Parameters {
			switch n := p.(type) {
			case *ast.PlainTypeNode:
				params = append(params, v.validateType(n))
				break
			case *ast.ExplicitVarDeclarationNode:
				t := v.validateType(n.DeclaredType)
				n.Resolved = t
				for _ = range n.Identifiers {
					params = append(params, t)
				}
				break
			}
		}
	}
	var results []typing.Type
	if node.Results != nil {
		for _, r := range node.Results {
			results = append(results, v.validateType(r))
		}
	}
	return &typing.Func{
		Params:  typing.NewTuple(params...),
		Results: typing.NewTuple(results...),
	}
}

func (v *Validator) validateDeclaration(node ast.Node) {
	switch n := node.(type) {
	case *ast.ClassDeclarationNode:
		v.validateClassDeclaration(n)
		break
	case *ast.ContractDeclarationNode:
		v.validateContractDeclaration(n)
		break
	case *ast.EnumDeclarationNode:
		v.validateEnumDeclaration(n)
		break
	case *ast.FuncDeclarationNode:
		v.validateFuncDeclaration(n)
		break
	case *ast.InterfaceDeclarationNode:
		v.validateInterfaceDeclaration(n)
		break
	case *ast.ExplicitVarDeclarationNode:
		v.validateVarDeclaration(n)
		break
	case *ast.EventDeclarationNode:
		v.validateEventDeclaration(n)
		break
	case *ast.TypeDeclarationNode:
		v.validateTypeDeclaration(n)
		break
	case *ast.LifecycleDeclarationNode:
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

func (v *Validator) validateVarDeclaration(node *ast.ExplicitVarDeclarationNode) {

	v.validateModifiers(ast.ExplicitVarDeclaration, node.Modifiers.Modifiers)

	typ := v.validateType(node.DeclaredType)
	typ.SetModifiers(&node.Modifiers)

	for _, id := range node.Identifiers {
		v.declareContextualVar(id, typ)
		//fmt.Printf("Declared: %s as %s\n", id, WriteType(typ))
	}
	node.Resolved = typ
}

func (v *Validator) validateGenerics(generics []*ast.GenericDeclarationNode) []*typing.Generic {
	genericTypes := make([]*typing.Generic, 0)
	for _, node := range generics {

		g := new(typing.Generic)

		v.declareContextualType(node.Identifier, g)

		var interfaces []*typing.Interface
		for _, ifc := range node.Implements {
			t := v.validatePlainType(ifc)
			if t != typing.Unknown() {
				if c, ok := t.(*typing.Interface); ok {
					interfaces = append(interfaces, c)
				} else {
					v.addError(g.Start.Line, errTypeRequired, makeName(ifc.Names), "interface")
				}
			}
		}
		var inherits []typing.Type
		for _, super := range node.Inherits {
			inherits = append(inherits, v.validatePlainType(super))
		}
		// enforce that all inheritors are the same type

		*g = typing.Generic{
			Identifier: node.Identifier,
			Interfaces: interfaces,
			Inherits:   inherits,
		}
		genericTypes = append(genericTypes, g)
	}

	for _, g := range genericTypes {
		v.validateGenericInherits(g.Inherits)
	}
	return genericTypes
}

func (v *Validator) validateGenericInherits(types []typing.Type) {
	var typ typing.Type
	for i, t := range types {
		switch t.(type) {
		case *typing.Contract, *typing.Class, *typing.Enum, *typing.Interface:
			break
		default:
			v.addError(errInvalidInheritance, typing.WriteType(t))
			break
		}
		if i == 0 {
			typ = t
		} else {
			if reflect.TypeOf(typing.ResolveUnderlying(t)) != reflect.TypeOf(typing.ResolveUnderlying(typ)) {
				v.addError(errIncompatibleInheritance, typing.WriteType(typ), typing.WriteType(t))
			}
		}
	}
}

func (v *Validator) validateClassesCancellation(parent *typing.Class, classes []*typing.Class) {
	props := map[string]bool{}
	for _, c := range classes {
		for k, _ := range c.Properties {
			// check if it already exists
			if props[k] {
				if parent.Cancelled == nil {
					parent.Cancelled = make(typing.CancellationMap)
				}
				parent.Cancelled[k] = true
			} else {
				props[k] = true
			}
		}
	}
}

func (v *Validator) validateContractsCancellation(parent *typing.Contract, contracts []*typing.Contract) {
	props := map[string]bool{}
	for _, c := range contracts {
		for k, _ := range c.Properties {
			// check if it already exists
			if props[k] {
				if parent.Cancelled == nil {
					parent.Cancelled = make(typing.CancellationMap)
				}
				parent.Cancelled[k] = true
			} else {
				props[k] = true
			}
		}
	}
}

func (v *Validator) validateEnumsCancellation(parent *typing.Enum, enums []*typing.Enum) {
	props := map[string]bool{}
	for _, e := range enums {
		for _, i := range e.Items {
			// check if it already exists
			if props[i] {
				if parent.Cancelled == nil {
					parent.Cancelled = make(typing.CancellationMap)
				}
				parent.Cancelled[i] = true
			} else {
				props[i] = true
			}
		}
	}
}

func (v *Validator) validateClassDeclaration(node *ast.ClassDeclarationNode) {

	v.validateModifiers(ast.ClassDeclaration, node.Modifiers.Modifiers)

	v.validateAnnotations(ast.ClassDeclaration, node.Modifiers.Annotations)

	generics := v.validateGenerics(node.Generics)

	var supers []*typing.Class
	for _, super := range node.Supers {
		t := v.validatePlainType(super)
		if t != typing.Unknown() {
			if c, ok := t.(*typing.Class); ok {
				supers = append(supers, c)
			} else {
				v.addError(errTypeRequired, makeName(super.Names), "class")
			}
		}
	}

	var interfaces []*typing.Interface
	for _, ifc := range node.Interfaces {
		t := v.validatePlainType(ifc)
		if t != typing.Unknown() {
			if c, ok := t.(*typing.Interface); ok {
				interfaces = append(interfaces, c)
			} else {
				v.addError(errTypeRequired, makeName(ifc.Names), "interface")
			}
		}
	}

	classType := &typing.Class{
		Name:       node.Identifier,
		Supers:     supers,
		Interfaces: interfaces,
		Generics:   generics,
		Mods:       &node.Modifiers,
	}

	node.Resolved = classType

	v.validateClassesCancellation(classType, supers)

	types, properties, lifecycles := v.validateScope(node, node.Body)

	classType.Types = types
	classType.Properties = properties
	classType.Lifecycles = lifecycles

	v.validateClassInterfaces(classType)

	v.declareContextualType(node.Identifier, classType)

}

func (v *Validator) validateEnumDeclaration(node *ast.EnumDeclarationNode) {

	v.validateModifiers(ast.EnumDeclaration, node.Modifiers.Modifiers)

	v.validateAnnotations(ast.EnumDeclaration, node.Modifiers.Annotations)

	var supers []*typing.Enum
	for _, super := range node.Inherits {
		t := v.validatePlainType(super)
		if c, ok := t.(*typing.Enum); ok {
			supers = append(supers, c)
		} else {
			v.addError(errTypeRequired, makeName(super.Names), "enum")
		}
	}

	list := node.Enums

	enumType := &typing.Enum{
		Name:   node.Identifier,
		Supers: supers,
		Items:  list,
		Mods:   &node.Modifiers,
	}

	v.validateEnumsCancellation(enumType, supers)

	node.Resolved = enumType

	v.declareContextualType(node.Identifier, enumType)
}

func (v *Validator) validateContractDeclaration(node *ast.ContractDeclarationNode) {

	v.validateModifiers(ast.ContractDeclaration, node.Modifiers.Modifiers)

	v.validateAnnotations(ast.ContractDeclaration, node.Modifiers.Annotations)

	var supers []*typing.Contract
	for _, super := range node.Supers {
		t := v.validatePlainType(super)
		if t != typing.Unknown() {
			if c, ok := t.(*typing.Contract); ok {
				supers = append(supers, c)
			} else {
				v.addError(errTypeRequired, makeName(super.Names), "contract")
			}
		}
	}

	var interfaces []*typing.Interface
	for _, ifc := range node.Interfaces {
		t := v.validatePlainType(ifc)
		if t != typing.Unknown() {
			if c, ok := t.(*typing.Interface); ok {
				interfaces = append(interfaces, c)
			} else {
				v.addError(errTypeRequired, makeName(ifc.Names), "interface")
			}
		}
	}

	generics := v.validateGenerics(node.Generics)

	contractType := &typing.Contract{
		Name:       node.Identifier,
		Generics:   generics,
		Supers:     supers,
		Interfaces: interfaces,

		Mods: &node.Modifiers,
	}

	node.Resolved = contractType

	v.validateContractsCancellation(contractType, supers)

	types, properties, lifecycles := v.validateScope(node, node.Body)

	contractType.Types = types
	contractType.Properties = properties
	contractType.Lifecycles = lifecycles

	v.validateContractInterfaces(contractType)

	v.declareContextualType(node.Identifier, contractType)
}

func (v *Validator) validateContractInterfaces(contract *typing.Contract) {
	for _, i := range contract.Interfaces {
		v.validateContractInterface(contract, i)
	}
}

func (v *Validator) validateClassInterfaces(class *typing.Class) {
	for _, i := range class.Interfaces {
		v.validateClassInterface(class, i)
	}
}

func hasContractFunction(contract *typing.Contract, name string, funcType *typing.Func) bool {
	if typ, ok := contract.Properties[name]; ok {
		// they share a name, now compare types
		if funcType.Compare(typ) {
			return true
		}
	}
	for _, sup := range contract.Supers {
		if hasContractFunction(sup, name, funcType) {
			return true
		}
	}
	return false
}

func hasClassFunction(class *typing.Class, name string, funcType *typing.Func) bool {
	if typ, ok := class.Properties[name]; ok {
		// they share a name, now compare types
		if funcType.Compare(typ) {
			return true
		}
	}
	for _, sup := range class.Supers {
		if hasClassFunction(sup, name, funcType) {
			return true
		}
	}
	return false
}

func (v *Validator) validateContractInterface(contract *typing.Contract, ifc *typing.Interface) {
	for f, t := range ifc.Funcs {
		if !hasContractFunction(contract, f, t) {
			v.addError(errUnimplementedInterface, contract.Name, ifc.Name, typing.WriteType(t))
		}
	}
	for _, super := range ifc.Supers {
		v.validateContractInterface(contract, super)
	}
}

func (v *Validator) validateClassInterface(class *typing.Class, ifc *typing.Interface) {
	for f, t := range ifc.Funcs {
		if !hasClassFunction(class, f, t) {
			v.addError(errUnimplementedInterface, class.Name, ifc.Name, typing.WriteType(t))
		}
	}
	for _, super := range ifc.Supers {
		v.validateClassInterface(class, super)
	}
}

func (v *Validator) validateInterfaceDeclaration(node *ast.InterfaceDeclarationNode) {

	v.validateModifiers(ast.InterfaceDeclaration, node.Modifiers.Modifiers)

	v.validateAnnotations(ast.InterfaceDeclaration, node.Modifiers.Annotations)

	var supers []*typing.Interface
	for _, super := range node.Supers {
		t := v.validatePlainType(super)
		if t != typing.Unknown() {
			if c, ok := t.(*typing.Interface); ok {
				supers = append(supers, c)
			} else {
				v.addError(errTypeRequired, makeName(super.Names), "interface")
			}
		}
	}

	funcs := map[string]*typing.Func{}
	for _, function := range node.Signatures {
		f, ok := v.validateContextualType(function).(*typing.Func)
		if ok {
			funcs[function.Identifier] = f
		} else {
			v.addError(errInvalidFuncType)
		}
	}

	generics := v.validateGenerics(node.Generics)

	interfaceType := &typing.Interface{
		Name:     node.Identifier,
		Generics: generics,
		Supers:   supers,
		Funcs:    funcs,
		Mods:     &node.Modifiers,
	}

	node.Resolved = interfaceType

	v.declareContextualType(node.Identifier, interfaceType)

}

func (v *Validator) validateContextualType(node ast.Node) typing.Type {
	return v.validateType(node)
}

func (v *Validator) declareContextualVar(name string, typ typing.Type) {
	if v.scope == nil {
		v.DeclareBuiltinOfType(name, typ)
	} else {
		v.DeclareVarOfType(name, typ)
	}
}

func (v *Validator) validateFuncDeclaration(node *ast.FuncDeclarationNode) {

	v.validateModifiers(ast.FuncDeclaration, node.Modifiers.Modifiers)

	v.validateAnnotations(ast.FuncDeclaration, node.Modifiers.Annotations)

	var params []typing.Type
	for _, node := range node.Signature.Parameters {
		// todo: check here?
		p := node.(*ast.ExplicitVarDeclarationNode)
		for _ = range p.Identifiers {
			typ := v.validateContextualType(p.DeclaredType)
			//TODO: declare this later
			//v.DeclareVarOfType(i, typ)
			params = append(params, typ)
		}
	}

	var results []typing.Type
	for _, r := range node.Signature.Results {
		results = append(results, v.validateType(r))
	}

	generics := v.validateGenerics(node.Generics)

	funcType := &typing.Func{
		Generics: generics,
		Params:   typing.NewTuple(params...),
		Results:  typing.NewTuple(results...),
		Mods:     &node.Modifiers,
	}

	node.Resolved = funcType
	v.declareContextualVar(node.Signature.Identifier, funcType)
	v.validateScope(node, node.Body)

}

func (v *Validator) validateAnnotations(typ ast.NodeType, annotations []*typing.Annotation) {
	// doesn't do anything yet?
}

func (v *Validator) validateModifiers(typ ast.NodeType, modifiers []string) {
	for _, mg := range v.modifierGroups {
		mg.reset()
	}
	for _, mod := range modifiers {
		for _, mg := range v.modifierGroups {
			if mg.has(mod) {
				if len(mg.selected) == mg.Maximum {
					v.addError(errMutuallyExclusiveModifiers)
				}
				mg.selected = append(mg.selected, mod)
			}
		}
	}
	for _, mg := range v.modifierGroups {
		if mg.requiredOn(typ) {
			if mg.selected == nil {
				v.addError(errRequiredModifier, mg.Name)
			}
		}
	}
}

func (v *Validator) processModifier(n, c token.Type) token.Type {
	if c == -1 {
		return n
	} else if n == c {
		v.addError(errDuplicateModifiers)
	} else {
		v.addError(errMutuallyExclusiveModifiers)
	}
	return c
}

func (v *Validator) validateEventDeclaration(node *ast.EventDeclarationNode) {

	v.validateModifiers(ast.EventDeclaration, node.Modifiers.Modifiers)

	var params []typing.Type
	for _, n := range node.Parameters {
		params = append(params, v.validateType(n.DeclaredType))
	}

	generics := v.validateGenerics(node.Generics)

	eventType := &typing.Event{
		Name:       node.Identifier,
		Generics:   generics,
		Parameters: typing.NewTuple(params...),
		Mods:       &node.Modifiers,
	}
	node.Resolved = eventType
	v.declareContextualVar(node.Identifier, eventType)
}

func (v *Validator) declareContextualType(name string, typ typing.Type) {
	if v.isParsingBuiltins {
		v.DeclareBuiltinType(name, typ)
	} else {
		v.DeclareType(name, typ)
	}
}

func (v *Validator) validateTypeDeclaration(node *ast.TypeDeclarationNode) {

	v.validateModifiers(ast.TypeDeclaration, node.Modifiers.Modifiers)

	typ := v.validateType(node.Value)
	node.Resolved = typ
	v.declareContextualType(node.Identifier, typ)
}

func (v *Validator) validateLifecycleDeclaration(node *ast.LifecycleDeclarationNode) {
	// TODO: enforce location
	var types []typing.Type
	for _, p := range node.Parameters {
		typ := v.validateType(p.DeclaredType)
		for _, i := range p.Identifiers {
			v.declareContextualVar(i, typ)
			types = append(types, typ)
		}
	}
	v.validateScope(node, node.Body)
	l := typing.Lifecycle{
		Type:       node.Category,
		Parameters: types,
	}

	v.declareLifecycle(node.Category, l)
}
