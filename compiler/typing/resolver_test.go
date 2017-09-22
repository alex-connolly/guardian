package typing

import (
	"testing"

	"github.com/end-r/goutil"
	"github.com/end-r/guardian/compiler/ast"
	"github.com/end-r/guardian/compiler/lexer"
)

func TestResolveLiteral(t *testing.T) {
	r := new(ast.LiteralNode)
	r.LiteralType = lexer.LiteralString
}

func TestResolveMapLiteral(t *testing.T) {
	m := new(ast.MapLiteralNode)
	m.Key = new(ast.ReferenceNode)
	m.Key.Names = []string{"string"}
}

func TestResolveArrayLiteral(t *testing.T) {
	r := new(ast.ArrayLiteralNode)
	r.LiteralType = lexer.LiteralString
	goutil.Assert()
}

func TestResolveBinaryExpression(t *testing.T) {
	b := new(ast.BinaryExpressionNode)
	r.LiteralType = lexer.LiteralString
	goutil.Assert()
}

func TestResolveUnaryExpression(t *testing.T) {
	b := new(ast.UnaryExpressionNode)
	r.LiteralType = lexer.LiteralString
	goutil.Assert()
}
