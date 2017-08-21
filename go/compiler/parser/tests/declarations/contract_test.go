package declarations

import (
	"testing"

	"github.com/end-r/guardian/go/compiler/ast"
	"github.com/end-r/guardian/go/compiler/parser"
	"github.com/end-r/guardian/go/util"
)

func TestParseContractDeclaration(t *testing.T) {
	p := parser.ParseString("contract Dog {")
	util.Assert(t, p.Scope.Type() == ast.ContractDeclaration, "wrong Scope type")
	n := p.Scope.(ast.ContractDeclarationNode)
	util.Assert(t, n.Identifier == "Dog", "wrong identifier")
	util.Assert(t, !n.IsAbstract, "wrong abstract")
}

func TestParseAbstractContractDeclaration(t *testing.T) {
	p := parser.ParseString("abstract contract Dog {")
	util.Assert(t, p.Scope.Type() == ast.ContractDeclaration, "wrong Scope type")
	n := p.Scope.(ast.ContractDeclarationNode)
	util.Assert(t, n.Identifier == "Dog", "wrong identifier")
	util.Assert(t, n.IsAbstract, "wrong abstract")
}

func TestParseContractDeclarationSingleSuper(t *testing.T) {
	p := parser.ParseString("abstract contract Dog inherits Animal {")
	util.Assert(t, p.Scope.Type() == ast.ContractDeclaration, "wrong Scope type")
	n := p.Scope.(ast.ContractDeclarationNode)
	util.Assert(t, n.Identifier == "Dog", "wrong identifier")
	util.Assert(t, n.IsAbstract, "wrong abstract")
	util.AssertNow(t, len(n.Supers) == 1, "wrong super length")
}

func TestParseContractDeclarationMultipleSupers(t *testing.T) {
	p := parser.ParseString("abstract contract Dog inherits Animal, Quadriped {")
	util.Assert(t, p.Scope.Type() == ast.ContractDeclaration, "wrong Scope type")
	n := p.Scope.(ast.ContractDeclarationNode)
	util.Assert(t, n.Identifier == "Dog", "wrong identifier")
	util.Assert(t, n.IsAbstract, "wrong abstract")
	util.AssertNow(t, len(n.Supers) == 2, "wrong super length")
}

func TestParseInvalidContractDeclaration(t *testing.T) {
	p := parser.ParseString("contract {")
	util.Assert(t, p.Scope.Type() == ast.ContractDeclaration, "wrong Scope type")
	util.Assert(t, len(p.Errs) == 1, "should throw one error")
	p = parser.ParseString("contract Hello contract {}")
	util.Assert(t, p.Scope.Type() == ast.ContractDeclaration, "wrong Scope type")
	util.Assert(t, len(p.Errs) == 2, "should throw 2 errors")
}
