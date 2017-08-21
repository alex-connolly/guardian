package expressions

import (
	"testing"

	"github.com/end-r/guardian/go/compiler/ast"
	"github.com/end-r/guardian/go/compiler/parser"
	"github.com/end-r/guardian/go/util"
)

func TestParseCompositeLiteralEmpty(t *testing.T) {
	p := parser.ParseString("Dog{}")
	util.AssertNow(t, p.Expression.Type() == ast.CompositeLiteral, "wrong node type")
	n := p.Expression.(ast.CompositeLiteralNode)
	util.AssertNow(t, n.Reference.Type() == ast.Reference, "wrong type type")
	r := n.Reference.(ast.ReferenceNode)
	util.AssertNow(t, len(r.Names) == 1, "wrong reference name length")
	util.Assert(t, r.Names[0] == "Dog", "wrong reference name")
}

func TestParseCompositeLiteralDeepReferenceEmpty(t *testing.T) {
	p := parser.ParseString("animals.Dog{}")
	util.AssertNow(t, p.Expression.Type() == ast.CompositeLiteral, "wrong node type")
	n := p.Expression.(ast.CompositeLiteralNode)
	util.AssertNow(t, n.Reference.Type() == ast.Reference, "wrong type type")
	r := n.Reference.(ast.ReferenceNode)
	util.AssertNow(t, len(r.Names) == 2, "wrong reference name length")
	util.Assert(t, r.Names[0] == "animals", "wrong reference name 0")
	util.Assert(t, r.Names[1] == "Dog", "wrong reference name 1")
}

func TestParseCompositeLiteralInline(t *testing.T) {
	p := parser.ParseString(`Dog{name: "Mr Woof"}`)
	util.AssertNow(t, p.Expression.Type() == ast.CompositeLiteral, "wrong node type")
	n := p.Expression.(ast.CompositeLiteralNode)
	util.AssertNow(t, n.Reference.Type() == ast.Reference, "wrong type type")
	r := n.Reference.(ast.ReferenceNode)
	util.AssertNow(t, len(r.Names) == 1, "wrong reference name length")
	util.Assert(t, r.Names[0] == "Dog", "wrong reference name 0")
	util.AssertNow(t, len(n.Fields) == 1, "wrong number of fields")
	value, ok := n.Fields["name"]
	util.AssertNow(t, ok, "correct field name not found")
	util.Assert(t, value.Type() == ast.Literal, "wrong value type")
	v := value.(ast.LiteralNode)
	util.Assert(t, v.Data == "Mr Woof", "wrong value data")
}
