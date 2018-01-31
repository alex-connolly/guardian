package validator

import (
	"github.com/end-r/guardian/token"

	"github.com/end-r/guardian/util"

	"github.com/end-r/guardian/typing"

	"github.com/end-r/guardian/ast"
)

func (v *Validator) resolveType(node ast.Node) typing.Type {
	if node == nil {
		return typing.Invalid()
	}
	switch t := node.(type) {
	case *ast.PlainTypeNode:
		return v.resolvePlainType(t)
	case *ast.MapTypeNode:
		return v.resolveMapType(t)
	case *ast.ArrayTypeNode:
		return v.resolveArrayType(t)
	case *ast.FuncTypeNode:
		return v.resolveFuncType(t)
	}
	return typing.Invalid()
}

func (v *Validator) resolvePlainType(node *ast.PlainTypeNode) typing.Type {
	typ, _ := v.isTypeVisible(node.Names[0])
	if typ == typing.Unknown() {
		v.addError(node.Start(), errTypeNotVisible, makeName(node.Names))
		return typ
	}
	for _, n := range node.Names[1:] {

		t, ok := v.getTypeType(node.Start(), typ, n)
		if !ok {
			v.addError(node.Start(), errInvalidTypeType, typing.WriteType(typ), n)
			break
		}
		typ = t
	}
	return typ
}

func (v *Validator) resolveArrayType(node *ast.ArrayTypeNode) *typing.Array {
	a := new(typing.Array)
	a.Length = node.Length
	a.Variable = node.Variable
	a.Value = v.resolveType(node.Value)
	return a
}

func (v *Validator) resolveMapType(node *ast.MapTypeNode) *typing.Map {
	m := new(typing.Map)
	m.Key = v.resolveType(node.Key)
	m.Value = v.resolveType(node.Value)
	return m
}

func (v *Validator) resolveFuncType(node *ast.FuncTypeNode) *typing.Func {
	f := new(typing.Func)
	f.Params = v.resolveTuple(node.Parameters)
	f.Results = v.resolveTuple(node.Results)
	return f
}

func (v *Validator) resolveTuple(nodes []ast.Node) *typing.Tuple {
	t := new(typing.Tuple)
	t.Types = make([]typing.Type, len(nodes))
	for i, n := range nodes {
		t.Types[i] = v.resolveType(n)
	}
	return t
}

func (v *Validator) resolveExpression(e ast.ExpressionNode) typing.Type {
	if e == nil {
		return typing.Invalid()
	}
	switch n := e.(type) {
	case *ast.LiteralNode:
		return v.resolveLiteralExpression(n)
	case *ast.BinaryExpressionNode:
		return v.resolveBinaryExpression(n)
	case *ast.UnaryExpressionNode:
		return v.resolveUnaryExpression(n)
	case *ast.CallExpressionNode:
		return v.resolveCallExpression(n)
	case *ast.ArrayLiteralNode:
		return v.resolveArrayLiteral(n)
	case *ast.MapLiteralNode:
		return v.resolveMapLiteral(n)
	case *ast.CompositeLiteralNode:
		return v.resolveCompositeLiteral(n)
	case *ast.ReferenceNode:
		return v.resolveReference(n)
	case *ast.IdentifierNode:
		return v.resolveIdentifier(n)
	case *ast.IndexExpressionNode:
		return v.resolveIndexExpression(n)
	case *ast.FuncLiteralNode:
		return v.resolveFuncLiteralExpression(n)
	case *ast.SliceExpressionNode:
		return v.resolveSliceExpression(n)
	case *ast.KeywordNode:
		return v.resolveKeywordExpression(n)
		// resolve types as well - for casting purposes
	}
	switch e.(type) {
	case *ast.PlainTypeNode, *ast.FuncTypeNode, *ast.ArrayTypeNode, *ast.MapTypeNode:
		return v.resolveType(e)

	}
	v.addError(e.Start(), errUnknownExpressionType)
	return typing.Invalid()
}

func (v *Validator) resolveKeywordExpression(n *ast.KeywordNode) typing.Type {
	t := v.validateType(n.TypeNode)
	args := v.ExpressionTuple(n.Arguments)
	switch a := t.(type) {
	case *typing.Class:
		constructors := a.Lifecycles[token.Constructor]
		if typing.NewTuple().Compare(args) && len(constructors) == 0 {
			return t
		}
		for _, c := range constructors {
			paramTuple := typing.NewTuple(c.Parameters...)
			if paramTuple.Compare(args) {
				return t
			}
		}
		v.addError(n.Start(), errInvalidConstructorCall, typing.WriteType(a), typing.WriteType(args))
		return t
	case *typing.Contract:
		constructors := a.Lifecycles[token.Constructor]
		if typing.NewTuple().Compare(args) && len(constructors) == 0 {
			return t
		}
		for _, c := range constructors {
			paramTuple := typing.NewTuple(c.Parameters...)
			if paramTuple.Compare(args) {
				return t
			}
		}
		v.addError(n.Start(), errInvalidConstructorCall, typing.WriteType(a), typing.WriteType(args))
		break
	}
	// TODO: error here?
	return typing.Invalid()
}

func (v *Validator) resolveThis(node *ast.IdentifierNode) (typing.Type, map[string]typing.Type) {
	for c := v.scope; c != nil; c = c.parent {
		switch a := c.context.(type) {
		case *ast.ClassDeclarationNode:
			return a.Resolved, c.variables
		case *ast.ContractDeclarationNode:
			return a.Resolved, c.variables
		}
	}
	v.addError(node.Start(), errInvalidThisContext)
	return typing.Invalid(), nil
}

func (v *Validator) resolveIdentifier(n *ast.IdentifierNode) typing.Type {

	if n.Name == "this" {
		t, _ := v.resolveThis(n)
		return t
	}

	// look up the identifier in scope
	t, ok := v.isVarVisible(n.Name)
	if t == typing.Unknown() || !ok {
		t, ok = v.isTypeVisible(n.Name)
		if t != nil {
			typing.AddModifier(t, "static")
		}
	}
	n.Resolved = t
	return t
}

func (v *Validator) resolveLiteralExpression(n *ast.LiteralNode) typing.Type {

	if literalResolver, ok := v.literals[n.LiteralType]; ok {
		t := literalResolver(v, n.Data)
		n.Resolved = t
		return n.Resolved
	}
	n.Resolved = typing.Invalid()
	return n.Resolved
}

func (v *Validator) resolveArrayLiteral(n *ast.ArrayLiteralNode) typing.Type {
	if n.Signature.Length > 0 {
		if n.Signature.Length != len(n.Data) {
			v.addError(n.Signature.Start(), errInvalidArrayLiteralLength, len(n.Data), n.Signature.Length)
		}
	}
	value := v.validateType(n.Signature.Value)
	for _, val := range n.Data {
		valueType := v.validateType(val)
		if typing.AssignableTo(value, valueType, false) {
			v.addError(val.Start(), errInvalidArrayLiteralValue, typing.WriteType(valueType), typing.WriteType(value))
		}
	}
	arrayType := &typing.Array{
		Value:    value,
		Length:   n.Signature.Length,
		Variable: n.Signature.Variable,
	}
	return arrayType
}

func (v *Validator) resolveCompositeLiteral(n *ast.CompositeLiteralNode) typing.Type {
	n.Resolved = v.resolvePlainType(n.TypeName)
	for f, exp := range n.Fields {
		switch cType := n.Resolved.(type) {
		case *typing.Class:
			if t, ok := v.getClassProperty(n.Start(), cType, f); ok {
				r := v.resolveExpression(exp)
				if !typing.AssignableTo(t, r, false) {
					v.addError(n.Start(), errInvalidCompositeLiteralFieldValue, typing.WriteType(n.Resolved), f, typing.WriteType(t), typing.WriteType(r))
				}
			} else {
				v.addError(n.Start(), errInvalidCompositeLiteralFieldName, typing.WriteType(cType), f)
			}
			break
		case *typing.Contract:
			if t, ok := v.getContractProperty(n.Start(), cType, f); ok {
				r := v.resolveExpression(exp)
				if !typing.AssignableTo(t, r, false) {
					v.addError(n.Start(), errInvalidCompositeLiteralFieldValue, typing.WriteType(n.Resolved), f, typing.WriteType(t), typing.WriteType(r))
				}
			} else {
				v.addError(n.Start(), errInvalidCompositeLiteralFieldName, typing.WriteType(cType), f)
			}
			break
		}
	}
	return n.Resolved
}

func (v *Validator) resolveFuncLiteralExpression(n *ast.FuncLiteralNode) typing.Type {

	v.openScope(nil, nil)

	generics := v.validateGenerics(n.Generics)

	params := make([]typing.Type, 0)
	for _, p := range n.Parameters {
		typ := v.validateType(p.DeclaredType)
		for _, i := range p.Identifiers {
			v.declareVar(p.Start(), i, typ)
			params = append(params, typ)
		}
	}

	results := make([]typing.Type, 0)
	for _, p := range n.Results {
		switch n := p.(type) {
		case *ast.ExplicitVarDeclarationNode:
			dec := v.validateType(n.DeclaredType)
			for _, i := range n.Identifiers {
				v.declareVar(p.Start(), i, dec)
				results = append(results, dec)
			}
			break
		default:
			typ := v.validateType(p)
			results = append(results, typ)
			break
		}
	}

	n.Resolved = &typing.Func{
		Generics: generics,
		Params:   typing.NewTuple(params...),
		Results:  typing.NewTuple(results...),
	}

	v.validateScope(n, n.Scope)

	v.closeScope()

	return n.Resolved
}

func isValidMapKey(key typing.Type) bool {
	switch t := typing.ResolveUnderlying(key).(type) {
	// numeric types
	case *typing.NumericType:
		return true
	case *typing.Array:
		return isValidMapKey(t.Value)
	}
	return false
}

func (v *Validator) resolveMapLiteral(n *ast.MapLiteralNode) typing.Type {

	key := v.validateType(n.Signature.Key)
	if !isValidMapKey(key) {
		v.addError(n.Signature.Key.Start(), errInvalidMapKey, typing.WriteType(key))
	}
	value := v.validateType(n.Signature.Value)
	for k, val := range n.Data {
		keyType := v.resolveExpression(k)
		valueType := v.resolveExpression(val)
		if !typing.AssignableTo(key, keyType, false) {
			v.addError(val.Start(), errInvalidMapLiteralKey, typing.WriteType(valueType), typing.WriteType(value))
		}
		if !typing.AssignableTo(value, valueType, false) {
			v.addError(val.Start(), errInvalidMapLiteralValue, typing.WriteType(valueType), typing.WriteType(value))
		}
	}
	mapType := &typing.Map{Key: key, Value: value}
	n.Resolved = mapType
	return n.Resolved
}

func (v *Validator) resolveIndexExpression(n *ast.IndexExpressionNode) typing.Type {
	// must be literal
	exprType := v.resolveExpression(n.Expression)
	// enforce that this must be an array/map type
	switch t := exprType.(type) {
	case *typing.Array:
		n.Resolved = t.Value
		break
	case *typing.Map:
		n.Resolved = t.Value
		break
	default:
		n.Resolved = typing.Invalid()
		break
	}
	return n.Resolved
}

func (v *Validator) resolveAsPlainType(e ast.ExpressionNode) (typing.Type, bool) {
	switch n := e.(type) {
	case *ast.IdentifierNode:
		return v.isTypeVisible(n.Name)
	case *ast.ReferenceNode:
		switch a := n.Parent.(type) {
		case *ast.IdentifierNode:
			if typ, ok := v.isTypeVisible(a.Name); ok {
				return v.resolveContextualReference(typ, n.Parent, n.Reference), true
			}
		}
	}
	return typing.Invalid(), false
}

func (v *Validator) replaceGeneric(t typing.Type, genDecs typing.TypeMap) typing.Type {
	switch a := t.(type) {
	case *typing.Generic:
		if typ, ok := genDecs[a.Identifier]; ok {
			return typ
		}
		return t
	case *typing.Array:
		a.Value = v.replaceGeneric(a.Value, genDecs)
		return a
	case *typing.Map:
		a.Key = v.replaceGeneric(a.Key, genDecs)
		a.Value = v.replaceGeneric(a.Value, genDecs)
		break
	case *typing.Func:
		for i, p := range a.Params.Types {
			a.Params.Types[i] = v.replaceGeneric(p, genDecs)
		}
		for i, p := range a.Results.Types {
			a.Results.Types[i] = v.replaceGeneric(p, genDecs)
		}
	}
	return t
}

func (v *Validator) resolveCallExpression(n *ast.CallExpressionNode) typing.Type {

	// attempt to resolve as if cast
	if left, ok := v.resolveAsPlainType(n.Call); ok {
		if len(n.Arguments) > 1 {
			v.addError(n.Call.Start(), errMultipleCast)
			n.Resolved = left
			return left
		}
		t := v.resolveExpression(n.Arguments[0])
		if t == typing.Unknown() || t == typing.Invalid() || t == nil {
			//TODO: change this error?
			v.addError(n.Arguments[0].Start(), errImpossibleCastToNonType)
			n.Resolved = left
			return left
		}
		if !v.vm.Castable(v, left, t, n.Arguments[0]) {
			v.addError(n.Arguments[0].Start(), errImpossibleCast, typing.WriteType(t), typing.WriteType(left))
		}
		n.Resolved = left
		return left
	}

	exprType := v.resolveExpression(n.Call)
	args := v.ExpressionTuple(n.Arguments)

	switch a := exprType.(type) {
	case *typing.Func:
		genDecs := make(typing.TypeMap)
		if len(a.Generics) > 0 {
			// implicit generics (takes first type)
			if len(a.Params.Types) != len(args.Types) {
				v.addError(n.Start(), errInvalidFuncCall, typing.WriteType(args), typing.WriteType(a))
			} else {
				for i, p := range a.Params.Types {
					switch g := p.(type) {
					case *typing.Generic:
						if t, ok := genDecs[g.Identifier]; ok {
							if !t.Compare(args.Types[i]) {
								v.addError(n.Start(), errInvalidFuncCall, typing.WriteType(args), typing.WriteType(a))
								n.Resolved = a.Results
								return a.Results
							}
						}
						if !g.Accepts(args.Types[i]) {
							v.addError(n.Start(), errInvalidFuncCall, typing.WriteType(args), typing.WriteType(a))
							n.Resolved = a.Results
							return a.Results
						}
						genDecs[g.Identifier] = args.Types[i]
						break
					default:
						break
					}
				}
			}
		} else {
			if !typing.AssignableTo(a.Params, args, false) {
				v.addError(n.Start(), errInvalidFuncCall, typing.WriteType(args), typing.WriteType(a))
			}
		}

		for i, r := range a.Results.Types {
			a.Results.Types[i] = v.replaceGeneric(r, genDecs)
		}
		n.Resolved = a.Results
		return a.Results
	case *typing.Event:
		if !typing.AssignableTo(a.Parameters, args, false) {
			v.addError(n.Start(), errInvalidFuncCall, typing.WriteType(args), typing.WriteType(a))
		}
		return typing.NewTuple()
	default:
		v.addError(n.Start(), errInvalidCall, typing.WriteType(exprType))
	}
	return typing.Invalid()
}

func (v *Validator) resolveSliceExpression(n *ast.SliceExpressionNode) typing.Type {
	// must be literal
	exprType := v.resolveExpression(n.Expression)
	// must be an array
	switch t := exprType.(type) {
	case *typing.Array:
		n.Resolved = t
		return n.Resolved
	}
	n.Resolved = typing.Invalid()
	return n.Resolved
}

func (v *Validator) resolveBinaryExpression(b *ast.BinaryExpressionNode) typing.Type {
	// rules for binary Expressions
	leftType := v.resolveExpression(b.Left)
	rightType := v.resolveExpression(b.Right)
	operatorFunc, ok := v.operators[b.Operator]
	if !ok {
		b.Resolved = typing.Invalid()
		return b.Resolved
	}
	t := operatorFunc(v, []typing.Type{leftType, rightType}, []ast.ExpressionNode{b.Left, b.Right})
	b.Resolved = t
	return b.Resolved
}

func (v *Validator) resolveUnaryExpression(n *ast.UnaryExpressionNode) typing.Type {
	operandType := v.resolveExpression(n.Operand)
	n.Resolved = operandType
	return operandType
}

func (v *Validator) determineType(t typing.Type, parent, exp ast.ExpressionNode) typing.Type {
	switch a := exp.(type) {
	case *ast.ReferenceNode:
		t = v.determineType(t, parent, a.Parent)
		return v.resolveContextualReference(t, parent, a.Reference)
	case *ast.IdentifierNode:
		return t
	case *ast.CallExpressionNode:
		switch f := t.(type) {
		case *typing.Func:
			return f.Results
		}
		break
	case *ast.IndexExpressionNode:
		switch f := t.(type) {
		case *typing.Map:
			return f.Value
		case *typing.Array:
			return f.Value
		}
	case *ast.SliceExpressionNode:
		switch f := t.(type) {
		case *typing.Array:
			return f
		default:
			break
		}
	default:
		v.addError(exp.Start(), errInvalidReference)
		return typing.Invalid()
	}
	return typing.Invalid()
}

func (v *Validator) findProperty(parent typing.Type, name string) (typing.Type, bool) {
	// we know the parent must be visible in scope
	var parentName string
	switch t := parent.(type) {
	case *typing.Func:
		parentName = t.Name
		break
	case *typing.Class:
		parentName = t.Name
		break
	case *typing.Contract:
		parentName = t.Name
		break
	default:
		return typing.Unknown(), false
	}
	var penultimate, ultimate *TypeScope
	for s := v.scope; s != nil; s = s.parent {
		if _, ok := s.types[parentName]; ok {
			if ultimate != nil {
				if ultimate.scopes != nil {
					for _, s := range ultimate.scopes {
						decl := s.GetDeclaration(name)
						if decl != nil {
							saved := v.scope
							v.scope = ultimate
							v.validateDeclaration(decl)
							v.scope = saved
							if t, ok := ultimate.variables[name]; ok {
								return t, ok
							}
							if t, ok := ultimate.types[name]; ok {
								typing.AddModifier(t, "static")
								return t, ok
							}
							return typing.Invalid(), false
						}
					}
				}
			}
		}
		ultimate = penultimate
		penultimate = s
	}
	return typing.Unknown(), false
}

func (v *Validator) resolveContextualReference(context typing.Type, parent, exp ast.ExpressionNode) typing.Type {
	if name, ok := getIdentifier(exp); ok {
		if t, ok := v.getTypeProperty(parent, exp, context, name); ok {
			if typing.HasModifier(context, "static") && !typing.HasModifier(t, "static") {
				v.addError(exp.Start(), errInvalidStaticReference)
			}
			return v.determineType(typing.ResolveUnderlying(t), parent, exp)
		} else {
			// TODO: what if the property just hasn't been discovered yet
			if t, ok := v.findProperty(context, name); ok {
				return t
			}
			v.addError(exp.Start(), errPropertyNotFound, typing.WriteType(context), name)
		}
	} else {
		v.addError(exp.Start(), errUnnamedReference)
	}
	return typing.Invalid()
}

func (v *Validator) resolveReference(n *ast.ReferenceNode) typing.Type {
	// must be reference
	context := v.resolveExpression(n.Parent)
	t := v.resolveContextualReference(context, n.Parent, n.Reference)
	n.Resolved = t
	return n.Resolved
}

func getIdentifier(exp ast.ExpressionNode) (string, bool) {
	switch exp.Type() {
	case ast.Identifier:
		i := exp.(*ast.IdentifierNode)
		return i.Name, true
	case ast.CallExpression:
		c := exp.(*ast.CallExpressionNode)
		return getIdentifier(c.Call)
	case ast.SliceExpression:
		s := exp.(*ast.SliceExpressionNode)
		return getIdentifier(s.Expression)
	case ast.IndexExpression:
		i := exp.(*ast.IndexExpressionNode)
		return getIdentifier(i.Expression)
	case ast.Reference:
		r := exp.(*ast.ReferenceNode)
		return getIdentifier(r.Parent)
	default:
		return "", false
	}
}

func (v *Validator) isCurrentContext(context typing.Type) bool {
	for c := v.scope; c != nil; c = c.parent {
		if c.context != nil {
			switch a := c.context.(type) {
			case *ast.ClassDeclarationNode:
				if a.Resolved.Compare(context) {
					return true
				}
				break
			case *ast.ContractDeclarationNode:
				if a.Resolved.Compare(context) {
					return true
				}
				break
			}
		}
	}
	return false
}

func (v *Validator) isCurrentContextOrSubclass(context typing.Type) bool {
	for c := v.scope; c != nil; c = c.parent {
		if c.context != nil {
			switch a := c.context.(type) {
			case *ast.ClassDeclarationNode:
				if a.Resolved.Compare(context) {
					return true
				}
				switch p := typing.ResolveUnderlying(a.Resolved).(type) {
				case *typing.Class:
					for _, c := range p.Supers {
						if c.Compare(context) {
							return true
						}
					}
				}

				break
			case *ast.ContractDeclarationNode:
				if a.Resolved.Compare(context) {
					return true
				}
				switch p := typing.ResolveUnderlying(a.Resolved).(type) {
				case *typing.Contract:
					for _, c := range p.Supers {
						if c.Compare(context) {
							return true
						}
					}
				}
				break
			}
		}
	}

	return false
}

func (v *Validator) checkVisible(loc util.Location, context, property typing.Type, name string) {
	if property.Modifiers() != nil {
		if property.Modifiers().HasModifier("private") {
			if !v.isCurrentContext(context) {
				v.addError(loc, errInvalidAccess, name, "private", typing.WriteType(property))
			}
		} else if property.Modifiers().HasModifier("protected") {
			if !v.isCurrentContextOrSubclass(context) {
				v.addError(loc, errInvalidAccess, name, "protected", typing.WriteType(property))
			}
		}
	}
}

func (v *Validator) getClassProperty(loc util.Location, class *typing.Class, name string) (typing.Type, bool) {
	for k, _ := range class.Cancelled {
		if k == name {
			v.addError(loc, errCancelledProperty, name, class.Name)
			return typing.Unknown(), false
		}
	}
	if p, has := class.Properties[name]; has {
		v.checkVisible(loc, class, p, name)
		return p, has
	}
	for _, super := range class.Supers {
		if c, ok := v.getClassProperty(loc, super, name); ok {
			return c, ok
		}
	}
	return nil, false
}

func (v *Validator) getContractProperty(loc util.Location, contract *typing.Contract, name string) (typing.Type, bool) {

	for k, _ := range contract.Cancelled {
		if k == name {
			v.addError(loc, errCancelledProperty, name, contract.Name)
			return typing.Unknown(), false
		}
	}

	if p, has := contract.Properties[name]; has {
		v.checkVisible(loc, contract, p, name)
		return p, has
	}
	for _, super := range contract.Supers {
		if c, ok := v.getContractProperty(loc, super, name); ok {
			return c, ok
		}
	}
	return nil, false
}

func (v *Validator) getPackageProperty(loc util.Location, pkg *typing.Package, name string) (typing.Type, bool) {
	if p, has := pkg.Variables[name]; has {
		v.checkVisible(loc, pkg, p, name)
		return p, has
	}
	return nil, false
}

func (v *Validator) getInterfaceProperty(loc util.Location, ifc *typing.Interface, name string) (typing.Type, bool) {

	for k, _ := range ifc.Cancelled {
		if k == name {
			v.addError(loc, errCancelledProperty, name, ifc.Name)
			return typing.Unknown(), false
		}
	}

	if p, has := ifc.Funcs[name]; has {
		return p, has
	}

	for _, super := range ifc.Supers {
		if c, ok := v.getInterfaceProperty(loc, super, name); ok {
			return c, ok
		}
	}
	return nil, false
}

func (v *Validator) getEnumProperty(loc util.Location, c *typing.Enum, name string) (typing.Type, bool) {
	for k, _ := range c.Cancelled {
		if k == name {
			v.addError(loc, errCancelledProperty, name, c.Name)
			t := typing.Unknown()
			typing.AddModifier(t, "static")
			return t, true
		}
	}

	for _, s := range c.Items {
		if s == name {
			// so that it can be referenced Day.Mon
			typing.AddModifier(c, "static")
			return c, true
		}
	}
	for _, s := range c.Supers {
		if a, ok := v.getEnumProperty(loc, s, name); ok {
			return a, ok
		}
	}
	return typing.Invalid(), false
}

func (v *Validator) checkThisProperty(parent ast.ExpressionNode, name string) (typing.Type, bool) {
	if parent != nil {
		switch p := parent.(type) {
		case *ast.IdentifierNode:
			if p.Name == "this" {
				_, vars := v.resolveThis(p)
				if t, ok := vars[name]; ok {
					return t, ok
				}
			}
		}
	}
	return typing.Invalid(), false
}

func (v *Validator) getPackageType(loc util.Location, p *typing.Package, name string) (typing.Type, bool) {
	t, ok := p.Types[name]
	return t, ok
}

func (v *Validator) getClassType(loc util.Location, c *typing.Class, name string) (typing.Type, bool) {
	t, ok := c.Types[name]
	return t, ok
}

func (v *Validator) getContractType(loc util.Location, c *typing.Contract, name string) (typing.Type, bool) {
	t, ok := c.Types[name]
	return t, ok
}

func (v *Validator) getTypeType(loc util.Location, t typing.Type, name string) (typing.Type, bool) {
	if t == nil {
		// TODO: do something?
		return typing.Invalid(), false
	}
	// only classes, interfaces, contracts and enums are subscriptable
	switch c := typing.ResolveUnderlying(t).(type) {
	case *typing.Class:
		return v.getClassType(loc, c, name)
	case *typing.Contract:
		return v.getContractType(loc, c, name)
	case *typing.Package:
		return v.getPackageType(loc, c, name)
	default:
		v.addError(loc, errInvalidSubscriptable, typing.WriteType(c))
		break
	}

	return typing.Invalid(), false
}

func (v *Validator) getTypeProperty(parent, exp ast.ExpressionNode, t typing.Type, name string) (typing.Type, bool) {
	if t == nil {
		// TODO: do something?
		return typing.Invalid(), false
	}
	// only classes, interfaces, contracts and enums are subscriptable
	switch c := typing.ResolveUnderlying(t).(type) {
	case *typing.Class:
		if t, ok := v.getClassProperty(exp.Start(), c, name); !ok {
			return v.checkThisProperty(parent, name)
		} else {
			return t, ok
		}
	case *typing.Contract:
		if t, ok := v.getContractProperty(exp.Start(), c, name); !ok {
			return v.checkThisProperty(parent, name)
		} else {
			return t, ok
		}
	case *typing.Interface:
		return v.getInterfaceProperty(exp.Start(), c, name)
	case *typing.Enum:
		return v.getEnumProperty(exp.Start(), c, name)
	case *typing.Package:
		return v.getPackageProperty(exp.Start(), c, name)
	case *typing.Tuple:
		if len(c.Types) == 1 {
			return v.getTypeProperty(parent, exp, c.Types[0], name)
		} else {
			v.addError(exp.Start(), errMultipleTypesInSingleValueContext)
		}
		break
	default:
		v.addError(exp.Start(), errInvalidSubscriptable, typing.WriteType(c))
		break
	}

	return typing.Invalid(), false
}
