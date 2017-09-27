package validator

import (
	"github.com/end-r/guardian/compiler/lexer"

	"github.com/end-r/guardian/compiler/ast"
)

func (v *Validator) ResolveExpression(e ast.ExpressionNode) Type {
	return resolvers[e.Type()](v, e)
}

type Resolver func(v *Validator, e ast.ExpressionNode) Type

var resolvers = map[ast.NodeType]Resolver{
	ast.Literal:      resolveLiteralExpression,
	ast.MapLiteral:   resolveMapLiteralExpression,
	ast.ArrayLiteral: resolveArrayLiteralExpression,
	/*ast.IndexExpression:  resolveIndexExpression,
	ast.CallExpression:   resolveCallExpression,
	ast.SliceExpression:  resolveSliceExpression,

	ast.BinaryExpression: resolveBinaryExpression,
	ast.UnaryExpression:  resolveUnaryExpression,
	ast.Reference:        resolveReference,*/
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
	//m := e.(ast.ArrayLiteralNode)

	arrayType := new(Array)
	//arrayType.Value = parseType(m.Key)
	return arrayType
}

func resolveMapLiteralExpression(v *Validator, e ast.ExpressionNode) Type {
	// must be literal
	//m := e.(ast.MapLiteralNode)
	mapType := new(Map)
	return mapType
}

/*
func resolveIndexExpression(v *Validator, e ast.ExpressionNode) Type {
	// must be literal
	i := e.(ast.IndexExpressionNode)
	exprType := v.ResolveExpression(i.Expression)
	// enforce that this must be an array type
	switch exprType.(type) {
	case Array:
		return exprType.(Array).Value
	case Map:
		return exprType.(Map).Value
	}
	return Invalid
}

func resolveCallExpression(v *Validator, e ast.ExpressionNode) Type {
	// must be literal
	c := e.(ast.CallExpressionNode)
	// return type of a call expression is always a tuple
	// tuple may be empty or single-valued
	call := v.ResolveExpression(c.Call)
	// enforce that this is a function pointer
	fn := call.(Func)
	return fn.Results
}

func resolveSliceExpression(v *Validator, e ast.ExpressionNode) Type {
	// must be literal
	s := e.(ast.SliceExpressionNode)
	exprType := v.ResolveExpression(s.Expression)
	// must be an array
	switch exprType.(type) {
	case Array:
	}
	return Invalid
}




func resolveBinaryExpression(v *Validator, e ast.ExpressionNode) Type {
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
	return Invalid
}

func resolveUnaryExpression(v *Validator, e ast.ExpressionNode) Type {
	// must be literal
	m := e.(ast.UnaryExpressionNode)
	operandType := ResolveExpression(m.Operand)

}

func resolveReference(v *Validator, e ast.ExpressionNode) Type {
	// must be literal
	m := e.(ast.ReferenceNode)
	// go up through table

}*/
