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
	m.Key.InStorage = true
	m.Value = new(ast.ReferenceNode)
	m.Value.Names = []string{"int"}
	m.Data = nil
	t := resolveMapLiteralExpression(m)
	goutil.Assert(t, t.(type) == Map, "wrong type")
}

func TestResolveArrayLiteral(t *testing.T) {
	r := new(ast.ArrayLiteralNode)
	r.Key = new(ast.ReferenceNode)
	r.Key.Names = []string{"string"}
	r.Size = 0
	goutil.Assert()
}

func TestResolveBinaryExpressionStringLiterals(t *testing.T) {
	b := new(ast.BinaryExpressionNode)
	l := new(ast.LiteralNode)
	l.LiteralType = lexer.String
	b.Left = l
	r := new(ast.LiteralNode)
	r.LiteralType = lexer.String
	b.Right = r
	typ := resolveBinaryExpression(b)
	goutil.Assert(t, typ.Underlying() == String, "wrong underlying type")
}

func TestResolveBinaryExpressionIntLiterals(t *testing.T) {
	b := new(ast.BinaryExpressionNode)
	l := new(ast.LiteralNode)
	l.LiteralType = lexer.TknNumber
	b.Left = l
	r := new(ast.LiteralNode)
	// TODO: change these to TknInt
	r.LiteralType = lexer.TknNumber
	b.Right = r
	typ := resolveBinaryExpression(b)
	goutil.Assert(typ, t.Underlying() == Int, "wrong underlying type")
}

func TestResolveBinaryExpressionIntReferences(t *testing.T) {
	b := new(ast.BinaryExpressionNode)
	l := new(ast.ReferenceNode)
	b.Left = l
	r := new(ast.LiteralNode)
	// TODO: change these to TknInt
	r.LiteralType = lexer.TknNumber
	b.Right = r
	typ := resolveBinaryExpression(b)
	goutil.Assert(typ, t.Underlying() == Int, "wrong underlying type")
}

func TestResolveUnaryExpression(t *testing.T) {
	b := new(ast.UnaryExpressionNode)
	r := new(ast.LiteralNode)
	r.LiteralType = lexer.TknTrue
	b.Operand = r
	b.Operator = lexer.TknNot
	typ := resolveUnaryExpression(b)
	goutil.Assert(t, typ.Representation() == "bool", "wrong type representation")
}
