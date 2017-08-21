package expressions

import (
	"testing"

	"github.com/end-r/guardian/go/compiler/ast"
	"github.com/end-r/guardian/go/compiler/parser"
	"github.com/end-r/guardian/go/util"
)

func TestParseCallExpressionEmpty(t *testing.T) {
	p := parser.ParseString("hi()")
	util.AssertNow(t, p.Expression.Type() == ast.CallExpression, "wrong node type")
	u := p.Expression.(ast.CallExpressionNode)
	util.AssertNow(t, u.Call.Type() == ast.Reference, "wrong call type")
	c := p.Expression.(ast.ReferenceNode)
	util.AssertNow(t, len(c.Names) == 1, "wrong reference length")
	util.Assert(t, c.Names[0] == "hi", "wrong reference name")
}

func TestParseCallExpressionEmptyReference(t *testing.T) {
	p := parser.ParseString("pkg.hi()")
	util.AssertNow(t, p.Expression.Type() == ast.CallExpression, "wrong node type")
	u := p.Expression.(ast.CallExpressionNode)
	util.AssertNow(t, u.Call.Type() == ast.Reference, "wrong call type")
	c := p.Expression.(ast.ReferenceNode)
	util.AssertNow(t, len(c.Names) == 2, "wrong reference length")
	util.Assert(t, c.Names[0] == "pkg", "wrong reference package name")
	util.Assert(t, c.Names[0] == "hi", "wrong reference name")
}

func TestParseCallExpressionLiterals(t *testing.T) {
	p := parser.ParseString(`hi("a", 6, true)`)
	util.AssertNow(t, p.Expression.Type() == ast.CallExpression, "wrong node type")
	u := p.Expression.(ast.CallExpressionNode)
	util.AssertNow(t, u.Call.Type() == ast.Reference, "wrong call type")
	c := p.Expression.(ast.ReferenceNode)
	util.AssertNow(t, len(c.Names) == 1, "wrong reference length")
	util.Assert(t, c.Names[0] == "hi", "wrong reference name")
}

func TestParseCallExpressionLiteralsReference(t *testing.T) {
	p := parser.ParseString(`pkg.hi("a", 6, true)`)
	util.AssertNow(t, p.Expression.Type() == ast.CallExpression, "wrong node type")
	u := p.Expression.(ast.CallExpressionNode)
	util.AssertNow(t, u.Call.Type() == ast.Reference, "wrong call type")
	c := p.Expression.(ast.ReferenceNode)
	util.AssertNow(t, len(c.Names) == 2, "wrong reference length")
	util.Assert(t, c.Names[0] == "pkg", "wrong reference package name")
	util.Assert(t, c.Names[0] == "hi", "wrong reference name")
}

func TestParseCallExpressionNestedEmpty(t *testing.T) {
	p := parser.ParseString(`hi(a(), b(), c())`)
	util.AssertNow(t, p.Expression.Type() == ast.CallExpression, "wrong node type")
	u := p.Expression.(ast.CallExpressionNode)
	util.AssertNow(t, u.Call.Type() == ast.Reference, "wrong call type")
	c := p.Expression.(ast.ReferenceNode)
	util.AssertNow(t, len(c.Names) == 1, "wrong reference length")
	util.Assert(t, c.Names[0] == "hi", "wrong reference name")
}

func TestParseCallExpressionNestedEmptyReference(t *testing.T) {
	p := parser.ParseString(`pkg.hi(a(), b(), c())`)
	util.AssertNow(t, p.Expression.Type() == ast.CallExpression, "wrong node type")
	u := p.Expression.(ast.CallExpressionNode)
	util.AssertNow(t, u.Call.Type() == ast.Reference, "wrong call type")
	c := p.Expression.(ast.ReferenceNode)
	util.AssertNow(t, len(c.Names) == 2, "wrong reference length")
	util.Assert(t, c.Names[0] == "pkg", "wrong reference package name")
	util.Assert(t, c.Names[0] == "hi", "wrong reference name")
}

func TestParseCallExpressionNestedFullReference(t *testing.T) {
	p := parser.ParseString(`pkg.hi(a("one", "two", "three"), b(four(), five()), c(six))`)
	util.AssertNow(t, p.Expression.Type() == ast.CallExpression, "wrong node type")
	u := p.Expression.(ast.CallExpressionNode)
	util.AssertNow(t, u.Call.Type() == ast.Reference, "wrong call type")
	c := p.Expression.(ast.ReferenceNode)
	util.AssertNow(t, len(c.Names) == 2, "wrong reference length")
	util.Assert(t, c.Names[0] == "pkg", "wrong reference package name")
	util.Assert(t, c.Names[0] == "hi", "wrong reference name")
}
