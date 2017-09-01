package parser

import (
	"testing"

	"github.com/end-r/goutil"
)

func TestParseReferenceSingle(t *testing.T) {
	p := createParser(`hello`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 1, "wrong token length")
	p.parseExpression()
}

func TestParseReferenceMultiple(t *testing.T) {
	p := createParser(`hello.aaa.bb`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 5, "wrong token length")
	p.parseExpression()
}

func TestParseConstant(t *testing.T) {
	p := createParser(`6`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 1, "wrong token length")
	p.parseExpression()
}
