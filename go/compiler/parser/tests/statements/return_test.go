package statements

import (
	"axia/guardian/go/compiler/ast"
	"axia/guardian/go/compiler/parser"
	"axia/guardian/go/util"
	"testing"
)

func TestEmptyReturnStatement(t *testing.T) {
	p := parser.ParseString("return")
	util.Assert(t, p.scope.Type() == ast.ReturnStatement, "wrong node type")
}

func TestSingleLiteralReturnStatement(t *testing.T) {
	p := parser.ParseString(`return "twenty"`)
	util.AssertNow(t, p.scope.Type() == ast.ReturnStatement, "wrong node type")
	u := p.scope.(ast.ReturnStatementNode)
	util.AssertNow(t, len(u.Results) == 1, "wrong result length")
	util.AssertNow(t, u.Results[0].Type() == ast.Literal, "wrong literal type")
}

func TestMultipleLiteralReturnStatement(t *testing.T) {
	p := parser.ParseString(`return "twenty", "thirty"`)
	util.AssertNow(t, p.scope.Type() == ast.ReturnStatement, "wrong node type")
	u := p.scope.(ast.ReturnStatementNode)
	util.AssertNow(t, len(u.Results) == 2, "wrong result length")
	util.AssertNow(t, u.Results[0].Type() == ast.Literal, "wrong result 0 type")
	util.AssertNow(t, u.Results[1].Type() == ast.Literal, "wrong result 1 type")
}

func TestSingleReferenceReturnStatement(t *testing.T) {
	p := parser.ParseString(`return twenty`)
	util.AssertNow(t, p.scope.Type() == ast.ReturnStatement, "wrong node type")
	u := p.scope.(ast.ReturnStatementNode)
	util.AssertNow(t, len(u.Results) == 1, "wrong result length")
	util.AssertNow(t, u.Results[0].Type() == ast.Reference, "wrong result 0 type")
}

func TestMultipleReferenceReturnStatement(t *testing.T) {
	p := parser.ParseString(`return twenty, thirty`)
	util.AssertNow(t, p.scope.Type() == ast.ReturnStatement, "wrong node type")
	u := p.scope.(ast.ReturnStatementNode)
	util.AssertNow(t, len(u.Results) == 2, "wrong result length")
	util.AssertNow(t, u.Results[0].Type() == ast.Reference, "wrong result 0 type")
	util.AssertNow(t, u.Results[1].Type() == ast.Reference, "wrong result 1 type")
}

func TestSingleCallReturnStatement(t *testing.T) {
	p := parser.ParseString(`return param()`)
	util.AssertNow(t, p.scope.Type() == ast.ReturnStatement, "wrong node type")
	u := p.scope.(ast.ReturnStatementNode)
	util.AssertNow(t, len(u.Results) == 1, "wrong result length")
	util.AssertNow(t, u.Results[0].Type() == ast.CallExpression, "wrong result 0 type")
}

func TestMultipleCallReturnStatement(t *testing.T) {
	p := parser.ParseString(`return a(param, "param"), b()`)
	util.AssertNow(t, p.scope.Type() == ast.ReturnStatement, "wrong node type")
	u := p.scope.(ast.ReturnStatementNode)
	util.AssertNow(t, len(u.Results) == 2, "wrong result length")
	util.AssertNow(t, u.Results[0].Type() == ast.CallExpression, "wrong result 0 type")
	util.AssertNow(t, u.Results[1].Type() == ast.CallExpression, "wrong result 1 type")
}

func TestSingleArrayLiterallReturnStatement(t *testing.T) {
	p := parser.ParseString(`return [int]{}`)
	util.AssertNow(t, p.scope.Type() == ast.ReturnStatement, "wrong node type")
	u := p.scope.(ast.ReturnStatementNode)
	util.AssertNow(t, len(u.Results) == 1, "wrong result length")
	util.AssertNow(t, u.Results[0].Type() == ast.ArrayLiteral, "wrong result 0 type")
}

func TestMultipleArrayLiteralReturnStatement(t *testing.T) {
	p := parser.ParseString(`return [string]{"one", "two"}, [Dog]{}`)
	util.AssertNow(t, p.scope.Type() == ast.ReturnStatement, "wrong node type")
	u := p.scope.(ast.ReturnStatementNode)
	util.AssertNow(t, len(u.Results) == 2, "wrong result length")
	util.AssertNow(t, u.Results[0].Type() == ast.ArrayLiteral, "wrong result 0 type")
	util.AssertNow(t, u.Results[1].Type() == ast.ArrayLiteral, "wrong result 1 type")
}

func TestSingleMapLiterallReturnStatement(t *testing.T) {
	p := parser.ParseString(`return map[string]int{"one":2, "two":3}`)
	util.AssertNow(t, p.scope.Type() == ast.ReturnStatement, "wrong node type")
	u := p.scope.(ast.ReturnStatementNode)
	util.AssertNow(t, len(u.Results) == 1, "wrong result length")
	util.AssertNow(t, u.Results[0].Type() == ast.ArrayLiteral, "wrong result 0 type")
}

func TestMultipleMapLiteralReturnStatement(t *testing.T) {
	p := parser.ParseString(`return map[string]int{"one":2, "two":3}, map[int]Dog{}`)
	util.AssertNow(t, p.scope.Type() == ast.ReturnStatement, "wrong node type")
	u := p.scope.(ast.ReturnStatementNode)
	util.AssertNow(t, len(u.Results) == 2, "wrong result length")
	util.AssertNow(t, u.Results[0].Type() == ast.ArrayLiteral, "wrong result 0 type")
	util.AssertNow(t, u.Results[1].Type() == ast.ArrayLiteral, "wrong result 1 type")
}

func TestSingleCompositeLiterallReturnStatement(t *testing.T) {
	p := parser.ParseString(`return Dog{}`)
	util.AssertNow(t, p.scope.Type() == ast.ReturnStatement, "wrong node type")
	u := p.scope.(ast.ReturnStatementNode)
	util.AssertNow(t, len(u.Results) == 1, "wrong result length")
	util.AssertNow(t, u.Results[0].Type() == ast.ArrayLiteral, "wrong result 0 type")
}

func TestMultipleCompositeLiteralReturnStatement(t *testing.T) {
	p := parser.ParseString(`return Cat{name:"Doggo"}, Dog{name:"Katter"}`)
	util.AssertNow(t, p.scope.Type() == ast.ReturnStatement, "wrong node type")
	u := p.scope.(ast.ReturnStatementNode)
	util.AssertNow(t, len(u.Results) == 2, "wrong result length")
	util.AssertNow(t, u.Results[0].Type() == ast.ArrayLiteral, "wrong result 0 type")
	util.AssertNow(t, u.Results[1].Type() == ast.ArrayLiteral, "wrong result 1 type")
}
