package validator

import (
	"fmt"

	"github.com/end-r/guardian/typing"

	"github.com/end-r/guardian/ast"
)

func (v *Validator) resolveType(node ast.Node) typing.Type {
	if node == nil {
		// ?
		return typing.Invalid()
	}
	switch node.Type() {
	case ast.PlainType:
		r := node.(*ast.PlainTypeNode)
		return v.resolvePlainType(r)
	case ast.MapType:
		m := node.(*ast.MapTypeNode)
		return v.resolveMapType(m)
	case ast.ArrayType:
		a := node.(*ast.ArrayTypeNode)
		return v.resolveArrayType(a)
	case ast.FuncType:
		f := node.(*ast.FuncTypeNode)
		return v.resolveFuncType(f)
	}
	return typing.Invalid()
}

func (v *Validator) resolvePlainType(node *ast.PlainTypeNode) typing.Type {
	return v.getNamedType(node.Names...)
}

func (v *Validator) resolveArrayType(node *ast.ArrayTypeNode) *typing.Array {
	a := new(typing.Array)
	a.Value = v.resolveType(node.Value)
	return a
}

func (v *Validator) resolveMapType(node *ast.MapTypeNode) *typing.Map {
	m := typing.Map{}
	m.Key = v.resolveType(node.Key)
	m.Value = v.resolveType(node.Value)
	return &m
}

func (v *Validator) resolveFuncType(node *ast.FuncTypeNode) *typing.Func {
	f := typing.Func{}
	f.Params = v.resolveTuple(node.Parameters)
	f.Results = v.resolveTuple(node.Results)
	return &f
}

func (v *Validator) resolveTuple(nodes []ast.Node) *typing.Tuple {
	t := typing.Tuple{}
	t.Types = make([]typing.Type, len(nodes))
	for i, n := range nodes {
		t.Types[i] = v.resolveType(n)
	}
	return &t
}

func (v *Validator) resolveExpression(e ast.ExpressionNode) typing.Type {
	if e == nil {
		return typing.Invalid()
	}
	resolvers := map[ast.NodeType]resolver{
		ast.Literal:          resolveLiteralExpression,
		ast.MapLiteral:       resolveMapLiteralExpression,
		ast.ArrayLiteral:     resolveArrayLiteralExpression,
		ast.FuncLiteral:      resolveFuncLiteralExpression,
		ast.IndexExpression:  resolveIndexExpression,
		ast.CallExpression:   resolveCallExpression,
		ast.SliceExpression:  resolveSliceExpression,
		ast.BinaryExpression: resolveBinaryExpression,
		ast.UnaryExpression:  resolveUnaryExpression,
		ast.Reference:        resolveReference,
		ast.Identifier:       resolveIdentifier,
		ast.CompositeLiteral: resolveCompositeLiteral,
		ast.Keyword:          resolveKeyword,
		ast.PlainType:        resolveUnknown,
		ast.FuncType:         resolveUnknown,
		ast.ArrayType:        resolveUnknown,
		ast.MapType:          resolveUnknown,
	}
	r, ok := resolvers[e.Type()]
	if !ok {
		v.addError(errUnknownExpressionType)
		return typing.Invalid()
	}
	return r(v, e)
}

type resolver func(v *Validator, e ast.ExpressionNode) typing.Type

func resolveUnknown(v *Validator, e ast.ExpressionNode) typing.Type {
	return typing.Unknown()
}

func resolveKeyword(v *Validator, e ast.ExpressionNode) typing.Type {
	i := e.(*ast.KeywordNode)
	// look up the identifier in scope
	t := v.resolveType(i.TypeNode)
	i.Resolved = t
	return t
}

func resolveIdentifier(v *Validator, e ast.ExpressionNode) typing.Type {
	i := e.(*ast.IdentifierNode)

	// look up the identifier in scope
	t := v.findVariable(i.Name)
	if t == typing.Unknown() {
		t = v.getNamedType(i.Name)
		t.MakeStatic()
	}
	i.Resolved = t
	return t
}

func resolveLiteralExpression(v *Validator, e ast.ExpressionNode) typing.Type {
	// must be literal
	l := e.(*ast.LiteralNode)
	literalResolver, ok := v.literals[l.LiteralType]
	if ok {
		t := literalResolver(v, l.Data)
		l.Resolved = t
		return l.Resolved
	}
	l.Resolved = typing.Invalid()
	return l.Resolved
}

func resolveArrayLiteralExpression(v *Validator, e ast.ExpressionNode) typing.Type {
	// must be literal
	m := e.(*ast.ArrayLiteralNode)
	keyType := v.resolveType(m.Signature.Value)
	arrayType := &typing.Array{
		Value:    keyType,
		Length:   m.Signature.Length,
		Variable: m.Signature.Variable,
	}
	return arrayType
}

func resolveCompositeLiteral(v *Validator, e ast.ExpressionNode) typing.Type {
	c := e.(*ast.CompositeLiteralNode)
	c.Resolved = v.getNamedType(c.TypeName)
	if c.Resolved == typing.Unknown() {
		c.Resolved = v.getDeclarationNode([]string{c.TypeName})
	}
	return c.Resolved
}

func resolveFuncLiteralExpression(v *Validator, e ast.ExpressionNode) typing.Type {
	// must be func literal
	f := e.(*ast.FuncLiteralNode)
	var params, results []typing.Type
	for _, p := range f.Parameters {
		typ := v.resolveType(p.DeclaredType)
		for _ = range p.Identifiers {
			params = append(params, typ)
		}
	}

	for _, r := range f.Results {
		results = append(results, v.resolveType(r))
	}
	f.Resolved = &typing.Func{
		Params:  typing.NewTuple(params...),
		Results: typing.NewTuple(results...),
	}

	return f.Resolved
}

func resolveMapLiteralExpression(v *Validator, e ast.ExpressionNode) typing.Type {
	// must be literal
	m := e.(*ast.MapLiteralNode)
	keyType := v.resolveType(m.Signature.Key)
	valueType := v.resolveType(m.Signature.Value)
	mapType := &typing.Map{Key: keyType, Value: valueType}
	m.Resolved = mapType
	return m.Resolved
}

func resolveIndexExpression(v *Validator, e ast.ExpressionNode) typing.Type {
	// must be literal
	i := e.(*ast.IndexExpressionNode)
	exprType := v.resolveExpression(i.Expression)
	// enforce that this must be an array/map type
	switch t := exprType.(type) {
	case *typing.Array:
		i.Resolved = t.Value
		break
	case *typing.Map:
		i.Resolved = t.Value
		break
	default:
		i.Resolved = typing.Invalid()
		break
	}
	return i.Resolved
}

/*
// attempts to resolve an expression component as a type name
// used in constructors e.g. Dog()
func (v *Validator) attemptToFindType(e ast.ExpressionNode) typing.Type {
	var names []string
	switch res := e.(type) {
	case ast.IdentifierNode:
		names = append(names, res.Name)
	case ast.ReferenceNode:
		var current ast.ExpressionNode
		for current = res; current != nil; current = res.Reference {
			switch a := current.(type) {
			case ast.ReferenceNode:
				if p, ok := a.Parent.(ast.IdentifierNode); ok {
					names = append(names, p.Name)
				} else {
					return typing.Unknown()
				}
				break
			case ast.IdentifierNode:
				names = append(names, a.Name)
				break
			default:
				return typing.Unknown()
			}
		}
		break
	default:
		return typing.Unknown()
	}
	return v.getNamedType(names...)
}*/
/*
func (v *Validator) resolveInContext(t typing.Type, property string) typing.Type {
	switch r := t.(type) {
	case typing.Class:
		t, ok := r.Types[property]
		if ok {
			return t
		}
		t, ok = r.Properties[property]
		if ok {
			return t
		}
		break
	case typing.Contract:
		t, ok := r.Types[property]
		if ok {
			return t
		}
		t, ok = r.Properties[property]
		if ok {
			return t
		}
		break
	case typing.Interface:
		t, ok := r.Funcs[property]
		if ok {
			return t
		}
		break
	case typing.Enum:
		for _, item := range r.Items {
			if item == property {
				return v.SmallestNumericType(typing.BitsNeeded(len(r.Items)), false)
			}
		}
		break
	}
	return typing.Unknown()
}*/

func resolveCallExpression(v *Validator, e ast.ExpressionNode) typing.Type {
	// must be call expression
	c := e.(*ast.CallExpressionNode)
	// return type of a call expression is always a tuple
	// tuple may be empty or single-valued
	call := v.resolveExpression(c.Call)
	var under typing.Type
	if call.Compare(typing.Unknown()) {
		// try to resolve as a type name
		switch n := c.Call.(type) {
		case *ast.IdentifierNode:
			under = v.getNamedType(n.Name)
		}

	} else {
		under = typing.ResolveUnderlying(call)
	}
	switch ctwo := under.(type) {
	case *typing.Func:
		c.Resolved = ctwo.Results
		return c.Resolved
	case *typing.Class:
		c.Resolved = ctwo
		return c.Resolved
	}
	v.addError(errCallExpressionNoFunc, typing.WriteType(call))
	c.Resolved = typing.Invalid()
	return c.Resolved

}

func resolveSliceExpression(v *Validator, e ast.ExpressionNode) typing.Type {
	// must be literal
	s := e.(*ast.SliceExpressionNode)
	exprType := v.resolveExpression(s.Expression)
	// must be an array
	switch t := exprType.(type) {
	case *typing.Array:
		s.Resolved = t
		return s.Resolved
	}
	s.Resolved = typing.Invalid()
	return s.Resolved
}

func resolveBinaryExpression(v *Validator, e ast.ExpressionNode) typing.Type {
	// must be literal
	b := e.(*ast.BinaryExpressionNode)
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

func resolveUnaryExpression(v *Validator, e ast.ExpressionNode) typing.Type {
	m := e.(*ast.UnaryExpressionNode)
	operandType := v.resolveExpression(m.Operand)
	m.Resolved = operandType
	return operandType
}

func (v *Validator) resolveContextualReference(context typing.Type, exp ast.ExpressionNode) typing.Type {
	// check if context is subscriptable
	if name, ok := getIdentifier(exp); ok {
		if t, ok := v.getTypeProperty(context, name); ok {
			if context.Static() && !t.Static() {
				v.addError(errInvalidStaticReference)
			}
			t = typing.ResolveUnderlying(t)
			switch a := exp.(type) {
			case *ast.ReferenceNode:
				context = v.resolveContextualReference(context, a.Parent)
				return v.resolveContextualReference(context, a.Reference)
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
				v.addError(errInvalidReference)
				return typing.Invalid()
			}
		} else {
			if context == nil {
				fmt.Println("NIL CONTEXT")
			} else {
				v.addError(errPropertyNotFound, typing.WriteType(context), name)
			}

		}
	} else {
		v.addError(errUnnamedReference)
	}
	return typing.Invalid()
}

func resolveReference(v *Validator, e ast.ExpressionNode) typing.Type {
	// must be reference
	m := e.(*ast.ReferenceNode)
	context := v.resolveExpression(m.Parent)
	t := v.resolveContextualReference(context, m.Reference)
	m.Resolved = t
	return m.Resolved
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

func (v *Validator) getPropertiesType(t typing.Type, names []string) (resolved typing.Type) {
	var working bool
	for _, name := range names {
		if !working {
			break
		}
		t, working = v.getTypeProperty(t, name)
	}
	return t
}

func (v *Validator) isCurrentContext(context typing.Type) bool {
	return v.context.Compare(context)
}

func (v *Validator) isCurrentContextOrSubclass(context typing.Type) bool {
	if v.isCurrentContext(context) {
		return true
	}
	switch a := context.(type) {
	case *typing.Class:
		if a.Supers != nil {
			for _, c := range a.Supers {
				if v.isCurrentContextOrSubclass(c) {
					return true
				}
			}
		}
		return false
	case *typing.Contract:
		if a.Supers != nil {
			for _, c := range a.Supers {
				if v.isCurrentContextOrSubclass(c) {
					return true
				}
			}
		}
		return false
	}
	return false
}

func (v *Validator) checkVisible(context, property typing.Type, name string) {
	if property != nil && property.Modifiers() != nil {
		if property.Modifiers().HasModifier("private") {
			if !v.isCurrentContext(context) {
				v.addError(errInvalidAccess, name, "private")
			}
		} else if property.Modifiers().HasModifier("protected") {
			if !v.isCurrentContextOrSubclass(context) {
				v.addError(errInvalidAccess, name, "protected")
			}
		}
	}
}

func (v *Validator) getClassProperty(class *typing.Class, name string) (typing.Type, bool) {
	for k, _ := range class.Cancelled {
		if k == name {
			v.addError(errCancelledProperty, name, class.Name)
			return typing.Invalid(), false
		}
	}
	if p, has := class.Properties[name]; has {
		v.checkVisible(class, p, name)
		return p, has
	}
	for _, super := range class.Supers {
		if c, ok := v.getClassProperty(super, name); ok {
			return c, ok
		}
	}
	return nil, false
}

func (v *Validator) getContractProperty(contract *typing.Contract, name string) (typing.Type, bool) {

	for k, _ := range contract.Cancelled {
		if k == name {
			v.addError(errCancelledProperty, name, contract.Name)
			return typing.Invalid(), false
		}
	}

	if p, has := contract.Properties[name]; has {
		v.checkVisible(contract, p, name)
		return p, has
	}
	for _, super := range contract.Supers {
		if c, ok := v.getContractProperty(super, name); ok {
			return c, ok
		}
	}
	return nil, false
}

func (v *Validator) getInterfaceProperty(ifc *typing.Interface, name string) (typing.Type, bool) {

	for k, _ := range ifc.Cancelled {
		if k == name {
			v.addError(errCancelledProperty, name, ifc.Name)
			return typing.Invalid(), false
		}
	}

	if p, has := ifc.Funcs[name]; has {
		return p, has
	}

	for _, super := range ifc.Supers {
		if c, ok := v.getInterfaceProperty(super, name); ok {
			return c, ok
		}
	}
	return nil, false
}

func (v *Validator) getEnumProperty(c *typing.Enum, name string) (typing.Type, bool) {
	for k, _ := range c.Cancelled {
		if k == name {
			v.addError(errCancelledProperty, name, c.Name)
			t := typing.Invalid()
			t.MakeStatic()
			return t, true
		}
	}

	for _, s := range c.Items {
		if s == name {
			t := v.SmallestNumericType(len(c.Items), false)
			// so that it can be referenced Day.Mon
			t.MakeStatic()
			return t, true
		}
	}
	for _, s := range c.Supers {
		if a, ok := v.getEnumProperty(s, name); ok {
			return a, ok
		}
	}
	return typing.Invalid(), false
}

func (v *Validator) getTypeProperty(t typing.Type, name string) (typing.Type, bool) {
	if t == nil {
		// TODO: do something?
		return typing.Invalid(), false
	}
	// only classes, interfaces, contracts and enums are subscriptable
	switch c := typing.ResolveUnderlying(t).(type) {
	case *typing.Class:
		return v.getClassProperty(c, name)
	case *typing.Contract:
		return v.getContractProperty(c, name)
	case *typing.Interface:
		return v.getInterfaceProperty(c, name)
	case *typing.Enum:
		return v.getEnumProperty(c, name)
	case *typing.Tuple:
		if len(c.Types) == 1 {
			return c.Types[0], true
		} else {
			v.addError(errMultipleTypesInSingleValueContext)
		}
		break
	default:
		v.addError(errInvalidSubscriptable, typing.WriteType(c))
		break
	}
	return typing.Invalid(), false
}
