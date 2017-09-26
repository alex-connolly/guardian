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
	goutil.Assert(t, v.ResolveExpression(r) == String, "wrong expression type")
}

func TestResolveLiteralExpressionBool(t *testing.T) {
	v := new(Validator)
	r := ast.LiteralNode{}
	r.LiteralType = lexer.TknTrue
	goutil.Assert(t, v.ResolveExpression(r) == Bool, "wrong true expression type")
	r.LiteralType = lexer.TknFalse
	goutil.Assert(t, v.ResolveExpression(r) == Bool, "wrong false expression type")
}

func TestResolveArrayLiteralExpression(t *testing.T) {
	v := new(Validator)
	a := ast.ArrayLiteralNode{}
	a.Key = ast.ReferenceNode{
		Names: []string{"dog"},
	}
	goutil.Assert(t, v.ResolveExpression(a).Base() == ArrayBase, "wrong base type")
}
