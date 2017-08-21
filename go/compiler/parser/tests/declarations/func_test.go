package declarations

import (
	"testing"

	"github.com/end-r/guardian/go/compiler/ast"
	"github.com/end-r/guardian/go/compiler/parser"
	"github.com/end-r/guardian/go/util"
)

func TestParseFuncDeclaration(t *testing.T) {

	p := parser.ParseString("main(){")
	util.AssertNow(t, p.Scope.Type() == ast.FuncDeclaration, "wrong Scope type")
	n := p.Scope.(ast.FuncDeclarationNode)
	util.AssertNow(t, n.Identifier == "main", "wrong identifier")
	util.AssertNow(t, !n.IsAbstract, "wrong abstract")
	util.AssertNow(t, n.Parameters == nil, "parameters should be nil")
	util.AssertNow(t, n.Results == nil, "results should be nil")

}

func TestParseBasicAbstractFuncDeclaration(t *testing.T) {
	p := parser.ParseString("abstract makeNoise(){")
	util.AssertNow(t, p.Scope.Type() == ast.FuncDeclaration, "wrong Scope type")
	n := p.Scope.(ast.FuncDeclarationNode)
	util.AssertNow(t, n.Identifier == "makeNoise", "wrong identifier")
	util.AssertNow(t, n.IsAbstract, "wrong abstract")
}

func TestParseBasicParameterisedFuncDeclaration(t *testing.T) {
	p := parser.ParseString("makeNoise(sound string){")
	util.AssertNow(t, p.Scope.Type() == ast.FuncDeclaration, "wrong Scope type")
	n := p.Scope.(ast.FuncDeclarationNode)
	util.AssertNow(t, n.Identifier == "makeNoise", "wrong identifier")
	util.AssertNow(t, !n.IsAbstract, "wrong abstract")
	util.AssertNow(t, len(n.Parameters) == 1, "wrong parameter length")
}

func TestParseFuncDeclarationMultipleParameters(t *testing.T) {
	p := parser.ParseString("makeNoise(sound string, times int){")
	util.AssertNow(t, p.Scope.Type() == ast.FuncDeclaration, "wrong Scope type")
	n := p.Scope.(ast.FuncDeclarationNode)
	util.AssertNow(t, n.Identifier == "makeNoise", "wrong identifier")
	util.AssertNow(t, !n.IsAbstract, "wrong abstract")
	util.AssertNow(t, len(n.Parameters) == 2, "wrong parameter length")
}

func TestParseFuncDeclarationMultipleParametersSameType(t *testing.T) {
	p := parser.ParseString("makeNoise(sound, loudSound string, volume int){")
	util.AssertNow(t, p.Scope.Type() == ast.FuncDeclaration, "wrong Scope type")
	n := p.Scope.(ast.FuncDeclarationNode)
	util.AssertNow(t, n.Identifier == "makeNoise", "wrong identifier")
	util.AssertNow(t, !n.IsAbstract, "wrong abstract")
	util.AssertNow(t, len(n.Parameters) == 2, "wrong parameter length")
}

func TestParseFuncDeclarationSingleReturn(t *testing.T) {
	p := parser.ParseString("makeNoise(sound, loudSound string, volume int) string {")
	util.AssertNow(t, p.Scope.Type() == ast.FuncDeclaration, "wrong Scope type")
	n := p.Scope.(ast.FuncDeclarationNode)
	util.AssertNow(t, n.Identifier == "makeNoise", "wrong identifier")
	util.AssertNow(t, !n.IsAbstract, "wrong abstract")
	util.AssertNow(t, len(n.Parameters) == 0, "wrong parameter length")
	util.AssertNow(t, len(n.Results) == 1, "wrong result length")
}

func TestParseFuncDeclarationMultipleReturns(t *testing.T) {
	p := parser.ParseString("makeNoise(sound, loudSound string, volume int) (string, int) {")
	util.AssertNow(t, p.Scope.Type() == ast.FuncDeclaration, "wrong Scope type")
	n := p.Scope.(ast.FuncDeclarationNode)
	util.AssertNow(t, n.Identifier == "makeNoise", "wrong identifier")
	util.AssertNow(t, !n.IsAbstract, "wrong abstract")
	util.AssertNow(t, len(n.Parameters) == 0, "wrong parameter length")
	util.AssertNow(t, len(n.Results) == 2, "wrong result length")
}
