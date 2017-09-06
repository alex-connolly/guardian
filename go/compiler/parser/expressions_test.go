package parser

import (
	"testing"

	"github.com/end-r/guardian/go/compiler/lexer"

	"github.com/end-r/guardian/go/compiler/ast"

	"github.com/end-r/goutil"
)

func TestParseReferenceSingle(t *testing.T) {
	p := createParser(`hello`)
	goutil.AssertNow(t, p.lexer != nil, "lexer should not be nil")
	goutil.AssertNow(t, len(p.lexer.Tokens) == 1, "wrong token length")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr should not be nil")
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
	goutil.AssertNow(t, expr != nil, "expr should not be nil")
	goutil.AssertNow(t, expr.Type() == ast.Reference, "wrong expr type")
	ref := expr.(ast.ReferenceNode)
	goutil.AssertNow(t, ref.Names != nil, "ref should not be nil")
	goutil.AssertNow(t, len(ref.Names) == 3, "wrong name length")
	goutil.Assert(t, ref.Names[0] == "hello", "wrong name data 0")
	goutil.Assert(t, ref.Names[1] == "aaa", "wrong name data 1")
	goutil.Assert(t, ref.Names[2] == "bb", "wrong name data 2")
}

func TestParseLiteralInteger(t *testing.T) {
	p := createParser(`6`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 1, "wrong token length")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr should not be nil")
	goutil.AssertNow(t, expr.Type() == ast.Literal, "wrong expr type")
	lit := expr.(ast.LiteralNode)
	goutil.AssertNow(t, lit.LiteralType == lexer.TknNumber, "wrong literal type")
}

func TestParseLiteralString(t *testing.T) {
	p := createParser(`"alex"`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 1, "wrong token length")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr should not be nil")
	goutil.AssertNow(t, expr.Type() == ast.Literal, "wrong expr type")
	lit := expr.(ast.LiteralNode)
	goutil.AssertNow(t, lit.LiteralType == lexer.TknString, "wrong literal type")
}

func TestParseMapLiteralEmpty(t *testing.T) {
	p := createParser("map[string]int{}")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expression shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.MapLiteral, "wrong node type")
	n := expr.(ast.MapLiteralNode)
	goutil.AssertNow(t, len(n.Key.Names) == 1, "wrong key name length")
	goutil.Assert(t, n.Key.Names[0] == "string", "wrong key name")
	goutil.AssertNow(t, len(n.Value.Names) == 1, "wrong value name length")
	goutil.Assert(t, n.Value.Names[0] == "int", "wrong value name")
	goutil.Assert(t, len(n.Data) == 0, "should be empty")
}

func TestParseMapLiteralSingle(t *testing.T) {
	p := createParser(`map[string]int{"Hi":3}`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expression shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.MapLiteral, "wrong node type")
	n := expr.(ast.MapLiteralNode)
	goutil.AssertNow(t, len(n.Key.Names) == 1, "wrong key name length")
	goutil.Assert(t, n.Key.Names[0] == "string", "wrong key name")
	goutil.AssertNow(t, len(n.Value.Names) == 1, "wrong value name length")
	goutil.Assert(t, n.Value.Names[0] == "int", "wrong value name")
	goutil.Assert(t, len(n.Data) == 1, "wrong data length")
}

func TestParseMapLiteralMultiple(t *testing.T) {
	p := createParser(`map[string]int{"Hi":3, "Byte":8}`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expression shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.MapLiteral, "wrong node type")
	n := expr.(ast.MapLiteralNode)
	goutil.AssertNow(t, len(n.Key.Names) == 1, "wrong key name length")
	goutil.Assert(t, n.Key.Names[0] == "string", "wrong key name")
	goutil.AssertNow(t, len(n.Value.Names) == 1, "wrong value name length")
	goutil.Assert(t, n.Value.Names[0] == "int", "wrong value name")
	goutil.Assert(t, len(n.Data) == 2, "wrong data length")
}

func TestParseSliceExpressionReferenceLowLiteral(t *testing.T) {
	p := createParser("slice[6:]")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.SliceExpression, "wrong node type")
	s := expr.(ast.SliceExpressionNode)
	goutil.AssertNow(t, s.Expression != nil, "expression shouldn't be nil")
	goutil.AssertNow(t, s.Expression.Type() == ast.Reference, "wrong expression type")
	r := s.Expression.(ast.ReferenceNode)
	goutil.AssertNow(t, len(r.Names) == 1, "wrong expr name length")
	goutil.AssertNow(t, r.Names[0] == "slice", "wrong expr name 0")
	goutil.AssertNow(t, s.Low.Type() == ast.Literal, "wrong low type")
	l := s.Low.(ast.LiteralNode)
	goutil.AssertNow(t, l.Data == "6", "wrong data")
}

func TestParseSliceExpressionReferenceLowReference(t *testing.T) {
	p := createParser("slice[low:]")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.SliceExpression, "wrong node type")
	s := expr.(ast.SliceExpressionNode)
	goutil.AssertNow(t, s.Expression.Type() == ast.Reference, "wrong expression type")
	r := s.Expression.(ast.ReferenceNode)
	goutil.AssertNow(t, len(r.Names) == 1, "wrong expr name length")
	goutil.AssertNow(t, r.Names[0] == "slice", "wrong expr name 0")
	goutil.AssertNow(t, s.Low.Type() == ast.Reference, "wrong low type")
	l := s.Low.(ast.ReferenceNode)
	goutil.AssertNow(t, len(l.Names) == 1, "wrong reference length")
	goutil.AssertNow(t, l.Names[0] == "low", "wrong ref 0")
}

func TestParseSliceExpressionReferenceLowCall(t *testing.T) {
	p := createParser("slice[low():]")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.SliceExpression, "wrong node type")
	s := expr.(ast.SliceExpressionNode)
	goutil.AssertNow(t, s.Expression.Type() == ast.Reference, "wrong expression type")
	r := s.Expression.(ast.ReferenceNode)
	goutil.AssertNow(t, len(r.Names) == 1, "wrong expr name length")
	goutil.AssertNow(t, r.Names[0] == "slice", "wrong expr name 0")
	goutil.AssertNow(t, s.Low.Type() == ast.CallExpression, "wrong low type")
	l := s.Low.(ast.CallExpressionNode)
	goutil.AssertNow(t, l.Arguments == nil, "call arguments should be nil")
	goutil.AssertNow(t, l.Call.Type() == ast.Reference, "wrong call type")
	c := l.Call.(ast.ReferenceNode)
	goutil.AssertNow(t, len(c.Names) == 1, "wrong call name length")
	goutil.AssertNow(t, c.Names[0] == "low", "wrong call name 0")
}

func TestParseSliceExpressionCallLowLiteral(t *testing.T) {
	p := createParser("getSlice()[6:]")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.SliceExpression, "wrong node type")
	s := expr.(ast.SliceExpressionNode)
	goutil.AssertNow(t, s.Expression.Type() == ast.CallExpression, "wrong expression type")
	c := s.Expression.(ast.CallExpressionNode)
	goutil.AssertNow(t, c.Arguments == nil, "arguments should be nil")
	goutil.AssertNow(t, c.Call.Type() == ast.Reference, "wrong call type")
	r := c.Call.(ast.ReferenceNode)
	goutil.AssertNow(t, len(r.Names) == 1, "wrong call name length")
	goutil.AssertNow(t, r.Names[0] == "getSlice", "wrong call name 0")
}

func TestParseSliceExpressionCallLowReference(t *testing.T) {
	p := createParser("getSlice()[low:]")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.SliceExpression, "wrong node type")
	s := expr.(ast.SliceExpressionNode)
	goutil.AssertNow(t, s.Expression.Type() == ast.CallExpression, "wrong expression type")
	c := s.Expression.(ast.CallExpressionNode)
	goutil.AssertNow(t, c.Arguments == nil, "arguments should be nil")
	goutil.AssertNow(t, c.Call.Type() == ast.Reference, "wrong call type")
	r := c.Call.(ast.ReferenceNode)
	goutil.AssertNow(t, len(r.Names) == 1, "wrong call name length")
	goutil.AssertNow(t, r.Names[0] == "getSlice", "wrong call name 0")
}

func TestParseSliceExpressionCallLowCall(t *testing.T) {
	p := createParser("getSlice()[getLow():]")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.SliceExpression, "wrong node type")
	s := expr.(ast.SliceExpressionNode)
	goutil.AssertNow(t, s.Expression.Type() == ast.CallExpression, "wrong expression type")
	c := s.Expression.(ast.CallExpressionNode)
	goutil.AssertNow(t, c.Arguments == nil, "arguments should be nil")
	goutil.AssertNow(t, c.Call.Type() == ast.Reference, "wrong call type")
	r := c.Call.(ast.ReferenceNode)
	goutil.AssertNow(t, len(r.Names) == 1, "wrong call name length")
	goutil.AssertNow(t, r.Names[0] == "getSlice", "wrong call name 0")
}

func TestParseSliceExpressionArrayLiteralowLiteral(t *testing.T) {
	p := createParser(`[string]{"a", "b", "c"}[6:]`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.Assert(t, expr.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionArrayLiteralowReference(t *testing.T) {
	p := createParser(`[string]{"a", "b", "c"}[low:]`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.Assert(t, expr.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionArrayLiteralowCall(t *testing.T) {
	p := createParser(`[string]{"a", "b", "c"}[getLow():]`)
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
	p := createParser(`[string]{"a", "b", "c"}[:6]`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.Assert(t, expr.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionArrayLiteralHighReference(t *testing.T) {
	p := createParser(`[string]{"a", "b", "c"}[:high]`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.Assert(t, expr.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionArrayLiteralHighCall(t *testing.T) {
	p := createParser(`[string]{"a", "b", "c"}[:getHigh()]`)
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
	p := createParser(`[string]{"a", "b", "c"}[1:6]`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.Assert(t, expr.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionArrayLiteralowHighReference(t *testing.T) {
	p := createParser(`[string]{"a", "b", "c"}[low:high]`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.Assert(t, expr.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionArrayLiteralowHighCall(t *testing.T) {
	p := createParser(`[string]{"a", "b", "c"}[getLow():getHigh()]`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.Assert(t, expr.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseCompositeLiteralEmpty(t *testing.T) {
	p := createParser("Dog{}")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.CompositeLiteral, "wrong node type")
	n := expr.(ast.CompositeLiteralNode)
	goutil.AssertNow(t, n.Reference.Type() == ast.Reference, "wrong type type")
	r := n.Reference.(ast.ReferenceNode)
	goutil.AssertNow(t, len(r.Names) == 1, "wrong reference name length")
	goutil.Assert(t, r.Names[0] == "Dog", "wrong reference name")
}

func TestParseCompositeLiteralDeepReferenceEmpty(t *testing.T) {
	p := createParser("animals.Dog{}")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.CompositeLiteral, "wrong node type")
	n := expr.(ast.CompositeLiteralNode)
	goutil.AssertNow(t, n.Reference.Type() == ast.Reference, "wrong type type")
	r := n.Reference.(ast.ReferenceNode)
	goutil.AssertNow(t, len(r.Names) == 2, "wrong reference name length")
	goutil.Assert(t, r.Names[0] == "animals", "wrong reference name 0")
	goutil.Assert(t, r.Names[1] == "Dog", "wrong reference name 1")
}

func TestParseCompositeLiteralInline(t *testing.T) {
	p := createParser(`Dog{name: "Mr Woof"}`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expression shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.CompositeLiteral, "wrong node type")
	n := expr.(ast.CompositeLiteralNode)
	goutil.AssertNow(t, n.Reference != nil, "reference shouldn't be nil")
	goutil.AssertNow(t, n.Reference.Type() == ast.Reference, "wrong reference type")
	r := n.Reference.(ast.ReferenceNode)
	goutil.AssertNow(t, len(r.Names) == 1, "wrong reference name length")
	goutil.Assert(t, r.Names[0] == "Dog", "wrong reference name 0")
	goutil.AssertNow(t, n.Fields != nil, "fields shouldn't be nil")
	goutil.AssertNow(t, len(n.Fields) == 1, "wrong number of fields")
}

func TestParseIndexExpressionReferenceReference(t *testing.T) {
	p := createParser(`array[index]`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.IndexExpression, "wrong expr type")
	indexExpr := expr.(ast.IndexExpressionNode)
	index := indexExpr.Index
	expression := indexExpr.Expression
	goutil.AssertNow(t, expression != nil, "expression shouldn't be nil")
	goutil.AssertNow(t, expression.Type() == ast.Reference, "wrong expression type")
	goutil.AssertNow(t, index.Type() == ast.Reference, "wrong index type")
}

func TestParseIndexExpressionReferenceLiteral(t *testing.T) {
	p := createParser(`array[6]`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.IndexExpression, "wrong expr type")
	indexExpr := expr.(ast.IndexExpressionNode)
	index := indexExpr.Index
	expression := indexExpr.Expression
	goutil.AssertNow(t, expression != nil, "expression shouldn't be nil")
	goutil.AssertNow(t, expression.Type() == ast.Reference, "wrong expression type")
	goutil.AssertNow(t, index.Type() == ast.Literal, "wrong index type")
}

func TestParseIndexExpressionReferenceCall(t *testing.T) {
	p := createParser(`array[getIndex()]`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.IndexExpression, "wrong expr type")
	indexExpr := expr.(ast.IndexExpressionNode)
	index := indexExpr.Index
	expression := indexExpr.Expression
	goutil.AssertNow(t, expression != nil, "expression shouldn't be nil")
	goutil.AssertNow(t, expression.Type() == ast.Reference, "wrong expression type")
	goutil.AssertNow(t, index.Type() == ast.CallExpression, "wrong index type")
}

func TestParseIndexExpressionReferenceIndex(t *testing.T) {
	p := createParser(`array[nested[0]]`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.IndexExpression, "wrong expr type")
	indexExpr := expr.(ast.IndexExpressionNode)
	index := indexExpr.Index
	expression := indexExpr.Expression
	goutil.AssertNow(t, expression != nil, "expression shouldn't be nil")
	goutil.AssertNow(t, expression.Type() == ast.Reference, "wrong expression type")
	goutil.AssertNow(t, index.Type() == ast.IndexExpression, "wrong index type")
}

func TestParseIndexExpressionCallReference(t *testing.T) {
	p := createParser(`getArray()[index]`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.IndexExpression, "wrong expr type")
	indexExpr := expr.(ast.IndexExpressionNode)
	index := indexExpr.Index
	expression := indexExpr.Expression
	goutil.AssertNow(t, expression != nil, "expression shouldn't be nil")
	goutil.AssertNow(t, expression.Type() == ast.CallExpression, "wrong expression type")
	goutil.AssertNow(t, index.Type() == ast.Reference, "wrong index type")
}

func TestParseIndexExpressionCallLiteral(t *testing.T) {
	p := createParser(`getArray()[5]`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.IndexExpression, "wrong expr type")
	indexExpr := expr.(ast.IndexExpressionNode)
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
	indexExpr := expr.(ast.IndexExpressionNode)
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
	indexExpr := expr.(ast.IndexExpressionNode)
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
	indexExpr := expr.(ast.IndexExpressionNode)
	index := indexExpr.Index
	expression := indexExpr.Expression
	goutil.AssertNow(t, expression != nil, "expression shouldn't be nil")
	goutil.AssertNow(t, expression.Type() == ast.IndexExpression, "wrong expression type")
	goutil.AssertNow(t, index.Type() == ast.Reference, "wrong index type")
}

func TestParseIndexExpressionIndexLiteral(t *testing.T) {
	p := createParser(`array[index][5]`)
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr shouldn't be nil")
	goutil.AssertNow(t, expr.Type() == ast.IndexExpression, "wrong expr type")
	indexExpr := expr.(ast.IndexExpressionNode)
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
	indexExpr := expr.(ast.IndexExpressionNode)
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
	indexExpr := expr.(ast.IndexExpressionNode)
	index := indexExpr.Index
	expression := indexExpr.Expression
	goutil.AssertNow(t, expression.Type() == ast.IndexExpression, "wrong expression type")
	goutil.AssertNow(t, index.Type() == ast.IndexExpression, "wrong index type")
}
