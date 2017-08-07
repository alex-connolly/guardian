package expressions

import (
	"axia/guardian/go/compiler/ast"
	"axia/guardian/go/compiler/parser"
	"axia/guardian/go/util"
	"testing"
)

func TestParseLiteralString(t *testing.T) {
	p := parser.ParseString(`"hello"`)
	util.AssertNow(t, p.Expression.Type() == ast.Literal, "wrong node type")
}

func TestParseLiteralInt(t *testing.T) {
	p := parser.ParseString(`6`)
	util.AssertNow(t, p.Expression.Type() == ast.Literal, "wrong node type")
}

func TestParseLiteralFloat(t *testing.T) {
	p := parser.ParseString(`6.5`)
	util.AssertNow(t, p.Expression.Type() == ast.Literal, "wrong node type")
}

func TestParseLiteralCharacter(t *testing.T) {
	p := parser.ParseString(`'h'`)
	util.AssertNow(t, p.Expression.Type() == ast.Literal, "wrong node type")
}
