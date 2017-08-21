package expressions

import (
	"testing"

	"github.com/end-r/guardian/go/compiler/ast"
	"github.com/end-r/guardian/go/compiler/parser"
	"github.com/end-r/guardian/go/util"
)

func TestParseIndexExpressionReferenceLiteral(t *testing.T) {
	p := parser.ParseString("array[5]")
	util.AssertNow(t, p.Expression.Type() == ast.IndexExpression, "wrong node type")
}

func TestParseIndexExpressionReferenceCall(t *testing.T) {
	p := parser.ParseString("array[getIndex()]")
	util.AssertNow(t, p.Expression.Type() == ast.IndexExpression, "wrong node type")
}

func TestParseIndexExpressionReferenceReference(t *testing.T) {
	p := parser.ParseString("array[index]")
	util.AssertNow(t, p.Expression.Type() == ast.IndexExpression, "wrong node type")
	n := p.Expression.(ast.IndexExpressionNode)
	util.AssertNow(t, n.Expression.Type() == ast.Reference, "wrong index type")
}

func TestParseIndexExpressionCallLiteral(t *testing.T) {
	p := parser.ParseString("getArray()[6]")
	util.AssertNow(t, p.Expression.Type() == ast.IndexExpression, "wrong node type")
}

func TestParseIndexExpressionCallCall(t *testing.T) {
	p := parser.ParseString("getArray()[getIndex()]")
	util.AssertNow(t, p.Expression.Type() == ast.IndexExpression, "wrong node type")
}

func TestParseIndexExpressionCallReference(t *testing.T) {
	p := parser.ParseString("getArray()[index]")
	util.AssertNow(t, p.Expression.Type() == ast.IndexExpression, "wrong node type")
}
