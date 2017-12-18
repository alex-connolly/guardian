package parser

import (
	"testing"

	"github.com/end-r/guardian/ast"
)

var result *ast.ScopeNode

func Benchmark1KFile(b *testing.B) {
	var a *ast.ScopeNode
	for n := 0; n < b.N; n++ {
		a, _ = ParseFile("tests/benchmarks/1k.grd")
	}
	result = a
}

func BenchmarkParseExpVarSimpleFuncType(b *testing.B) {
	for n := 0; n < b.N; n++ {
		p := createParser("a func()")
		parseExplicitVarDeclaration(p)
	}
}

func BenchmarkParseExpVarComplexFuncType(b *testing.B) {
	for n := 0; n < b.N; n++ {
		p := createParser("a func(c int, b func(string, string) (int, int), d int) (int, float, int)")
		parseExplicitVarDeclaration(p)
	}
}
