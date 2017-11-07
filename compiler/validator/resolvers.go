package validator

import (
	"github.com/end-r/guardian/compiler/lexer"

	"github.com/end-r/guardian/compiler/ast"
)

func (v *Validator) resolveType(node ast.Node) Type {
	switch node.Type() {
	case ast.PlainType:
		r := node.(ast.PlainTypeNode)
		return v.resolvePlainType(r)
	case ast.MapType:
		m := node.(ast.MapTypeNode)
		return v.resolveMapType(m)
	case ast.ArrayType:
		a := node.(ast.ArrayTypeNode)
		return v.resolveArrayType(a)
	case ast.FuncType:
		f := node.(ast.FuncTypeNode)
		return v.resolveFuncType(f)
	}
	return standards[Invalid]
}

func (v *Validator) resolvePlainType(node ast.PlainTypeNode) Type {
	return v.getNamedType(node.Names...)
}

func (v *Validator) resolveArrayType(node ast.ArrayTypeNode) Array {
	a := Array{}
	a.Value = v.resolveType(node.Value)
	return a
}

func (v *Validator) resolveMapType(node ast.MapTypeNode) Map {
	m := Map{}
	m.Key = v.resolveType(node.Key)
	m.Value = v.resolveType(node.Value)
	return m
}

func (v *Validator) resolveFuncType(node ast.FuncTypeNode) Func {
	f := Func{}
	f.Params = v.resolveTuple(node.Parameters)
	f.Results = v.resolveTuple(node.Results)
	return f
}

func (v *Validator) resolveTuple(nodes []ast.Node) Tuple {
	t := Tuple{}
	t.types = make([]Type, len(nodes))
	for i, n := range nodes {
		t.types[i] = v.resolveType(n)
	}
	return t
}

func (v *Validator) resolveExpression(e ast.ExpressionNode) Type {
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
	}
	return resolvers[e.Type()](v, e)
}

type resolver func(v *Validator, e ast.ExpressionNode) Type

func resolveIdentifier(v *Validator, e ast.ExpressionNode) Type {
	i := e.(ast.IdentifierNode)
	// look up the identifier in scope
	typ := v.findVariable(i.Name)
	return typ
}

func resolveLiteralExpression(v *Validator, e ast.ExpressionNode) Type {
	// must be literal
	l := e.(ast.LiteralNode)
	switch l.LiteralType {
	case lexer.TknString:
		return standards[String]
	case lexer.TknTrue, lexer.TknFalse:
		return standards[Bool]
	case lexer.TknInteger:
		return standards[Int]
	case lexer.TknFloat:
		return standards[Float]
	}
	return standards[Invalid]
}

func resolveArrayLiteralExpression(v *Validator, e ast.ExpressionNode) Type {
	// must be literal
	m := e.(ast.ArrayLiteralNode)
	keyType := v.resolveType(m.Signature.Value)
	arrayType := NewArray(keyType)
	return arrayType
}

func resolveCompositeLiteral(v *Validator, e ast.ExpressionNode) Type {
	c := e.(ast.CompositeLiteralNode)
	return v.getNamedType(c.TypeName)
}

func resolveFuncLiteralExpression(v *Validator, e ast.ExpressionNode) Type {
	// must be func literal
	f := e.(ast.FuncLiteralNode)
	var params, results []Type
	for _, p := range f.Signature.Parameters {
		params = append(params, v.resolveType(p))
	}

	for _, r := range f.Signature.Results {
		results = append(results, v.resolveType(r))
	}
	return NewFunc(NewTuple(params...), NewTuple(results...))
}

func resolveMapLiteralExpression(v *Validator, e ast.ExpressionNode) Type {
	// must be literal
	m := e.(ast.MapLiteralNode)
	keyType := v.resolveType(m.Signature.Key)
	valueType := v.resolveType(m.Signature.Value)
	mapType := NewMap(keyType, valueType)
	return mapType
}

func resolveIndexExpression(v *Validator, e ast.ExpressionNode) Type {
	// must be literal
	i := e.(ast.IndexExpressionNode)
	exprType := v.resolveExpression(i.Expression)
	// enforce that this must be an array/map type
	switch exprType.(type) {
	case Array:
		return exprType.(Array).Value
	case Map:
		return exprType.(Map).Value
	}
	return standards[Invalid]
}

func resolveCallExpression(v *Validator, e ast.ExpressionNode) Type {
	// must be call expression
	c := e.(ast.CallExpressionNode)
	// return type of a call expression is always a tuple
	// tuple may be empty or single-valued
	call := v.resolveExpression(c.Call)
	// enforce that this is a function pointer (at whatever depth)
	under := resolveUnderlying(call)
	fn, ok := under.(Func)
	if ok {
		return fn.Results
	} else {
		v.addError(errCallExpressionNoFunc, WriteType(under))
		return standards[Invalid]
	}

}

func resolveUnderlying(t Type) Type {
	for al, ok := t.(Aliased); ok; al, ok = t.(Aliased) {
		t = al.underlying
	}
	return t
}

func resolveSliceExpression(v *Validator, e ast.ExpressionNode) Type {
	// must be literal
	s := e.(ast.SliceExpressionNode)
	exprType := v.resolveExpression(s.Expression)
	// must be an array
	switch exprType.(type) {
	case Array:
	}
	return standards[Invalid]
}

func resolveBinaryExpression(v *Validator, e ast.ExpressionNode) Type {
	// must be literal
	b := e.(ast.BinaryExpressionNode)
	// rules for binary Expressions
	//leftType := v.resolveExpression(b.Left)
	//rightType := v.resolveExpression(b.Right)

	switch b.Operator {
	case lexer.TknAdd:
		// can be numeric or a string
		if v.resolveExpression(b.Left).compare(standards[String]) {
			return standards[String]
		} else {
			return standards[Int]
		}
	case lexer.TknSub, lexer.TknDiv, lexer.TknMul, lexer.TknMod:
		// must be numeric
		return standards[Int]
	case lexer.TknGeq, lexer.TknLeq, lexer.TknLss, lexer.TknGtr:
		// must be numeric
		return standards[Bool]
	case lexer.TknEql, lexer.TknNeq:
		// don't have to be numeric
		return standards[Bool]
	case lexer.TknShl, lexer.TknShr, lexer.TknAnd, lexer.TknOr, lexer.TknXor:
		// must be numeric
		return standards[Int]
	case lexer.TknLogicalAnd, lexer.TknLogicalOr:
		// must be boolean
		return standards[Bool]
	}

	// else it is a type which is not defined for binary operators
	return standards[Invalid]
}

func resolveUnaryExpression(v *Validator, e ast.ExpressionNode) Type {
	m := e.(ast.UnaryExpressionNode)
	operandType := v.resolveExpression(m.Operand)
	return operandType
}

func (v *Validator) resolveContextualReference(context Type, exp ast.ExpressionNode) Type {
	// check if context is subscriptable
	if isSubscriptable(context) {
		if name, ok := getIdentifier(exp); ok {
			if _, ok := getPropertyType(context, name); ok {
				if exp.Type() == ast.Reference {
					a := exp.(ast.ReferenceNode)
					context := v.resolveExpression(a.Parent)
					return v.resolveContextualReference(context, a.Reference)
				} else {
					return v.resolveExpression(exp)
				}
			} else {
				v.addError(errPropertyNotFound, WriteType(context), name)
			}
		} else {
			v.addError(errUnnamedReference)
		}
	} else {
		v.addError(errInvalidSubscriptable, WriteType(context))
	}
	return standards[Invalid]
}

func resolveReference(v *Validator, e ast.ExpressionNode) Type {
	// must be reference
	m := e.(ast.ReferenceNode)
	context := v.resolveExpression(m.Parent)
	return v.resolveContextualReference(context, m.Reference)
}

func getIdentifier(exp ast.ExpressionNode) (string, bool) {
	switch exp.Type() {
	case ast.Identifier:
		i := exp.(ast.IdentifierNode)
		return i.Name, true
	case ast.CallExpression:
		c := exp.(ast.CallExpressionNode)
		return getIdentifier(c.Call)
	case ast.SliceExpression:
		s := exp.(ast.SliceExpressionNode)
		return getIdentifier(s.Expression)
	case ast.IndexExpression:
		i := exp.(ast.IndexExpressionNode)
		return getIdentifier(i.Expression)
	case ast.Reference:
		r := exp.(ast.ReferenceNode)
		return getIdentifier(r.Parent)
	default:
		return "", false
	}
}

func getPropertiesType(t Type, names []string) (Type, bool) {
	var working bool
	for _, name := range names {
		if !working {
			break
		}
		t, working = getPropertyType(t, name)
	}
	return t, working
}

func getPropertyType(t Type, name string) (Type, bool) {
	// only classes, interfaces, contracts and enums are subscriptable
	switch c := t.(type) {
	case Class:
		p, has := c.Properties[name]
		return p, has
	case Contract:
		p, has := c.Properties[name]
		return p, has
	case Interface:
		p, has := c.Funcs[name]
		return p, has
	case Enum:
		_, has := c.Items[name]
		return standards[Int], has
	}
	return standards[Invalid], false
}

func isSubscriptable(t Type) bool {
	// only classes, interfaces and enums are subscriptable
	switch t.(type) {
	case Class, Interface, Enum, Contract:
		return true
	}
	return false
}
