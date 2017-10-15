package validator

import (
	"github.com/end-r/guardian/compiler/lexer"

	"github.com/end-r/guardian/compiler/ast"
)

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
	}
	return resolvers[e.Type()](v, e)
}

type resolver func(v *Validator, e ast.ExpressionNode) Type

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
	/*b := e.(ast.BinaryExpressionNode)
	// rules for binary Expressions
	leftType := v.resolveExpression(b.Left)
	rightType := v.resolveExpression(b.Left)
	switch leftType.(type) {
	case standards[String]:
		// TODO: error if rightType != string
		return standards[String]
	case standards[Int]:
		// TODO: error if rightType != string
		return standards[Int]
	}*/
	// else it is a type which is not defined for binary operators
	return standards[Invalid]
}

func resolveUnaryExpression(v *Validator, e ast.ExpressionNode) Type {
	// must be literal
	//m := e.(ast.UnaryExpressionNode)
	//operandType := v.resolveExpression(m.Operand)
	return standards[Invalid]
}

func resolveReference(v *Validator, e ast.ExpressionNode) Type {
	// must be reference
	m := e.(ast.ReferenceNode)
	// go up through table
	return v.findReference(m.Names...)
}
