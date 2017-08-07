package expressions

import (
	"axia/guardian/go/compiler/ast"
	"axia/guardian/go/compiler/lexer"
	"axia/guardian/go/compiler/parser"
	"axia/guardian/go/util"
	"testing"
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
