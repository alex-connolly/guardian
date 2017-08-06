package declarations

import (
	"axia/guardian/go/compiler/ast"
	"axia/guardian/go/util"
	"testing"
)

func TestParseContractDeclaration(t *testing.T) {
	p := ParseString("contract Dog {")
	util.Assert(t, p.scope.Type() == ast.ContractDeclaration, "wrong scope type")
	n := p.scope.(ast.ContractDeclarationNode)
	util.Assert(t, n.Identifier == "Dog", "wrong identifier")
	util.Assert(t, !n.IsAbstract, "wrong abstract")
}

func TestParseAbstractContractDeclaration(t *testing.T) {
	p := ParseString("abstract contract Dog {")
	util.Assert(t, p.scope.Type() == ast.ContractDeclaration, "wrong scope type")
	n := p.scope.(ast.ContractDeclarationNode)
	util.Assert(t, n.Identifier == "Dog", "wrong identifier")
	util.Assert(t, n.IsAbstract, "wrong abstract")
}

func TestParseContractDeclarationSingleSuper(t *testing.T) {
	p := ParseString("abstract contract Dog inherits Animal {")
	util.Assert(t, p.scope.Type() == ast.ContractDeclaration, "wrong scope type")
	n := p.scope.(ast.ContractDeclarationNode)
	util.Assert(t, n.Identifier == "Dog", "wrong identifier")
	util.Assert(t, n.IsAbstract, "wrong abstract")
	util.AssertNow(t, len(n.Supers) == 1, "wrong super length")
}

func TestParseContractDeclarationMultipleSupers(t *testing.T) {
	p := ParseString("abstract contract Dog inherits Animal, Quadriped {")
	util.Assert(t, p.scope.Type() == ast.ContractDeclaration, "wrong scope type")
	n := p.scope.(ast.ContractDeclarationNode)
	util.Assert(t, n.Identifier == "Dog", "wrong identifier")
	util.Assert(t, n.IsAbstract, "wrong abstract")
	util.AssertNow(t, len(n.Supers) == 2, "wrong super length")
}

func TestParseInvalidContractDeclaration(t *testing.T) {
	p := ParseString("contract {")
	util.Assert(t, p.scope.Type() == ast.ContractDeclaration, "wrong scope type")
	util.Assert(t, len(p.errs) == 1, "should throw one error")
	p = ParseString("contract Hello contract {}")
	util.Assert(t, p.scope.Type() == ast.ContractDeclaration, "wrong scope type")
	util.Assert(t, len(p.errs) == 2, "should throw 2 errors")
}
