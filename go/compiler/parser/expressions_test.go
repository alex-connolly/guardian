package parser

import (
	"testing"

	"github.com/end-r/guardian/go/compiler/lexer"

	"github.com/end-r/guardian/go/compiler/ast"

	"github.com/end-r/goutil"
)

func TestParseReferenceSingle(t *testing.T) {
	p := createParser(`hello`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 1, "wrong token length")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr.Type() == ast.Reference, "wrong expr type")
	ref := expr.(ast.ReferenceNode)
	goutil.AssertNow(t, ref.Names != nil, "ref should not be nil")
	goutil.AssertNow(t, len(ref.Names) == 1, "wrong name length")
	goutil.Assert(t, ref.Names[0] == "hello", "wrong name data")
}

func TestParseReferenceMultiple(t *testing.T) {
	p := createParser(`hello.aaa.bb`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 5, "wrong token length")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr.Type() == ast.Reference, "wrong expr type")
	ref := expr.(ast.ReferenceNode)
	goutil.AssertNow(t, ref.Names != nil, "ref should not be nil")
	goutil.AssertNow(t, len(ref.Names) == 3, "wrong name length")
	goutil.Assert(t, ref.Names[0] == "hello", "wrong name data 0")
	goutil.Assert(t, ref.Names[0] == "aaa", "wrong name data 0")
	goutil.Assert(t, ref.Names[0] == "bb", "wrong name data 0")
}

func TestParseLiteralInteger(t *testing.T) {
	p := createParser(`6`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 1, "wrong token length")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr.Type() == ast.Literal, "wrong expr type")
	lit := expr.(ast.LiteralNode)
	goutil.AssertNow(t, lit.LiteralType == lexer.TknNumber, "wrong literal type")
}

func TestParseLiteralString(t *testing.T) {
	p := createParser(`"alex"`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 1, "wrong token length")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr.Type() == ast.Literal, "wrong expr type")
	lit := expr.(ast.LiteralNode)
	goutil.AssertNow(t, lit.LiteralType == lexer.TknString, "wrong literal type")
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

func TestParseIndexExpressionReferenceReference(t *testing.T) {
	p := createParser(`array[index]`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 4, "wrong token length")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr.Type() == ast.IndexExpression, "wrong expr type")
	indexExpr := expr.(ast.IndexExpressionNode)
	index := indexExpr.Index
	expression := indexExpr.Expression
	goutil.AssertNow(t, expression.Type() == ast.Reference, "wrong expression type")
	goutil.AssertNow(t, index.Type() == ast.Reference, "wrong index type")
}

func TestParseIndexExpressionReferenceLiteral(t *testing.T) {
	p := createParser(`array[6]`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 4, "wrong token length")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr.Type() == ast.IndexExpression, "wrong expr type")
	indexExpr := expr.(ast.IndexExpressionNode)
	index := indexExpr.Index
	expression := indexExpr.Expression
	goutil.AssertNow(t, expression.Type() == ast.Reference, "wrong expression type")
	goutil.AssertNow(t, index.Type() == ast.Literal, "wrong index type")
}

func TestParseIndexExpressionReferenceCall(t *testing.T) {
	p := createParser(`array[getIndex()]`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 6, "wrong token length")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr.Type() == ast.IndexExpression, "wrong expr type")
	indexExpr := expr.(ast.IndexExpressionNode)
	index := indexExpr.Index
	expression := indexExpr.Expression
	goutil.AssertNow(t, expression.Type() == ast.Reference, "wrong expression type")
	goutil.AssertNow(t, index.Type() == ast.CallExpression, "wrong index type")
}

func TestParseIndexExpressionReferenceIndex(t *testing.T) {
	p := createParser(`array[nested[0]]`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 7, "wrong token length")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr.Type() == ast.IndexExpression, "wrong expr type")
	indexExpr := expr.(ast.IndexExpressionNode)
	index := indexExpr.Index
	expression := indexExpr.Expression
	goutil.AssertNow(t, expression.Type() == ast.Reference, "wrong expression type")
	goutil.AssertNow(t, index.Type() == ast.IndexExpression, "wrong index type")
}

func TestParseIndexExpressionCallReference(t *testing.T) {
	p := createParser(`getArray()[index]`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 6, "wrong token length")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr.Type() == ast.IndexExpression, "wrong expr type")
	indexExpr := expr.(ast.IndexExpressionNode)
	index := indexExpr.Index
	expression := indexExpr.Expression
	goutil.AssertNow(t, expression.Type() == ast.CallExpression, "wrong expression type")
	goutil.AssertNow(t, index.Type() == ast.Reference, "wrong index type")
}

func TestParseIndexExpressionCallLiteral(t *testing.T) {
	p := createParser(`getArray()[5]`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 6, "wrong token length")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr.Type() == ast.IndexExpression, "wrong expr type")
	indexExpr := expr.(ast.IndexExpressionNode)
	index := indexExpr.Index
	expression := indexExpr.Expression
	goutil.AssertNow(t, expression.Type() == ast.CallExpression, "wrong expression type")
	goutil.AssertNow(t, index.Type() == ast.Literal, "wrong index type")
}

func TestParseIndexExpressionCallCall(t *testing.T) {
	p := createParser(`getArray()[getIndex()]`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 8, "wrong token length")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr.Type() == ast.IndexExpression, "wrong expr type")
	indexExpr := expr.(ast.IndexExpressionNode)
	index := indexExpr.Index
	expression := indexExpr.Expression
	goutil.AssertNow(t, expression.Type() == ast.CallExpression, "wrong expression type")
	goutil.AssertNow(t, index.Type() == ast.CallExpression, "wrong index type")
}

func TestParseIndexExpressionCallIndex(t *testing.T) {
	p := createParser(`getArray()[nested[0]]`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 9, "wrong token length")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr.Type() == ast.IndexExpression, "wrong expr type")
	indexExpr := expr.(ast.IndexExpressionNode)
	index := indexExpr.Index
	expression := indexExpr.Expression
	goutil.AssertNow(t, expression.Type() == ast.CallExpression, "wrong expression type")
	goutil.AssertNow(t, index.Type() == ast.IndexExpression, "wrong index type")
}
