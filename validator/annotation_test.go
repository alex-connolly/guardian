package validator

import (
	"testing"

	"github.com/end-r/guardian/typing"

	"github.com/end-r/guardian/ast"

	"github.com/end-r/goutil"
	"github.com/end-r/guardian/parser"
)

func TestBinaryExpressionAnnotation(t *testing.T) {
	e := parser.ParseExpression("5 == 5")
	v := NewValidator(NewTestVM())
	v.resolveExpression(e)
	goutil.AssertNow(t, e.Type() == ast.BinaryExpression, "wrong node type")
	b := e.(*ast.BinaryExpressionNode)
	goutil.AssertNow(t, b.Resolved != nil, "resolved nil")
	goutil.AssertNow(t, b.Resolved.Compare(typing.Boolean()), "wrong resolved type")
}
