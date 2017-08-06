package declarations

import (
	"axia/guardian/go/compiler/ast"
	"axia/guardian/go/util"
	"testing"
)

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

func TestParseBasicParameterisedFuncDeclaration(t *testing.T) {
	p := ParseString("makeNoise(sound string){")
	util.AssertNow(t, p.scope.Type() == ast.FuncDeclaration, "wrong scope type")
	n := p.scope.(ast.FuncDeclarationNode)
	util.AssertNow(t, n.Identifier == "makeNoise", "wrong identifier")
	util.AssertNow(t, !n.IsAbstract, "wrong abstract")
	util.AssertNow(t, len(n.Parameters) == 1, "wrong parameter length")
}
