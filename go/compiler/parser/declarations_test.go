package parser

import (
	"axia/guardian/go/compiler/ast"
	"axia/guardian/go/util"
	"fmt"
	"testing"
)

func TestParseTypeDeclaration(t *testing.T) {
	ParseString("type Bowl int")
	//util.Assert(t, p.scope.Type() == , "not a type declaration")
}

func TestParseFuncDeclaration(t *testing.T) {

	p := ParseString("main(){")
	util.AssertNow(t, p.scope.Type() == ast.FuncDeclaration, "wrong scope type")
	n := p.scope.(ast.FuncDeclarationNode)
	util.AssertNow(t, n.Identifier == "main", "wrong identifier")
	util.AssertNow(t, !n.IsAbstract, "wrong abstract")
	util.AssertNow(t, n.Parameters == nil, "parameters should be nil")
	util.AssertNow(t, n.Results == nil, "results should be nil")

}

func TestParseBasicAbstractFuncDeclaration(t *testing.T) {
	p := ParseString("abstract makeNoise(){")
	util.AssertNow(t, p.scope.Type() == ast.FuncDeclaration, "wrong scope type")
	n := p.scope.(ast.FuncDeclarationNode)
	util.AssertNow(t, n.Identifier == "makeNoise", "wrong identifier")
	util.AssertNow(t, n.IsAbstract, "wrong abstract")
}

func TestParseBasicParameterisedFunctDeclaration(t *testing.T) {
	p := ParseString("makeNoise(sound string){")
	util.AssertNow(t, p.scope.Type() == ast.FuncDeclaration, "wrong scope type")
	n := p.scope.(ast.FuncDeclarationNode)
	util.AssertNow(t, n.Identifier == "makeNoise", "wrong identifier")
	util.AssertNow(t, !n.IsAbstract, "wrong abstract")
	util.AssertNow(t, len(n.Parameters) == 1, "wrong parameter length")
}

func TestParseBasicClassDeclaration(t *testing.T) {
	p := ParseString("class Dog {")
	util.AssertNow(t, p.scope.Type() == ast.ClassDeclaration, "wrong scope type")
	n := p.scope.(ast.ClassDeclarationNode)
	util.Assert(t, n.Identifier == "Dog", "wrong identifier")
	util.Assert(t, !n.IsAbstract, "wrong abstract")
}

func TestParseAbstractClassDeclaration(t *testing.T) {
	p := ParseString("abstract class Dog {")
	util.AssertNow(t, p.scope.Type() == ast.ClassDeclaration, "wrong scope type")
	n := p.scope.(ast.ClassDeclarationNode)
	util.Assert(t, n.Identifier == "Dog", "wrong identifier")
	util.Assert(t, n.IsAbstract, "wrong abstract")
}

func TestParseInheritingClassDeclaration(t *testing.T) {
	p := ParseString("class Dog inherits Animal, Thing {")
	util.AssertNow(t, p.scope.Type() == ast.ClassDeclaration, "wrong scope type")
	n := p.scope.(ast.ClassDeclarationNode)
	util.Assert(t, n.Identifier == "Dog", "wrong identifier")
	util.Assert(t, !n.IsAbstract, "wrong abstract")
	util.AssertNow(t, len(n.Supers) == 2, "wrong number of super classes")
	util.AssertNow(t, len(n.Interfaces) == 0, "wrong number of interfaces")
}

func TestParseInterfacingClassDeclaration(t *testing.T) {
	p := ParseString("class Dog is Noisy, Walkable {")
	util.AssertNow(t, p.scope.Type() == ast.ClassDeclaration, "wrong scope type")
	n := p.scope.(ast.ClassDeclarationNode)
	util.Assert(t, n.Identifier == "Dog", "wrong identifier")
	util.Assert(t, !n.IsAbstract, "wrong abstract")
	util.AssertNow(t, len(n.Supers) == 0, "wrong number of super classes")
	util.AssertNow(t, len(n.Interfaces) == 2, fmt.Sprintf("wrong number of interfaces (%d)", len(n.Interfaces)))
}

func TestParseMultipleInterfacesFirstClassDeclaration(t *testing.T) {
	p := ParseString("class Dog is Noisy, Walkable inherits Animal, Thing {")
	util.AssertNow(t, p.scope.Type() == ast.ClassDeclaration, "wrong scope type")
	n := p.scope.(ast.ClassDeclarationNode)
	util.Assert(t, n.Identifier == "Dog", "wrong identifier")
	util.Assert(t, !n.IsAbstract, "wrong abstract")
	util.AssertNow(t, len(n.Supers) == 2, "wrong number of super classes")
	util.AssertNow(t, len(n.Interfaces) == 2, fmt.Sprintf("wrong number of interfaces (%d)", len(n.Interfaces)))
}

func TestParseSingleSuperFirstClassDeclaration(t *testing.T) {
	p := ParseString("class Dog inherits Animal is Noisy, Walkable {")
	util.AssertNow(t, p.scope.Type() == ast.ClassDeclaration, "wrong scope type")
	n := p.scope.(ast.ClassDeclarationNode)
	util.Assert(t, n.Identifier == "Dog", "wrong identifier")
	util.Assert(t, !n.IsAbstract, "wrong abstract")
	util.AssertNow(t, len(n.Supers) == 1, "wrong number of super classes")
	util.AssertNow(t, len(n.Interfaces) == 2, fmt.Sprintf("wrong number of interfaces (%d)", len(n.Interfaces)))
}

func TestParseSingleInterfaceFirstClassDeclaration(t *testing.T) {
	p := ParseString("class Dog is Noisy inherits Animal, Thing {")
	util.AssertNow(t, p.scope.Type() == ast.ClassDeclaration, "wrong scope type")
	n := p.scope.(ast.ClassDeclarationNode)
	util.Assert(t, n.Identifier == "Dog", "wrong identifier")
	util.Assert(t, !n.IsAbstract, "wrong abstract")
	util.AssertNow(t, len(n.Supers) == 2, "wrong number of super classes")
	util.AssertNow(t, len(n.Interfaces) == 1, fmt.Sprintf("wrong number of interfaces (%d)", len(n.Interfaces)))
}

func TestParseFullClassDeclaration(t *testing.T) {
	p := ParseString("abstract class Dog inherits Animal, Thing is Noisy, Walkable {")
	util.AssertNow(t, p.scope.Type() == ast.ClassDeclaration, "wrong scope type")
	n := p.scope.(ast.ClassDeclarationNode)
	util.Assert(t, n.Identifier == "Dog", "wrong identifier")
	util.Assert(t, n.IsAbstract, "wrong abstract")
	util.AssertNow(t, len(n.Supers) == 2, "wrong number of super classes")
	util.AssertNow(t, len(n.Interfaces) == 2, fmt.Sprintf("wrong number of interfaces (%d)", len(n.Interfaces)))
}

func TestParseMultilineClassDeclaration(t *testing.T) {
	p := ParseString(`abstract class Dog
		inherits Animal, Thing
		is Noisy, Walkable {`)
	util.AssertNow(t, p.scope.Type() == ast.ClassDeclaration, "wrong scope type")
	n := p.scope.(ast.ClassDeclarationNode)
	util.Assert(t, n.Identifier == "Dog", "wrong identifier")
	util.Assert(t, n.IsAbstract, "wrong abstract")
	util.AssertNow(t, len(n.Supers) == 2, "wrong number of super classes")
	util.AssertNow(t, len(n.Interfaces) == 2, fmt.Sprintf("wrong number of interfaces (%d)", len(n.Interfaces)))
}

func TestParseInterfaceDeclaration(t *testing.T) {
	p := ParseString("interface Dog {")
	util.AssertNow(t, p.scope.Type() == ast.InterfaceDeclaration, "wrong scope type")
	n := p.scope.(ast.InterfaceDeclarationNode)
	util.Assert(t, n.Identifier == "Dog", "wrong identifier")
	util.Assert(t, n.IsAbstract == false, "wrong abstract")
}

func TestParseAbstractInterfaceDeclaration(t *testing.T) {
	p := ParseString("abstract interface Dog {")
	util.AssertNow(t, p.scope.Type() == ast.InterfaceDeclaration, "wrong scope type")
	n := p.scope.(ast.InterfaceDeclarationNode)
	util.Assert(t, n.Identifier == "Dog", "wrong identifier")
	util.Assert(t, !n.IsAbstract == false, "wrong abstract")
}

func TestParseInterfaceDeclarationSingleSuper(t *testing.T) {
	p := ParseString("interface Dog inherits Animal {")
	util.AssertNow(t, p.scope.Type() == ast.InterfaceDeclaration, "wrong scope type")
	n := p.scope.(ast.InterfaceDeclarationNode)
	util.Assert(t, n.Identifier == "Dog", "wrong identifier")
	util.Assert(t, !n.IsAbstract, "wrong abstract")
	util.AssertNow(t, len(n.Supers) == 1, "wrong super length")
}

func TestParseInterfaceDeclarationMultipleSupers(t *testing.T) {
	p := ParseString("interface Dog inherits Animal, Quadriped {")
	util.AssertNow(t, p.scope.Type() == ast.InterfaceDeclaration, "wrong scope type")
	n := p.scope.(ast.InterfaceDeclarationNode)
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
