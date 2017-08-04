package parser

import (
	"axia/guardian/go/compiler/ast"
	"axia/guardian/go/util"
	"testing"
)

func TestParseTypeDeclaration(t *testing.T) {
	p := ParseString("type Bowl int")
	util.Assert(t, isTypeDeclaration(p), "not a type declaration")
}

func TestParseFuncDeclaration(t *testing.T) {
	p := ParseString("main(){}")
	util.AssertNow(t, p.scope.Type() == ast.FuncDeclaration, "wrong scope type")
}

func TestParseClassDeclaration(t *testing.T) {
	p := ParseString("class Dog {}")
	util.AssertNow(t, p.scope.Type() == ast.ClassDeclaration, "wrong scope type")
}

func TestParseInterfaceDeclaration(t *testing.T) {
	p := ParseString("interface Dog {}")
	util.AssertNow(t, p.scope.Type() == ast.InterfaceDeclaration, "wrong scope type")
	n := p.scope.(ast.InterfaceDeclarationNode)
	util.Assert(t, n.Identifier == "Dog", "wrong identifier")
	util.Assert(t, n.IsAbstract == false, "wrong abstract")
}

func TestParseFullInterfaceDeclaration(t *testing.T) {
	p := ParseString(`interface Dog {
			walk(int position) bool
		}`)
	util.Assert(t, p != nil, "parser is nil")

}

func TestParseContractDeclaration(t *testing.T) {
	p := ParseString("contract Dog {}")
	util.Assert(t, p.scope.Type() == ast.ContractDeclaration, "wrong scope type")
	p = ParseString("contract Dog inherits Animal {}")
	util.Assert(t, p.scope.Type() == ast.ContractDeclaration, "wrong scope type")
	p = ParseString("contract Dog<T inherits Cat> inherits Animal<T> {}")
	util.Assert(t, p.scope.Type() == ast.ContractDeclaration, "wrong scope type")
}

func TestParseInvalidContractDeclaration(t *testing.T) {
	p := ParseString("contract {}")
	util.Assert(t, p.scope.Type() == ast.ContractDeclaration, "wrong scope type")
	util.Assert(t, len(p.errs) == 1, "should throw one error")
	p = ParseString("contract Hello contract {}")
	util.Assert(t, p.scope.Type() == ast.ContractDeclaration, "wrong scope type")
	util.Assert(t, len(p.errs) == 2, "should throw 2 errors")
}
