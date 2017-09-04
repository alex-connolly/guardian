package parser

import (
	"testing"

	"github.com/end-r/guardian/go/compiler/lexer"

	"github.com/end-r/guardian/go/compiler/ast"

	"github.com/end-r/goutil"
)

func TestParseReferenceSingle(t *testing.T) {
	p := createParser(`hello`)
	goutil.AssertNow(t, p.lexer != nil, "lexer should not be nil")
	goutil.AssertNow(t, len(p.lexer.Tokens) == 1, "wrong token length")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr should not be nil")
	goutil.AssertNow(t, expr.Type() == ast.Reference, "wrong expr type")
	ref := expr.(ast.ReferenceNode)
	goutil.AssertNow(t, ref.Names != nil, "ref should not be nil")
	goutil.AssertNow(t, len(ref.Names) == 1, "wrong name length")
	goutil.Assert(t, ref.Names[0] == "hello", "wrong name data")
}

func TestParseReferenceMultiple(t *testing.T) {
	p := createParser(`hello.aaa.bb`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 5, "wrong token length")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr should not be nil")
	goutil.AssertNow(t, expr.Type() == ast.Reference, "wrong expr type")
	ref := expr.(ast.ReferenceNode)
	goutil.AssertNow(t, ref.Names != nil, "ref should not be nil")
	goutil.AssertNow(t, len(ref.Names) == 3, "wrong name length")
	goutil.Assert(t, ref.Names[0] == "hello", "wrong name data 0")
	goutil.Assert(t, ref.Names[1] == "aaa", "wrong name data 1")
	goutil.Assert(t, ref.Names[2] == "bb", "wrong name data 2")
}

func TestParseLiteralInteger(t *testing.T) {
	p := createParser(`6`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 1, "wrong token length")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr should not be nil")
	goutil.AssertNow(t, expr.Type() == ast.Literal, "wrong expr type")
	lit := expr.(ast.LiteralNode)
	goutil.AssertNow(t, lit.LiteralType == lexer.TknNumber, "wrong literal type")
}

func TestParseLiteralString(t *testing.T) {
	p := createParser(`"alex"`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 1, "wrong token length")
	expr := p.parseExpression()
	goutil.AssertNow(t, expr != nil, "expr should not be nil")
	goutil.AssertNow(t, expr.Type() == ast.Literal, "wrong expr type")
	lit := expr.(ast.LiteralNode)
	goutil.AssertNow(t, lit.LiteralType == lexer.TknString, "wrong literal type")
}
