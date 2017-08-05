package parser

import (
	"axia/guardian/go/compiler/ast"
	"axia/guardian/go/compiler/lexer"
	"axia/guardian/go/util"
	"testing"
)

func TestParseUnaryReferencePrefixExpression(t *testing.T) {
	p := createParser("!me")
	n := p.parseExpression()
	util.AssertNow(t, n.Type() == ast.UnaryExpression, "wrong node type")
	u := n.(ast.UnaryExpressionNode)
	util.Assert(t, u.Operator == lexer.TknNot, "wrong operator")
}

func TestParseUnaryReferencePostfixExpression(t *testing.T) {
	p := createParser("me++")
	n := p.parseExpression()
	util.AssertNow(t, n.Type() == ast.UnaryExpression, "wrong node type")
	u := n.(ast.UnaryExpressionNode)
	util.Assert(t, u.Operator == lexer.TknIncrement, "wrong operator")
}

func TestParseBinaryLiteralExpression(t *testing.T) {

	// two literal nodes
	p := createParser("3 + 5")
	n := p.parseExpression()
	util.AssertNow(t, n.Type() == ast.BinaryExpression, "wrong node type")
	u := n.(ast.BinaryExpressionNode)
	util.Assert(t, u.Operator == lexer.TknAdd, "wrong operator")
}

func TestParseBinaryReferenceExpression(t *testing.T) {

	// two reference nodes
	p := createParser("a - b")
	n := p.parseExpression()
	util.AssertNow(t, n.Type() == ast.BinaryExpression, "wrong node type")
	u := n.(ast.BinaryExpressionNode)
	util.Assert(t, u.Operator == lexer.TknSub, "wrong operator")
}

func TestParseBinaryCallExpression(t *testing.T) {

	// two call nodes
	p := createParser("a() - b()")
	n := p.parseExpression()
	util.AssertNow(t, n.Type() == ast.BinaryExpression, "wrong node type")
	u := n.(ast.BinaryExpressionNode)
	util.Assert(t, u.Operator == lexer.TknSub, "wrong operator")

}

func TestParseMultipleBinaryExpression(t *testing.T) {
	p := createParser("a() - b() - c()")
	n := p.parseExpression()
	util.AssertNow(t, n.Type() == ast.BinaryExpression, "wrong node type")
	u := n.(ast.BinaryExpressionNode)
	util.Assert(t, u.Operator == lexer.TknSub, "wrong operator")
}

func TestParseMultipleBinaryExpressionBracketed(t *testing.T) {
	p := createParser("a() - (b() - c())")
	n := p.parseExpression()
	util.AssertNow(t, n.Type() == ast.BinaryExpression, "wrong node type")
	u := n.(ast.BinaryExpressionNode)
	util.Assert(t, u.Operator == lexer.TknSub, "wrong operator")
}

func TestParseIndexExpression(t *testing.T) {
	p := createParser("array[5]")
	n := p.parseExpression()
	util.AssertNow(t, n.Type() == ast.IndexExpression, "wrong node type")
	u := n.(ast.IndexExpressionNode)
	util.Assert(t, u.Index.Type() == ast.Literal, "wrong index type")

	p = createParser(`array[getNumber("hi")]`)
	n = p.parseExpression()
	util.AssertNow(t, n.Type() == ast.IndexExpression, "wrong node type")
	u = n.(ast.IndexExpressionNode)
	util.Assert(t, u.Index.Type() == ast.CallExpression, "wrong index type")
}

func TestParseSliceExpression(t *testing.T) {
	p := createParser("array[5:]")
	n := p.parseExpression()
	util.AssertNow(t, n.Type() == ast.SliceExpression, "wrong node type")
	u := n.(ast.SliceExpressionNode)
	util.Assert(t, u.Low.Type() == ast.Literal, "wrong slice type")

	p = createParser("array[:5]")
	n = p.parseExpression()
	util.AssertNow(t, n.Type() == ast.SliceExpression, "wrong node type")
	u = n.(ast.SliceExpressionNode)
	util.Assert(t, u.High.Type() == ast.Literal, "wrong slice type")

	p = createParser("array[1:10]")
	n = p.parseExpression()
	util.AssertNow(t, n.Type() == ast.SliceExpression, "wrong node type")
	u = n.(ast.SliceExpressionNode)
	util.Assert(t, u.Low.Type() == ast.Literal, "wrong slice low type")
	util.Assert(t, u.High.Type() == ast.Literal, "wrong slice high type")
}
