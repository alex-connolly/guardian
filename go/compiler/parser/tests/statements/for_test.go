package statements

import (
	"axia/guardian/go/compiler/parser"
	"testing"

	"github.com/end-r/guardian/go/compiler/ast"
	"github.com/end-r/guardian/go/util"
)

func TestSinglePartExpressionFor(t *testing.T) {
	p := parser.ParseString("for x > 5 { }")
	util.AssertNow(t, p.scope.Type() == ast.ForStatement, "wrong node type")
	u := p.scope.(ast.ForStatementNode)
	util.AssertNow(t, u.Init == nil, "should not have init node")
	util.AssertNow(t, u.Post == nil, "should not have post node")

}

func TestInitExpressionFor(t *testing.T) {
	p := parser.ParseString("for x := 2; x > 5 { }")
	util.AssertNow(t, p.scope.Type() == ast.ForStatement, "wrong node type")
	u := p.scope.(ast.ForStatementNode)
	util.AssertNow(t, u.Post == nil, "should not have post node")
}

func TestFullExpressionFor(t *testing.T) {
	p := parser.ParseString("for x := 0; x < 5; x++ { }")
	util.AssertNow(t, p.scope.Type() == ast.ForStatement, "wrong node type")
	u := p.scope.(ast.ForStatementNode)
}
