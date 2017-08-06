package parser

import (
	"axia/guardian/go/compiler/ast"
	"axia/guardian/go/compiler/lexer"
	"axia/guardian/go/util"
	"testing"
)

func TestParseUnaryLiteralPrefixExpression(t *testing.T) {
	p := createParser("++5")
	n := p.parseExpression()
	util.AssertNow(t, n.Type() == ast.UnaryExpression, "wrong node type")
	u := n.(ast.UnaryExpressionNode)
	util.Assert(t, u.Operator == lexer.TknIncrement, "wrong operator")
}

func TestParseUnaryReferencePrefixExpression(t *testing.T) {
	p := createParser("!me")
	n := p.parseExpression()
	util.AssertNow(t, n.Type() == ast.UnaryExpression, "wrong node type")
	u := n.(ast.UnaryExpressionNode)
	util.Assert(t, u.Operator == lexer.TknNot, "wrong operator")

	p = createParser("++me")
	n = p.parseExpression()
	util.AssertNow(t, n.Type() == ast.UnaryExpression, "wrong node type")
	u = n.(ast.UnaryExpressionNode)
	util.Assert(t, u.Operator == lexer.TknIncrement, "wrong operator")
}

func TestParseUnaryLiteralPostfixExpression(t *testing.T) {
	p := createParser("5++")
	n := p.parseExpression()
	util.AssertNow(t, n.Type() == ast.UnaryExpression, "wrong node type")
	u := n.(ast.UnaryExpressionNode)
	util.Assert(t, u.Operator == lexer.TknIncrement, "wrong operator")
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

func TestParseCallExpressionEmpty(t *testing.T) {
	p := createParser("hi()")
	n := p.parseExpression()
	util.AssertNow(t, n.Type() == ast.CallExpression, "wrong node type")
	u := n.(ast.CallExpressionNode)
	util.AssertNow(t, u.Call.Type() == ast.Reference, "wrong call type")
	c := n.(ast.ReferenceNode)
	util.AssertNow(t, len(c.Names) == 1, "wrong reference length")
	util.Assert(t, c.Names[0] == "hi", "wrong reference name")
}

func TestParseCallExpressionEmptyReference(t *testing.T) {
	p := createParser("pkg.hi()")
	n := p.parseExpression()
	util.AssertNow(t, n.Type() == ast.CallExpression, "wrong node type")
	u := n.(ast.CallExpressionNode)
	util.AssertNow(t, u.Call.Type() == ast.Reference, "wrong call type")
	c := n.(ast.ReferenceNode)
	util.AssertNow(t, len(c.Names) == 2, "wrong reference length")
	util.Assert(t, c.Names[0] == "pkg", "wrong reference package name")
	util.Assert(t, c.Names[0] == "hi", "wrong reference name")
}

func TestParseCallExpressionLiterals(t *testing.T) {
	p := createParser(`hi("a", 6, true)`)
	n := p.parseExpression()
	util.AssertNow(t, n.Type() == ast.CallExpression, "wrong node type")
	u := n.(ast.CallExpressionNode)
	util.AssertNow(t, u.Call.Type() == ast.Reference, "wrong call type")
	c := n.(ast.ReferenceNode)
	util.AssertNow(t, len(c.Names) == 1, "wrong reference length")
	util.Assert(t, c.Names[0] == "hi", "wrong reference name")
}

func TestParseCallExpressionLiteralsReference(t *testing.T) {
	p := createParser(`pkg.hi("a", 6, true)`)
	n := p.parseExpression()
	util.AssertNow(t, n.Type() == ast.CallExpression, "wrong node type")
	u := n.(ast.CallExpressionNode)
	util.AssertNow(t, u.Call.Type() == ast.Reference, "wrong call type")
	c := n.(ast.ReferenceNode)
	util.AssertNow(t, len(c.Names) == 2, "wrong reference length")
	util.Assert(t, c.Names[0] == "pkg", "wrong reference package name")
	util.Assert(t, c.Names[0] == "hi", "wrong reference name")
}

func TestParseCallExpressionNestedEmpty(t *testing.T) {
	p := createParser(`hi(a(), b(), c())`)
	n := p.parseExpression()
	util.AssertNow(t, n.Type() == ast.CallExpression, "wrong node type")
	u := n.(ast.CallExpressionNode)
	util.AssertNow(t, u.Call.Type() == ast.Reference, "wrong call type")
	c := n.(ast.ReferenceNode)
	util.AssertNow(t, len(c.Names) == 1, "wrong reference length")
	util.Assert(t, c.Names[0] == "hi", "wrong reference name")
}

func TestParseCallExpressionNestedEmptyReference(t *testing.T) {
	p := createParser(`pkg.hi(a(), b(), c())`)
	n := p.parseExpression()
	util.AssertNow(t, n.Type() == ast.CallExpression, "wrong node type")
	u := n.(ast.CallExpressionNode)
	util.AssertNow(t, u.Call.Type() == ast.Reference, "wrong call type")
	c := n.(ast.ReferenceNode)
	util.AssertNow(t, len(c.Names) == 2, "wrong reference length")
	util.Assert(t, c.Names[0] == "pkg", "wrong reference package name")
	util.Assert(t, c.Names[0] == "hi", "wrong reference name")
}

func TestParseCallExpressionNestedFullReference(t *testing.T) {
	p := createParser(`pkg.hi(a("one", "two", "three"), b(four(), five()), c(six))`)
	n := p.parseExpression()
	util.AssertNow(t, n.Type() == ast.CallExpression, "wrong node type")
	u := n.(ast.CallExpressionNode)
	util.AssertNow(t, u.Call.Type() == ast.Reference, "wrong call type")
	c := n.(ast.ReferenceNode)
	util.AssertNow(t, len(c.Names) == 2, "wrong reference length")
	util.Assert(t, c.Names[0] == "pkg", "wrong reference package name")
	util.Assert(t, c.Names[0] == "hi", "wrong reference name")
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
