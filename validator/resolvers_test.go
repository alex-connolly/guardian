package validator

import (
	"fmt"
	"testing"

	"github.com/end-r/guardian/ast"
	"github.com/end-r/guardian/parser"

	"github.com/end-r/goutil"
)

func TestResolveLiteralExpressionboolean(t *testing.T) {
	v := NewValidator(NewTestVM())
	p := parser.ParseExpression("true")
	goutil.AssertNow(t, p != nil, "expr should not be nil")
	goutil.AssertNow(t, p.Type() == ast.Literal, "wrong expression type")
	a := p.(ast.LiteralNode)
	goutil.Assert(t, v.resolveExpression(a).Compare(standards[boolean]), "wrong true expression type")
}

func TestResolveCallExpression(t *testing.T) {
	v := NewValidator(NewTestVM())
	fn := new(Func)
	fn.Params = NewTuple(standards[boolean], standards[boolean])
	fn.Results = NewTuple(standards[boolean])
	v.DeclareVarOfType("hello", NewFunc(NewTuple(), NewTuple(standards[boolean])))
	p := parser.ParseExpression("hello(5, 5)")
	goutil.AssertNow(t, p.Type() == ast.CallExpression, "wrong expression type")
	a := p.(ast.CallExpressionNode)
	resolved, ok := v.resolveExpression(a).(Tuple)
	goutil.Assert(t, ok, "wrong base type")
	goutil.Assert(t, len(resolved.types) == 1, "wrong type length")
	goutil.Assert(t, fn.Results.Compare(resolved), "should be equal")
}

func TestResolveArrayLiteralExpression(t *testing.T) {
	v := NewValidator(NewTestVM())
	v.DeclareType("dog", standards[unknown])
	p := parser.ParseExpression("[]dog{}")
	goutil.AssertNow(t, p.Type() == ast.ArrayLiteral, "wrong expression type")
	a := p.(ast.ArrayLiteralNode)
	_, ok := v.resolveExpression(a).(Array)
	goutil.Assert(t, ok, "wrong base type")
}

func TestResolveArrayLiteralSliceExpressionCopy(t *testing.T) {
	v := NewValidator(NewTestVM())
	v.DeclareType("dog", standards[unknown])
	p := parser.ParseExpression("[]dog{}[:]")
	goutil.AssertNow(t, p.Type() == ast.SliceExpression, "wrong expression type")
	_, ok := v.resolveExpression(p).(Array)
	goutil.Assert(t, ok, "wrong base type")
}

func TestResolveArrayLiteralSliceExpressionLower(t *testing.T) {
	v := NewValidator(NewTestVM())
	v.DeclareType("dog", standards[unknown])
	p := parser.ParseExpression("[]dog{}[6:]")
	goutil.AssertNow(t, p.Type() == ast.SliceExpression, "wrong expression type")
	_, ok := v.resolveExpression(p).(Array)
	goutil.Assert(t, ok, "wrong base type")
}

func TestResolveArrayLiteralSliceExpressionUpper(t *testing.T) {
	v := NewValidator(NewTestVM())
	v.DeclareType("dog", standards[unknown])
	p := parser.ParseExpression("[]dog{}[:10]")
	goutil.AssertNow(t, p.Type() == ast.SliceExpression, "wrong expression type")
	_, ok := v.resolveExpression(p).(Array)
	goutil.Assert(t, ok, "wrong base type")
}

func TestResolveArrayLiteralSliceExpressionBoth(t *testing.T) {
	v := NewValidator(NewTestVM())
	v.DeclareType("dog", standards[unknown])
	p := parser.ParseExpression("[]dog{}[6:10]")
	goutil.AssertNow(t, p.Type() == ast.SliceExpression, "wrong expression type")
	_, ok := v.resolveExpression(p).(Array)
	goutil.Assert(t, ok, "wrong base type")
}

func TestResolveMapLiteralExpression(t *testing.T) {
	v := NewValidator(NewTestVM())
	v.DeclareType("dog", standards[unknown])
	v.DeclareType("cat", standards[unknown])
	p := parser.ParseExpression("map[dog]cat{}")
	goutil.AssertNow(t, p.Type() == ast.MapLiteral, "wrong expression type")
	m, ok := v.resolveExpression(p).(Map)
	goutil.AssertNow(t, ok, "wrong base type")
	goutil.Assert(t, m.Key.Compare(standards[unknown]), fmt.Sprintf("wrong key: %s", WriteType(m.Key)))
	goutil.Assert(t, m.Value.Compare(standards[unknown]), fmt.Sprintf("wrong val: %s", WriteType(m.Value)))
}

func TestResolveIndexExpressionArrayLiteral(t *testing.T) {
	v := NewValidator(NewTestVM())
	v.DeclareVarOfType("cat", standards[boolean])
	p := parser.ParseExpression("[]cat{}[0]")
	goutil.AssertNow(t, p.Type() == ast.IndexExpression, "wrong expression type")
	b := p.(ast.IndexExpressionNode)
	resolved := v.resolveExpression(b)
	goutil.AssertNow(t, resolved.Compare(v.getNamedType("cat")), "wrong expression type")
}

func TestResolveIndexExpressionMapLiteral(t *testing.T) {
	v := NewValidator(NewTestVM())
	v.DeclareType("dog", standards[unknown])
	v.DeclareType("cat", standards[unknown])
	p := parser.ParseExpression(`map[dog]cat{}["hi"]`)
	goutil.AssertNow(t, p.Type() == ast.IndexExpression, "wrong expression type")
	ok := v.resolveExpression(p).Compare(v.getNamedType("cat"))
	goutil.AssertNow(t, ok, "wrong type returned")

}

func TestResolveBinaryExpressionSimpleNumeric(t *testing.T) {
	p := parser.ParseExpression("5 + 5")
	goutil.AssertNow(t, p.Type() == ast.BinaryExpression, "wrong expression type")
	b := p.(ast.BinaryExpressionNode)
	v := NewValidator(NewTestVM())
	resolved := v.resolveExpression(b)
	goutil.AssertNow(t, resolved.Compare(v.smallestNumericType(0, false)), fmt.Sprintf("wrong expression type: %s", WriteType(resolved)))
}

func TestResolveBinaryExpressionConcatenation(t *testing.T) {
	p := parser.ParseExpression(`"a" + "b"`)
	goutil.AssertNow(t, p.Type() == ast.BinaryExpression, "wrong expression type")
	b := p.(ast.BinaryExpressionNode)
	v := NewValidator(NewTestVM())
	resolved := v.resolveExpression(b)
	goutil.AssertNow(t, resolved.Compare(v.getNamedType("string")), "wrong expression type")
}

func TestResolveBinaryExpressionEql(t *testing.T) {
	p := parser.ParseExpression("5 == 5")
	goutil.AssertNow(t, p.Type() == ast.BinaryExpression, "wrong expression type")
	b := p.(ast.BinaryExpressionNode)
	v := NewValidator(NewTestVM())
	resolved := v.resolveExpression(b)
	goutil.AssertNow(t, resolved.Compare(standards[boolean]), "wrong expression type")
}

func TestResolveBinaryExpressionGeq(t *testing.T) {
	p := parser.ParseExpression("5 >= 5")
	goutil.AssertNow(t, p.Type() == ast.BinaryExpression, "wrong expression type")
	b := p.(ast.BinaryExpressionNode)
	v := NewValidator(NewTestVM())
	resolved := v.resolveExpression(b)
	goutil.AssertNow(t, resolved.Compare(standards[boolean]), "wrong expression type")
}

func TestResolveBinaryExpressionLeq(t *testing.T) {
	p := parser.ParseExpression("5 <= 5")
	goutil.AssertNow(t, p.Type() == ast.BinaryExpression, "wrong expression type")
	b := p.(ast.BinaryExpressionNode)
	v := NewValidator(NewTestVM())
	resolved := v.resolveExpression(b)
	goutil.AssertNow(t, resolved.Compare(standards[boolean]), "wrong expression type")
}

func TestResolveBinaryExpressionLss(t *testing.T) {
	p := parser.ParseExpression("5 < 5")
	goutil.AssertNow(t, p.Type() == ast.BinaryExpression, "wrong expression type")
	b := p.(ast.BinaryExpressionNode)
	v := NewValidator(NewTestVM())
	resolved := v.resolveExpression(b)
	goutil.AssertNow(t, resolved.Compare(standards[boolean]), "wrong expression type")
}

func TestResolveBinaryExpressionGtr(t *testing.T) {
	p := parser.ParseExpression("5 > 5")
	goutil.AssertNow(t, p.Type() == ast.BinaryExpression, "wrong b expression type")
	b := p.(ast.BinaryExpressionNode)
	v := NewValidator(NewTestVM())
	resolved := v.resolveExpression(b)
	goutil.AssertNow(t, resolved.Compare(standards[boolean]), fmt.Sprintf("wrong resolved expression type: %s", WriteType(resolved)))
}

func TestResolveBinaryExpressionCast(t *testing.T) {
	p := parser.ParseExpression("5 as uint8")
	goutil.AssertNow(t, p.Type() == ast.BinaryExpression, "wrong expression type")

}
