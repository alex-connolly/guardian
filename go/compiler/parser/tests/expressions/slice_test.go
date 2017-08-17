package expressions

import (
	"axia/guardian/go/compiler/ast"
	"axia/guardian/go/compiler/parser"
	"axia/guardian/go/util"
	"testing"
)

func TestParseSliceExpressionReferenceLowLiteral(t *testing.T) {
	p := parser.ParseString("slice[6:]")
	util.Assert(t, p.Expression.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionReferenceLowReference(t *testing.T) {
	p := parser.ParseString("slice[low:]")
	util.Assert(t, p.Expression.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionReferenceLowCall(t *testing.T) {
	p := parser.ParseString("slice[low():]")
	util.Assert(t, p.Expression.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionCallLowLiteral(t *testing.T) {
	p := parser.ParseString("getSlice()[6:]")
	util.Assert(t, p.Expression.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionCallLowReference(t *testing.T) {
	p := parser.ParseString("getSlice()[low:]")
	util.Assert(t, p.Expression.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionCallLowCall(t *testing.T) {
	p := parser.ParseString("getSlice()[getLow():]")
	util.Assert(t, p.Expression.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionLiteralLowLiteral(t *testing.T) {
	p := parser.ParseString(`"hello"[6:]`)
	util.Assert(t, p.Expression.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionLiteralLowReference(t *testing.T) {
	p := parser.ParseString(`"hello"[low:]`)
	util.Assert(t, p.Expression.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionLiteralLowCall(t *testing.T) {
	p := parser.ParseString(`"hello"[getLow():]`)
	util.Assert(t, p.Expression.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionArrayLiteralLowLiteral(t *testing.T) {
	p := parser.ParseString(`[string]{"a", "b", "c"}[6:]`)
	util.Assert(t, p.Expression.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionArrayLiteralLowReference(t *testing.T) {
	p := parser.ParseString(`[string]{"a", "b", "c"}[low:]`)
	util.Assert(t, p.Expression.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionArrayLiteralLowCall(t *testing.T) {
	p := parser.ParseString(`[string]{"a", "b", "c"}[getLow():]`)
	util.Assert(t, p.Expression.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionReferenceHighLiteral(t *testing.T) {
	p := parser.ParseString("slice[:6]")
	util.Assert(t, p.Expression.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionReferenceHighReference(t *testing.T) {
	p := parser.ParseString("slice[:high]")
	util.Assert(t, p.Expression.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionReferenceHighCall(t *testing.T) {
	p := parser.ParseString("slice[:high()]")
	util.Assert(t, p.Expression.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionCallHighLiteral(t *testing.T) {
	p := parser.ParseString("getSlice()[:6]")
	util.Assert(t, p.Expression.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionCallHighReference(t *testing.T) {
	p := parser.ParseString("getSlice()[:high]")
	util.Assert(t, p.Expression.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionCallHighCall(t *testing.T) {
	p := parser.ParseString("getSlice()[:getHigh()]")
	util.Assert(t, p.Expression.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionLiteralHighLiteral(t *testing.T) {
	p := parser.ParseString(`"hello"[:6]`)
	util.Assert(t, p.Expression.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionLiteralHighReference(t *testing.T) {
	p := parser.ParseString(`"hello"[:high]`)
	util.Assert(t, p.Expression.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionLiteralHighCall(t *testing.T) {
	p := parser.ParseString(`"hello"[:getHigh()]`)
	util.Assert(t, p.Expression.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionArrayLiteralHighLiteral(t *testing.T) {
	p := parser.ParseString(`[string]{"a", "b", "c"}[:6]`)
	util.Assert(t, p.Expression.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionArrayLiteralHighReference(t *testing.T) {
	p := parser.ParseString(`[string]{"a", "b", "c"}[:high]`)
	util.Assert(t, p.Expression.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionArrayLiteralHighCall(t *testing.T) {
	p := parser.ParseString(`[string]{"a", "b", "c"}[:getHigh()]`)
	util.Assert(t, p.Expression.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionReferenceLowHighLiteral(t *testing.T) {
	p := parser.ParseString("slice[2:6]")
	util.Assert(t, p.Expression.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionReferenceLowHighReference(t *testing.T) {
	p := parser.ParseString("slice[low:high]")
	util.Assert(t, p.Expression.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionReferenceLowHighCall(t *testing.T) {
	p := parser.ParseString("slice[low():high()]")
	util.Assert(t, p.Expression.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionCallLowHighLiteral(t *testing.T) {
	p := parser.ParseString("getSlice()[3:6]")
	util.Assert(t, p.Expression.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionCallLowHighReference(t *testing.T) {
	p := parser.ParseString("getSlice()[low:high]")
	util.Assert(t, p.Expression.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionCallLowHighCall(t *testing.T) {
	p := parser.ParseString("getSlice()[getLow():getHigh()]")
	util.Assert(t, p.Expression.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionLiteralLowHighLiteral(t *testing.T) {
	p := parser.ParseString(`"hello"[2:6]`)
	util.Assert(t, p.Expression.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionLiteralLowHighReference(t *testing.T) {
	p := parser.ParseString(`"hello"[low:high]`)
	util.Assert(t, p.Expression.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionLiteralLowHighCall(t *testing.T) {
	p := parser.ParseString(`"hello"[getLow():getHigh()]`)
	util.Assert(t, p.Expression.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionArrayLiteralLowHighLiteral(t *testing.T) {
	p := parser.ParseString(`[string]{"a", "b", "c"}[1:6]`)
	util.Assert(t, p.Expression.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionArrayLiteralLowHighReference(t *testing.T) {
	p := parser.ParseString(`[string]{"a", "b", "c"}[low:high]`)
	util.Assert(t, p.Expression.Type() == ast.SliceExpression, "wrong node type")
}

func TestParseSliceExpressionArrayLiteralLowHighCall(t *testing.T) {
	p := parser.ParseString(`[string]{"a", "b", "c"}[getLow():getHigh()]`)
	util.Assert(t, p.Expression.Type() == ast.SliceExpression, "wrong node type")
}
