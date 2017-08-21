package statements

import (
	"testing"

	"github.com/end-r/guardian/go/compiler/ast"
	"github.com/end-r/guardian/go/compiler/parser"
	"github.com/end-r/guardian/go/util"
)

func TestSimpleLiteralAssignmentStatement(t *testing.T) {
	p := parser.ParseString("x = 5")
	util.AssertNow(t, p.scope.Type() == ast.AssignmentStatement, "statement not recognised")
	n := p.scope.(ast.AssignmentStatementNode)
	util.AssertNow(t, len(n.Left) == 1, "should be one left value")
	util.AssertNow(t, n.Left[0].Type() == ast.Reference, "wrong left type")
	l := p.scope.(ast.ReferenceNode)
	util.Assert(t, len(l.Names) == 1 && l.Names[0] == "x", "wrong left name")

}

func TestMultiToSingleLiteralAssignmentStatement(t *testing.T) {
	p := parser.ParseString("x, y = 5")
	util.AssertNow(t, p.scope.Type() == ast.AssignmentStatement, "statement not recognised")
	n := p.scope.(ast.AssignmentStatementNode)
	util.AssertNow(t, len(n.Left) == 1, "should be two left values")
	util.AssertNow(t, n.Left[0].Type() == ast.Reference, "wrong left type")
	l := p.scope.(ast.ReferenceNode)
	util.Assert(t, len(l.Names) == 1 && l.Names[0] == "x", "wrong left name")
}

func TestMultiLiteralAssignmentStatement(t *testing.T) {
	p := parser.ParseString("x, y = 5, 3")
	util.AssertNow(t, p.scope.Type() == ast.AssignmentStatement, "statement not recognised")
}

func TestSimpleReferenceAssignmentStatement(t *testing.T) {
	p := parser.ParseString("x = a")
	util.AssertNow(t, p.scope.Type() == ast.AssignmentStatement, "statement not recognised")
}

func TestMultiToSingleReferenceAssignmentStatement(t *testing.T) {
	p := parser.ParseString("x, y = a")
	util.AssertNow(t, p.scope.Type() == ast.AssignmentStatement, "statement not recognised")
}

func TestMultiReferenceAssignmentStatement(t *testing.T) {
	p := parser.ParseString("x, y = a, b")
	util.AssertNow(t, p.scope.Type() == ast.AssignmentStatement, "statement not recognised")
}

func TestSimpleCallAssignmentStatement(t *testing.T) {
	p := parser.ParseString("x = a()")
	util.AssertNow(t, p.scope.Type() == ast.AssignmentStatement, "statement not recognised")
}

func TestMultiToSingleCallAssignmentStatement(t *testing.T) {
	p := parser.ParseString("x, y = ab()")
	util.AssertNow(t, p.scope.Type() == ast.AssignmentStatement, "statement not recognised")
}

func TestMultiCallAssignmentStatement(t *testing.T) {
	p := parser.ParseString("x, y = a(), b()")
	util.AssertNow(t, p.scope.Type() == ast.AssignmentStatement, "statement not recognised")
}

func TestSimpleCompositeLiteralAssignmentStatement(t *testing.T) {
	p := parser.ParseString("x = Dog{}")
	util.AssertNow(t, p.scope.Type() == ast.AssignmentStatement, "statement not recognised")
}

func TestMultiToSingleCompositeLiteralAssignmentStatement(t *testing.T) {
	p := parser.ParseString("x, y = Dog{}")
	util.AssertNow(t, p.scope.Type() == ast.AssignmentStatement, "statement not recognised")
}

func TestMultiCompositeLiteralAssignmentStatement(t *testing.T) {
	p := parser.ParseString("x, y = Dog{}, Cat{}")
	util.AssertNow(t, p.scope.Type() == ast.AssignmentStatement, "statement not recognised")
}

func TestSimpleArrayLiteralAssignmentStatement(t *testing.T) {
	p := parser.ParseString("x = [int]{3, 5}")
	util.AssertNow(t, p.scope.Type() == ast.AssignmentStatement, "statement not recognised")
}

func TestMultiToSingleArrayLiteralAssignmentStatement(t *testing.T) {
	p := parser.ParseString("x, y = [int]{3, 5}")
	util.AssertNow(t, p.scope.Type() == ast.AssignmentStatement, "statement not recognised")
}

func TestMultiArrayLiteralAssignmentStatement(t *testing.T) {
	p := parser.ParseString("x, y = [int]{1, 2}, [int]{}")
	util.AssertNow(t, p.scope.Type() == ast.AssignmentStatement, "statement not recognised")
}

func TestSimpleMapLiteralAssignmentStatement(t *testing.T) {
	p := parser.ParseString("x = [int]{3, 5}")
	util.AssertNow(t, p.scope.Type() == ast.AssignmentStatement, "statement not recognised")
}

func TestMultiToSingleMapLiteralAssignmentStatement(t *testing.T) {
	p := parser.ParseString("x, y = [int]{3, 5}")
	util.AssertNow(t, p.scope.Type() == ast.AssignmentStatement, "statement not recognised")
}

func TestMultiMapLiteralAssignmentStatement(t *testing.T) {
	p := parser.ParseString(`x, y = map[int]string{1:"A", 2:"B"}, map[string]int{"A":3, "B": 4}`)
	util.AssertNow(t, p.scope.Type() == ast.AssignmentStatement, "statement not recognised")

}
