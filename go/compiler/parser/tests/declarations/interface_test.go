package declarations

import (
	"axia/guardian/go/compiler/ast"
	"axia/guardian/go/util"
	"testing"
)

func TestParseInterfaceDeclaration(t *testing.T) {
	p := ParseString("interface Dog {")
	util.AssertNow(t, p.Scope.Type() == ast.InterfaceDeclaration, "wrong Scope type")
	n := p.Scope.(ast.InterfaceDeclarationNode)
	util.Assert(t, n.Identifier == "Dog", "wrong identifier")
	util.Assert(t, n.IsAbstract == false, "wrong abstract")
}

func TestParseAbstractInterfaceDeclaration(t *testing.T) {
	p := ParseString("abstract interface Dog {")
	util.AssertNow(t, p.Scope.Type() == ast.InterfaceDeclaration, "wrong Scope type")
	n := p.Scope.(ast.InterfaceDeclarationNode)
	util.Assert(t, n.Identifier == "Dog", "wrong identifier")
	util.Assert(t, !n.IsAbstract == false, "wrong abstract")
}

func TestParseInterfaceDeclarationSingleSuper(t *testing.T) {
	p := ParseString("interface Dog inherits Animal {")
	util.AssertNow(t, p.Scope.Type() == ast.InterfaceDeclaration, "wrong Scope type")
	n := p.Scope.(ast.InterfaceDeclarationNode)
	util.Assert(t, n.Identifier == "Dog", "wrong identifier")
	util.Assert(t, !n.IsAbstract, "wrong abstract")
	util.AssertNow(t, len(n.Supers) == 1, "wrong super length")
}

func TestParseInterfaceDeclarationMultipleSupers(t *testing.T) {
	p := ParseString("interface Dog inherits Animal, Quadriped {")
	util.AssertNow(t, p.Scope.Type() == ast.InterfaceDeclaration, "wrong Scope type")
	n := p.Scope.(ast.InterfaceDeclarationNode)
	util.Assert(t, n.Identifier == "Dog", "wrong identifier")
	util.Assert(t, !n.IsAbstract, "wrong abstract")
	util.AssertNow(t, len(n.Supers) == 2, "wrong super length")
}

func TestParseFullInterfaceDeclaration(t *testing.T) {
	p := ParseString(`interface Dog {
			walk(int position) bool
		}`)
	util.Assert(t, p != nil, "parser is nil")
}
