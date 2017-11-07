package validator

import (
	"fmt"
	"testing"

	"github.com/end-r/guardian/compiler/parser"

	"github.com/end-r/guardian/compiler/ast"

	"github.com/end-r/goutil"
)

func TestResolveLiteralExpressionString(t *testing.T) {
	v := NewValidator()
	p := parser.ParseExpression(`"hi"`)
	goutil.AssertNow(t, p.Type() == ast.Literal, "wrong expression type")
	a := p.(ast.LiteralNode)
	goutil.Assert(t, v.resolveExpression(a).compare(standards[String]), "wrong true expression type")
}

func TestResolveLiteralExpressionBool(t *testing.T) {
	v := NewValidator()
	p := parser.ParseExpression("true")
	goutil.AssertNow(t, p != nil, "expr should not be nil")
	goutil.AssertNow(t, p.Type() == ast.Literal, "wrong expression type")
	a := p.(ast.LiteralNode)
	goutil.Assert(t, v.resolveExpression(a).compare(standards[Bool]), "wrong true expression type")
}

func TestResolveCallExpression(t *testing.T) {
	v := NewValidator()
	fn := new(Func)
	fn.Params = NewTuple(standards[Int], standards[Int])
	fn.Results = NewTuple(standards[Int])
	v.DeclareVarOfType("hello", NewFunc(NewTuple(), NewTuple(standards[Int])))
	p := parser.ParseExpression("hello(5, 5)")
	goutil.AssertNow(t, p.Type() == ast.CallExpression, "wrong expression type")
	a := p.(ast.CallExpressionNode)
	resolved, ok := v.resolveExpression(a).(Tuple)
	goutil.Assert(t, ok, "wrong base type")
	goutil.Assert(t, len(resolved.types) == 1, "wrong type length")
	goutil.Assert(t, fn.Results.compare(resolved), "should be equal")
}

func TestResolveArrayLiteralExpression(t *testing.T) {
	v := NewValidator()
	v.DeclareType("dog", standards[String])
	p := parser.ParseExpression("[dog]{}")
	goutil.AssertNow(t, p.Type() == ast.ArrayLiteral, "wrong expression type")
	a := p.(ast.ArrayLiteralNode)
	_, ok := v.resolveExpression(a).(Array)
	goutil.Assert(t, ok, "wrong base type")
}

func TestResolveMapLiteralExpression(t *testing.T) {
	v := NewValidator()
	v.DeclareType("dog", standards[String])
	v.DeclareType("cat", standards[String])
	p := parser.ParseExpression("map[dog]cat{}")
	goutil.AssertNow(t, p.Type() == ast.MapLiteral, "wrong expression type")
	m, ok := v.resolveExpression(p).(Map)
	goutil.AssertNow(t, ok, "wrong base type")
	goutil.Assert(t, m.Key.compare(standards[String]), fmt.Sprintf("wrong key: %s", WriteType(m.Key)))
	goutil.Assert(t, m.Value.compare(standards[String]), fmt.Sprintf("wrong val: %s", WriteType(m.Value)))
}

func TestResolveIndexExpressionArrayLiteral(t *testing.T) {
	v := NewValidator()
	v.DeclareVarOfType("cat", standards[Int])
	p := parser.ParseExpression("[]cat{}[0]")
	goutil.AssertNow(t, p.Type() == ast.IndexExpression, "wrong expression type")
	b := p.(ast.IndexExpressionNode)
	resolved := v.resolveExpression(b)
	goutil.AssertNow(t, resolved.compare(v.getNamedType("cat")), "wrong expression type")
}

func TestResolveBinaryExpressionSimpleNumeric(t *testing.T) {
	p := parser.ParseExpression("5 + 5")
	goutil.AssertNow(t, p.Type() == ast.BinaryExpression, "wrong expression type")
	b := p.(ast.BinaryExpressionNode)
	v := NewValidator()
	resolved := v.resolveExpression(b)
	goutil.AssertNow(t, resolved.compare(standards[Int]), "wrong expression type")
}

func TestResolveBinaryExpressionConcatenation(t *testing.T) {
	p := parser.ParseExpression(`"a" + "b"`)
	goutil.AssertNow(t, p.Type() == ast.BinaryExpression, "wrong expression type")
	b := p.(ast.BinaryExpressionNode)
	v := NewValidator()
	resolved := v.resolveExpression(b)
	goutil.AssertNow(t, resolved.compare(standards[String]), "wrong expression type")
}

func TestResolveBinaryExpressionEql(t *testing.T) {
	p := parser.ParseExpression("5 == 5")
	goutil.AssertNow(t, p.Type() == ast.BinaryExpression, "wrong expression type")
	b := p.(ast.BinaryExpressionNode)
	v := NewValidator()
	resolved := v.resolveExpression(b)
	goutil.AssertNow(t, resolved.compare(standards[Bool]), "wrong expression type")
}

func TestResolveBinaryExpressionGeq(t *testing.T) {
	p := parser.ParseExpression("5 >= 5")
	goutil.AssertNow(t, p.Type() == ast.BinaryExpression, "wrong expression type")
	b := p.(ast.BinaryExpressionNode)
	v := NewValidator()
	resolved := v.resolveExpression(b)
	goutil.AssertNow(t, resolved.compare(standards[Bool]), "wrong expression type")
}

func TestResolveBinaryExpressionLeq(t *testing.T) {
	p := parser.ParseExpression("5 <= 5")
	goutil.AssertNow(t, p.Type() == ast.BinaryExpression, "wrong expression type")
	b := p.(ast.BinaryExpressionNode)
	v := NewValidator()
	resolved := v.resolveExpression(b)
	goutil.AssertNow(t, resolved.compare(standards[Bool]), "wrong expression type")
}

func TestResolveBinaryExpressionLss(t *testing.T) {
	p := parser.ParseExpression("5 < 5")
	goutil.AssertNow(t, p.Type() == ast.BinaryExpression, "wrong expression type")
	b := p.(ast.BinaryExpressionNode)
	v := NewValidator()
	resolved := v.resolveExpression(b)
	goutil.AssertNow(t, resolved.compare(standards[Bool]), "wrong expression type")
}

func TestResolveBinaryExpressionGtr(t *testing.T) {
	p := parser.ParseExpression("5 > 5")
	goutil.AssertNow(t, p.Type() == ast.BinaryExpression, "wrong expression type")
	b := p.(ast.BinaryExpressionNode)
	v := NewValidator()
	resolved := v.resolveExpression(b)
	goutil.AssertNow(t, resolved.compare(standards[Bool]), "wrong expression type")
}
