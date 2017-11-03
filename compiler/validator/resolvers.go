package validator

import (
	"fmt"

	"github.com/end-r/guardian/compiler/lexer"

	"github.com/end-r/guardian/compiler/ast"
)

func (v *Validator) resolveType(node ast.Node) Type {
	if node != nil {
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
	}
	return standards[Invalid]
}

func (v *Validator) resolvePlainType(node ast.PlainTypeNode) Type {
	return v.findReference(node.Names...)
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

func (v *Validator) validateType(node ast.Node) {
	if node == nil {
		fmt.Println("val node nil")
	} else {
		switch node.Type() {
		case ast.PlainType:
			fmt.Println("plain type")
			typ := node.(ast.PlainTypeNode)
			v.requireVisibleType(typ.Names...)
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

}

func (v *Validator) resolveExpression(e ast.ExpressionNode) Type {
	resolvers := map[ast.NodeType]resolver{
		ast.Literal:          resolveLiteralExpression,
		ast.MapLiteral:       resolveMapLiteralExpression,
		ast.ArrayLiteral:     resolveArrayLiteralExpression,
		ast.IndexExpression:  resolveIndexExpression,
		ast.CallExpression:   resolveCallExpression,
		ast.SliceExpression:  resolveSliceExpression,
		ast.BinaryExpression: resolveBinaryExpression,
		ast.UnaryExpression:  resolveUnaryExpression,
		ast.Reference:        resolveReference,
		ast.Identifier:       resolveIdentifier,
	}
	return resolvers[e.Type()](v, e)
}

type resolver func(v *Validator, e ast.ExpressionNode) Type

func resolveIdentifier(v *Validator, e ast.ExpressionNode) Type {
	i := e.(ast.IdentifierNode)
	// look up the identifier in scope
	return v.findReference(i.Name)
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
	// enforce that this must be an array type
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
	fn := resolveUnderlying(call).(Func)
	return fn.Results
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
	leftType := v.resolveExpression(b.Left)
	rightType := v.resolveExpression(b.Left)
	if leftType.compare(standards[String]) {
		if !rightType.compare(standards[String]) {
			v.addError(errInvalidBinaryOpTypes, "placeholder", WriteType(leftType), WriteType(rightType))
		}
		return standards[String]
	} else if leftType.compare(standards[Int]) {
		if !rightType.compare(standards[Int]) {
			v.addError(errInvalidBinaryOpTypes, "placeholder", WriteType(leftType), WriteType(rightType))
		}
		return standards[Int]
	}

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
	case lexer.TknGeq, lexer.TknLeq, lexer.TknLss, lexer.TknGtr, lexer.TknEql:
		// must be numeric
		return standards[Int]
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
			if _, ok := getProperty(context, name); ok {
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
	case ast.CallExpression:
		c := exp.(ast.CallExpressionNode)
		return getIdentifier(c.Call)
	case ast.SliceExpression:
		s := exp.(ast.SliceExpressionNode)
		return getIdentifier(s.Expression)
	case ast.IndexExpression:
		i := exp.(ast.IndexExpressionNode)
		return getIdentifier(i.Expression)
	case ast.Identifier:
		i := exp.(ast.IdentifierNode)
		return i.Name, true
	case ast.Reference:
		r := exp.(ast.ReferenceNode)
		return getIdentifier(r.Parent)
	default:
		return "", false
	}
}

func getProperty(t Type, name string) (Type, bool) {
	// only classes, interfaces and enums are subscriptable
	c, ok := t.(Class)
	if ok {
		p, has := c.Properties[name]
		return p, has
	}
	i, ok := t.(Interface)
	if ok {
		p, has := i.Funcs[name]
		return p, has
	}
	e, ok := t.(Enum)
	if ok {
		_, has := e.Items[name]
		return standards[Int], has
	}
	return standards[Invalid], false
}

func isSubscriptable(t Type) bool {
	// only classes, interfaces and enums are subscriptable
	_, ok := t.(Class)
	if ok {
		return true
	}
	_, ok = t.(Interface)
	if ok {
		return true
	}
	_, ok = t.(Enum)
	if ok {
		return true
	}
	return false
}
