package declarations

import (
	"axia/guardian/go/compiler/ast"
	"axia/guardian/go/compiler/parser"
	"axia/guardian/go/util"
	"fmt"
	"testing"
)

func TestParseBasicClassDeclaration(t *testing.T) {
	p := parser.ParseString("class Dog {")
	util.AssertNow(t, p.Scope.Type() == ast.ClassDeclaration, "wrong Scope type")
	n := p.Scope.(ast.ClassDeclarationNode)
	util.Assert(t, n.Identifier == "Dog", "wrong identifier")
	util.Assert(t, !n.IsAbstract, "wrong abstract")
}

func TestParseAbstractClassDeclaration(t *testing.T) {
	p := parser.ParseString("abstract class Dog {")
	util.AssertNow(t, p.Scope.Type() == ast.ClassDeclaration, "wrong Scope type")
	n := p.Scope.(ast.ClassDeclarationNode)
	util.Assert(t, n.Identifier == "Dog", "wrong identifier")
	util.Assert(t, n.IsAbstract, "wrong abstract")
}

func TestParseInheritingClassDeclaration(t *testing.T) {
	p := parser.ParseString("class Dog inherits Animal, Thing {")
	util.AssertNow(t, p.Scope.Type() == ast.ClassDeclaration, "wrong Scope type")
	n := p.Scope.(ast.ClassDeclarationNode)
	util.Assert(t, n.Identifier == "Dog", "wrong identifier")
	util.Assert(t, !n.IsAbstract, "wrong abstract")
	util.AssertNow(t, len(n.Supers) == 2, "wrong number of super classes")
	util.AssertNow(t, len(n.Interfaces) == 0, "wrong number of interfaces")
}

func TestParseInterfacingClassDeclaration(t *testing.T) {
	p := parser.ParseString("class Dog is Noisy, Walkable {")
	util.AssertNow(t, p.Scope.Type() == ast.ClassDeclaration, "wrong Scope type")
	n := p.Scope.(ast.ClassDeclarationNode)
	util.Assert(t, n.Identifier == "Dog", "wrong identifier")
	util.Assert(t, !n.IsAbstract, "wrong abstract")
	util.AssertNow(t, len(n.Supers) == 0, "wrong number of super classes")
	util.AssertNow(t, len(n.Interfaces) == 2, fmt.Sprintf("wrong number of interfaces (%d)", len(n.Interfaces)))
}

func TestParseMultipleInterfacesFirstClassDeclaration(t *testing.T) {
	p := parser.ParseString("class Dog is Noisy, Walkable inherits Animal, Thing {")
	util.AssertNow(t, p.Scope.Type() == ast.ClassDeclaration, "wrong Scope type")
	n := p.Scope.(ast.ClassDeclarationNode)
	util.Assert(t, n.Identifier == "Dog", "wrong identifier")
	util.Assert(t, !n.IsAbstract, "wrong abstract")
	util.AssertNow(t, len(n.Supers) == 2, "wrong number of super classes")
	util.AssertNow(t, len(n.Interfaces) == 2, fmt.Sprintf("wrong number of interfaces (%d)", len(n.Interfaces)))
}

func TestParseSingleSuperFirstClassDeclaration(t *testing.T) {
	p := parser.ParseString("class Dog inherits Animal is Noisy, Walkable {")
	util.AssertNow(t, p.Scope.Type() == ast.ClassDeclaration, "wrong Scope type")
	n := p.Scope.(ast.ClassDeclarationNode)
	util.Assert(t, n.Identifier == "Dog", "wrong identifier")
	util.Assert(t, !n.IsAbstract, "wrong abstract")
	util.AssertNow(t, len(n.Supers) == 1, "wrong number of super classes")
	util.AssertNow(t, len(n.Interfaces) == 2, fmt.Sprintf("wrong number of interfaces (%d)", len(n.Interfaces)))
}

func TestParseSingleInterfaceFirstClassDeclaration(t *testing.T) {
	p := parser.ParseString("class Dog is Noisy inherits Animal, Thing {")
	util.AssertNow(t, p.Scope.Type() == ast.ClassDeclaration, "wrong Scope type")
	n := p.Scope.(ast.ClassDeclarationNode)
	util.Assert(t, n.Identifier == "Dog", "wrong identifier")
	util.Assert(t, !n.IsAbstract, "wrong abstract")
	util.AssertNow(t, len(n.Supers) == 2, "wrong number of super classes")
	util.AssertNow(t, len(n.Interfaces) == 1, fmt.Sprintf("wrong number of interfaces (%d)", len(n.Interfaces)))
}

func TestParseFullClassDeclaration(t *testing.T) {
	p := parser.ParseString("abstract class Dog inherits Animal, Thing is Noisy, Walkable {")
	util.AssertNow(t, p.Scope.Type() == ast.ClassDeclaration, "wrong Scope type")
	n := p.Scope.(ast.ClassDeclarationNode)
	util.Assert(t, n.Identifier == "Dog", "wrong identifier")
	util.Assert(t, n.IsAbstract, "wrong abstract")
	util.AssertNow(t, len(n.Supers) == 2, "wrong number of super classes")
	util.AssertNow(t, len(n.Interfaces) == 2, fmt.Sprintf("wrong number of interfaces (%d)", len(n.Interfaces)))
}

func TestParseMultilineClassDeclaration(t *testing.T) {
	p := parser.ParseString(`abstract class Dog
		inherits Animal, Thing
		is Noisy, Walkable {`)
	util.AssertNow(t, p.Scope.Type() == ast.ClassDeclaration, "wrong Scope type")
	n := p.Scope.(ast.ClassDeclarationNode)
	util.Assert(t, n.Identifier == "Dog", "wrong identifier")
	util.Assert(t, n.IsAbstract, "wrong abstract")
	util.AssertNow(t, len(n.Supers) == 2, "wrong number of super classes")
	util.AssertNow(t, len(n.Interfaces) == 2, fmt.Sprintf("wrong number of interfaces (%d)", len(n.Interfaces)))
}
