package typing

import (
	"axia/guardian/compiler/lexer"

	"github.com/end-r/guardian/compiler/ast"
)

func ResolveExpression(e ast.ExpressionNode) Type {
	return resolvers[e.Type()]()
}

type Resolver func(e ast.ExpressionNode) Type

var resolvers = map[ast.NodeType]Resolver{
	ast.Literal:          resolveLiteralExpression,
	ast.IndexExpression:  resolveIndexExpression,
	ast.CallExpression:   resolveCallExpression,
	ast.SliceExpression:  resolveSliceExpression,
	ast.MapLiteral:       resolveMapLiteralExpression,
	ast.ArrayLiteral:     resolveArrayLiteralExpression,
	ast.BinaryExpression: resolveBinaryExpression,
	ast.UnaryExpression:  resolveUnaryExpression,
	ast.Reference:        resolveReference,
}

func resolveLiteralExpression(e ast.ExpressionNode) Type {
	// must be literal
	l := e.(ast.LiteralNode)
	switch l.LiteralType {
	case lexer.TknString:
		return String
	case lexer.TknTrue, lexer.TknFalse:
		return Bool
	case lexer.TknNumber:
		return Int
	}
}

func resolveIndexExpression(e ast.ExpressionNode) Type {
	// must be literal
	i := e.(ast.IndexExpressionNode)
	exprType := ResolveExpression(i.Expression)
	// enforce that this must be an array type
	switch exprType {
	case Array:
		return exprType.(Array).key
	case Map:
		return exprType.(Map).value
	}
	return InvalidType
}

func resolveCallExpression(e ast.ExpressionNode) Type {
	// must be literal
	c := e.(ast.CallExpressionNode)
	// return type of a call expression is always a tuple
	// tuple may be empty or single-valued
	call := ResolveExpression(c.Call)
	// enforce that this is a function pointer
	fn := call.(Function)
	return createTuple(fn.Results)
}

func resolveSliceExpression(e ast.ExpressionNode) Type {
	// must be literal
	s := e.(ast.SliceExpressionNode)
	exprType := ResolveExpression(i.Expression)
	// must be an array
	switch exprType {
	case Array:
		a := exprType.(Array)
		return createArray(a.key)
	}
	return InvalidType
}

func resolveMapLiteralExpression(e ast.ExpressionNode) Type {
	// must be literal
	m := e.(ast.MapLiteralNode)
	mapType := new(Map)
	mapType.key = parseType(m.Key)
	mapType.value = parseType(m.Value)
	return mapType
}

func resolveArrayLiteralExpression(e ast.ExpressionNode) Type {
	// must be literal
	m := e.(ast.ArrayLiteralNode)

	arrayType := new(Array)
	arrayType.key = parseType(m.Key)
	return arrayType
}

func resolveBinaryExpression(e ast.ExpressionNode) Type {
	// must be literal
	b := e.(ast.BinaryExpressionNode)
	// rules for binary Expressions
	leftType := ResolveExpression(b.Left)
	rightType := ResolveExpression(b.Left)
	switch leftType.Underlying() {
	case String:
		// TODO: error if rightType != string
		return String
	case Int:
		// TODO: error if rightType != string
		return Int
	}
	// else it is a type which is not defined for binary operators
	return InvalidType
}

func resolveUnaryExpression(e ast.ExpressionNode) Type {
	// must be literal
	m := e.(ast.UnaryExpressionNode)
	operandType := ResolveExpression(m.Operand)

}

func resolveReference(e ast.ExpressionNode) Type {
	// must be literal
	m := e.(ast.ReferenceNode)
	// go up through table

}
