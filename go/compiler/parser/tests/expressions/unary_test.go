package expressions

import (
	"testing"

	"github.com/end-r/guardian/go/compiler/ast"
	"github.com/end-r/guardian/go/compiler/lexer"
	"github.com/end-r/guardian/go/compiler/parser"
	"github.com/end-r/guardian/go/util"
)

func TestParseUnaryLiteralPrefixExpression(t *testing.T) {
	p := parser.ParseString("++5")
	util.AssertNow(t, p.Expression.Type() == ast.UnaryExpression, "wrong node type")
	u := p.Expression.(ast.UnaryExpressionNode)
	util.Assert(t, u.Operator == lexer.TknIncrement, "wrong operator")
}

func TestParseUnaryReferencePrefixExpression(t *testing.T) {
	p := parser.ParseString("!me")
	util.AssertNow(t, p.Expression.Type() == ast.UnaryExpression, "wrong node type")
	u := p.Expression.(ast.UnaryExpressionNode)
	util.Assert(t, u.Operator == lexer.TknNot, "wrong operator")

	p = parser.ParseString("++me")
	util.AssertNow(t, p.Expression.Type() == ast.UnaryExpression, "wrong node type")
	u = p.Expression.(ast.UnaryExpressionNode)
	util.Assert(t, u.Operator == lexer.TknIncrement, "wrong operator")
}

func TestParseUnaryLiteralPostfixExpression(t *testing.T) {
	p := parser.ParseString("5++")
	util.AssertNow(t, p.Expression.Type() == ast.UnaryExpression, "wrong node type")
	u := p.Expression.(ast.UnaryExpressionNode)
	util.Assert(t, u.Operator == lexer.TknIncrement, "wrong operator")
}

func TestParseUnaryReferencePostfixExpression(t *testing.T) {
	p := parser.ParseString("me++")
	util.AssertNow(t, p.Expression.Type() == ast.UnaryExpression, "wrong node type")
	u := p.Expression.(ast.UnaryExpressionNode)
	util.Assert(t, u.Operator == lexer.TknIncrement, "wrong operator")
}
