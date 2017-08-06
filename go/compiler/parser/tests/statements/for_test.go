package statements

import (
	"axia/guardian/go/compiler/ast"
	"axia/guardian/go/util"
	"testing"
)

func TestBooleanExpressionFor(t *testing.T) {
	p := ParseString("for x > 5 { }")
	util.AssertNow(t, p.scope.Type() == ast.ForStatement, "wrong node type")
}

func TestInitExpressionFor(t *testing.T) {
	p := ParseString("for x := 2; x > 5 { }")
	util.AssertNow(t, p.scope.Type() == ast.ForStatement, "wrong node type")
}

func TestFullExpressionFor(t *testing.T) {
	p := ParseString("for x := 0; x < 5; x++ { }")
	util.AssertNow(t, p.scope.Type() == ast.ForStatement, "wrong node type")
}
