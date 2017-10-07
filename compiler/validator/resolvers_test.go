package validator

import (
	"fmt"
	"testing"

	"github.com/end-r/guardian/compiler/lexer"

	"github.com/end-r/guardian/compiler/ast"

	"github.com/end-r/goutil"
)

func TestResolveLiteralExpressionString(t *testing.T) {
	v := NewValidator()
	r := ast.LiteralNode{}
	r.LiteralType = lexer.TknString
	goutil.Assert(t, v.resolveExpression(r) == standards[String], "wrong expression type")
}

func TestResolveLiteralExpressionBool(t *testing.T) {
	v := NewValidator()
	r := ast.LiteralNode{}
	r.LiteralType = lexer.TknTrue
	goutil.Assert(t, v.resolveExpression(r) == standards[Bool], "wrong true expression type")
	r.LiteralType = lexer.TknFalse
	goutil.Assert(t, v.resolveExpression(r) == standards[Bool], "wrong false expression type")
}

func TestResolveCallExpression(t *testing.T) {
	v := NewValidator()
	fn := new(Func)
	fn.Params = NewTuple(standards[Int], standards[Int])
	fn.Results = NewTuple(standards[Int])
	v.DeclareType("hello", NewFunc(NewTuple(), NewTuple(standards[Int])))
	c := ast.CallExpressionNode{}
	r := ast.ReferenceNode{}
	r.Names = []string{"hello"}
	c.Call = r
	c.Arguments = []ast.ExpressionNode{
		ast.LiteralNode{
			LiteralType: lexer.TknNumber,
		},
		ast.LiteralNode{
			LiteralType: lexer.TknNumber,
		},
	}
	tuple, ok := v.resolveExpression(c).(Tuple)
	goutil.Assert(t, ok, "wrong base type")
	goutil.Assert(t, len(tuple.types) == 1, "wrong type length")
}

func TestResolveArrayLiteralExpression(t *testing.T) {
	v := NewValidator()

	a := ast.ArrayLiteralNode{}
	a.Key = ast.ReferenceNode{
		Names: []string{"dog"},
	}
	_, ok := v.resolveExpression(a).(Array)
	goutil.Assert(t, ok, "wrong base type")
}

func TestResolveMapLiteralExpression(t *testing.T) {
	v := NewValidator()
	a := ast.MapLiteralNode{}
	v.DeclareType("dog", standards[String])
	v.DeclareType("cat", standards[String])
	a.Key = ast.ReferenceNode{
		Names: []string{"dog"},
	}
	a.Value = ast.ReferenceNode{
		Names: []string{"cat"},
	}
	m, ok := v.resolveExpression(a).(Map)
	goutil.AssertNow(t, ok, "wrong base type")
	goutil.Assert(t, WriteType(m.Key) == "dog", fmt.Sprintf("wrong key written: %s", WriteType(m.Key)))
	goutil.Assert(t, WriteType(m.Value) == "cat", fmt.Sprintf("wrong val written: %s", WriteType(m.Value)))
}

func TestResolveIndexExpressionArrayLiteral(t *testing.T) {
	//v := NewValidator()
	a := ast.ArrayLiteralNode{}
	a.Key = ast.ReferenceNode{
		Names: []string{"cat"},
	}
	e := ast.IndexExpressionNode{}
	e.Expression = a
	e.Index = createIndex()
}

func createIndex() ast.ExpressionNode {
	idx := ast.LiteralNode{}
	idx.LiteralType = lexer.TknNumber
	return idx
}
