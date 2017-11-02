package validator

import (
	"github.com/end-r/guardian/compiler/lexer"

	"github.com/end-r/guardian/compiler/ast"
)

func (v *Validator) resolveType(node ast.Node) Type {
	switch node.Type() {
	case ast.Reference:
		r := node.(ast.ReferenceNode)
		return v.resolveReference(r)
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

func (v *Validator) resolveReference(node ast.ReferenceNode) Type {
	return standards[Invalid]
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

	} else {
		switch node.Type() {
		case ast.PlainType:
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
	case lexer.TknNumber:
		return standards[Int]
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

func resolveReference(v *Validator, e ast.ExpressionNode) Type {
	// must be reference
	m := e.(ast.ReferenceNode)
	// go up through table
	return v.resolveReference(m)
}
