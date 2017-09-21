package typing

import (
	"github.com/end-r/guardian/compiler/ast"
)

func ResolveExpression(e ast.ExpressionNode) Type {
	return resolvers[e.Type()]()
}

type Resolver func(e ast.ExpressionNode) Type

var resolvers = map[ast.NodeType]Resolver{
	ast.Literal:         resolveLiteralExpression,
	ast.IndexExpression: resolveIndexExpression,
	ast.CallExpression:  resolveCallExpression,
	ast.SliceExpression: resolveSliceExpression,
}

func resolveLiteralExpression(e ast.ExpressionNode) Type {
	// must be literal
	l := e.(ast.LiteralNode)
}

func resolveIndexExpression(e ast.ExpressionNode) Type {
	// must be literal
	i := e.(ast.IndexExpressionNode)
}

func resolveCallExpression(e ast.ExpressionNode) Type {
	// must be literal
	c := e.(ast.CallExpressionNode)
}

func resolveSliceExpression(e ast.ExpressionNode) Type {
	// must be literal
	s := e.(ast.SliceExpressionNode)
}

func resolveMapLiteralExpression(e ast.ExpressionNode) Type {
	// must be literal
	m := e.(ast.MapLiteralNode)
}
