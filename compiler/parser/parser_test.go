package parser

import (
	"testing"

	"github.com/end-r/goutil"
)

// mini-parser tests belong here

func TestHasTokens(t *testing.T) {
	p := createParser("this is data")
	goutil.Assert(t, len(p.lexer.Tokens) == 3, "wrong length of tokens")
	goutil.Assert(t, p.hasTokens(0), "should have 0 tokens")
	goutil.Assert(t, p.hasTokens(1), "should have 1 tokens")
	goutil.Assert(t, p.hasTokens(2), "should have 2 tokens")
	goutil.Assert(t, p.hasTokens(3), "should have 3 tokens")
	goutil.Assert(t, !p.hasTokens(4), "should not have 4 tokens")
}

func TestParseIdentifier(t *testing.T) {
	p := createParser("identifier")
	goutil.Assert(t, p.parseIdentifier() == "identifier", "wrong identifier")
	p = createParser("")
	goutil.Assert(t, p.parseIdentifier() == "", "empty should be nil")
	p = createParser("{")
	goutil.Assert(t, p.parseIdentifier() == "", "wrong token should be nil")
}
