package validator

import (
	"testing"

	"github.com/end-r/guardian/compiler/lexer"

	"github.com/end-r/guardian/compiler/ast"

	"github.com/end-r/goutil"
)

func TestResolveLiteralExpressionString(t *testing.T) {
	v := new(Validator)
	r := ast.LiteralNode{}
	r.LiteralType = lexer.TknString
	goutil.Assert(t, v.ResolveExpression(r) == standards[String], "wrong expression type")
}

func TestResolveLiteralExpressionBool(t *testing.T) {
	v := new(Validator)
	r := ast.LiteralNode{}
	r.LiteralType = lexer.TknTrue
	goutil.Assert(t, v.ResolveExpression(r) == standards[Bool], "wrong true expression type")
	r.LiteralType = lexer.TknFalse
	goutil.Assert(t, v.ResolveExpression(r) == standards[Bool], "wrong false expression type")
}

func TestResolveArrayLiteralExpression(t *testing.T) {
	v := new(Validator)
	a := ast.ArrayLiteralNode{}
	a.Key = ast.ReferenceNode{
		Names: []string{"dog"},
	}
	_, ok := v.ResolveExpression(a).(Array)
	goutil.Assert(t, ok, "wrong base type")
}

func TestResolveMapLiteralExpression(t *testing.T) {
	v := new(Validator)
	a := ast.MapLiteralNode{}
	a.Key = ast.ReferenceNode{
		Names: []string{"dog"},
	}
	a.Value = ast.ReferenceNode{
		Names: []string{"cat"},
	}
	_, ok := v.ResolveExpression(a).(Map)
	goutil.Assert(t, ok, "wrong base type")
}

func TestResolveIndexExpressionArrayLiteral(t *testing.T) {
	v := new(Validator)
	a := ast.ArrayLiteralNode{}
	a.Key = ast.ReferenceNode{
		Names: []string{"cat"},
	}
	e := ast.IndexExpressionNode{}
	e.Expression = a
	e.Index = createIndex()
	m := v.ResolveExpression(e)
	goutil.Assert(t, WriteType(m) == "[cat]", "wrong base type")
}

func createIndex() ast.ExpressionNode {
	idx := ast.LiteralNode{}
	idx.LiteralType = lexer.TknNumber
	return idx
}
