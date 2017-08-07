package expressions

import (
	"axia/guardian/go/compiler/ast"
	"axia/guardian/go/compiler/lexer"
	"axia/guardian/go/compiler/parser"
	"axia/guardian/go/util"
	"testing"
)

func TestParseBinaryLiteralExpression(t *testing.T) {

	// two literal nodes
	p := parser.ParseString("3 + 5")
	util.AssertNow(t, p.Expression.Type() == ast.BinaryExpression, "wrong node type")
	u := p.Expression.(ast.BinaryExpressionNode)
	util.Assert(t, u.Operator == lexer.TknAdd, "wrong operator")
}

func TestParseBinaryReferenceExpression(t *testing.T) {

	// two reference nodes
	p := parser.ParseString("a - b")
	util.AssertNow(t, p.Expression.Type() == ast.BinaryExpression, "wrong node type")
	u := p.Expression.(ast.BinaryExpressionNode)
	util.Assert(t, u.Operator == lexer.TknSub, "wrong operator")
}

func TestParseBinaryCallExpression(t *testing.T) {

	// two call nodes
	p := parser.ParseString("a() - b()")
	util.AssertNow(t, p.Expression.Type() == ast.BinaryExpression, "wrong node type")
	u := p.Expression.(ast.BinaryExpressionNode)
	util.Assert(t, u.Operator == lexer.TknSub, "wrong operator")
}

func TestParseMultipleBinaryExpression(t *testing.T) {
	p := parser.ParseString("a() - b() - c()")
	util.AssertNow(t, p.Expression.Type() == ast.BinaryExpression, "wrong node type")
	u := p.Expression.(ast.BinaryExpressionNode)
	util.Assert(t, u.Operator == lexer.TknSub, "wrong operator")
}

func TestParseMultipleBinaryExpressionBracketed(t *testing.T) {
	p := parser.ParseString("a() - (b() - c())")
	util.AssertNow(t, p.Expression.Type() == ast.BinaryExpression, "wrong node type")
	u := p.Expression.(ast.BinaryExpressionNode)
	util.Assert(t, u.Operator == lexer.TknSub, "wrong operator")
}
