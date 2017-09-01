package expressions

import (
	"testing"

	"github.com/end-r/guardian/go/compiler/ast"

	"github.com/end-r/goutil"
)

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

func TestParseIndexExpressionIndexReference(t *testing.T) {
	p := createParser(`array[0][index]`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 6, "wrong token length")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr.Type() == ast.IndexExpression, "wrong expr type")
	indexExpr := expr.(ast.IndexExpressionNode)
	index := indexExpr.Index
	expression := indexExpr.Expression
	goutil.AssertNow(t, expression.Type() == ast.IndexExpression, "wrong expression type")
	goutil.AssertNow(t, index.Type() == ast.Reference, "wrong index type")
}

func TestParseIndexExpressionIndexLiteral(t *testing.T) {
	p := createParser(`array[index][5]`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 6, "wrong token length")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr.Type() == ast.IndexExpression, "wrong expr type")
	indexExpr := expr.(ast.IndexExpressionNode)
	index := indexExpr.Index
	expression := indexExpr.Expression
	goutil.AssertNow(t, expression.Type() == ast.IndexExpression, "wrong expression type")
	goutil.AssertNow(t, index.Type() == ast.Literal, "wrong index type")
}

func TestParseIndexExpressionIndexCall(t *testing.T) {
	p := createParser(`array[4][getIndex()]`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 8, "wrong token length")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr.Type() == ast.IndexExpression, "wrong expr type")
	indexExpr := expr.(ast.IndexExpressionNode)
	index := indexExpr.Index
	expression := indexExpr.Expression
	goutil.AssertNow(t, expression.Type() == ast.IndexExpression, "wrong expression type")
	goutil.AssertNow(t, index.Type() == ast.CallExpression, "wrong index type")
}

func TestParseIndexExpressionIndexIndex(t *testing.T) {
	p := createParser(`array[nested[0]][nested[0]]`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 9, "wrong token length")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr.Type() == ast.IndexExpression, "wrong expr type")
	indexExpr := expr.(ast.IndexExpressionNode)
	index := indexExpr.Index
	expression := indexExpr.Expression
	goutil.AssertNow(t, expression.Type() == ast.IndexExpression, "wrong expression type")
	goutil.AssertNow(t, index.Type() == ast.IndexExpression, "wrong index type")
}

func TestParseIndexExpressionLiteralReference(t *testing.T) {
	p := createParser(`"hello"[index]`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 6, "wrong token length")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr.Type() == ast.IndexExpression, "wrong expr type")
	indexExpr := expr.(ast.IndexExpressionNode)
	index := indexExpr.Index
	expression := indexExpr.Expression
	goutil.AssertNow(t, expression.Type() == ast.Literal, "wrong expression type")
	goutil.AssertNow(t, index.Type() == ast.Reference, "wrong index type")
}

func TestParseIndexExpressionLiteralLiteral(t *testing.T) {
	p := createParser(`"hello"[5]`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 6, "wrong token length")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr.Type() == ast.IndexExpression, "wrong expr type")
	indexExpr := expr.(ast.IndexExpressionNode)
	index := indexExpr.Index
	expression := indexExpr.Expression
	goutil.AssertNow(t, expression.Type() == ast.Literal, "wrong expression type")
	goutil.AssertNow(t, index.Type() == ast.Literal, "wrong index type")
}

func TestParseIndexExpressionLiteralCall(t *testing.T) {
	p := createParser(`"hello"[getIndex()]`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 8, "wrong token length")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr.Type() == ast.IndexExpression, "wrong expr type")
	indexExpr := expr.(ast.IndexExpressionNode)
	index := indexExpr.Index
	expression := indexExpr.Expression
	goutil.AssertNow(t, expression.Type() == ast.IndexExpression, "wrong expression type")
	goutil.AssertNow(t, index.Type() == ast.Literal, "wrong index type")
}

func TestParseIndexExpressionLiteralIndex(t *testing.T) {
	p := createParser(`"hello"[nested[0]]`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 9, "wrong token length")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr.Type() == ast.IndexExpression, "wrong expr type")
	indexExpr := expr.(ast.IndexExpressionNode)
	index := indexExpr.Index
	expression := indexExpr.Expression
	goutil.AssertNow(t, expression.Type() == ast.Literal, "wrong expression type")
	goutil.AssertNow(t, index.Type() == ast.IndexExpression, "wrong index type")
}
