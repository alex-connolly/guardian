package parser

import (
	"fmt"
	"testing"

	"github.com/end-r/guardian/token"

	"github.com/end-r/guardian/ast"

	"github.com/end-r/goutil"
)

func TestParseIdentifierSingle(t *testing.T) {
	p := createParser(`hello`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 1, "wrong token length")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr should not be nil")
	goutil.AssertNow(t, expr.Type() == ast.Identifier, "wrong expr type")
}

func TestParseReference(t *testing.T) {
	p := createParser(`hello.aaa.bb`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 5, "wrong token length")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr should not be nil")
	goutil.AssertNow(t, expr.Type() == ast.Reference, "wrong expr type")
}

func TestParseLiteralInteger(t *testing.T) {
	p := createParser(`6`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 1, "wrong token length")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr should not be nil")
	goutil.AssertNow(t, expr.Type() == ast.Literal, "wrong expr type")
	lit := expr.(*ast.LiteralNode)
	goutil.AssertNow(t, lit.LiteralType == token.Integer, "wrong literal type")
}

func TestParseLiteralString(t *testing.T) {
	p := createParser(`"alex"`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 1, "wrong token length")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr should not be nil")
	goutil.AssertNow(t, expr.Type() == ast.Literal, "wrong expr type")
	lit := expr.(*ast.LiteralNode)
	goutil.AssertNow(t, lit.LiteralType == token.String, "wrong literal type")
}

func TestParseMapLiteralEmpty(t *testing.T) {
	p := createParser("map[string]int{}")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expression shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.MapLiteral, "wrong node type")

}

func TestParseMapLiteralSingle(t *testing.T) {
	p := createParser(`map[string]int{"Hi":3}`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expression shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.MapLiteral, "wrong node type")

}

func TestParseMapLiteralMultiple(t *testing.T) {
	p := createParser(`map[string]int{"Hi":3, "Byte":8}`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expression shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.MapLiteral, "wrong node type")
	n := expr.(*ast.MapLiteralNode)

	goutil.Assert(t, len(n.Data) == 2, "wrong data length")
}

func TestParseSliceExpressionReferenceLowLiteral(t *testing.T) {
	p := createParser("slice[6:]")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.SliceExpression, "wrong node type")
	s := expr.(*ast.SliceExpressionNode)
	goutil.AssertNow(t, s.Expression != nil, "expression shouldn't be nil")
	goutil.AssertNow(t, s.Expression.Type() == ast.Identifier, "wrong expression type")
	goutil.AssertNow(t, s.Low.Type() == ast.Literal, "wrong low type")
	l := s.Low.(*ast.LiteralNode)
	goutil.AssertNow(t, l.Data == "6", "wrong data")
}

func TestParseSliceExpressionReferenceLowReference(t *testing.T) {
	p := createParser("slice[low:]")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.SliceExpression, "wrong node type")
	s := expr.(*ast.SliceExpressionNode)
	goutil.AssertNow(t, s.Expression.Type() == ast.Identifier, "wrong expression type")
}

func TestParseSliceExpressionReferenceLowCall(t *testing.T) {
	p := createParser("slice[low():]")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.SliceExpression, "wrong node type")
	s := expr.(*ast.SliceExpressionNode)
	goutil.AssertNow(t, s.Expression.Type() == ast.Identifier, "wrong expression type")
	goutil.AssertNow(t, s.Low.Type() == ast.CallExpression, "wrong low type")
	l := s.Low.(*ast.CallExpressionNode)
	goutil.AssertNow(t, l.Arguments == nil, "call arguments should be nil")
	goutil.AssertNow(t, l.Call.Type() == ast.Identifier, "wrong call type")
}

func TestParseSliceExpressionCallLowLiteral(t *testing.T) {
	p := createParser("getSlice()[6:]")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.SliceExpression, "wrong node type")
	s := expr.(*ast.SliceExpressionNode)
	goutil.AssertNow(t, s.Expression.Type() == ast.CallExpression, "wrong expression type")
	c := s.Expression.(*ast.CallExpressionNode)
	goutil.AssertNow(t, c.Arguments == nil, "arguments should be nil")
	goutil.AssertNow(t, c.Call.Type() == ast.Identifier, "wrong call type")
}

func TestParseSliceExpressionCallLowReference(t *testing.T) {
	p := createParser("getSlice()[low:]")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.SliceExpression, "wrong node type")
	s := expr.(*ast.SliceExpressionNode)
	goutil.AssertNow(t, s.Expression.Type() == ast.CallExpression, "wrong expression type")
	c := s.Expression.(*ast.CallExpressionNode)
	goutil.AssertNow(t, c.Arguments == nil, "arguments should be nil")
	goutil.AssertNow(t, c.Call.Type() == ast.Identifier, "wrong call type")
}

func TestParseSliceExpressionCallLowCall(t *testing.T) {
	p := createParser("getSlice()[getLow():]")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.SliceExpression, "wrong node type")
	s := expr.(*ast.SliceExpressionNode)
	goutil.AssertNow(t, s.Expression.Type() == ast.CallExpression, "wrong expression type")
	c := s.Expression.(*ast.CallExpressionNode)
	goutil.AssertNow(t, c.Arguments == nil, "arguments should be nil")
	goutil.AssertNow(t, c.Call.Type() == ast.Identifier, "wrong call type")
}

func TestParseSliceExpressionArrayLiteralowLiteral(t *testing.T) {
	p := createParser(`[]string{"a", "b", "c"}[6:]`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.Assert(t, expr.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionArrayLiteralowReference(t *testing.T) {
	p := createParser(`[]string{"a", "b", "c"}[low:]`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.Assert(t, expr.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionArrayLiteralowCall(t *testing.T) {
	p := createParser(`[]string{"a", "b", "c"}[getLow():]`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.Assert(t, expr.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionReferenceHighLiteral(t *testing.T) {
	p := createParser("slice[:6]")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.Assert(t, expr.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionReferenceHighReference(t *testing.T) {
	p := createParser("slice[:high]")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.Assert(t, expr.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionReferenceHighCall(t *testing.T) {
	p := createParser("slice[:high()]")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.Assert(t, expr.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionCallHighLiteral(t *testing.T) {
	p := createParser("getSlice()[:6]")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.Assert(t, expr.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionCallHighReference(t *testing.T) {
	p := createParser("getSlice()[:high]")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.Assert(t, expr.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionCallHighCall(t *testing.T) {
	p := createParser("getSlice()[:getHigh()]")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.Assert(t, expr.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionArrayLiteralHighLiteral(t *testing.T) {
	p := createParser(`[]string{"a", "b", "c"}[:6]`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.Assert(t, expr.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionArrayLiteralHighReference(t *testing.T) {
	p := createParser(`[]string{"a", "b", "c"}[:high]`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.Assert(t, expr.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionArrayLiteralHighCall(t *testing.T) {
	p := createParser(`[]string{"a", "b", "c"}[:getHigh()]`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.Assert(t, expr.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionReferenceLowHighLiteral(t *testing.T) {
	p := createParser("slice[2:6]")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.Assert(t, expr.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionReferenceLowHighReference(t *testing.T) {
	p := createParser("slice[low:high]")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.Assert(t, expr.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionReferenceLowHighCall(t *testing.T) {
	p := createParser("slice[low():high()]")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.Assert(t, expr.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionCallLowHighLiteral(t *testing.T) {
	p := createParser("getSlice()[3:6]")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.Assert(t, expr.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionCallLowHighReference(t *testing.T) {
	p := createParser("getSlice()[low:high]")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.Assert(t, expr.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionCallLowHighCall(t *testing.T) {
	p := createParser("getSlice()[getLow():getHigh()]")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.Assert(t, expr.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionArrayLiteralowHighLiteral(t *testing.T) {
	p := createParser(`[]string{"a", "b", "c"}[1:6]`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.Assert(t, expr.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionArrayLiteralowHighReference(t *testing.T) {
	p := createParser(`[]string{"a", "b", "c"}[low:high]`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.Assert(t, expr.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionArrayLiteralowHighCall(t *testing.T) {
	p := createParser(`[]string{"a", "b", "c"}[getLow():getHigh()]`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.Assert(t, expr.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseCompositeLiteralEmpty(t *testing.T) {
	p := createParser("Dog{}")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.CompositeLiteral, "wrong node type")
}

func TestParseCompositeLiteralDeepReferenceEmpty(t *testing.T) {
	p := createParser("animals.Dog{}")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.Reference, "wrong node type")
}

func TestParseCompositeLiteralInline(t *testing.T) {
	p := createParser(`Dog{name: "Mr Woof"}`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expression shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.CompositeLiteral, "wrong node type")
	n := expr.(*ast.CompositeLiteralNode)
	goutil.AssertNow(t, n.Fields != nil, "fields shouldn't be nil")
	goutil.AssertNow(t, len(n.Fields) == 1, "wrong number of fields")
}

func TestParseCompositeLiteralMixed(t *testing.T) {
	expr := ParseExpression(`Proposal{
		name: proposalNames[i],
		voteCount: 0,
	}`)
	goutil.Assert(t, expr.Type() == ast.CompositeLiteral, "wrong expr type")
}

func TestParseCallExpressionCompositeLiteral(t *testing.T) {
	expr := ParseExpression(`append(proposals, Proposal{
		name: proposalNames[i],
		voteCount: 0,
	}`)
	goutil.Assert(t, expr.Type() == ast.CallExpression, "wrong expr type")
}

func TestParseCompositeLiteralMultiline(t *testing.T) {
	p := createParser(`Dog{
		name: "Mr Woof",
		age: 17,
		weight: calculateWeight(),
		colour: estimateColour("blueish"),
		height: 180,
		}`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expression shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.CompositeLiteral, "wrong node type")
}

func TestParseIndexExpressionReferenceReference(t *testing.T) {
	p := createParser(`array[index]`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.IndexExpression, "wrong expr type")
	indexExpr := expr.(*ast.IndexExpressionNode)
	index := indexExpr.Index
	expression := indexExpr.Expression
	goutil.AssertNow(t, expression != nil, "expression shouldn't be nil")
	goutil.AssertNow(t, expression.Type() == ast.Identifier, "wrong expression type")
	goutil.AssertNow(t, index.Type() == ast.Identifier, "wrong index type")
}

func TestParseIndexExpressionReferenceLiteral(t *testing.T) {
	p := createParser(`array[6]`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.IndexExpression, "wrong expr type")
	indexExpr := expr.(*ast.IndexExpressionNode)
	index := indexExpr.Index
	expression := indexExpr.Expression
	goutil.AssertNow(t, expression != nil, "expression shouldn't be nil")
	goutil.AssertNow(t, expression.Type() == ast.Identifier, "wrong expression type")
	goutil.AssertNow(t, index.Type() == ast.Literal, "wrong index type")
}

func TestParseIndexExpressionReferenceCall(t *testing.T) {
	p := createParser(`array[getIndex()]`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.IndexExpression, "wrong expr type")
	indexExpr := expr.(*ast.IndexExpressionNode)
	index := indexExpr.Index
	expression := indexExpr.Expression
	goutil.AssertNow(t, expression != nil, "expression shouldn't be nil")
	goutil.AssertNow(t, expression.Type() == ast.Identifier, "wrong expression type")
	goutil.AssertNow(t, index.Type() == ast.CallExpression, "wrong index type")
}

func TestParseIndexExpressionReferenceIndex(t *testing.T) {
	p := createParser(`array[nested[0]]`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.IndexExpression, "wrong expr type")
	indexExpr := expr.(*ast.IndexExpressionNode)
	index := indexExpr.Index
	expression := indexExpr.Expression
	goutil.AssertNow(t, expression != nil, "expression shouldn't be nil")
	goutil.AssertNow(t, expression.Type() == ast.Identifier, "wrong expression type")
	goutil.AssertNow(t, index.Type() == ast.IndexExpression, "wrong index type")
}

func TestParseIndexExpressionCallReference(t *testing.T) {
	p := createParser(`getArray()[index]`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.IndexExpression, "wrong expr type")
	indexExpr := expr.(*ast.IndexExpressionNode)
	index := indexExpr.Index
	expression := indexExpr.Expression
	goutil.AssertNow(t, expression != nil, "expression shouldn't be nil")
	goutil.AssertNow(t, expression.Type() == ast.CallExpression, "wrong expression type")
	goutil.AssertNow(t, index.Type() == ast.Identifier, "wrong index type")
}

func TestParseIndexExpressionCallLiteral(t *testing.T) {
	p := createParser(`getArray()[5]`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.IndexExpression, "wrong expr type")
	indexExpr := expr.(*ast.IndexExpressionNode)
	index := indexExpr.Index
	expression := indexExpr.Expression
	goutil.AssertNow(t, expression != nil, "expression shouldn't be nil")
	goutil.AssertNow(t, expression.Type() == ast.CallExpression, "wrong expression type")
	goutil.AssertNow(t, index.Type() == ast.Literal, "wrong index type")
}

func TestParseIndexExpressionCallCall(t *testing.T) {
	p := createParser(`getArray()[getIndex()]`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.IndexExpression, "wrong expr type")
	indexExpr := expr.(*ast.IndexExpressionNode)
	index := indexExpr.Index
	expression := indexExpr.Expression
	goutil.AssertNow(t, expression != nil, "expression shouldn't be nil")
	goutil.AssertNow(t, expression.Type() == ast.CallExpression, "wrong expression type")
	goutil.AssertNow(t, index.Type() == ast.CallExpression, "wrong index type")
}

func TestParseIndexExpressionCallIndex(t *testing.T) {
	p := createParser(`getArray()[nested[0]]`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.IndexExpression, "wrong expr type")
	indexExpr := expr.(*ast.IndexExpressionNode)
	index := indexExpr.Index
	expression := indexExpr.Expression
	goutil.AssertNow(t, expression != nil, "expression shouldn't be nil")
	goutil.AssertNow(t, expression.Type() == ast.CallExpression, "wrong expression type")
	goutil.AssertNow(t, index.Type() == ast.IndexExpression, "wrong index type")
}

func TestParseIndexExpressionIndexReference(t *testing.T) {
	p := createParser(`array[0][index]`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.IndexExpression, "wrong expr type")
	indexExpr := expr.(*ast.IndexExpressionNode)
	index := indexExpr.Index
	expression := indexExpr.Expression
	goutil.AssertNow(t, expression != nil, "expression shouldn't be nil")
	goutil.AssertNow(t, expression.Type() == ast.IndexExpression, "wrong expression type")
	goutil.AssertNow(t, index.Type() == ast.Identifier, "wrong index type")
}

func TestParseIndexExpressionIndexLiteral(t *testing.T) {
	p := createParser(`array[index][5]`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.IndexExpression, "wrong expr type")
	indexExpr := expr.(*ast.IndexExpressionNode)
	index := indexExpr.Index
	expression := indexExpr.Expression
	goutil.AssertNow(t, expression != nil, "expression shouldn't be nil")
	goutil.AssertNow(t, expression.Type() == ast.IndexExpression, "wrong expression type")
	goutil.AssertNow(t, index.Type() == ast.Literal, "wrong index type")
}

func TestParseIndexExpressionIndexCall(t *testing.T) {
	p := createParser(`array[4][getIndex()]`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.IndexExpression, "wrong expr type")
	indexExpr := expr.(*ast.IndexExpressionNode)
	index := indexExpr.Index
	expression := indexExpr.Expression
	goutil.AssertNow(t, expression != nil, "expression shouldn't be nil")
	goutil.AssertNow(t, expression.Type() == ast.IndexExpression, "wrong expression type")
	goutil.AssertNow(t, index.Type() == ast.CallExpression, "wrong index type")
}

func TestParseIndexExpressionIndexIndex(t *testing.T) {
	p := createParser(`array[nested[0]][nested[0]]`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.IndexExpression, "wrong expr type")
	indexExpr := expr.(*ast.IndexExpressionNode)
	index := indexExpr.Index
	expression := indexExpr.Expression
	goutil.AssertNow(t, expression.Type() == ast.IndexExpression, "wrong expression type")
	goutil.AssertNow(t, index.Type() == ast.IndexExpression, "wrong index type")
}

func TestParseBinaryExpressionLiteralLiteral(t *testing.T) {
	p := createParser(`6 + 4`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.BinaryExpression, "wrong expr type")
	b := expr.(*ast.BinaryExpressionNode)
	goutil.AssertNow(t, b.Left.Type() == ast.Literal, "wrong left type")
	goutil.AssertNow(t, b.Right.Type() == ast.Literal, "wrong right type")
	goutil.AssertNow(t, b.Operator == token.Add, "wrong operator")
}

func TestParseBinaryExpressionLiteralLiteralBracketed(t *testing.T) {
	p := createParser(`(6 + 4)`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.BinaryExpression, "wrong expr type")
	b := expr.(*ast.BinaryExpressionNode)
	goutil.AssertNow(t, b.Left.Type() == ast.Literal, "wrong left type")
	goutil.AssertNow(t, b.Right.Type() == ast.Literal, "wrong right type")
	goutil.AssertNow(t, b.Operator == token.Add, "wrong operator")
}

func TestParseBinaryExpressionReferenceReference(t *testing.T) {
	p := createParser(`a - b`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.BinaryExpression, "wrong expr type")
	b := expr.(*ast.BinaryExpressionNode)
	goutil.AssertNow(t, b.Left.Type() == ast.Identifier, "wrong left type")
	goutil.AssertNow(t, b.Right.Type() == ast.Identifier, "wrong right type")
	goutil.AssertNow(t, b.Operator == token.Sub, "wrong operator")
}

func TestParseBinaryExpressionIndexIndex(t *testing.T) {
	p := createParser(`a[0] * b[1]`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.BinaryExpression, "wrong expr type")
	b := expr.(*ast.BinaryExpressionNode)
	goutil.AssertNow(t, b.Left.Type() == ast.IndexExpression, "wrong left type")
	goutil.AssertNow(t, b.Right.Type() == ast.IndexExpression, "wrong right type")
	goutil.AssertNow(t, b.Operator == token.Mul, "wrong operator")
}

func TestParseBinaryExpressionCallCall(t *testing.T) {
	p := createParser(`a(0) / b(0)`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.BinaryExpression, "wrong expr type")
	b := expr.(*ast.BinaryExpressionNode)
	goutil.AssertNow(t, b.Left.Type() == ast.CallExpression, "wrong left type")
	goutil.AssertNow(t, b.Right.Type() == ast.CallExpression, "wrong right type")
	goutil.AssertNow(t, b.Operator == token.Div, "wrong operator")
}

func TestParseUnaryExpressionReference(t *testing.T) {
	p := createParser(`!me`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.UnaryExpression, "wrong expr type")
	u := expr.(*ast.UnaryExpressionNode)
	goutil.AssertNow(t, u.Operand.Type() == ast.Identifier, "wrong left type")
	goutil.AssertNow(t, u.Operator == token.Not, "wrong operator")
}

func TestParseUnaryExpressionIndex(t *testing.T) {
	p := createParser(`!me[0]`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.UnaryExpression, "wrong expr type")
	u := expr.(*ast.UnaryExpressionNode)
	goutil.AssertNow(t, u.Operand.Type() == ast.IndexExpression, "wrong left type")
	goutil.AssertNow(t, u.Operator == token.Not, "wrong operator")
}

func TestParseChainedExpressionSimple(t *testing.T) {
	p := createParser("5 + 4")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.BinaryExpression, "wrong expr type")
	b := expr.(*ast.BinaryExpressionNode)
	goutil.AssertNow(t, b.Left.Type() == ast.Literal, "wrong left type")
	goutil.AssertNow(t, b.Right.Type() == ast.Literal, "wrong right type")
	goutil.AssertNow(t, b.Operator == token.Add, "wrong operator")
}

func TestParseBinaryExpressionComparative(t *testing.T) {
	p := createParser("5 > 4")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.BinaryExpression, "wrong expr type")
	b := expr.(*ast.BinaryExpressionNode)
	goutil.AssertNow(t, b.Left.Type() == ast.Literal, "wrong left type")
	goutil.AssertNow(t, b.Right.Type() == ast.Literal, "wrong right type")
	goutil.AssertNow(t, b.Operator == token.Gtr, "wrong operator")
}

func TestParseChainedExpressionThreeLiterals(t *testing.T) {
	p := createParser("5 + 4 - 3")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.BinaryExpression, "wrong expr type")
	b := expr.(*ast.BinaryExpressionNode)
	goutil.AssertNow(t, b.Left.Type() == ast.BinaryExpression, "wrong left type")
	goutil.AssertNow(t, b.Right.Type() == ast.Literal, "wrong right type")
	goutil.AssertNow(t, b.Operator == token.Sub, "wrong operator")
}

func TestParseChainedExpressionLiteralsSingleBracket(t *testing.T) {
	p := createParser("5 + (4 - 3)")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.BinaryExpression, "wrong expr type")
	b := expr.(*ast.BinaryExpressionNode)
	goutil.AssertNow(t, b.Left.Type() == ast.Literal, "wrong left type")
	goutil.AssertNow(t, b.Right.Type() == ast.BinaryExpression, "wrong right type")
	goutil.AssertNow(t, b.Operator == token.Add, "wrong operator")
}

func TestParseChainedExpressionLiteralsDoubleBracket(t *testing.T) {
	p := createParser("(5 + 2) + (4 - 3)")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.BinaryExpression, "wrong expr type")
	b := expr.(*ast.BinaryExpressionNode)
	goutil.AssertNow(t, b.Left.Type() == ast.BinaryExpression, "wrong left type")
	goutil.AssertNow(t, b.Right.Type() == ast.BinaryExpression, "wrong right type")
	goutil.AssertNow(t, b.Operator == token.Add, "wrong operator")
}

func TestParseChainedExpressionLiteralsExpectPrecedence(t *testing.T) {
	p := createParser("5 + 4 * 3")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.BinaryExpression, "wrong expr type")
	b := expr.(*ast.BinaryExpressionNode)
	goutil.AssertNow(t, b.Left.Type() == ast.Literal, "wrong left type")
	goutil.AssertNow(t, b.Right.Type() == ast.BinaryExpression, "wrong right type")
	goutil.AssertNow(t, b.Operator == token.Add, "wrong operator")
}

func TestParseChainedExpressionLiteralsOverridePrecedence(t *testing.T) {
	p := createParser("(5 + 4) * 3")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.BinaryExpression, "wrong expr type")
	b := expr.(*ast.BinaryExpressionNode)
	goutil.AssertNow(t, b.Left.Type() == ast.BinaryExpression, "wrong left type")
	goutil.AssertNow(t, b.Right.Type() == ast.Literal, "wrong right type")
	goutil.AssertNow(t, b.Operator == token.Mul, "wrong operator")
}

func TestParseChainedExpressionReferencesOverridePrecedence(t *testing.T) {
	p := createParser("(a + b) * c")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.BinaryExpression, "wrong expr type")
	b := expr.(*ast.BinaryExpressionNode)
	goutil.AssertNow(t, b.Left.Type() == ast.BinaryExpression, "wrong left type")
	goutil.AssertNow(t, b.Right.Type() == ast.Identifier, "wrong right type")
	goutil.AssertNow(t, b.Operator == token.Mul, "wrong operator")
}

func TestParseChainedExpressionCallsOverridePrecedence(t *testing.T) {
	p := createParser("(a() + b(1)) * c(1, 2)")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.BinaryExpression, "wrong expr type")
	b := expr.(*ast.BinaryExpressionNode)
	goutil.AssertNow(t, b.Left.Type() == ast.BinaryExpression, "wrong left type")
	goutil.AssertNow(t, b.Right.Type() == ast.CallExpression, "wrong right type")
	goutil.AssertNow(t, b.Operator == token.Mul, "wrong operator")
}

func TestParseHighlyChainedExpressionCallsOverridePrecedence(t *testing.T) {
	p := createParser("(a() + b(1)) * (c(1, 2) + d(1, 2, 3))")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.BinaryExpression, "wrong expr type")
	b := expr.(*ast.BinaryExpressionNode)
	goutil.AssertNow(t, b.Left.Type() == ast.BinaryExpression, "wrong left type")
	goutil.AssertNow(t, b.Right.Type() == ast.BinaryExpression, "wrong right type")
	goutil.AssertNow(t, b.Operator == token.Mul, "wrong operator")
}

func TestParseCallExpressionSingleParameter(t *testing.T) {
	p := createParser("do(6)")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.CallExpression, "wrong expr type")
	c := expr.(*ast.CallExpressionNode)
	goutil.AssertNow(t, c.Call.Type() == ast.Identifier, "wrong call type")
	goutil.AssertNow(t, len(c.Arguments) == 1, "wrong arg length")
}

func TestParseCallExpressionMultipleParameters(t *testing.T) {
	p := createParser("do(6, 5)")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.CallExpression, "wrong expr type")
	c := expr.(*ast.CallExpressionNode)
	goutil.AssertNow(t, c.Call.Type() == ast.Identifier, "wrong call type")
	goutil.AssertNow(t, len(c.Arguments) == 2, "wrong arg length")
}

func TestParseReferenceExpressionIndexExpression(t *testing.T) {
	p := createParser("data[5].hello")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.Reference, "wrong expr type")
}

func TestParsenIndexExpressionAddition(t *testing.T) {
	p := createParser("proposals[p] + 99")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.BinaryExpression, "wrong expr type")
}

func TestParsenIndexExpressionAdditionReference(t *testing.T) {
	p := createParser("proposals[p].voteCount + 99")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.BinaryExpression, "wrong expr type")
}

func TestParseReferenceExpressionIndexExpressionComparison(t *testing.T) {
	p := createParser("proposals[p].voteCount > 99")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.BinaryExpression, "wrong expr type")
}

func TestParseFuncLiteralSingleParameter(t *testing.T) {
	expr := ParseExpression("func (a string) int { return 0 }")
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.FuncLiteral, "wrong expr type")
	f := expr.(*ast.FuncLiteralNode)
	goutil.AssertNow(t, len(f.Parameters) == 1, "wrong parameter length")
	goutil.AssertNow(t, len(f.Results) == 1, fmt.Sprintf("wrong result length: %d", len(f.Results)))
}

func TestParseArrayLiteralEmpty(t *testing.T) {
	expr := ParseExpression("[]string{}")
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.ArrayLiteral, "wrong expr type")
	f := expr.(*ast.ArrayLiteralNode)
	goutil.AssertNow(t, f.Signature != nil, "non-null sig")
}

func TestParseArrayLiteralSingle(t *testing.T) {
	expr := ParseExpression(`[]string{"a"}`)
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.ArrayLiteral, "wrong expr type")
	f := expr.(*ast.ArrayLiteralNode)
	goutil.AssertNow(t, f.Signature != nil, "non-null sig")
}

func TestParseArrayLiteralMultiple(t *testing.T) {
	expr := ParseExpression(`[]string{"a", "b"}`)
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.ArrayLiteral, "wrong expr type")
	f := expr.(*ast.ArrayLiteralNode)
	goutil.AssertNow(t, f.Signature != nil, "non-null sig")
}

func TestParseArrayLiteralMultipleSpaced(t *testing.T) {
	expr := ParseExpression(`[]string{
		"a,
		"b"
		}`)
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.ArrayLiteral, "wrong expr type")
	f := expr.(*ast.ArrayLiteralNode)
	goutil.AssertNow(t, f.Signature != nil, "non-null sig")
}

func TestParseArrayLiteralSingleSpaced(t *testing.T) {
	expr := ParseExpression(`[]string{
		"a"
		}`)
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.ArrayLiteral, "wrong expr type")
	f := expr.(*ast.ArrayLiteralNode)
	goutil.AssertNow(t, f.Signature != nil, "non-null sig")
}

func TestParseSimpleExpression(t *testing.T) {
	p := createParser("a {}")
	expr := p.parseSimpleExpression()
	goutil.AssertNow(t, expr.Type() == ast.Identifier, "wrong type")
	i := expr.(*ast.IdentifierNode)
	goutil.AssertNow(t, i.Name == "a", "wrong name")
}

func TestParseCallExpressionCompositeLiteralParameter(t *testing.T) {
	_, errs := ParseString(`contributions.push(
		Contribution {
			amount: msg.value,
			contributor: msg.sender,
		}
	)`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestParseCallExpressionSequential(t *testing.T) {
	_, errs := ParseString(`
		func auctionEnd(){
			assert(now() >= auctionEnd)
		}
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

/*
func TestParseCastExpressionPlainType(t *testing.T) {
	expr := ParseExpression(`a.b(5)`)
	goutil.AssertNow(t, expr.Type() == ast.CallExpression, "wrong node type")
}

func TestParseCastExpressionMapType(t *testing.T) {
	expr := ParseExpression(`map[string]string(a.b)`)
	goutil.AssertNow(t, expr.Type() == ast.CallExpression, "wrong node type")
}

func TestParseCastExpressionArrayType(t *testing.T) {
	expr := ParseExpression(`[]string(a[b])`)
	goutil.AssertNow(t, expr.Type() == ast.CallExpression, "wrong node type")
}

func TestParseCastExpressionInvalidType(t *testing.T) {
	expr := ParseExpression(`6(5)`)
	goutil.AssertNow(t, expr == nil, "should be nil")
}

func TestParseCastExpressionInvalidTypeChained(t *testing.T) {
	expr := ParseExpression(`6 + 1(5)`)
	goutil.AssertNow(t, expr == nil, "should be nil")
}x

func TestParseCastExpressionPlainTypeChained(t *testing.T) {
	expr := ParseExpression(`a.b + 1(5)`)
	goutil.AssertNow(t, expr.Type() == ast.BinaryExpression, "wrong node type")
	b := expr.(*ast.BinaryExpressionNode)
	goutil.AssertNow(t, b.Left != nil, "left is nil")
	goutil.AssertNow(t, b.Right != nil, "right is nil")
}*/

func TestParseBinaryExpressionUnfinished(t *testing.T) {
	expr := ParseExpression(`5 +`)
	goutil.AssertNow(t, expr == nil, "should be nil")
}

func TestParseBinaryExpressionDoubleOperator(t *testing.T) {
	expr := ParseExpression(`5 + + 5`)
	goutil.AssertNow(t, expr == nil, "should be nil")
}

func TestParseBinaryExpressionDoubleOperatorCloseBracket(t *testing.T) {
	expr := ParseExpression(`5 + ) + 5`)
	goutil.AssertNow(t, expr == nil, "should be nil")
}

func TestParseBinaryExpressionDoubleOperatorOpenBracket(t *testing.T) {
	expr := ParseExpression(`5 + ( + 5`)
	goutil.AssertNow(t, expr == nil, "should be nil")
}

func TestParseBinaryExpressionDoubleOperatorFullBrackets(t *testing.T) {
	expr := ParseExpression(`(5 + ) (+ 5)`)
	goutil.AssertNow(t, expr == nil, "should be nil")
}

func TestParseBinaryExpressionUnmatchedOpenBracket(t *testing.T) {
	expr := ParseExpression(`(5 + `)
	goutil.AssertNow(t, expr == nil, "should be nil")
}

func TestParseBinaryExpressionUnmatchedCloseBracket(t *testing.T) {
	expr := ParseExpression(`5 + )`)
	goutil.AssertNow(t, expr == nil, "should be nil")
}

func TestParseMapLiteralMultipleLinesNoTrailingComma(t *testing.T) {
	expr := ParseExpression(`map[string]int {
		"hi": 1,
		"bye": 2
	}`)
	goutil.AssertNow(t, expr != nil, "should not be nil")
	goutil.AssertNow(t, expr.Type() == ast.MapLiteral, "wrong node type")
	m := expr.(*ast.MapLiteralNode)
	goutil.AssertNow(t, len(m.Data) == 2, "wrong data length")
}

func TestParseMapLiteralMultipleLinesTrailingComma(t *testing.T) {
	expr := ParseExpression(`map[string]int {
		"hi": 1,
		"bye": 2,
	}`)
	goutil.AssertNow(t, expr != nil, "should not be nil")
	goutil.AssertNow(t, expr.Type() == ast.MapLiteral, "wrong node type")
	m := expr.(*ast.MapLiteralNode)
	goutil.AssertNow(t, len(m.Data) == 2, "wrong data length")
}

func TestParseMapLiteralSingleLine(t *testing.T) {
	expr := ParseExpression(`map[string]int { "hi": 1, "bye": 2 }`)
	goutil.AssertNow(t, expr != nil, "should not be nil")
	goutil.AssertNow(t, expr.Type() == ast.MapLiteral, "wrong node type")
	m := expr.(*ast.MapLiteralNode)
	goutil.AssertNow(t, len(m.Data) == 2, "wrong data length")
}
