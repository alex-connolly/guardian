package expressions

import (
	"axia/guardian/go/compiler/ast"
	"axia/guardian/go/compiler/parser"
	"axia/guardian/go/util"
	"testing"
)

func TestParseArrayLiteralEmpty(t *testing.T) {
	p := parser.ParseString("[int]{}")
	util.AssertNow(t, p.Expression.Type() == ast.ArrayLiteral, "wrong node type")
	n := p.Expression.(ast.ArrayLiteralNode)
	util.AssertNow(t, len(n.Key.Names) == 1, "name length is one")
	util.AssertNow(t, len(n.Data) == 0, "wrong data length")
}

func TestParseArrayLiteralSingle(t *testing.T) {
	p := parser.ParseString("[int]{3}")
	util.Assert(t, p.Expression.Type() == ast.ArrayLiteral, "wrong node type")
	n := p.Expression.(ast.ArrayLiteralNode)
	util.AssertNow(t, len(n.Key.Names) == 1, "name length is one")
	util.AssertNow(t, len(n.Data) == 1, "wrong data length")
}

func TestParseArrayLiteralMultiple(t *testing.T) {
	p := parser.ParseString("[int]{3, 4}")
	util.Assert(t, p.Expression.Type() == ast.ArrayLiteral, "wrong node type")
	n := p.Expression.(ast.ArrayLiteralNode)
	util.AssertNow(t, len(n.Key.Names) == 1, "name length is one")
	util.AssertNow(t, len(n.Data) == 2, "wrong data length")
}
