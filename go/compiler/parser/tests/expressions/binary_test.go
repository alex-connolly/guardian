package expressions

import (
	"testing"

	"github.com/end-r/goutil"
	"github.com/end-r/guardian/go/compiler/ast"
	"github.com/end-r/guardian/go/compiler/lexer"
	"github.com/end-r/guardian/go/compiler/parser"
	"github.com/end-r/guardian/go/util"
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

func TestParseBinaryExpressionStringConcat(t *testing.T) {
	p := createParser(`"alex" + "bob"`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 3, "wrong token length")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr.Type() == ast.BinaryExpression, "wrong expr type")
	bin := expr.(ast.BinaryExpressionNode)
	left := bin.Left
	right := bin.Right
	goutil.AssertNow(t, left.Type() == ast.Literal, "wrong left type")
	goutil.AssertNow(t, right.Type() == ast.Literal, "wrong right type")
}

func TestParseBinaryExpressionEmptyCallExpressions(t *testing.T) {
	p := createParser(`call() + call()`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 7, "wrong token length")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr.Type() == ast.BinaryExpression, "wrong expr type")
	bin := expr.(ast.BinaryExpressionNode)
	left := bin.Left
	right := bin.Right
	goutil.AssertNow(t, left.Type() == ast.CallExpression, "wrong left type")
	goutil.AssertNow(t, right.Type() == ast.CallExpression, "wrong right type")
}

func TestParseBinaryExpressionFullCallExpressions(t *testing.T) {
	p := createParser(`call(15, 29) + call("Alex")`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 11, "wrong token length")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr.Type() == ast.BinaryExpression, "wrong expr type")
	bin := expr.(ast.BinaryExpressionNode)
	left := bin.Left
	right := bin.Right
	goutil.AssertNow(t, left.Type() == ast.CallExpression, "wrong left type")
	goutil.AssertNow(t, right.Type() == ast.CallExpression, "wrong right type")
}

func TestParseBinaryExpressionNested(t *testing.T) {
	p := createParser(`5 + 5 + 3`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 5, "wrong token length")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr.Type() == ast.BinaryExpression, "wrong expr type")
	bin := expr.(ast.BinaryExpressionNode)
	left := bin.Left
	right := bin.Right
	goutil.AssertNow(t, left.Type() == ast.BinaryExpression, "wrong left type")
	goutil.AssertNow(t, right.Type() == ast.CallExpression, "wrong right type")
}
