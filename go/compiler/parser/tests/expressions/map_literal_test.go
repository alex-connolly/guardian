package expressions

import (
	"axia/guardian/go/compiler/ast"
	"axia/guardian/go/compiler/parser"
	"axia/guardian/go/util"
	"testing"
)

func TestParseMapLiteralEmpty(t *testing.T) {
	p := parser.ParseString("map[string]int{}")
	util.AssertNow(t, p.Expression.Type() == ast.MapLiteral, "wrong node type")
	n := p.Expression.(ast.MapLiteralNode)
	util.AssertNow(t, len(n.Key.Names) == 1, "wrong key name length")
	util.Assert(t, n.Key.Names[0] == "string", "wrong key name")
	util.AssertNow(t, len(n.Value.Names) == 1, "wrong value name length")
	util.Assert(t, n.Value.Names[0] == "int", "wrong value name")
	util.Assert(t, len(n.Data) == 0, "should be empty")
}

func TestParseMapLiteralSingle(t *testing.T) {
	p := parser.ParseString(`map[string]int{"Hi":3}`)
	util.AssertNow(t, p.Expression.Type() == ast.MapLiteral, "wrong node type")
	n := p.Expression.(ast.MapLiteralNode)
	util.AssertNow(t, len(n.Key.Names) == 1, "wrong key name length")
	util.Assert(t, n.Key.Names[0] == "string", "wrong key name")
	util.AssertNow(t, len(n.Value.Names) == 1, "wrong value name length")
	util.Assert(t, n.Value.Names[0] == "int", "wrong value name")
	util.Assert(t, len(n.Data) == 1, "wrong data length")
}

func TestParseMapLiteralMultiple(t *testing.T) {
	p := parser.ParseString(`map[string]int{"Hi":3, "Byte":8}`)
	util.AssertNow(t, p.Expression.Type() == ast.MapLiteral, "wrong node type")
	n := p.Expression.(ast.MapLiteralNode)
	util.AssertNow(t, len(n.Key.Names) == 1, "wrong key name length")
	util.Assert(t, n.Key.Names[0] == "string", "wrong key name")
	util.AssertNow(t, len(n.Value.Names) == 1, "wrong value name length")
	util.Assert(t, n.Value.Names[0] == "int", "wrong value name")
	util.Assert(t, len(n.Data) == 2, "wrong data length")
}
