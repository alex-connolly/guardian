package statements

import (
	"axia/guardian/go/compiler/ast"
	"axia/guardian/go/util"
	"testing"
)

func TestSinglePartExpressionFor(t *testing.T) {
	p := ParseString("for x > 5 { }")
	util.AssertNow(t, p.scope.Type() == ast.ForStatement, "wrong node type")
	u := p.scope.(ast.ForStatementNode)
	util.AssertNow(t, u.Init == nil, "should not have init node")
}

func TestInitExpressionFor(t *testing.T) {
	p := ParseString("for x := 2; x > 5 { }")
	util.AssertNow(t, p.scope.Type() == ast.ForStatement, "wrong node type")
}

func TestFullExpressionFor(t *testing.T) {
	p := ParseString("for x := 0; x < 5; x++ { }")
	util.AssertNow(t, p.scope.Type() == ast.ForStatement, "wrong node type")
}
