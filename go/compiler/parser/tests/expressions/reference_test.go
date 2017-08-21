package expressions

import (
	"fmt"
	"testing"

	"github.com/end-r/guardian/go/compiler/ast"
	"github.com/end-r/guardian/go/compiler/parser"
	"github.com/end-r/guardian/go/util"
)

func TestParseReference(t *testing.T) {
	p := parser.ParseString("hello")
	util.AssertNow(t, p.Expression.Type() == ast.Reference, "Wrong node type")
	n := p.Expression.(ast.ReferenceNode)
	util.AssertNow(t, len(n.Names) == 1, "wrong name length")
	util.Assert(t, n.Names[0] == "hello", "wrong reference 0 name")
}

func TestParseImportedReference(t *testing.T) {
	p := parser.ParseString("pkg.hello")
	util.AssertNow(t, p.Expression.Type() == ast.Reference, "Wrong node type")
	n := p.Expression.(ast.ReferenceNode)
	util.AssertNow(t, len(n.Names) == 2, "wrong name length")
	util.Assert(t, n.Names[0] == "pkg", "wrong reference 0 name")
	util.Assert(t, n.Names[1] == "hello", "wrong reference 0 name")
}

func TestParseImportedReferenceArbitraryLength(t *testing.T) {
	p := parser.ParseString("a.b.c.d.e.pkg.hello")
	util.AssertNow(t, p.Expression.Type() == ast.Reference, "Wrong node type")
	n := p.Expression.(ast.ReferenceNode)
	util.AssertNow(t, len(n.Names) == 7, "wrong name length")
	for i, s := range []string{"a", "b", "c", "d", "e", "pkg", "hello"} {
		util.AssertNow(t, n.Names[i] == s, fmt.Sprintf("wrong reference %d name", i))
	}
}
